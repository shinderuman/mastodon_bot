package mastodon

import (
	"context"
	"fmt"
	"log"
	"slices"

	strip "github.com/grokify/html-strip-tags-go"
	"github.com/mattn/go-mastodon"

	"mastodon_bot/internal/config"
)

type Client struct {
	client *mastodon.Client
}

// NewMastodonClient creates a new Mastodon client with the provided configuration
func NewMastodonClient(cfg *config.MastodonConfig) *Client {
	mastodonClient := mastodon.NewClient(&mastodon.Config{
		Server:       cfg.Server,
		ClientID:     cfg.ClientID,
		ClientSecret: cfg.ClientSecret,
		AccessToken:  cfg.AccessToken,
	})

	return &Client{client: mastodonClient}
}

// StreamAndProcessPosts streams Mastodon posts and processes them with the given processor function
func (c *Client) StreamAndProcessPosts(ctx context.Context, processor func(string) string, targetUsers []string) error {
	webSocketClient := c.client.NewWSClient()
	eventQueue, err := webSocketClient.StreamingWSUser(ctx)
	if err != nil {
		return fmt.Errorf("failed to start streaming: %w", err)
	}

	log.Println("Started streaming Mastodon posts")

	for event := range eventQueue {
		updateEvent, ok := event.(*mastodon.UpdateEvent)
		if !ok {
			continue
		}

		if updateEvent.Status.Sensitive {
			log.Printf("Skipping sensitive post from @%s", updateEvent.Status.Account.Username)
			continue
		}

		// 特定ユーザーのみに反応
		if !slices.Contains(targetUsers, updateEvent.Status.Account.Username) {
			continue
		}

		// HTMLタグを除去
		cleanedContent := strip.StripTags(updateEvent.Status.Content)
		fmt.Printf("\x1b[37m[%s] \x1b[35m%-20s: \x1b[33m%s\n",
			updateEvent.Status.CreatedAt.Local().Format("15:04:05"),
			updateEvent.Status.Account.Acct,
			cleanedContent)

		// テキストを処理
		processedText := processor(cleanedContent)
		if processedText == "" {
			log.Printf("Processor returned empty result for post from @%s", updateEvent.Status.Account.Username)
			continue
		}

		// 投稿
		if err := c.PostProcessedToot(ctx, processedText); err != nil {
			log.Printf("Failed to post processed content: %v", err)
		} else {
			log.Printf("Successfully posted processed content: %s", processedText)
		}
	}

	return nil
}

// PostProcessedToot posts the processed text as an unlisted toot
func (c *Client) PostProcessedToot(ctx context.Context, text string) error {
	_, err := c.client.PostStatus(ctx, &mastodon.Toot{
		Status:     text,
		Visibility: "unlisted",
	})
	if err != nil {
		return fmt.Errorf("failed to post toot: %w", err)
	}
	return nil
}

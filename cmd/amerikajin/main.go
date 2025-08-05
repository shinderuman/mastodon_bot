package main

import (
	"context"
	"log"

	"mastodon_bot/internal/config"
	"mastodon_bot/internal/mastodon"
	"mastodon_bot/internal/processor"
)

func main() {
	cfg, err := config.LoadConfigFromFile("config.json")
	if err != nil {
		log.Fatal("Failed to load configuration:", err)
	}

	ctx := context.Background()

	// 翻訳プロセッサーを初期化
	amerikajinTranslator, err := processor.NewAmerikajinTranslator(ctx)
	if err != nil {
		log.Fatal("Failed to initialize translation client:", err)
	}
	defer amerikajinTranslator.Close()

	// Mastodonクライアントを初期化
	mastodonClient := mastodon.NewMastodonClient(&cfg.Mastodon)

	// ストリーミング開始
	log.Println("Amerikajin bot started - English translation mode")
	if err := mastodonClient.StreamAndProcessPosts(ctx, amerikajinTranslator.ProcessTextToEnglish, cfg.Mastodon.TargetUsers); err != nil {
		log.Fatal("Streaming error:", err)
	}
}

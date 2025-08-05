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

	// 擬音語プロセッサーを初期化
	gienkunProcessor := processor.NewGienkunProcessor()

	// Mastodonクライアントを初期化
	mastodonClient := mastodon.NewMastodonClient(&cfg.Mastodon)

	// ストリーミング開始
	log.Println("Gienkun bot started - Onomatopoeia conversion mode")
	if err := mastodonClient.StreamAndProcessPosts(ctx, gienkunProcessor.ProcessTextToOnomatopoeia, cfg.Mastodon.TargetUsers); err != nil {
		log.Fatal("Streaming error:", err)
	}
}

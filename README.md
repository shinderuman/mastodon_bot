# Mastodon Bot Collection

A collection of Mastodon bots that process and transform text from specific users' posts using various text processing techniques including morphological analysis, translation, and text transformation.

## Features

- **アメリカジンbot**: Translates Japanese text to English using Google Translate API
- **原始人くんbot**: Transforms text using primitive/caveman-style speech patterns
- **魏延くんbot**: Converts text to onomatopoeia using morphological analysis (kagome tokenizer)

## Architecture

The project follows Go's standard project layout:

```
mastodon_bot/
├── cmd/                     # Command-line applications
│   ├── amerikajin/         # アメリカジンbot
│   ├── genshijin/          # 原始人くんbot
│   └── gienkun/            # 魏延くんbot
├── internal/               # Private application code
│   ├── config/            # Configuration management
│   ├── mastodon/          # Mastodon client wrapper
│   └── processor/         # Text processing modules
├── config.json.example    # Configuration template
└── go.mod                 # Go module definition
```

## Setup

1. Clone the repository
2. Copy the configuration template:
   ```bash
   cp config.json.example config.json
   ```
3. Edit `config.json` with your Mastodon credentials and target users
4. Install dependencies:
   ```bash
   go mod tidy
   ```

## Configuration

Edit `config.json`:

```json
{
    "mastodon": {
        "server": "https://your-mastodon-server.com",
        "clientId": "your_client_id_here",
        "clientSecret": "your_client_secret_here",
        "accessToken": "your_access_token_here",
        "targetUsers": ["target_username"]
    }
}
```

## Usage

Run individual bots:

```bash
# アメリカジンbot
go run ./cmd/amerikajin

# 原始人くんbot
go run ./cmd/genshijin

# 魏延くんbot
go run ./cmd/gienkun
```

Or build binaries:

```bash
go build ./cmd/amerikajin
go build ./cmd/genshijin
go build ./cmd/gienkun
```

## License

MIT License

---

# Mastodon Bot コレクション

特定ユーザーの投稿に対して、形態素解析、翻訳、テキスト変換などの様々なテキスト処理技術を使用してテキストを処理・変換するMastodonボットのコレクションです。

## 機能

- **アメリカジンbot**: Google Translate APIを使用して日本語テキストを英語に翻訳
- **原始人くんbot**: 原始人風の話し方にテキストを変換
- **魏延くんbot**: 形態素解析（kagomeトークナイザー）を使用してテキストを擬音語に変換

## アーキテクチャ

プロジェクトはGoの標準的なプロジェクトレイアウトに従っています：

```
mastodon_bot/
├── cmd/                     # コマンドラインアプリケーション
│   ├── amerikajin/         # アメリカジンbot
│   ├── genshijin/          # 原始人くんbot
│   └── gienkun/            # 魏延くんbot
├── internal/               # プライベートアプリケーションコード
│   ├── config/            # 設定管理
│   ├── mastodon/          # Mastodonクライアントラッパー
│   └── processor/         # テキスト処理モジュール
├── config.json.example    # 設定テンプレート
└── go.mod                 # Goモジュール定義
```

## セットアップ

1. リポジトリをクローン
2. 設定テンプレートをコピー：
   ```bash
   cp config.json.example config.json
   ```
3. Mastodonの認証情報とターゲットユーザーで`config.json`を編集
4. 依存関係をインストール：
   ```bash
   go mod tidy
   ```

## 設定

`config.json`を編集：

```json
{
    "mastodon": {
        "server": "https://your-mastodon-server.com",
        "clientId": "your_client_id_here",
        "clientSecret": "your_client_secret_here",
        "accessToken": "your_access_token_here",
        "targetUsers": ["target_username"]
    }
}
```

## 使用方法

個別のボットを実行：

```bash
# アメリカジンbot
go run ./cmd/amerikajin

# 原始人くんbot
go run ./cmd/genshijin

# 魏延くんbot
go run ./cmd/gienkun
```

またはバイナリをビルド：

```bash
go build ./cmd/amerikajin
go build ./cmd/genshijin
go build ./cmd/gienkun
```

## ライセンス

MIT License
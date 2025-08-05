package processor

import (
	"strings"
	"unicode"

	"github.com/ikawaha/kagome/tokenizer"
)

type Gienkun struct {
	tokenizer tokenizer.Tokenizer
}

// NewGienkunProcessor creates a new Gienkun onomatopoeia processor with morphological analyzer
func NewGienkunProcessor() *Gienkun {
	return &Gienkun{
		tokenizer: tokenizer.New(),
	}
}

// ProcessTextToOnomatopoeia converts text to onomatopoeia using morphological analysis
func (g *Gienkun) ProcessTextToOnomatopoeia(text string) string {
	cleanedText := strings.TrimSpace(text)
	if cleanedText == "" {
		return ""
	}

	tokens := g.tokenizer.Tokenize(cleanedText)
	var onomatopoeiaItems []string

	// カナ変換用の特殊ケース定義
	kanaConverter := unicode.SpecialCase{
		// ひらがなをカタカナに変換
		unicode.CaseRange{
			Lo: 0x3041, // ぁ
			Hi: 0x3093, // ん
			Delta: [unicode.MaxCase]rune{
				0x30a1 - 0x3041, // UpperCase でカタカナに変換
				0,               // LowerCase では変換しない
				0x30a1 - 0x3041, // TitleCase でカタカナに変換
			},
		},
		// カタカナをひらがなに変換
		unicode.CaseRange{
			Lo: 0x30a1, // ァ
			Hi: 0x30f3, // ン
			Delta: [unicode.MaxCase]rune{
				0,               // UpperCase では変換しない
				0x3041 - 0x30a1, // LowerCase でひらがなに変換
				0,               // TitleCase では変換しない
			},
		},
	}

	for _, token := range tokens {
		features := token.Features()

		// 未知語でない場合の処理
		if token.Class != tokenizer.UNKNOWN {
			if len(features) == 0 || token.Surface == "BOS" || token.Surface == "EOS" {
				continue
			}
			// 助詞や読みが不明な語をスキップ
			if features[0] == "助詞" || features[6] == "*" {
				continue
			}
		}

		// カタカナに変換して擬音語として追加
		onomatopoeiaText := strings.ToUpperSpecial(kanaConverter, token.Surface)
		onomatopoeiaItems = append(onomatopoeiaItems, onomatopoeiaText)
	}

	return strings.Join(onomatopoeiaItems, "…")
}

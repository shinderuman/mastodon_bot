package processor

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/translate"
	"golang.org/x/text/language"
)

type Amerikajin struct {
	client *translate.Client
}

// NewAmerikajinTranslator creates a new Amerikajin translator with Google Translate client
func NewAmerikajinTranslator(ctx context.Context) (*Amerikajin, error) {
	translateClient, err := translate.NewClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create Google Translate client: %w", err)
	}

	return &Amerikajin{client: translateClient}, nil
}

func (a *Amerikajin) Close() error {
	return a.client.Close()
}

// ProcessTextToEnglish translates Japanese text to English
func (a *Amerikajin) ProcessTextToEnglish(text string) string {
	if text == "" {
		log.Println("Empty text provided for translation")
		return ""
	}

	translationResponse, err := a.client.Translate(context.Background(), []string{text}, language.English, nil)
	if err != nil {
		log.Printf("Failed to translate text '%s': %v", text, err)
		return ""
	}

	if len(translationResponse) == 0 {
		log.Printf("No translation results returned for text: %s", text)
		return ""
	}

	translatedText := translationResponse[0].Text
	log.Printf("Successfully translated text: '%s' -> '%s'", text, translatedText)
	return translatedText
}

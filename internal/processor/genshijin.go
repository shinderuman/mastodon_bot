package processor

import (
	"github.com/mattn/genshijin"
)

type Genshijin struct{}

// NewGenshijinProcessor creates a new Genshijin text processor
func NewGenshijinProcessor() *Genshijin {
	return &Genshijin{}
}

// ProcessTextToPrimitive converts text to primitive/caveman-style speech
func (g *Genshijin) ProcessTextToPrimitive(text string) string {
	if text == "" {
		return ""
	}
	primitiveText := genshijin.Shaberu(text)
	return primitiveText
}

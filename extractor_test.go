package gec

import (
	"testing"
)

func TestExtractTitle(t *testing.T) {
	opt := NewOption()
	extractor := newExtractor(opt)

	actuals := []string{
		extractor.ExtractTitle(""),
		extractor.ExtractTitle("<html><head><title>hoge</title></head><html>"),
	}
	expecteds := []string{
		"",
		"hoge",
	}

	for i := range actuals {
		if actuals[i] != expecteds[i] {
			t.Errorf("expected %v, but got %v", expecteds[i], actuals[i])
		}
	}
}

func TestExtractContent(t *testing.T) {
	opt := NewOption()
	extractor := newExtractor(opt)

	actuals := []string{
		extractor.ExtractContent(""),
		extractor.ExtractContent("<html><head><title>hoge</title></head><html>"),
	}
	expecteds := []string{
		"",
		"",
	}

	for i := range actuals {
		if actuals[i] != expecteds[i] {
			t.Errorf("expected %v, but got %v", expecteds[i], actuals[i])
		}
	}
}

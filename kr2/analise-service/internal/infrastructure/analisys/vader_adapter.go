package analysis

import (
	"context"

	"github.com/Babushkin05/software-dev-course/kr2/analise-service/internal/application/ports/output"
	"github.com/jonreiter/govader"
)

type VaderAdapter struct {
	analyzer *govader.SentimentIntensityAnalyzer
}

func NewVaderAdapter() output.SentimentAnalyzer {
	a := govader.NewSentimentIntensityAnalyzer()
	return &VaderAdapter{analyzer: a}
}

func (v *VaderAdapter) AnalyzeSentiment(ctx context.Context, text string) (float64, float64, float64, float64, error) {
	scores := v.analyzer.PolarityScores(text)
	return scores.Positive, scores.Negative, scores.Neutral, scores.Compound, nil
}

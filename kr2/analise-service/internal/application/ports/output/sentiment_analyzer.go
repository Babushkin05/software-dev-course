package output

import "context"

type SentimentAnalyzer interface {
	AnalyzeSentiment(ctx context.Context, text string) (positive, negative, neutral, compound float64, err error)
}

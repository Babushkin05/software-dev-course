package entities

import (
	"fmt"
)

type AnalysisResult struct {
	Positive float64
	Negative float64
	Neutral  float64
	Compound float64
}

func (r AnalysisResult) Summary() string {
	return fmt.Sprintf(
		"Positive: %.2f%%, Negative: %.2f%%, Neutral: %.2f%%, Compound: %.4f",
		r.Positive*100,
		r.Negative*100,
		r.Neutral*100,
		r.Compound,
	)
}

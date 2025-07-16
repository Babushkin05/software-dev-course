package services

import (
	"context"
	"fmt"

	"github.com/Babushkin05/software-dev-course/kr2/analise-service/internal/application/ports/input"
	"github.com/Babushkin05/software-dev-course/kr2/analise-service/internal/application/ports/output"
	"github.com/Babushkin05/software-dev-course/kr2/analise-service/internal/domain/entities"
)

type AnalysisService struct {
	storingClient     output.StoringClient
	sentimentAnalyzer output.SentimentAnalyzer
	wordCloudGen      output.WordCloudGenerator
}

func NewAnalysisService(
	storingClient output.StoringClient,
	sentimentAnalyzer output.SentimentAnalyzer,
	wordCloudGen output.WordCloudGenerator,
) input.AnalysisUseCase {
	return &AnalysisService{
		storingClient:     storingClient,
		sentimentAnalyzer: sentimentAnalyzer,
		wordCloudGen:      wordCloudGen,
	}
}

// Analyze делает анализ текста и возвращает текстовую сводку
func (s *AnalysisService) Analyze(ctx context.Context, fileID string) (string, string, error) {
	filename, content, err := s.storingClient.DownloadFile(ctx, fileID)
	if err != nil {
		return "", "", fmt.Errorf("failed to download file: %w", err)
	}

	pos, neg, neu, comp, err := s.sentimentAnalyzer.AnalyzeSentiment(ctx, string(content))
	if err != nil {
		return "", "", fmt.Errorf("sentiment analysis failed: %w", err)
	}

	result := entities.AnalysisResult{
		Positive: pos,
		Negative: neg,
		Neutral:  neu,
		Compound: comp,
	}

	outputFile := fmt.Sprintf("analyze_%s.txt", filename)
	return outputFile, result.Summary(), nil
}

// GenerateWordCloud возвращает PNG изображение облака слов
func (s *AnalysisService) GenerateWordCloud(ctx context.Context, fileID string) (string, []byte, error) {
	filename, content, err := s.storingClient.DownloadFile(ctx, fileID)
	if err != nil {
		return "", nil, fmt.Errorf("failed to download file: %w", err)
	}

	imgData, err := s.wordCloudGen.GenerateImage(ctx, string(content))
	if err != nil {
		return "", nil, fmt.Errorf("failed to generate word cloud: %w", err)
	}

	outputFile := fmt.Sprintf("cloud_%s.png", filename)
	return outputFile, imgData, nil
}

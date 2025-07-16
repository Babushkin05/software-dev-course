package analysis

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Babushkin05/software-dev-course/kr2/analise-service/internal/application/ports/output"
)

type WordCloudAdapter struct {
	client *http.Client
}

func NewWordCloudAdapter() output.WordCloudGenerator {
	return &WordCloudAdapter{
		client: &http.Client{},
	}
}

func (w *WordCloudAdapter) GenerateImage(ctx context.Context, text string) ([]byte, error) {
	payload := map[string]interface{}{
		"format":     "png",
		"width":      1000,
		"height":     1000,
		"fontFamily": "sans-serif",
		"fontScale":  15,
		"scale":      "linear",
		"text":       text,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", "https://quickchart.io/wordcloud", bytes.NewReader(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := w.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to perform wordcloud request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("wordcloud API returned %d: %s", resp.StatusCode, string(body))
	}

	imgData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read wordcloud response: %w", err)
	}

	return imgData, nil
}

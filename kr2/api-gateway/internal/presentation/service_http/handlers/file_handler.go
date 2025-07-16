package handlers

import (
	"net/http"

	"github.com/Babushkin05/software-dev-course/kr2/api-gateway/internal/domain/ports/input"

	"github.com/gin-gonic/gin"
)

type FileHandler struct {
	Usecase input.FileUsecase
}

func NewFileHandler(usecase input.FileUsecase) *FileHandler {
	return &FileHandler{Usecase: usecase}
}

func (h *FileHandler) RegisterRoutes(r *gin.Engine) {
	r.POST("/upload", h.Upload)
	r.POST("/analyze/:file_id", h.Analyze)
	r.GET("/download/:file_id", h.Download)
	r.GET("/wordcloud/:file_id", h.WordCloud)
}

func (h *FileHandler) Upload(c *gin.Context) {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "file not provided"})
		return
	}

	file, err := fileHeader.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to open file"})
		return
	}
	defer file.Close()

	fileID, err := h.Usecase.UploadFile(c.Request.Context(), file, fileHeader.Filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "upload failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"file_id": fileID})
}

func (h *FileHandler) Analyze(c *gin.Context) {
	fileID := c.Param("file_id")

	content, filename, err := h.Usecase.AnalyzeFile(c.Request.Context(), fileID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "analysis failed"})
		return
	}

	c.Header("Content-Disposition", "attachment; filename="+filename)
	c.Data(http.StatusOK, "text/plain", content)
}

func (h *FileHandler) Download(c *gin.Context) {
	fileID := c.Param("file_id")

	content, filename, err := h.Usecase.DownloadFile(c.Request.Context(), fileID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "download failed"})
		return
	}

	c.Header("Content-Disposition", "attachment; filename="+filename)
	c.Data(http.StatusOK, "text/plain", content)
}

func (h *FileHandler) WordCloud(c *gin.Context) {
	fileID := c.Param("file_id")

	content, filename, err := h.Usecase.GetWordCloud(c.Request.Context(), fileID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "word cloud generation failed"})
		return
	}

	c.Header("Content-Disposition", "attachment; filename="+filename)
	c.Data(http.StatusOK, "application/pdf", content)
}

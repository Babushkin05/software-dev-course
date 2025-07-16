package postgres

import (
	"database/sql"
	"fmt"

	"github.com/Babushkin05/software-dev-course/kr2/storing-service/internal/domain/entities"
)

type FileRepo struct {
	db *sql.DB
}

func NewFileRepo(db *sql.DB) *FileRepo {
	return &FileRepo{db: db}
}

func (r *FileRepo) SaveMetadata(file *entities.File) error {
	query := `
		INSERT INTO files (id, filename, size, created_at)
		VALUES ($1, $2, $3, $4)
	`
	_, err := r.db.Exec(query, file.ID, file.Filename, file.Size, file.CreatedAt)
	return err
}

func (r *FileRepo) GetMetadata(fileID string) (*entities.File, error) {
	query := `
		SELECT id, filename, size, created_at
		FROM files
		WHERE id = $1
	`

	row := r.db.QueryRow(query, fileID)

	var file entities.File
	err := row.Scan(&file.ID, &file.Filename, &file.Size, &file.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("file with ID %s not found", fileID)
		}
		return nil, err
	}

	return &file, nil
}

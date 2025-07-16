package entities

import "time"

type File struct {
	ID        string    // UUID
	Filename  string    // Имя файла при загрузке
	Size      int64     // Размер файла в байтах
	CreatedAt time.Time // Время загрузки
}

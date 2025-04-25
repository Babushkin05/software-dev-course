package storage_test

import (
	"testing"

	"github.com/Babushkin05/software-dev-course/hw2/internal/domain"
	"github.com/Babushkin05/software-dev-course/hw2/internal/infrastructure/storage"
	"github.com/stretchr/testify/assert"
)

func TestInMemoryEnclosureRepo(t *testing.T) {
	// Создаем тестовые данные
	testEnclosure := &domain.Enclosure{
		ID:       "e1",
		Type:     domain.Predator,
		Capacity: 5,
	}

	t.Run("Create and Get", func(t *testing.T) {
		repo := storage.NewInMemoryEnclosureRepo()

		// Сохраняем вольер
		err := repo.Save(testEnclosure)
		assert.NoError(t, err)

		// Получаем вольер
		enclosure, err := repo.GetByID("e1")
		assert.NoError(t, err)
		assert.Equal(t, testEnclosure, enclosure)

		// Проверяем несуществующий вольер
		_, err = repo.GetByID("e99")
		assert.ErrorIs(t, err, domain.ErrEnclosureNotFound)
	})

	t.Run("Update", func(t *testing.T) {
		repo := storage.NewInMemoryEnclosureRepo()
		_ = repo.Save(testEnclosure)

		// Обновляем вольер
		updatedEnclosure := *testEnclosure
		updatedEnclosure.Capacity = 10
		err := repo.Save(&updatedEnclosure)
		assert.NoError(t, err)

		// Проверяем обновление
		enclosure, err := repo.GetByID("e1")
		assert.NoError(t, err)
		assert.Equal(t, 10, enclosure.Capacity)
	})

	t.Run("Delete", func(t *testing.T) {
		repo := storage.NewInMemoryEnclosureRepo()
		_ = repo.Save(testEnclosure)

		// Удаляем вольер
		err := repo.Delete("e1")
		assert.NoError(t, err)

		// Проверяем что удалилось
		_, err = repo.GetByID("e1")
		assert.ErrorIs(t, err, domain.ErrEnclosureNotFound)

		// Удаление несуществующего вольера
		err = repo.Delete("e99")
		assert.NoError(t, err) // Удаление несуществующего не должно возвращать ошибку
	})

	t.Run("List", func(t *testing.T) {
		repo := storage.NewInMemoryEnclosureRepo()

		// Пустой список
		list, err := repo.List()
		assert.NoError(t, err)
		assert.Empty(t, list)

		// Добавляем вольеры
		enclosures := []*domain.Enclosure{
			{ID: "e1", Type: domain.Predator, Capacity: 5},
			{ID: "e2", Type: domain.Herbivore, Capacity: 10},
		}

		for _, e := range enclosures {
			_ = repo.Save(e)
		}

		// Проверяем список
		list, err = repo.List()
		assert.NoError(t, err)
		assert.Len(t, list, 2)
		assert.ElementsMatch(t, enclosures, list)
	})

	t.Run("ConcurrentAccess", func(t *testing.T) {
		repo := storage.NewInMemoryEnclosureRepo()
		const numRoutines = 100

		// Тест на конкурентные записи
		t.Run("ParallelWrites", func(t *testing.T) {
			for i := 0; i < numRoutines; i++ {
				go func(id int) {
					e := &domain.Enclosure{
						ID:   string(rune(id)),
						Type: domain.Predator,
					}
					_ = repo.Save(e)
				}(i)
			}
		})

		// Тест на конкурентное чтение
		t.Run("ParallelReads", func(t *testing.T) {
			for i := 0; i < numRoutines; i++ {
				go func(id int) {
					_, _ = repo.GetByID(string(rune(id % 10)))
				}(i)
			}
		})

		// Тест на конкурентное удаление
		t.Run("ParallelDeletes", func(t *testing.T) {
			for i := 0; i < numRoutines; i++ {
				go func(id int) {
					_ = repo.Delete(string(rune(id % 10)))
				}(i)
			}
		})
	})
}

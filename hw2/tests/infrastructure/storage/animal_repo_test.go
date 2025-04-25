package storage_test

import (
	"testing"
	"time"

	"github.com/Babushkin05/software-dev-course/hw2/internal/domain"
	"github.com/Babushkin05/software-dev-course/hw2/internal/infrastructure/storage"
	"github.com/stretchr/testify/assert"
)

func TestInMemoryAnimalRepo(t *testing.T) {
	// Создаем тестовые данные
	testAnimal := &domain.Animal{
		ID:        "1",
		Name:      "Simba",
		Species:   "Lion",
		BirthDate: time.Now(),
		Gender:    domain.Male,
	}

	t.Run("Create and Get", func(t *testing.T) {
		repo := storage.NewInMemoryAnimalRepo()

		// Сохраняем животное
		err := repo.Save(testAnimal)
		assert.NoError(t, err)

		// Получаем животное
		animal, err := repo.GetByID("1")
		assert.NoError(t, err)
		assert.Equal(t, testAnimal, animal)

		// Проверяем несуществующее животное
		_, err = repo.GetByID("999")
		assert.ErrorIs(t, err, domain.ErrAnimalNotFound)
	})

	t.Run("Update", func(t *testing.T) {
		repo := storage.NewInMemoryAnimalRepo()
		_ = repo.Save(testAnimal)

		// Обновляем животное
		updatedAnimal := *testAnimal
		updatedAnimal.Name = "Mufasa"
		err := repo.Save(&updatedAnimal)
		assert.NoError(t, err)

		// Проверяем обновление
		animal, err := repo.GetByID("1")
		assert.NoError(t, err)
		assert.Equal(t, "Mufasa", animal.Name)
	})

	t.Run("Delete", func(t *testing.T) {
		repo := storage.NewInMemoryAnimalRepo()
		_ = repo.Save(testAnimal)

		// Удаляем животное
		err := repo.Delete("1")
		assert.NoError(t, err)

		// Проверяем что удалилось
		_, err = repo.GetByID("1")
		assert.ErrorIs(t, err, domain.ErrAnimalNotFound)

		// Удаление несуществующего животного
		err = repo.Delete("999")
		assert.NoError(t, err) // Удаление несуществующего не должно возвращать ошибку
	})

	t.Run("List", func(t *testing.T) {
		repo := storage.NewInMemoryAnimalRepo()

		// Пустой список
		list, err := repo.List()
		assert.NoError(t, err)
		assert.Empty(t, list)

		// Добавляем животных
		animals := []*domain.Animal{
			{ID: "1", Name: "Simba"},
			{ID: "2", Name: "Zoe"},
		}

		for _, a := range animals {
			_ = repo.Save(a)
		}

		// Проверяем список
		list, err = repo.List()
		assert.NoError(t, err)
		assert.Len(t, list, 2)
		assert.ElementsMatch(t, animals, list)
	})

	t.Run("ConcurrentAccess", func(t *testing.T) {
		repo := storage.NewInMemoryAnimalRepo()
		const numRoutines = 100

		// Тест на конкурентные записи
		t.Run("ParallelWrites", func(t *testing.T) {
			for i := 0; i < numRoutines; i++ {
				go func(id int) {
					a := &domain.Animal{
						ID:   string(rune(id)),
						Name: "Animal_" + string(rune(id)),
					}
					_ = repo.Save(a)
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

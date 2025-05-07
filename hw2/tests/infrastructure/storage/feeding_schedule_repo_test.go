package storage_test

import (
	"testing"
	"time"

	"github.com/Babushkin05/software-dev-course/hw2/internal/domain"
	"github.com/Babushkin05/software-dev-course/hw2/internal/infrastructure/storage"
	"github.com/stretchr/testify/assert"
)

func TestInMemoryFeedingScheduleRepo(t *testing.T) {
	now := time.Now()
	testSchedule := &domain.FeedingSchedule{
		ID:       "s1",
		AnimalID: "a1",
		Time:     now.Add(2 * time.Hour),
		FoodType: "Meat",
		IsDone:   false,
	}

	t.Run("Add and ListAll", func(t *testing.T) {
		repo := storage.NewInMemoryFeedingScheduleRepo()

		// Добавляем расписание
		err := repo.Add(testSchedule)
		assert.NoError(t, err)

		// Проверяем список всех расписаний
		schedules, err := repo.ListAll()
		assert.NoError(t, err)
		assert.Len(t, schedules, 1)
		assert.Equal(t, testSchedule, schedules[0])
	})

	t.Run("ListByAnimal", func(t *testing.T) {
		repo := storage.NewInMemoryFeedingScheduleRepo()

		// Добавляем тестовые данные
		schedules := []*domain.FeedingSchedule{
			{ID: "s1", AnimalID: "a1", Time: now},
			{ID: "s2", AnimalID: "a2", Time: now},
			{ID: "s3", AnimalID: "a1", Time: now.Add(1 * time.Hour)},
		}

		for _, s := range schedules {
			_ = repo.Add(s)
		}

		// Получаем расписания для животного a1
		result, err := repo.ListByAnimal("a1")
		assert.NoError(t, err)
		assert.Len(t, result, 2)
		assert.Equal(t, "s1", result[0].ID)
		assert.Equal(t, "s3", result[1].ID)

		// Проверяем пустой результат
		result, err = repo.ListByAnimal("a99")
		assert.NoError(t, err)
		assert.Empty(t, result)
	})

	t.Run("MarkDone", func(t *testing.T) {
		repo := storage.NewInMemoryFeedingScheduleRepo()
		_ = repo.Add(testSchedule)

		// Помечаем как выполненное
		err := repo.MarkDone("s1")
		assert.NoError(t, err)

		// Проверяем изменение
		schedules, _ := repo.ListAll()
		assert.True(t, schedules[0].IsDone)

		// Пытаемся пометить несуществующее расписание
		err = repo.MarkDone("s99")
		assert.ErrorIs(t, err, domain.ErrScheduleNotFound)
	})

	t.Run("ConcurrentAccess", func(t *testing.T) {
		repo := storage.NewInMemoryFeedingScheduleRepo()
		const numRoutines = 100

		// Параллельные добавления
		t.Run("ParallelAdds", func(t *testing.T) {
			for i := 0; i < numRoutines; i++ {
				go func(id int) {
					s := &domain.FeedingSchedule{
						ID:       string(rune(id)),
						AnimalID: "a1",
					}
					_ = repo.Add(s)
				}(i)
			}
		})

		// Параллельные пометки выполнения
		t.Run("ParallelMarkDone", func(t *testing.T) {
			for i := 0; i < numRoutines; i++ {
				go func(id int) {
					_ = repo.MarkDone(string(rune(id % 10)))
				}(i)
			}
		})

		// Параллельные запросы
		t.Run("ParallelLists", func(t *testing.T) {
			for i := 0; i < numRoutines; i++ {
				go func() {
					_, _ = repo.ListAll()
					_, _ = repo.ListByAnimal("a1")
				}()
			}
		})
	})

	t.Run("EdgeCases", func(t *testing.T) {
		repo := storage.NewInMemoryFeedingScheduleRepo()

		t.Run("EmptyRepo", func(t *testing.T) {
			// ListAll для пустого репозитория
			schedules, err := repo.ListAll()
			assert.NoError(t, err)
			assert.Empty(t, schedules)

			// ListByAnimal для пустого репозитория
			schedules, err = repo.ListByAnimal("a1")
			assert.NoError(t, err)
			assert.Empty(t, schedules)

			// MarkDone для пустого репозитория
			err = repo.MarkDone("s1")
			assert.ErrorIs(t, err, domain.ErrScheduleNotFound)
		})
	})
}

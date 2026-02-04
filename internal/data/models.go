package data

import (
	"database/sql"
	"errors"
)

// Определяем кастомную ошибку ErrRecordNotFound. Мы будем возвращать её из нашего метода Get(),
// когда будем искать фильм, которого не существует в нашей базе данных.
var (
	ErrRecordNotFound = errors.New("record not found")
)

// Создаем структуру Models, которая оборачивает MovieModel. Мы будем добавлять сюда другие модели,
// такие как UserModel и PermissionModel, по мере развития проекта.
type Models struct {
	Movies MovieModel
}

// Для удобства использования мы также добавляем метод New(), который возвращает структуру Models,
// содержащую инициализированную MovieModel.
func NewModels(db *sql.DB) Models {
	return Models{
		Movies: MovieModel{DB: db},
	}
}

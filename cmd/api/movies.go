package main

import (
	"fmt"
	"greenlight/internal/data"
	"net/http"
	"time"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	// Объявляем анонимную структуру для хранения информации, которую мы ожидаем
	// получить в теле HTTP-запроса (обратите внимание, что имена полей и типы
	// в структуре являются подмножеством структуры Movie, которую мы создали ранее).
	// Эта структура будет нашей *целевой точкой декодирования*.
	var input struct {
		Title   string       `json:"title"`
		Year    int32        `json:"year"`
		Runtime data.Runtime `json:"runtime"`
		Genres  []string     `json:"genres"`
	}

	// Инициализируем новый экземпляр json.Decoder, который читает из тела запроса,
	// и затем используем метод Decode() для декодирования содержимого тела в структуру input.
	// Важно: обратите внимание, что при вызове Decode() мы передаем *указатель*
	// на структуру input. Если во время декодирования произошла ошибка, мы используем
	// нашу общую вспомогательную функцию errorResponse(), чтобы отправить клиенту
	// ответ 400 Bad Request с сообщением об ошибке.
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	// Выводим содержимое структуры input в HTTP-ответ.
	fmt.Fprintf(w, "%+v\n", input)
}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	movie := data.Movie{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "Casablanca",
		Runtime:   102,
		Genres:    []string{"drama", "romance", "war"},
		Version:   1,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"movie": movie}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

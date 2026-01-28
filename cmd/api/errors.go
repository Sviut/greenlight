package main

import (
	"fmt"
	"net/http"
)

func (app *application) logError(r *http.Request, err error) {
	var (
		method = r.Method
		uri    = r.URL.RequestURI()
	)

	app.logger.Error(err.Error(), "method", method, "uri", uri)
}

// Метод errorResponse() — это универсальный хелпер для отправки JSON-сообщений об ошибках
// клиенту с заданным статус-кодом. Мы используем тип any для параметра message,
// что дает нам больше гибкости.
func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, status int, message any) {
	env := envelope{"error": message}

	// Записываем ответ с помощью хелпера writeJSON(). Если это вернет ошибку,
	// логируем её и отправляем пустой ответ с кодом 500 Internal Server Error.
	err := app.writeJSON(w, status, env, nil)
	if err != nil {
		app.logError(r, err)
		w.WriteHeader(500)
	}
}

// Метод serverErrorResponse() будет использоваться при возникновении непредвиденных проблем
// во время выполнения. Он логирует ошибку и отправляет код 500 и JSON-ответ.
func (app *application) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logError(r, err)

	message := "the server encountered a problem and could not process your request"
	app.errorResponse(w, r, http.StatusInternalServerError, message)
}

// Метод notFoundResponse() используется для отправки кода 404 Not Found и JSON-ответа.
func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request) {
	message := "the requested resource could not be found"
	app.errorResponse(w, r, http.StatusNotFound, message)
}

// Метод methodNotAllowedResponse() используется для отправки кода 405 Method Not Allowed.
func (app *application) methodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("the %s method is not supported for this resource", r.Method)
	app.errorResponse(w, r, http.StatusMethodNotAllowed, message)
}

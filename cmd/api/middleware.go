package main

import (
	"fmt"
	"net/http"
)

func (app *application) recoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Создаем отложенную функцию (которая всегда будет запущена в случае паники,
		// так как Go разматывает стек).
		defer func() {
			// Используем встроенную функцию recover, чтобы проверить, была паника или нет.
			if err := recover(); err != nil {
				// Если была паника, устанавливаем заголовок "Connection: close" в ответе.
				// Это действует как триггер, заставляющий HTTP-сервер Go автоматически
				// закрыть текущее соединение после отправки ответа.
				w.Header().Set("Connection", "close")

				// Значение, возвращаемое recover(), имеет тип any, поэтому мы используем
				// fmt.Errorf(), чтобы нормализовать его в ошибку и вызвать нашу
				// вспомогательную функцию serverErrorResponse(). В свою очередь, она
				// запишет ошибку в лог с использованием нашего кастомного типа Logger
				// на уровне ERROR и отправит клиенту ответ 500 Internal Server Error.
				app.serverErrorResponse(w, r, fmt.Errorf("%s", err))
			}
		}()

		next.ServeHTTP(w, r)
	})
}

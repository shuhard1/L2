package event

import (
	"github.com/shuhard1/L2/develop/dev11/internal/render"
)

// Event - структура, описывающая мироприятия
type Event struct {
	EventID     int         `json:"event_id"`
	UserID      int         `json:"user_id"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Date        render.Date `json:"date"`
}

// ErrorEvent - структура, с помощью которой выводим оишбку на страницу
type ErrorEvent struct {
	Error string `json:"error"`
}

// Result - структура, с помощью которой выводим успешный результат запроса на страницу
type Result struct {
	Message string  `json:"message"`
	Events  []Event `json:"result"`
}

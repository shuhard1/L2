package server

// В данном пакете хранится логика работы сервера,
// вспомогательные функции для парсинга и валидации параметров методов, а также логгер запросов

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/shuhard1/L2/develop/dev11/internal/event"
	"github.com/shuhard1/L2/develop/dev11/pkg/store"
)

// Структура для работы сервера с хранилищем событий
type handler struct {
	store store.Store
}

// NewHandler - функция-конструктор для создания экземпляра handler
func NewHandler() handler {
	return handler{
		store: store.NewStore(),
	}
}

// Routing - метод для регистрации всех хендлеров в мультиплексоре
func (h *handler) Routing(mux *http.ServeMux) {
	mux.HandleFunc("/create_event", h.createEvent)
	mux.HandleFunc("/update_event", h.updateEvent)
	mux.HandleFunc("/delete_event", h.deleteEvent)
	mux.HandleFunc("/events_for_day", h.eventsForDay)
	mux.HandleFunc("/events_for_week", h.eventsForWeek)
	mux.HandleFunc("/events_for_month", h.eventsForMonth)

}

// Метод для парсинга
func (h *handler) unmarshalJSON(r *http.Request) (*event.Event, error) {
	var event event.Event
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	err = json.Unmarshal(data, &event)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

// Метод для возвращения сервером успешного ответа
func (h *handler) successfulResponse(w http.ResponseWriter, message string, events []event.Event) {
	w.Header().Set("Content-Type", "application/json")
	result := event.Result{
		Message: message,
		Events:  events,
	}
	data, err := json.MarshalIndent(result, "", "\t")
	if err != nil {
		h.errorResponse(w, err, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

// Метод для возвращения сервером JSONа с описанием ошибки
func (h *handler) errorResponse(w http.ResponseWriter, err error, status int) {
	w.Header().Set("Content-Type", "application/json")
	errMessage := err.Error()
	data, er := json.MarshalIndent(event.ErrorEvent{Error: errMessage}, "", "\t")
	if er != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error marshaling another error:" + errMessage))
		return
	}
	w.WriteHeader(status)
	w.Write(data)
}

// Метод для создания события
func (h *handler) createEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		e, err := h.unmarshalJSON(r)
		if err != nil {
			h.errorResponse(w, err, http.StatusBadRequest)
			return
		}
		h.store.CreateEvent(e)
		h.successfulResponse(w, "Event successfully created!", []event.Event{*e})
		return
	}
	err := fmt.Errorf("invalid method passed %s, expected POST", r.Method)
	h.errorResponse(w, err, http.StatusBadRequest)
}

// Метод для изменения события
func (h *handler) updateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		e, err := h.unmarshalJSON(r)
		if err != nil {
			h.errorResponse(w, err, http.StatusBadRequest)
			return
		}
		err = h.store.UpdateEvent(e)
		if err != nil {
			h.errorResponse(w, err, http.StatusBadRequest)
			return
		}
		h.successfulResponse(w, "Event changed successfully!", []event.Event{*e})
		return
	}
	err := fmt.Errorf("invalid method passed %s, expected POST", r.Method)
	h.errorResponse(w, err, http.StatusBadRequest)
}

// Метод для удаления события
func (h *handler) deleteEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			h.errorResponse(w, err, http.StatusBadRequest)
			return
		}
		defer r.Body.Close()
		var e *event.Event
		err = json.Unmarshal(data, &e)
		if err != nil {
			h.errorResponse(w, err, http.StatusBadRequest)
			return
		}
		e, err = h.store.DeleteEvent(e.EventID)
		if err != nil {
			h.errorResponse(w, err, http.StatusBadRequest)
			return
		}
		h.successfulResponse(w, "Event deleted successfully!", []event.Event{*e})
		return
	}
	err := fmt.Errorf("invalid method passed %s, expected POST", r.Method)
	h.errorResponse(w, err, http.StatusBadRequest)
}

// Метод для показа событий пользователя в определённый день
func (h *handler) eventsForDay(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// Получаем id пользователя и дату из запроса
		id, err := strconv.Atoi(r.URL.Query().Get("user_id"))
		if err != nil {
			h.errorResponse(w, err, http.StatusBadRequest)
			return
		}
		dateStr := r.URL.Query().Get("date")
		dateTime, err := time.Parse("2006-01-02", dateStr)
		if err != nil {
			h.errorResponse(w, err, http.StatusBadRequest)
			return
		}
		events := h.store.EventsForDay(id, dateTime)
		if len(events) == 0 {
			h.successfulResponse(w, "No events scheduled for this day.", events)
			return
		}
		h.successfulResponse(w, "List of events successfully received!", events)
		return
	}
	err := fmt.Errorf("invalid method passed %s, expected GET", r.Method)
	h.errorResponse(w, err, http.StatusBadRequest)
}

// Метод для показа событий пользователя за неделю
func (h *handler) eventsForWeek(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// Получаем id пользователя и дату из запроса
		id, err := strconv.Atoi(r.URL.Query().Get("user_id"))
		if err != nil {
			h.errorResponse(w, err, http.StatusBadRequest)
			return
		}
		dateStr := r.URL.Query().Get("date")
		dateTime, err := time.Parse("2006-01-02", dateStr)
		if err != nil {
			h.errorResponse(w, err, http.StatusBadRequest)
			return
		}
		events := h.store.EventsForWeek(id, dateTime)
		if len(events) == 0 {
			h.successfulResponse(w, "No events scheduled for this week.", events)
			return
		}
		h.successfulResponse(w, "List of events successfully received!", events)
		return
	}
	err := fmt.Errorf("invalid method passed %s, expected GET", r.Method)
	h.errorResponse(w, err, http.StatusBadRequest)
}

// Метод для показа событий пользователя за месяц
func (h *handler) eventsForMonth(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// Получаем id пользователя и дату из запроса
		id, err := strconv.Atoi(r.URL.Query().Get("user_id"))
		if err != nil {
			h.errorResponse(w, err, http.StatusBadRequest)
			return
		}
		dateStr := r.URL.Query().Get("date")
		dateTime, err := time.Parse("2006-01-02", dateStr)
		if err != nil {
			dateTime, err = time.Parse("2006-01", dateStr)
			if err != nil {
				h.errorResponse(w, err, http.StatusBadRequest)
				return
			}
		}
		events := h.store.EventsForMonth(id, dateTime)
		if len(events) == 0 {
			h.successfulResponse(w, "No events scheduled for this month.", events)
			return
		}
		h.successfulResponse(w, "List of events successfully received!", events)
		return
	}
	err := fmt.Errorf("invalid method passed %s, expected GET", r.Method)
	h.errorResponse(w, err, http.StatusBadRequest)
}

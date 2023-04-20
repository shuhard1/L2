package store

import (
	"fmt"
	"sync"
	"time"

	"github.com/shuhard1/L2/develop/dev11/internal/event"
)

// Store - хранилище событий
type Store struct {
	mu            *sync.RWMutex
	storageEvents map[int]*event.Event
	eventNumber   int
}

// NewStore - Фунция-конструктор для store
func NewStore() Store {
	return Store{
		storageEvents: make(map[int]*event.Event),
		eventNumber:   1,
		mu:            new(sync.RWMutex),
	}
}

// CreateEvent - Метод для создания события, ему присваивается свой уникальный ID
// и далее оно кладётся в хранилище
func (s *Store) CreateEvent(e *event.Event) {
	s.mu.Lock()
	e.EventID = s.eventNumber
	s.storageEvents[e.EventID] = e
	s.eventNumber++
	s.mu.Unlock()
}

// UpdateEvent - Метод для обновления события
func (s *Store) UpdateEvent(e *event.Event) error {
	s.mu.Lock()
	if _, ok := s.storageEvents[e.EventID]; !ok {
		return fmt.Errorf("еvent with event id %d was not found", e.EventID)
	}
	s.storageEvents[e.EventID] = e
	s.mu.Unlock()
	return nil
}

// DeleteEvent - Метод для удаления события
func (s *Store) DeleteEvent(id int) (event *event.Event, err error) {
	s.mu.Lock()
	if _, ok := s.storageEvents[id]; !ok {
		return nil, fmt.Errorf("еvent with event id %d was not found", id)
	}
	event = s.storageEvents[id]
	delete(s.storageEvents, id)
	s.mu.Unlock()
	return event, nil
}

// EventsForDay - Метод для получения событий конкретного пользователя в указанный день
func (s *Store) EventsForDay(userID int, date time.Time) []event.Event {
	events := make([]event.Event, 0)
	s.mu.RLock()
	// Пробегаемся по хранилищу и при совпадении даты и ID пользователя добавляем событие в массив
	for _, e := range s.storageEvents {
		if e.Date.Day() == date.Day() && e.Date.Month() == date.Month() && e.Date.Year() == date.Year() && e.UserID == userID {
			events = append(events, *e)
		}
	}
	s.mu.RUnlock()
	return events
}

// EventsForWeek - Метод для получения событий конкретного пользователя в указанную неделю аналогично EventsForDay
func (s *Store) EventsForWeek(userID int, date time.Time) []event.Event {
	events := make([]event.Event, 0)
	s.mu.RLock()
	yearOne, weekOne := date.ISOWeek()
	// Пробегаемся по хранилищу и при совпадении даты и ID пользователя добавляем событие в массив
	for _, e := range s.storageEvents {
		// Получаем номер года и недели события, затем сравниваем их с теми, которые в запросе
		yearTwo, weekTwo := e.Date.ISOWeek()
		if yearOne == yearTwo && weekOne == weekTwo && e.UserID == userID {
			events = append(events, *e)
		}
	}
	s.mu.RUnlock()
	return events
}

// EventsForMonth - Метод для получения событий конкретного пользователя в указанный месяц аналогично EventsForDay
func (s *Store) EventsForMonth(userID int, date time.Time) []event.Event {
	events := make([]event.Event, 0)
	s.mu.RLock()
	// Пробегаемся по хранилищу и при совпадении даты и ID пользователя добавляем событие в массив
	for _, e := range s.storageEvents {
		if date.Year() == e.Date.Year() && date.Month() == e.Date.Month() && e.UserID == userID {
			events = append(events, *e)
		}
	}
	s.mu.RUnlock()
	return events
}

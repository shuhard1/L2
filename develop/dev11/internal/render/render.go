package render

import (
	"errors"
	"strings"
	"time"
)

// Date - отдельная структура для переопределения метода UnmarshalJSON
// из пакета json, чтобы правильно декодиравать даты, подаваемой клиентом в запросах
type Date struct {
	time.Time
}

// UnmarshalJSON - переопределённый метода из пакета json для парсинга переданной клиентом даты
func (t *Date) UnmarshalJSON(data []byte) error {
	if string(data) == "" || string(data) == `""` {
		*t = Date{time.Now()}
		return nil
	}

	timeStr := strings.ReplaceAll(string(data), `"`, "")
	parsedTime, err := time.Parse("2006-01-02T15:04", timeStr)
	if err != nil {
		parsedTime, err = time.Parse("2006-01-02T15:04:00Z", timeStr)
		if err != nil {
			parsedTime, err = time.Parse("2006-01-02", timeStr)
			if err != nil {
				return errors.New("wrong date format, example 2006-01-02T15:04")
			}
		}
	}
	*t = Date{parsedTime}
	return nil
}

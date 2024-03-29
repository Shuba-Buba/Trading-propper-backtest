package storage

import (
	"encoding/json"
	"log"
	"path"
	"time"

	"github.com/Shuba-Buba/Trading-propper-backtest/common/types"
)

type EventsKeeper struct {
	dao Dao
}

func getTableName(t time.Time) string {
	return t.Format("2006-01-02T15:04")
}

func MakeEventsKeeper(symbol string) *EventsKeeper {
	dao := MakePlainFileDao(path.Join("data", symbol))
	return &EventsKeeper{dao}
}

func (t *EventsKeeper) Save(event types.Event) {
	table := getTableName(event.Timestamp)
	b, err := json.Marshal(event)
	if err != nil {
		log.Fatal(err)
	}
	t.dao.Append(table, string(b))
}

// Возвращает события в указанном промежутке с точностью до минуты
func (t *EventsKeeper) GetEvents(from time.Time, to time.Time) chan types.Event {
	tables := t.dao.GetAllTables()
	from_table := getTableName(from)
	to_table := getTableName(to)

	ch := make(chan types.Event)
	go func() {
		defer close(ch)
		for _, table := range tables {

			if from_table <= table && table <= to_table {
				for row := range t.dao.GetRows(table) {
					event := types.Event{}
					err := json.Unmarshal([]byte(row), &event)
					if err != nil {
						log.Fatal(err)
					}
					ch <- event
				}
			}
		}
	}()

	return ch
}

package daysteps

import (
	"errors"
	"fmt"
	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
	"log"
	"strconv"
	"strings"
	"time"
)

type DaySteps struct {
	// TODO: добавить поля
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

// Parse - парсит строку с данными формата "678,0h50m" и записывает данные в соответствующие поля структуры DaySteps.
func (ds *DaySteps) Parse(datastring string) (err error) {
	// TODO: реализовать функцию

	sliceData := strings.Split(datastring, ",")
	if len(sliceData) != 2 {
		return errors.New("incorrect string composition: <> 2")
	}
	numberOfSteps, err := strconv.Atoi(sliceData[0])
	if err != nil {
		return fmt.Errorf("steps conversion error: %w", err)
	}
	if numberOfSteps <= 0 {
		err := errors.New("number of steps <= 0")
		return fmt.Errorf("steps number <= 0: %w", err)
	}

	duration, err := time.ParseDuration(sliceData[1])
	if err != nil {
		return err
	} else if duration <= 0 {
		err := errors.New("set time <= 0")
		return fmt.Errorf("duration: %w", err)
	}

	ds.Steps = numberOfSteps
	ds.Duration = duration

	return nil

}

// ActionInfo - метод формирует и возвращает строку с данными DaySteps о прогулке в формате:
// Количество шагов: 792.
// Дистанция составила 0.51 км.
// Вы сожгли 221.33 ккал.
func (ds DaySteps) ActionInfo() (string, error) {
	// TODO: реализовать функцию

	distanceKm := spentenergy.Distance(ds.Steps, ds.Height)

	numberCalStep, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		log.Printf("incorrect format: %v", err)
		return "", err
	}
	message := fmt.Sprintf("Количество шагов: %d.\n"+
		"Дистанция составила %.2f км.\n"+
		"Вы сожгли %.2f ккал.\n", ds.Steps, distanceKm, numberCalStep)
	return message, err
}

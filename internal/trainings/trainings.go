package trainings

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

// Training - структура содержит все необходимые данные о тренировке:
// количество шагов, тип тренировки, длительность тренировки,
// а также данные из структуры personaldata.Personal, то есть имя, вес и рост пользователя.
type Training struct {
	// TODO: добавить поля
	personaldata.Personal
	Steps        int
	TrainingType string
	Duration     time.Duration
}

// Training - метод принимает строку с данными формата "3456,Ходьба,3h00m".
// Запсиывает в структуру Training шаги, тип тренировки, время.
func (t *Training) Parse(datastring string) (err error) {
	// TODO: реализовать функцию
	sliceData := strings.Split(datastring, ",")
	if len(sliceData) != 3 {
		return errors.New("incorrect string composition: <> 3")
	}
	t.TrainingType = sliceData[1] //тип тренировки

	numberOfSteps, err := strconv.Atoi(sliceData[0])
	if err != nil {
		return fmt.Errorf("steps conversion error: %w", err)
	}
	if numberOfSteps <= 0 {
		err := errors.New("number of steps <= 0")
		return fmt.Errorf("steps number <=0: %w", err)
	}
	t.Steps = numberOfSteps // кол-во шагов

	duration, err := time.ParseDuration(sliceData[2])
	if err != nil {
		return fmt.Errorf("time conversion error: %w", err)
	} else if duration <= 0 {
		err := errors.New("set time <= 0")
		return fmt.Errorf("duration: %w", err)
	}
	t.Duration = duration // время тренировки
	return nil
}

// ActionInfo - метод формирует и возвращает строку с данными о тренировке
// (Тип тренировки: Бег
// Длительность: 0.75 ч.
// Дистанция: 10.00 км.
// Скорость: 13.34 км/ч
// Сожгли калорий: 18621.75),
// исходя из того, какой тип тренировки был передан.
func (t Training) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	var (
		distanceActive   float64
		speedActive      float64
		numberOfCalories float64
		err              error
	)

	distanceActive = spentenergy.Distance(t.Steps, t.Height)           // дистанция
	speedActive = spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration) // средняя скорость

	switch t.TrainingType {
	case "Ходьба":
		numberOfCalories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			log.Println(err)
		}
	case "Бег":
		numberOfCalories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			log.Println(err)
		}
	default:
		return "", errors.New("unknown type of training")
	}

	message := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\n"+
		"Дистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		t.TrainingType, t.Duration.Hours(), distanceActive, speedActive, numberOfCalories)
	return message, nil
}

package spentenergy

import (
	"errors"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

// Функция рассчитывает количество калорий по входным данным:
//
//	кол-во шагов, вес, рост, время активности
//
// Возвращает: кол-во калорий (float64) при ходьбе, ошибку
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if duration <= 0 {
		return 0, errors.New("WalkingSpentCalories: incorrect time")
	} else if steps <= 0 || height <= 0 || weight <= 0 {
		return 0, errors.New("walkingSpentCalories: steps, weight, height <= 0")
	}

	averageSpeed := MeanSpeed(steps, height, duration)
	numberCal := weight * averageSpeed * duration.Minutes() / float64(minInH)
	return numberCal * walkingCaloriesCoefficient, nil
}

// Функция рассчитывает количество калорий по входным данным:
//
//	кол-во шагов, вес, рост, время активности
//
// Возвращает: кол-во калорий (float64) при беге, ошибку
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if duration <= 0 {
		return 0, errors.New("RunningSpentCalories: incorrect time")
	} else if steps <= 0 || height <= 0 || weight <= 0 {
		return 0, errors.New("runningSpentCalories: steps, weight, height <= 0")
	}

	averageSpeed := MeanSpeed(steps, height, duration)
	return weight * averageSpeed * duration.Minutes() / float64(minInH), nil
}

// Функция принимает количество шагов steps,
//
//	рост пользователя height и продолжительность активности duration
//	и возвращает среднюю скорость.
func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration.Seconds() <= 0 {
		return 0
	}
	return Distance(steps, height) / duration.Hours()
}

// Функция принимает количество шагов и рост пользователя в метрах,
//
//	а возвращает дистанцию в километрах(float64).
func Distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	return height * stepLengthCoefficient * float64(steps) / float64(mInKm)
}

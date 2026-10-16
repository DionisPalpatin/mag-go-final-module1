package spentenergy

import (
	"fmt"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps < 0 || duration < 0 || weight <= 0 || height <= 0 {
		return 0, fmt.Errorf(
			"ошибка во входных данных, вес и рост не могут быть меньше или равны 0, "+
				"шаги и длительность не могут быть меньше 0: "+
				"steps = %d, длительность = %s, weight = %f, height = %f",
			steps, duration, weight, height,
		)
	}

	meanSpeed := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	spentCalories := walkingCaloriesCoefficient * weight * meanSpeed * durationInMinutes / minInH

	return spentCalories, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps < 0 || duration < 0 || weight <= 0 || height <= 0 {
		return 0, fmt.Errorf(
			"ошибка во входных данных, вес и рост не могут быть меньше или равны 0, "+
				"шаги и длительность не могут быть меньше 0: "+
				"steps = %d, длительность = %s, weight = %f, height = %f",
			steps, duration, weight, height,
		)
	}

	meanSpeed := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	spentCalories := weight * meanSpeed * durationInMinutes / minInH

	return spentCalories, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration == 0 {
		return 0
	}

	dist := Distance(steps, height)

	return dist / duration.Hours()
}

func Distance(steps int, height float64) float64 {
	return float64(steps) * height * stepLengthCoefficient / mInKm
}

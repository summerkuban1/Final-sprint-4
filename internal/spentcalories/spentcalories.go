package spentcalories

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	parts := strings.Split(data, ",")
	if len(parts) != 3 {
		return 0, "", 0, fmt.Errorf("некорректный формат данных")
	}

	stepsString := parts[0]
	trainingType := parts[1]
	durationString := parts[2]

	if stepsString != strings.TrimSpace(stepsString) {
		return 0, "", 0, fmt.Errorf("некорректный формат данных")
	}

	if durationString != strings.TrimSpace(durationString) {
		return 0, "", 0, fmt.Errorf("некорректный формат данных")
	}

	steps, err := strconv.Atoi(stepsString)
	if err != nil {
		return 0, "", 0, err
	}

	if steps <= 0 {
		return 0, "", 0, fmt.Errorf("количество шагов должно быть больше 0")
	}

	duration, err := time.ParseDuration(durationString)
	if err != nil {
		return 0, "", 0, err
	}

	if duration <= 0 {
		return 0, "", 0, fmt.Errorf("продолжительность тренировки должна быть больше 0")
	}

	return steps, trainingType, duration, nil
}

func distance(steps int, height float64) float64 {
	if steps <= 0 || height <= 0 {
		return 0
	}

	stepLength := height * stepLengthCoefficient
	distanceMeters := float64(steps) * stepLength

	return distanceMeters / mInKm
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	if steps <= 0 || height <= 0 || duration <= 0 {
		return 0
	}

	distanceKm := distance(steps, height)

	return distanceKm / duration.Hours()
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	steps, trainingType, duration, err := parseTraining(data)
	if err != nil {
		log.Println(err)
		return "", err
	}

	var calories float64

	switch trainingType {
	case "Бег":
		calories, err = RunningSpentCalories(steps, weight, height, duration)
	case "Ходьба":
		calories, err = WalkingSpentCalories(steps, weight, height, duration)
	default:
		return "", fmt.Errorf("неизвестный тип тренировки")
	}

	if err != nil {
		log.Println(err)
		return "", err
	}

	distanceKm := distance(steps, height)
	speed := meanSpeed(steps, height, duration)

	return fmt.Sprintf(
		"Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		trainingType,
		duration.Hours(),
		distanceKm,
		speed,
		calories,
	), nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, fmt.Errorf("количество шагов должно быть больше 0")
	}

	if weight <= 0 {
		return 0, fmt.Errorf("вес должен быть больше 0")
	}

	if height <= 0 {
		return 0, fmt.Errorf("рост должен быть больше 0")
	}

	if duration <= 0 {
		return 0, fmt.Errorf("продолжительность тренировки должна быть больше 0")
	}

	speed := meanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()

	calories := weight * speed * durationInMinutes / minInH

	return calories, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, fmt.Errorf("количество шагов должно быть больше 0")
	}

	if weight <= 0 {
		return 0, fmt.Errorf("вес должен быть больше 0")
	}

	if height <= 0 {
		return 0, fmt.Errorf("рост должен быть больше 0")
	}

	if duration <= 0 {
		return 0, fmt.Errorf("продолжительность тренировки должна быть больше 0")
	}

	speed := meanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()

	calories := weight * speed * durationInMinutes / minInH
	calories *= walkingCaloriesCoefficient

	return calories, nil
}
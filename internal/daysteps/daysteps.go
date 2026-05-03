package daysteps

import (
	"errors"
	"strconv"
	"strings"
	"time"
	"fmt"
	"log"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	parts := strings.Split(data, ",")
	if len(parts) != 2 {
		return 0, 0, errors.New("неправильный формат данных")
	}
	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, err
	}
	duration, err := time.ParseDuration(parts[1])
	if err != nil {
		return 0, 0, err
	}
	if steps <= 0 {
		return 0, 0, errors.New("количество шагов должно быть больше нуля")
	}
	return steps, duration, nil
}


func DayActionInfo(data string, weight, height float64) string {
	steps, duration, err := parsePackage(data)
	if err != nil {
		log.Println("Ошибка при парсинге данных:", err)
		return ""
	}
	if steps <= 0 {
		log.Println("Количество шагов должно быть больше нуля.")
		return ""
	}
	if weight <= 0 || height <= 0 {
		log.Println("Вес и рост должны быть больше нуля.")
		return "Ошибка: вес и рост должны быть больше нуля."
	}
	distance := float64(steps) * stepLength / mInKm
	calories, err := spentcalories.WalkingSpentCalories(steps, weight, height, duration)
	if err != nil {
		log.Println(err)
		return ""
	}
	result := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.", steps, distance, calories)
	return result
}

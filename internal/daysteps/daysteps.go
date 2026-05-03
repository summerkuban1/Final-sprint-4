package daysteps

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
	// Коэффициенты для расчета калорий
	walkingCaloriesCoefficient = 0.05 // Ходьба (калории за минуту)
	runningCaloriesCoefficient = 0.1  // Бег (калории за минуту)
)

// Функция для парсинга строки с данными (шаги и продолжительность)
func parsePackage(data string) (int, time.Duration, error) {
	parts := strings.Split(data, ",")
	if len(parts) != 2 {
		return 0, 0, errors.New("неправильный формат данных, ожидается <шаги>,<время>")
	}
	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, fmt.Errorf("не удалось преобразовать количество шагов: %v", err)
	}
	duration, err := time.ParseDuration(parts[1])
	if err != nil {
		return 0, 0, fmt.Errorf("не удалось распарсить продолжительность времени: %v", err)
	}
	if steps <= 0 {
		return 0, 0, errors.New("количество шагов должно быть больше нуля")
	}
	return steps, duration, nil
}

// Функция для расчета калорий (ходьба или бег)
func calculateCalories(steps int, weight, height float64, duration time.Duration, activityType string) (float64, error) {
	if weight <= 0 || height <= 0 {
		return 0, errors.New("вес и рост должны быть больше нуля")
	}

	// Вычисляем количество минут в продолжительности
	durationInMinutes := duration.Minutes()

	// Расчет калорий в зависимости от типа активности
	var calories float64
	switch activityType {
	case "walking":
		calories = walkingCaloriesCoefficient * weight * durationInMinutes
	case "running":
		calories = runningCaloriesCoefficient * weight * durationInMinutes
	default:
		return 0, errors.New("неизвестный тип активности")
	}

	return calories, nil
}

// Функция для получения информации о действиях пользователя
func DayActionInfo(data string, weight, height float64, activityType string) string {
	// Парсим входные данные
	steps, duration, err := parsePackage(data)
	if err != nil {
		log.Println("Ошибка при парсинге данных:", err)
		return fmt.Sprintf("Ошибка: %v", err) // Возвращаем ошибку с пояснением
	}

	// Проверка на корректность веса и роста
	if weight <= 0 || height <= 0 {
		log.Println("Вес и рост должны быть больше нуля.")
		return "Ошибка: вес и рост должны быть больше нуля."
	}

	// Вычисление пройденной дистанции
	distance := float64(steps) * stepLength / mInKm

	// Рассчитываем калории, с учетом ошибок
	calories, err := calculateCalories(steps, weight, height, duration, activityType)
	if err != nil {
		log.Println("Ошибка при расчете калорий:", err)
		return fmt.Sprintf("Ошибка при расчете калорий: %v", err)
	}

	// Формируем и возвращаем строку с результатами
	result := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.", steps, distance, calories)
	return result
}
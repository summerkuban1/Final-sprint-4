package daysteps

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

const (
	// Длина одного шага в метрах.
	stepLength = 0.65

	// Количество метров в одном километре.
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	parts := strings.Split(data, ",")
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("некорректный формат данных")
	}

	stepsString := parts[0]
	durationString := parts[1]

	if stepsString != strings.TrimSpace(stepsString) {
		return 0, 0, fmt.Errorf("некорректный формат данных")
	}

	if durationString != strings.TrimSpace(durationString) {
		return 0, 0, fmt.Errorf("некорректный формат данных")
	}

	steps, err := strconv.Atoi(stepsString)
	if err != nil {
		return 0, 0, err
	}

	if steps <= 0 {
		return 0, 0, fmt.Errorf("количество шагов должно быть больше 0")
	}

	duration, err := time.ParseDuration(durationString)
	if err != nil {
		return 0, 0, err
	}

	if duration <= 0 {
		return 0, 0, fmt.Errorf("продолжительность тренировки должна быть больше 0")
	}

	return steps, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	steps, duration, err := parsePackage(data)
	if err != nil {
		log.Println(err)
		return ""
	}

	distanceMeters := float64(steps) * stepLength
	distanceKm := distanceMeters / mInKm

	calories, err := spentcalories.WalkingSpentCalories(steps, weight, height, duration)
	if err != nil {
		log.Println(err)
		return ""
	}

	return fmt.Sprintf(
		"Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		steps,
		distanceKm,
		calories,
	)
}
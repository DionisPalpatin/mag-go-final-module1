package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration

	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	values := strings.Split(datastring, ",")
	if len(values) != 3 {
		err = fmt.Errorf("в строке %s не хватает данных либо формат строки неправильный", datastring)
		return
	}

	t.TrainingType = values[1]
	if t.Steps, err = strconv.Atoi(values[0]); err != nil {
		err = fmt.Errorf("ошибка при получении количества шагов: %w", err)
		return
	} else if t.Steps <= 0 {
		err = errors.New("некорректное количество шагов")
		return
	}

	if t.Duration, err = time.ParseDuration(values[2]); err != nil {
		err = fmt.Errorf("ошибка при получении длительности тренировки: %w", err)
		return
	} else if t.Duration <= 0 {
		err = errors.New("некорректная продолжительность тренировки")
		return
	}

	return nil
}

func (t Training) ActionInfo() (string, error) {
	var calories float64 = 0
	var err error = nil
	switch t.TrainingType {
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	default:
		err = errors.New("неизвестный тип тренировки")
	}

	if err != nil {
		return "", err
	}

	dist := spentenergy.Distance(t.Steps, t.Height)
	speed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	res_str := fmt.Sprintf(
		"Тип тренировки: %s\n"+
			"Длительность: %.2f ч.\n"+
			"Дистанция: %.2f км.\n"+
			"Скорость: %.2f км/ч\n"+
			"Сожгли калорий: %.2f\n",
		t.TrainingType, t.Duration.Hours(), dist, speed, calories,
	)

	return res_str, nil
}

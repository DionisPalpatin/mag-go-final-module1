package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	values := strings.Split(datastring, ",")
	if len(values) != 2 {
		err = fmt.Errorf("в строке %s не хватает данных либо формат строки неправильный", datastring)
		return
	}

	if ds.Steps, err = strconv.Atoi(values[0]); err != nil {
		err = fmt.Errorf("ошибка при получении количества шагов: %w", err)
		return
	} else if ds.Steps <= 0 {
		err = errors.New("некорректное количество шагов")
		return
	}

	if ds.Duration, err = time.ParseDuration(values[1]); err != nil {
		err = fmt.Errorf("ошибка при получении длительности тренировки: %w", err)
		return
	} else if ds.Duration <= 0 {
		err = errors.New("некорректная продолжительность тренировки")
		return
	}

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	if ds.Steps <= 0 {
		return "", errors.New("некорректное количество шагов")
	} else if ds.Duration <= 0 {
		return "", errors.New("некорректная продолжительность тренировки")
	}

	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}

	dist := spentenergy.Distance(ds.Steps, ds.Height)

	res_str := fmt.Sprintf(
		"Количество шагов: %d.\n"+
			"Дистанция составила %.2f км.\n"+
			"Вы сожгли %.2f ккал.\n",
		ds.Steps, dist, calories,
	)

	return res_str, nil
}

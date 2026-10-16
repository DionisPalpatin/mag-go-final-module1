package trainings

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration

	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	values := strings.Split(datastring, ",")
	if len(values) < 3 {
		err = fmt.Errorf("в строке %s не хватает данных либо формат строки неправильный", datastring)
		return
	}

	t.TrainingType = values[1]
	if t.Steps, err = strconv.Atoi(values[0]); err != nil {
		err = fmt.Errorf("ошибка при получении количества шагов: %w", err)
		return
	}
	if t.Duration, err = time.ParseDuration(values[2]); err != nil {
		err = fmt.Errorf("ошибка при получении длительности тренировки: %w", err)
		return
	}

	return nil
}

func (t Training) ActionInfo() (string, error) {
	fmt.Printf(
		"Тип тренировки: Бег" +
			"Длительность: 0.75 ч." +
			"Дистанция: 10.00 км." +
			"Скорость: 13.34 км/ч" +
			"Сожгли калорий: 18621.75",
	)
}

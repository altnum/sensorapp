package executiontimer

import (
	"errors"
	"strconv"

	"github.com/spf13/cobra"
)

var getFlagValueString func(string) (string, error)
var parseInt = strconv.ParseInt

type ITimer interface {
	SetTotalDuration(*cobra.Command) error
	SetDeltaDuration(*cobra.Command) error
	GetTotalDuration() int64
	GetDeltaDuration() int64
	SetUpTimer(*cobra.Command) error
}

type Timer struct {
	total_duration int64
	delta_duration int64
}

func (t *Timer) SetUpTimer(cmd *cobra.Command) error {
	t.total_duration = 1
	t.delta_duration = 0
	getFlagValueString = cmd.Flags().GetString

	err := t.SetTotalDuration(cmd)
	if err != nil {
		return err
	}

	err = t.SetDeltaDuration(cmd)
	if err != nil {
		return err
	}

	return nil
}

func (t *Timer) SetTotalDuration(cmd *cobra.Command) error {
	var durationStr string
	var err error

	if durationStr, err = getFlagValueString("total_duration"); err == nil {
		t.total_duration, err = parseInt(durationStr, 0, 64)
		if err != nil || t.total_duration < 1 {
			return errors.New("invalid input for 'total_duration' command")
		}
	}

	return err
}

func (t *Timer) SetDeltaDuration(cmd *cobra.Command) error {
	var deltaDurationStr string
	var err error

	if deltaDurationStr, err = getFlagValueString("delta_duration"); err == nil {
		t.delta_duration, err = parseInt(deltaDurationStr, 0, 64)
		if err != nil || t.delta_duration < 0 {
			return errors.New("invalid input for 'delta_duration' command")
		}

		if t.total_duration <= t.delta_duration {
			return errors.New("'total_duration' value is smaller or equal to 'delta_duration' value")
		}
	}

	return err
}

func (t *Timer) GetTotalDuration() int64 {
	return t.total_duration
}

func (t *Timer) GetDeltaDuration() int64 {
	return t.delta_duration
}

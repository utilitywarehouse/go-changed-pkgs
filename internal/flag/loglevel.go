package flag

import (
	"fmt"
	"log/slog"
	"slices"
	"strings"

	"github.com/urfave/cli/v2"
)

// SlogLevelValue is a [cli.Generic] value that can be
// used in a [cli.GenericFlag], for example
//
//	app := &cli.App{
//		Flags: []cli.Flag{
//			&cli.GenericFlag{
//				Name:  "log-level",
//				Value: &flag.SlogLevelValue{},
//			},
//		},
//		Action: func(ctx *cli.Context) error {
//			level = ctx.Value("log-level").(slog.Level)
//			return nil
//		},
//	}
//
// It expects the provided argument to be a valid [slog.Level] name, e.g. for
// the above example `--log-level info`, otherwise an error will be raised when
// setting the flag.
type SlogLevelValue struct {
	slog.Level
}

var (
	levelNames []string = []string{
		"debug",
		"info",
		"warn",
		"error",
	}

	// SlogLevelValueUsage can be used as the `Usage` value for a [cli.Flag] that
	// uses [SlogLevelValue].
	SlogLevelValueUsage = fmt.Sprintf(
		"The level to log at. Valid values are: %s",
		strings.Join(levelNames, ", "),
	)
)

// SlogLevelValueFlag ready to use log-level [cli.Flag] using
// [SlogLevelValue]. It defaults to [slog.LevelWarn].
type SlogLevelValueFlag struct {
	*cli.GenericFlag
}

func NewSlogLevelValueFlag() *SlogLevelValueFlag {
	return &SlogLevelValueFlag{
		&cli.GenericFlag{
			Name:  "log-level",
			Usage: SlogLevelValueUsage,
			Value: &SlogLevelValue{Level: slog.LevelWarn},
		},
	}
}

func (v *SlogLevelValue) Set(value string) error {
	if slices.Contains(levelNames, value) {
		level := new(slog.Level)
		level.UnmarshalText([]byte(value)) //nolint:errcheck // value is a valid level
		v.Level = *level
		return nil
	}

	return fmt.Errorf(
		"invalid level %s: must be one of: %s",
		value,
		strings.Join(levelNames, ", "),
	)
}

func (v *SlogLevelValue) Get() any {
	return v.Level
}

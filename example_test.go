package errors_test

import (
	"fmt"

	"git.corout.in/golibs/fields"
	"git.corout.in/golibs/slog"
	"git.corout.in/golibs/slog/zaplogger"

	"git.corout.in/golibs/errors"
)

func Example_errorsUsage() {
	err := fmt.Errorf("Some error")

	err = errors.Ctx().
		Int("test-int-key", 10).
		Str("test-str-key", "blah-blah").
		Just(err)

	log := zaplogger.New(slog.DebugLevel, true)

	log.Error("Test error", err)
	log.Errorf("Test error with arg: %s", "test arg", err)
	log.Error(
		"Test error with additional fields",
		fields.Str("additional", "field"),
		fields.Bool("is_worked", true),
		err,
	)

	err = errors.Wrap(err, "test wrap wrapped error")
	log.Error("Have new wrapped error", err)

	err = errors.Ctx().Str("injecting", "this some injecting on next level").Just(err)
	log.Error("Next level injected fields", err)
}
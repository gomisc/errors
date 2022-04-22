package errors_test

import (
	"fmt"

	"git.dev.cloud.mts.ru/billing/common/pkg/infra/errors"
	"git.dev.cloud.mts.ru/billing/common/pkg/infra/fields"
	"git.dev.cloud.mts.ru/billing/common/pkg/infra/logger/zap"
)

func Example_errorsUsage() {
	err := fmt.Errorf("Some error")

	err = errors.Ctx().
		Int("test-int-key", 10).
		Str("test-str-key", "blah-blah").
		Just(err)

	log := zap.NewColored(5)

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

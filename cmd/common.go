package cmd

import (
	"github.com/hedon-go-road/ddd-archetype-go/pkg/log"
	"github.com/hedon-go-road/ddd-archetype-go/pkg/safe"
)

func init() {
	startSafe()
}

func startSafe() {
	safe.Callback(func(err any, stack []byte) {
		log.Error().
			Any("err", err).
			Str("stack", string(stack)).
			Msg("safe occurs panic")
	})
}

func StopSafe() {
	safe.Wait()
}

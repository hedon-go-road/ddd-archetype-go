package race

import (
	"testing"
)

func TestConcurrentOpMap(t *testing.T) {
	// use go test -race to check for potential concurrency issues if the test will run to related code.
	ConcurrentOpMap()
}

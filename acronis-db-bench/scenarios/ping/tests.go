package ping

import (
	"context"

	"github.com/acronis/perfkit/benchmark"

	"github.com/acronis/perfkit/acronis-db-bench/engine"
)

func init() {
	tests := []*engine.TestDesc{
		// Ping tests
		&TestPing,
	}

	scenario := &engine.TestScenario{
		Name:   "ping",
		Tests:  tests,
		Tables: make(map[string]engine.TestTable), // Empty map since ping test doesn't use tables
	}

	if err := engine.RegisterTestScenario(scenario); err != nil {
		panic(err)
	}
}

// TestPing tests just ping DB
var TestPing = engine.TestDesc{
	Name:        "ping",
	Metric:      "ping/sec",
	Description: "just ping DB",
	Category:    engine.TestOther,
	IsReadonly:  true,
	IsDBRTest:   false,
	Databases:   engine.ALL,
	LauncherFunc: func(b *benchmark.Benchmark, testDesc *engine.TestDesc) {
		worker := func(b *benchmark.Benchmark, c *engine.DBConnector, testDesc *engine.TestDesc, batch int) (loops int) { //nolint:revive
			if err := c.Database.Ping(context.Background()); err != nil {
				return 0
			}

			return 1
		}
		engine.TestGeneric(b, testDesc, worker, 0)
	},
}

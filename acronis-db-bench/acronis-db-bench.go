// Package main provides benchmarking utilities for databases.
// It includes various tests and options for database performance analysis.
package main

import (
	_ "net/http/pprof"

	// Engine
	"github.com/acronis/perfkit/acronis-db-bench/engine"

	// List of database drivers
	_ "github.com/acronis/perfkit/db/es/es8"         // elasticsearch 8 driver
	_ "github.com/acronis/perfkit/db/es/opensearch"  // opensearch driver
	_ "github.com/acronis/perfkit/db/sql/cassandra"  // cassandra driver
	_ "github.com/acronis/perfkit/db/sql/clickhouse" // clickhouse driver
	_ "github.com/acronis/perfkit/db/sql/mssql"      // mssql driver
	_ "github.com/acronis/perfkit/db/sql/mysql"      // mysql driver
	_ "github.com/acronis/perfkit/db/sql/postgres"   // postgres driver
	_ "github.com/acronis/perfkit/db/sql/sqlite"     // sqlite driver

	// List of test groups
	_ "github.com/acronis/perfkit/acronis-db-bench/test-groups/basic-scenarios"          // basic-scenarios
	_ "github.com/acronis/perfkit/acronis-db-bench/test-groups/composite-retrieval"      // composite-retreival
	_ "github.com/acronis/perfkit/acronis-db-bench/test-groups/json-search"              // json-search
	_ "github.com/acronis/perfkit/acronis-db-bench/test-groups/large-objects-operations" // large-objects-operations
	_ "github.com/acronis/perfkit/acronis-db-bench/test-groups/logs-search"              // logs-search
	_ "github.com/acronis/perfkit/acronis-db-bench/test-groups/ping"                     // ping
	_ "github.com/acronis/perfkit/acronis-db-bench/test-groups/sample-dwh"               // sample-dwh
	_ "github.com/acronis/perfkit/acronis-db-bench/test-groups/select-one"               // select-one
	_ "github.com/acronis/perfkit/acronis-db-bench/test-groups/sequence-operations"      // sequence-operations
	_ "github.com/acronis/perfkit/acronis-db-bench/test-groups/tenant-and-cti"           // tenant-and-cti
	_ "github.com/acronis/perfkit/acronis-db-bench/test-groups/time-series-operations"   // time-series-operations
	_ "github.com/acronis/perfkit/acronis-db-bench/test-groups/vector-search"            // vector-search

	// List of suites
	_ "github.com/acronis/perfkit/acronis-db-bench/suites/all-basic" // all-basic
)

func main() {
	engine.Main()
}

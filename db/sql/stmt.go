package sql

import (
	"database/sql"

	"github.com/acronis/perfkit/db"
)

// Statement defines the interface for prepared statement operations
type Statement interface {
	// Exec executes a prepared statement with the given arguments.
	// Returns a Result interface and error.
	// Usage example:
	//   stmt.Exec(2, 2, "test") // for most dialects
	//   stmt.Exec(sql.Named("Origin", 2), sql.Named("Type", 2), sql.Named("Name", "test")) // for MSSQL
	Exec(args ...any) (db.Result, error)

	// Close releases resources associated with the prepared statement.
	// Should be called when the statement is no longer needed.
	// Typically used with defer:
	//   defer stmt.Close()
	Close() error
}

// Stmt implements the sqlStatement interface by wrapping a *sql.Stmt
type Stmt struct {
	Stmt *sql.Stmt // underlying SQL prepared statement
}

// Exec executes the prepared statement with the given arguments.
// Implements the sqlStatement interface.
// Args can be either positional parameters (?, $1, etc.) or named parameters (@Name)
// depending on the SQL dialect being used.
// Returns the sql.Result from the underlying statement execution.
func (s *Stmt) Exec(args ...any) (db.Result, error) {
	return s.Stmt.Exec(args...)
}

// Close releases the database resources associated with the prepared statement.
// Should be called when done with the statement to prevent resource leaks.
// Implements the sqlStatement interface.
func (s *Stmt) Close() error {
	return s.Stmt.Close()
}

// Prepare creates a prepared statement for the given query.
// Part of the db.DatabaseAccessor interface.
// The query format depends on the dialect:
//   - MSSQL: Uses @param named parameters
//   - PostgreSQL: Uses $1, $2, etc. positional parameters
//   - Others: Uses ? positional parameters
//
// Example queries:
//   - MSSQL: "INSERT INTO table (col1, col2) VALUES (@Param1, @Param2)"
//   - PostgreSQL: "INSERT INTO table (col1, col2) VALUES ($1, $2)"
//   - Others: "INSERT INTO table (col1, col2) VALUES (?, ?)"
func (g *sqlGateway) Prepare(query string) (db.Stmt, error) {
	return g.rw.prepareContext(g.ctx, query)
}

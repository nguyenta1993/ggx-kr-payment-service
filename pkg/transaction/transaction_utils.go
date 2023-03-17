package transactionutils

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// Transaction is an interface that models the standard transaction in
// `database/sql`.
//
// To ensure `TxFn` funcs cannot commit or rollback a transaction (which is
// handled by `WithTransaction`), those methods are not included here.
// Have to define Sqlx with Content
type Transaction interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
	NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error)
}

// A Txfn is a function that will be called with an initialized `Transaction` object
// that can be used for executing statements and queries against a database.
type TxFn func(Transaction) error

// WithTransaction creates a new transaction and handles rollback/commit based on the
// error object returned by the `TxFn`
func WithTransaction(db *sqlx.DB, ctx context.Context, fn TxFn) (err error) {
	tx, err := db.BeginTxx(ctx, nil)
	if err != nil {
		return
	}

	defer func() {
		if p := recover(); p != nil {
			// a panic occurred, rollback and repanic
			err = tx.Rollback()
			panic(p)
		} else if err != nil {
			// something went wrong, rollback
			errRollback := tx.Rollback()
			if errRollback != nil {
				err = errRollback
			}
		} else {
			// all good, commit
			err = tx.Commit()
		}
	}()

	err = fn(tx)
	return err
}

package helper

import "database/sql"

func CommitOrRollback(tx *sql.Tx) {
	err := recover()

	if err != nil {
		errRollback := tx.Rollback()
		LogErrorAndPanic(errRollback)
	} else {
		errCommit := tx.Commit()
		LogErrorAndPanic(errCommit)
	}
}

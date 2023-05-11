// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package repository

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.createPaymentStmt, err = db.PrepareContext(ctx, createPayment); err != nil {
		return nil, fmt.Errorf("error preparing query CreatePayment: %w", err)
	}
	if q.createTransactionStmt, err = db.PrepareContext(ctx, createTransaction); err != nil {
		return nil, fmt.Errorf("error preparing query CreateTransaction: %w", err)
	}
	if q.deleteExpiredTokensStmt, err = db.PrepareContext(ctx, deleteExpiredTokens); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteExpiredTokens: %w", err)
	}
	if q.deleteTokenStmt, err = db.PrepareContext(ctx, deleteToken); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteToken: %w", err)
	}
	if q.deleteTokensByCredentialStmt, err = db.PrepareContext(ctx, deleteTokensByCredential); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteTokensByCredential: %w", err)
	}
	if q.getPaymentStmt, err = db.PrepareContext(ctx, getPayment); err != nil {
		return nil, fmt.Errorf("error preparing query GetPayment: %w", err)
	}
	if q.getPaymentByExternalIDStmt, err = db.PrepareContext(ctx, getPaymentByExternalID); err != nil {
		return nil, fmt.Errorf("error preparing query GetPaymentByExternalID: %w", err)
	}
	if q.getPendingTransactionsStmt, err = db.PrepareContext(ctx, getPendingTransactions); err != nil {
		return nil, fmt.Errorf("error preparing query GetPendingTransactions: %w", err)
	}
	if q.getTokenStmt, err = db.PrepareContext(ctx, getToken); err != nil {
		return nil, fmt.Errorf("error preparing query GetToken: %w", err)
	}
	if q.getTransactionStmt, err = db.PrepareContext(ctx, getTransaction); err != nil {
		return nil, fmt.Errorf("error preparing query GetTransaction: %w", err)
	}
	if q.getTransactionByPaymentIDSourceWalletAndMintStmt, err = db.PrepareContext(ctx, getTransactionByPaymentIDSourceWalletAndMint); err != nil {
		return nil, fmt.Errorf("error preparing query GetTransactionByPaymentIDSourceWalletAndMint: %w", err)
	}
	if q.getTransactionByReferenceStmt, err = db.PrepareContext(ctx, getTransactionByReference); err != nil {
		return nil, fmt.Errorf("error preparing query GetTransactionByReference: %w", err)
	}
	if q.getTransactionsByPaymentIDStmt, err = db.PrepareContext(ctx, getTransactionsByPaymentID); err != nil {
		return nil, fmt.Errorf("error preparing query GetTransactionsByPaymentID: %w", err)
	}
	if q.markPaymentsExpiredStmt, err = db.PrepareContext(ctx, markPaymentsExpired); err != nil {
		return nil, fmt.Errorf("error preparing query MarkPaymentsExpired: %w", err)
	}
	if q.markTransactionsAsExpiredStmt, err = db.PrepareContext(ctx, markTransactionsAsExpired); err != nil {
		return nil, fmt.Errorf("error preparing query MarkTransactionsAsExpired: %w", err)
	}
	if q.storeTokenStmt, err = db.PrepareContext(ctx, storeToken); err != nil {
		return nil, fmt.Errorf("error preparing query StoreToken: %w", err)
	}
	if q.updatePaymentStatusStmt, err = db.PrepareContext(ctx, updatePaymentStatus); err != nil {
		return nil, fmt.Errorf("error preparing query UpdatePaymentStatus: %w", err)
	}
	if q.updateTransactionByReferenceStmt, err = db.PrepareContext(ctx, updateTransactionByReference); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateTransactionByReference: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.createPaymentStmt != nil {
		if cerr := q.createPaymentStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createPaymentStmt: %w", cerr)
		}
	}
	if q.createTransactionStmt != nil {
		if cerr := q.createTransactionStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createTransactionStmt: %w", cerr)
		}
	}
	if q.deleteExpiredTokensStmt != nil {
		if cerr := q.deleteExpiredTokensStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteExpiredTokensStmt: %w", cerr)
		}
	}
	if q.deleteTokenStmt != nil {
		if cerr := q.deleteTokenStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteTokenStmt: %w", cerr)
		}
	}
	if q.deleteTokensByCredentialStmt != nil {
		if cerr := q.deleteTokensByCredentialStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteTokensByCredentialStmt: %w", cerr)
		}
	}
	if q.getPaymentStmt != nil {
		if cerr := q.getPaymentStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getPaymentStmt: %w", cerr)
		}
	}
	if q.getPaymentByExternalIDStmt != nil {
		if cerr := q.getPaymentByExternalIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getPaymentByExternalIDStmt: %w", cerr)
		}
	}
	if q.getPendingTransactionsStmt != nil {
		if cerr := q.getPendingTransactionsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getPendingTransactionsStmt: %w", cerr)
		}
	}
	if q.getTokenStmt != nil {
		if cerr := q.getTokenStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getTokenStmt: %w", cerr)
		}
	}
	if q.getTransactionStmt != nil {
		if cerr := q.getTransactionStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getTransactionStmt: %w", cerr)
		}
	}
	if q.getTransactionByPaymentIDSourceWalletAndMintStmt != nil {
		if cerr := q.getTransactionByPaymentIDSourceWalletAndMintStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getTransactionByPaymentIDSourceWalletAndMintStmt: %w", cerr)
		}
	}
	if q.getTransactionByReferenceStmt != nil {
		if cerr := q.getTransactionByReferenceStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getTransactionByReferenceStmt: %w", cerr)
		}
	}
	if q.getTransactionsByPaymentIDStmt != nil {
		if cerr := q.getTransactionsByPaymentIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getTransactionsByPaymentIDStmt: %w", cerr)
		}
	}
	if q.markPaymentsExpiredStmt != nil {
		if cerr := q.markPaymentsExpiredStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing markPaymentsExpiredStmt: %w", cerr)
		}
	}
	if q.markTransactionsAsExpiredStmt != nil {
		if cerr := q.markTransactionsAsExpiredStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing markTransactionsAsExpiredStmt: %w", cerr)
		}
	}
	if q.storeTokenStmt != nil {
		if cerr := q.storeTokenStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing storeTokenStmt: %w", cerr)
		}
	}
	if q.updatePaymentStatusStmt != nil {
		if cerr := q.updatePaymentStatusStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updatePaymentStatusStmt: %w", cerr)
		}
	}
	if q.updateTransactionByReferenceStmt != nil {
		if cerr := q.updateTransactionByReferenceStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateTransactionByReferenceStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                                               DBTX
	tx                                               *sql.Tx
	createPaymentStmt                                *sql.Stmt
	createTransactionStmt                            *sql.Stmt
	deleteExpiredTokensStmt                          *sql.Stmt
	deleteTokenStmt                                  *sql.Stmt
	deleteTokensByCredentialStmt                     *sql.Stmt
	getPaymentStmt                                   *sql.Stmt
	getPaymentByExternalIDStmt                       *sql.Stmt
	getPendingTransactionsStmt                       *sql.Stmt
	getTokenStmt                                     *sql.Stmt
	getTransactionStmt                               *sql.Stmt
	getTransactionByPaymentIDSourceWalletAndMintStmt *sql.Stmt
	getTransactionByReferenceStmt                    *sql.Stmt
	getTransactionsByPaymentIDStmt                   *sql.Stmt
	markPaymentsExpiredStmt                          *sql.Stmt
	markTransactionsAsExpiredStmt                    *sql.Stmt
	storeTokenStmt                                   *sql.Stmt
	updatePaymentStatusStmt                          *sql.Stmt
	updateTransactionByReferenceStmt                 *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                           tx,
		tx:                           tx,
		createPaymentStmt:            q.createPaymentStmt,
		createTransactionStmt:        q.createTransactionStmt,
		deleteExpiredTokensStmt:      q.deleteExpiredTokensStmt,
		deleteTokenStmt:              q.deleteTokenStmt,
		deleteTokensByCredentialStmt: q.deleteTokensByCredentialStmt,
		getPaymentStmt:               q.getPaymentStmt,
		getPaymentByExternalIDStmt:   q.getPaymentByExternalIDStmt,
		getPendingTransactionsStmt:   q.getPendingTransactionsStmt,
		getTokenStmt:                 q.getTokenStmt,
		getTransactionStmt:           q.getTransactionStmt,
		getTransactionByPaymentIDSourceWalletAndMintStmt: q.getTransactionByPaymentIDSourceWalletAndMintStmt,
		getTransactionByReferenceStmt:                    q.getTransactionByReferenceStmt,
		getTransactionsByPaymentIDStmt:                   q.getTransactionsByPaymentIDStmt,
		markPaymentsExpiredStmt:                          q.markPaymentsExpiredStmt,
		markTransactionsAsExpiredStmt:                    q.markTransactionsAsExpiredStmt,
		storeTokenStmt:                                   q.storeTokenStmt,
		updatePaymentStatusStmt:                          q.updatePaymentStatusStmt,
		updateTransactionByReferenceStmt:                 q.updateTransactionByReferenceStmt,
	}
}
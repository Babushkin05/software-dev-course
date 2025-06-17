// internal/service/account_storage.go
package service

import (
	"context"
	"database/sql"

	"github.com/Babushkin05/software-dev-course/kr3/payments-service/internal/db"
	"github.com/Babushkin05/software-dev-course/kr3/payments-service/internal/model"
)

type AccountStorage interface {
	Create(ctx context.Context, userID string) (*model.Account, error)
	AddBalance(ctx context.Context, userID string, amount int64) (int64, error)
	GetBalance(ctx context.Context, userID string) (int64, error)
	Withdraw(ctx context.Context, userID string, amount int64) (int64, error)

	// Inbox
	SaveInboxMessage(msg db.InboxMessage) error
	FetchUnprocessedInboxMessages(limit int) ([]db.InboxMessage, error)
	MarkInboxMessageProcessed(messageID string) error

	// Outbox
	SaveOutboxMessage(ctx context.Context, msg db.OutboxMessage) error
	SaveOutboxMessageTx(ctx context.Context, tx *sql.Tx, msg db.OutboxMessage) error
	FetchUnsentOutboxMessages(limit int) ([]db.OutboxMessage, error)
	MarkOutboxMessageSent(messageID string) error

	// Tx
	BeginTx(ctx context.Context) (*sql.Tx, error)
	RollbackTx(tx *sql.Tx)
	CommitTx(tx *sql.Tx) error
	WithdrawTx(ctx context.Context, tx *sql.Tx, userID string, amount int64) (int64, error)
}

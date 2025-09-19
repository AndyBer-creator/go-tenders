package storage

import (
	"context"

	"github.com/AndyBer-creator/go-tenders/model"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type Storage interface {
	GetUserBids(ctx echo.Context, params model.GetUserBidsParams) error

	// Создание нового предложения (POST /bids/new)
	CreateBid(ctx echo.Context, bid model.Bid) error

	// Редактирование параметров предложения (PATCH /bids/{bidId}/edit)
	EditBid(ctx echo.Context, bidId model.BidId, params model.EditBidParams) error

	// Отправка отзыва по предложению (PUT /bids/{bidId}/feedback)
	SubmitBidFeedback(ctx echo.Context, bidId model.BidId, params model.SubmitBidFeedbackParams) error

	// Откат версии предложения (PUT /bids/{bidId}/rollback/{version})
	RollbackBid(ctx echo.Context, bidId model.BidId, version int32, params model.RollbackBidParams) error

	// Получение текущего статуса предложения (GET /bids/{bidId}/status)
	GetBidStatus(ctx echo.Context, bidId model.BidId, params model.GetBidStatusParams) error

	// Изменение статуса предложения (PUT /bids/{bidId}/status)
	UpdateBidStatus(ctx echo.Context, bidId model.BidId, params model.UpdateBidStatusParams) error

	// Отправка решения по предложению (PUT /bids/{bidId}/submit_decision)
	SubmitBidDecision(ctx echo.Context, bidId model.BidId, params model.SubmitBidDecisionParams) error

	// Получение списка предложений для тендера (GET /bids/{tenderId}/list)
	GetBidsForTender(ctx echo.Context, tenderId model.TenderId, params model.GetBidsForTenderParams) error

	// Просмотр отзывов на прошлые предложения (GET /bids/{tenderId}/reviews)
	GetBidReviews(ctx echo.Context, tenderId model.TenderId, params model.GetBidReviewsParams) error

	// Проверка доступности сервера (GET /ping)
	CheckServer(ctx echo.Context) error

	// Получение списка тендеров (GET /tenders)
	GetTenders(ctx echo.Context, params model.GetTendersParams) error

	// Получить тендеры пользователя (GET /tenders/my)
	GetUserTenders(ctx echo.Context, params model.GetUserTendersParams) error

	// Создание нового тендера (POST /tenders/new)
	CreateTender(ctx echo.Context) error

	// Редактирование тендера (PATCH /tenders/{tenderId}/edit)
	EditTender(ctx echo.Context, tenderId model.TenderId, params model.EditTenderParams) error

	// Откат версии тендера (PUT /tenders/{tenderId}/rollback/{version})
	RollbackTender(ctx echo.Context, tenderId model.TenderId, version int32, params model.RollbackTenderParams) error

	// Получение текущего статуса тендера (GET /tenders/{tenderId}/status)
	GetTenderStatus(ctx echo.Context, tenderId model.TenderId, params model.GetTenderStatusParams) error

	// Изменение статуса тендера (PUT /tenders/{tenderId}/status)
	UpdateTenderStatus(ctx echo.Context, tenderId model.TenderId, params model.UpdateTenderStatusParams) error
}

type PostgresStorage struct {
	db *sqlx.DB
}

func NewPostgresStorage(db *sqlx.DB) *PostgresStorage {
	return &PostgresStorage{db: db}
}

func (s *PostgresStorage) GetUserBids(ctx context.Context, userId string, limit, offset int) ([]model.Bid, error) {
	query := `
        SELECT id, title, description, created_at, updated_at
        FROM bids
        WHERE user_id = $1
        ORDER BY created_at DESC
        LIMIT $2 OFFSET $3
    `
	rows, err := s.db.QueryContext(ctx, query, userId, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bids []model.Bid
	for rows.Next() {
		var b model.Bid
		if err := rows.Scan(&b.Id, &b.Title, &b.Amount, &b.CreatedAt, &b.UpdatedAt); err != nil {
			return nil, err
		}
		bids = append(bids, b)
	}
	return bids, rows.Err()
}

func (s *PostgresStorage) CreateBid(ctx context.Context, bid model.Bid) error {
	query := `INSERT INTO bids (id, name, description) VALUES ($1, $2, $3)`
	_, err := s.db.ExecContext(ctx, query, bid.Id, bid.Name, bid.Description)
	return err
}

func (s *PostgresStorage) EditBid(ctx context.Context, bidId string, params model.EditBidParams) error {
	query := `
        UPDATE bids
        SET title = $1,
            description = $2
        WHERE id = $3
    `
	_, err := s.db.ExecContext(ctx, query, params.Title, params.Description, bidId)
	return err
}

func (s *PostgresStorage) SubmitBidFeedback(ctx context.Context, bidId string, params model.SubmitBidFeedbackParams) error {
	query := `
        INSERT INTO bid_feedback (bid_id, feedback_text, rating, created_at, updated_at)
        VALUES ($1, $2, $3, NOW(), NOW())
        ON CONFLICT (bid_id) DO UPDATE SET
            feedback_text = EXCLUDED.feedback_text,
            rating = EXCLUDED.rating,
            updated_at = NOW()
    `
	_, err := s.db.ExecContext(ctx, query, bidId, params.FeedbackText, params.Rating)
	return err
}

func (s *PostgresStorage) RollbackBid(ctx context.Context, bidId string, version int32, params model.RollbackBidParams) error {
	return s.execTx(ctx, func(q *Queries) error {
		// Копируем текущую версию в историю
		err := q.ArchiveCurrentBid(ctx, bidId)
		if err != nil {
			return err
		}
		// Восстанавливаем выбранную версию из истории
		err = q.RestoreBidVersion(ctx, bidId, version)
		if err != nil {
			return err
		}
		return nil
	})
}

func (s *PostgresStorage) GetBidStatus(ctx context.Context, bidId string) (*model.BidStatus, error) {
	query := `SELECT status, updated_at FROM bids WHERE id = $1`
	var status model.BidStatus
	err := s.db.QueryRowContext(ctx, query, bidId).Scan(&status.Status, &status.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &status, nil
}

func (s *PostgresStorage) UpdateBidStatus(ctx context.Context, bidId string, params model.UpdateBidStatusParams) error {
	query := `
        UPDATE bids
        SET status = $1,
            updated_at = NOW()
        WHERE id = $2
    `
	_, err := s.db.ExecContext(ctx, query, params.Status, bidId)
	return err
}

func (s *PostgresStorage) SubmitBidDecision(ctx context.Context, bidId string, params model.SubmitBidDecisionParams) error {
	query := `
        INSERT INTO bid_decisions (bid_id, decision, decided_at)
        VALUES ($1, $2, NOW())
        ON CONFLICT (bid_id) DO UPDATE SET
            decision = EXCLUDED.decision,
            decided_at = NOW()
    `
	_, err := s.db.ExecContext(ctx, query, bidId, params.Decision)
	return err
}

func (s *PostgresStorage) GetBidsForTender(ctx context.Context, tenderId string, limit, offset int) ([]model.Bid, error) {
	query := `
        SELECT id, title, description, amount, user_id, created_at, updated_at
        FROM bids
        WHERE tender_id = $1
        ORDER BY created_at DESC
        LIMIT $2 OFFSET $3
    `
	rows, err := s.db.QueryContext(ctx, query, tenderId, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bids []model.Bid
	for rows.Next() {
		var b model.Bid
		if err := rows.Scan(&b.Id, &b.Title, &b.Description, &b.Amount, &b.UserId, &b.CreatedAt, &b.UpdatedAt); err != nil {
			return nil, err
		}
		bids = append(bids, b)
	}
	return bids, rows.Err()
}

func (s *PostgresStorage) GetBidReviews(ctx context.Context, tenderId string, limit, offset int) ([]model.BidReview, error) {
	query := `
        SELECT id, bid_id, review_text, rating, created_at, updated_at
        FROM bid_reviews
        WHERE tender_id = $1
        ORDER BY created_at DESC
        LIMIT $2 OFFSET $3
    `
	rows, err := s.db.QueryContext(ctx, query, tenderId, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reviews []model.BidReview
	for rows.Next() {
		var r model.BidReview
		if err := rows.Scan(&r.Id, &r.BidId, &r.ReviewText, &r.Rating, &r.CreatedAt, &r.UpdatedAt); err != nil {
			return nil, err
		}
		reviews = append(reviews, r)
	}
	return reviews, rows.Err()
}

func (s *PostgresStorage) GetTenders(ctx context.Context, limit, offset int) ([]model.Tender, error) {
	query := `
        SELECT id, title, description, created_at, updated_at
        FROM tenders
        ORDER BY created_at DESC
        LIMIT $1 OFFSET $2
    `
	rows, err := s.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tenders []model.Tender
	for rows.Next() {
		var t model.Tender
		if err := rows.Scan(&t.Id, &t.Title, &t.Description, &t.CreatedAt, &t.UpdatedAt); err != nil {
			return nil, err
		}
		tenders = append(tenders, t)
	}
	return tenders, rows.Err()
}

func (s *PostgresStorage) GetUserTenders(ctx context.Context, userId string, limit, offset int) ([]model.Tender, error) {
	query := `
        SELECT id, title, description, created_at, updated_at
        FROM tenders
        WHERE user_id = $1
        ORDER BY created_at DESC
        LIMIT $2 OFFSET $3
    `
	rows, err := s.db.QueryContext(ctx, query, userId, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tenders []model.Tender
	for rows.Next() {
		var t model.Tender
		if err := rows.Scan(&t.Id, &t.Title, &t.Description, &t.CreatedAt, &t.UpdatedAt); err != nil {
			return nil, err
		}
		tenders = append(tenders, t)
	}
	return tenders, rows.Err()
}

func (s *PostgresStorage) CreateTender(ctx context.Context, tender model.Tender) error {
	query := `
        INSERT INTO tenders (id, title, description, created_at, updated_at)
        VALUES ($1, $2, $3, NOW(), NOW())
    `
	_, err := s.db.ExecContext(ctx, query, tender.Id, tender.Title, tender.Description)
	return err
}

func (s *PostgresStorage) EditTender(ctx context.Context, tenderId string, params model.EditTenderParams) error {
	query := `
        UPDATE tenders
        SET title = $1,
            description = $2,
            updated_at = NOW()
        WHERE id = $3
    `
	_, err := s.db.ExecContext(ctx, query, params.Title, params.Description, tenderId)
	return err
}

func (s *PostgresStorage) RollbackTender(ctx context.Context, tenderId string, version int32, params model.RollbackTenderParams) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		}
	}()

	_, err = tx.ExecContext(ctx, `
        INSERT INTO tenders_history (tender_id, version, data, archived_at)
        SELECT id, version, row_to_json(tenders.*), NOW()
        FROM tenders WHERE id = $1`, tenderId)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.ExecContext(ctx, `
        UPDATE tenders t
        SET data = h.data,
            version = h.version,
            updated_at = NOW()
        FROM tenders_history h
        WHERE t.id = $1 AND h.tender_id = $1 AND h.version = $2
    `, tenderId, version)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (s *PostgresStorage) GetTenderStatus(ctx context.Context, tenderId string) (*model.TenderStatus, error) {
	query := `SELECT status, updated_at FROM tenders WHERE id = $1`
	var status model.TenderStatus
	err := s.db.QueryRowContext(ctx, query, tenderId).Scan(&status.Status, &status.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &status, nil
}

func (s *PostgresStorage) UpdateTenderStatus(ctx context.Context, tenderId string, params model.UpdateBidStatusParams) error {
	query := `
        UPDATE tenders
        SET status = $1,
            updated_at = NOW()
        WHERE id = $2
    `
	_, err := s.db.ExecContext(ctx, query, params.Status, tenderId)
	return err
}

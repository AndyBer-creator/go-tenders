package storage

import (
	"go-tenders/model"

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

func (s *PostgresStorage) CreateBid(ctx echo.Context, bid model.Bid) error {
	query := `INSERT INTO bids (id, title, amount) VALUES ($1, $2, $3)`
	_, err := s.db.ExecContext(ctx.Request().Context(), query, bid.Id, bid.Name, bid.Description)
	return err
}

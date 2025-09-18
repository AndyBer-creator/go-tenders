package server

import (
	"go-tenders/api"
	"go-tenders/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

// storage storage.Storage
// заглушки ----------------------------------------------
type Storage interface {
	GetBids(limit, offset int, username string) ([]model.Bid, error)
	CreateBid(bid model.Bid) error
	UpdateBid(bidId model.BidId, params model.EditBidParams) error
	SaveBidFeedback(bidId model.BidId, params model.SubmitBidFeedbackParams) error
	// другие методы по мере необходимости
}

type Logger interface {
	Info(args ...interface{})
	Error(args ...interface{})
}

type Config struct {
	Port int
}

// -------------------------------------------------------
type Server struct {
	storage Storage // интерфейс доступа к данным (репозиторий)
	logger  Logger  // интерфейс логгера
	config  *Config // конфигурация приложения
}

var _ api.ServerInterface = (*Server)(nil)

func NewServer(storage Storage, logger Logger, cfg *Config) *Server {
	return &Server{
		storage: storage,
		logger:  logger,
		config:  cfg,
	}
}
func (s *Server) CheckServer(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "pong")
}

func (s *Server) GetUserBids(ctx echo.Context, params model.GetUserBidsParams) error {
	var username string
	if params.Username != nil {
		username = string(*params.Username)
	}
	var limit int
	if params.Limit != nil {
		limit = int(*params.Limit)
	}
	var offset int
	if params.Offset != nil {
		offset = int(*params.Offset)
	}

	bids, err := s.storage.GetBids(limit, offset, username)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Ошибка при получении предложений")
	}

	return ctx.JSON(http.StatusOK, bids)
}
func (s *Server) CreateBid(ctx echo.Context, bid model.Bid) error {
	return ctx.NoContent(http.StatusNotImplemented)
}

func (s *Server) EditBid(ctx echo.Context, bidId model.BidId, params model.EditBidParams) error {
	return ctx.NoContent(http.StatusNoContent)
}

func (s *Server) SubmitBidFeedback(ctx echo.Context, bidId model.BidId, params model.SubmitBidFeedbackParams) error {
	return ctx.NoContent(http.StatusNoContent)
}

func (s *Server) RollbackBid(ctx echo.Context, bidId model.BidId, version int32, params model.RollbackBidParams) error {
	return ctx.NoContent(http.StatusNoContent)
}

func (s *Server) GetBidStatus(ctx echo.Context, bidId model.BidId, params model.GetBidStatusParams) error {
	return ctx.NoContent(http.StatusNoContent)
}

func (s *Server) UpdateBidStatus(ctx echo.Context, bidId model.BidId, params model.UpdateBidStatusParams) error {
	return ctx.NoContent(http.StatusNoContent)
}

func (s *Server) SubmitBidDecision(ctx echo.Context, bidId model.BidId, params model.SubmitBidDecisionParams) error {
	return ctx.NoContent(http.StatusNoContent)
}

func (s *Server) GetBidsForTender(ctx echo.Context, tenderId model.TenderId, params model.GetBidsForTenderParams) error {
	return ctx.NoContent(http.StatusNoContent)
}

func (s *Server) GetBidReviews(ctx echo.Context, tenderId model.TenderId, params model.GetBidReviewsParams) error {
	return ctx.NoContent(http.StatusNoContent)
}

func (s *Server) GetTenders(ctx echo.Context, params model.GetTendersParams) error {
	return ctx.NoContent(http.StatusNoContent)
}

func (s *Server) GetUserTenders(ctx echo.Context, params model.GetUserTendersParams) error {
	return ctx.NoContent(http.StatusNoContent)
}

func (s *Server) CreateTender(ctx echo.Context) error {
	return ctx.NoContent(http.StatusNoContent)
}

func (s *Server) EditTender(ctx echo.Context, tenderId model.TenderId, params model.EditTenderParams) error {
	return ctx.NoContent(http.StatusNoContent)
}

func (s *Server) RollbackTender(ctx echo.Context, tenderId model.TenderId, version int32, params model.RollbackTenderParams) error {
	return ctx.NoContent(http.StatusNoContent)
}

func (s *Server) GetTenderStatus(ctx echo.Context, tenderId model.TenderId, params model.GetTenderStatusParams) error {
	return ctx.NoContent(http.StatusNoContent)
}

func (s *Server) UpdateTenderStatus(ctx echo.Context, tenderId model.TenderId, params model.UpdateTenderStatusParams) error {
	return ctx.NoContent(http.StatusNoContent)
}

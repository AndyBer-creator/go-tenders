package server

import (
	"net/http"

	"go-tenders/api"
	"go-tenders/config"
	"go-tenders/model"
	"go-tenders/storage"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Logger интерфейс для логирования
type Logger interface {
	Info(args ...interface{})
	Error(args ...interface{})
}

// Config структура конфигурации сервера
type Config struct {
	Host string
	Port int
}

// Server структура вашего сервера с зависимостями
type Server struct {
	storage storage.Storage
	logger  Logger
	config  *config.Config
}

// Проверка соответствия интерфейсу api.ServerInterface
var _ api.ServerInterface = (*Server)(nil)

// Конструктор сервера
func NewServer(storage storage.Storage, logger Logger, cfg *config.Config) *Server {
	return &Server{
		storage: storage,
		logger:  logger,
		config:  cfg,
	}
}

// Метод запуска HTTP сервера
func (s *Server) Start(address string) error {
	e := echo.New()

	// Добавляем middleware для логирования и восстановления после паники
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Регистрируем обработчики API с префиксом "/api/v1"
	api.RegisterHandlersWithBaseURL(e, s, "/api/v1")

	s.logger.Info("Server starting at ", address)
	return e.Start(address)
}

// Реализация метода проверки сервера
func (s *Server) CheckServer(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "pong")
}

func (s *Server) GetUserBids(ctx echo.Context, params model.GetUserBidsParams) error {
	username := ""
	if params.Username != nil {
		username = *params.Username
	}
	limit := 0
	if params.Limit != nil {
		limit = *params.Limit
	}
	offset := 0
	if params.Offset != nil {
		offset = *params.Offset
	}

	bids, err := s.storage.GetBids(limit, offset, username)
	if err != nil {
		s.logger.Error("GetUserBids error: ", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get user bids")
	}

	return ctx.JSON(http.StatusOK, bids)
}

func (s *Server) CreateBid(ctx echo.Context, bid model.Bid) error {
	err := s.storage.CreateBid(bid)
	if err != nil {
		s.logger.Error("CreateBid error: ", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create bid")
	}
	return ctx.JSON(http.StatusCreated, bid)
}

func (s *Server) EditBid(ctx echo.Context, bidId model.BidId, params model.EditBidParams) error {
	err := s.storage.UpdateBid(bidId, params)
	if err != nil {
		s.logger.Error("EditBid error: ", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to edit bid")
	}
	return ctx.NoContent(http.StatusNoContent)
}

func (s *Server) SubmitBidFeedback(ctx echo.Context, bidId model.BidId, params model.SubmitBidFeedbackParams) error {
	err := s.storage.SaveBidFeedback(bidId, params)
	if err != nil {
		s.logger.Error("SubmitBidFeedback error: ", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to submit bid feedback")
	}
	return ctx.NoContent(http.StatusNoContent)
}

func (s *Server) RollbackBid(ctx echo.Context, bidId model.BidId, version int32, params model.RollbackBidParams) error {
	err := s.storage.RollbackBid(bidId, version, params)
	if err != nil {
		s.logger.Error("RollbackBid error: ", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to rollback bid")
	}
	return ctx.NoContent(http.StatusNoContent)
}

func (s *Server) GetBidStatus(ctx echo.Context, bidId model.BidId, params model.GetBidStatusParams) error {
	status, err := s.storage.GetBidStatus(bidId, params)
	if err != nil {
		s.logger.Error("GetBidStatus error: ", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get status bid")
	}
	return ctx.JSON(http.StatusOK, status)
}

func (s *Server) UpdateBidStatus(ctx echo.Context, bidId model.BidId, params model.UpdateBidStatusParams) error {
	err := s.storage.UpdateBidStatus(didId, params)
	if err != nil {
		s.logger.Error("UpdateBidStatus error: ", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to update status bid")
	}
	return ctx.NoContent(http.StatusNoContent)
}

func (s *Server) SubmitBidDecision(ctx echo.Context, bidId model.BidId, params model.SubmitBidDecisionParams) error {
	err := s.storage.SubmitBidDecision(didId, params)
	if err != nil {
		s.logger.Error("SubmitBidDecision error: ", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to submit decision bid")
	}
	return ctx.NoContent(http.StatusNoContent)
}

func (s *Server) GetBidsForTender(ctx echo.Context, tenderId model.TenderId, params model.GetBidsForTenderParams) error {
	bids, err := s.storage.GetBidsForTender(tenderId, params)
	if err != nil {
		s.logger.Error("GetBidsForTender error: ", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get bids for tender")
	}
	return ctx.JSON(http.StatusOK, bids)
}

func (s *Server) GetBidReviews(ctx echo.Context, tenderId model.TenderId, params model.GetBidReviewsParams) error {
	authorUsername := ""
	if params.AuthorUsername != nil {
		authorUsername = *params.AuthorUsername
	}
	requesterUsername := ""
	if params.RequesterUsername != nil {
		requesterUsername = *params.RequesterUsername
	}

	limit := 0
	if params.Limit != nil {
		limit = *params.Limit
	}

	offset := 0
	if params.Offset != nil {
		offset = *params.Offset
	}

	reviews, err := s.storage.GetBidReviews(tenderId, authorUsername, requesterUsername, limit, offset)
	if err != nil {
		s.logger.Error("GetBidReviews error: ", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get bid reviews")
	}

	return ctx.JSON(http.StatusOK, reviews)
}

func (s *Server) GetTenders(ctx echo.Context, params model.GetTendersParams) error {
	limit := 0
	if params.Limit != nil {
		limit = *params.Limit
	}
	offset := 0
	if params.Offset != nil {
		offset = *params.Offset
	}
	serviceType := ""
	if params.ServiceType != nil {
		serviceType = *params.ServiceType
	}

	tenders, err := s.storage.GetTenders(limit, offset, serviceType)
	if err != nil {
		s.logger.Error("GetTenders error: ", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get tenders")
	}
	return ctx.JSON(http.StatusOK, tenders)
}

func (s *Server) GetUserTenders(ctx echo.Context, params model.GetUserTendersParams) error {
	limit := 0
	if params.Limit != nil {
		limit = *params.Limit
	}
	offset := 0
	if params.Offset != nil {
		offset = *params.Offset
	}
	username := ""
	if params.Username != nil {
		username = *params.Username
	}

	tenders, err := s.storage.GetUserTenders(limit, offset, username)
	if err != nil {
		s.logger.Error("GetUserTenders error: ", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get user tenders")
	}

	return ctx.JSON(http.StatusOK, tenders)
}

func (s *Server) CreateTender(ctx echo.Context) error {
	var tender model.Tender
	if err := ctx.Bind(&tender); err != nil {
		s.logger.Error("CreateTender bind error: ", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
	}

	err := s.storage.CreateTender(tender)
	if err != nil {
		s.logger.Error("CreateTender error: ", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create tender")
	}

	return ctx.JSON(http.StatusCreated, tender)
}

func (s *Server) EditTender(ctx echo.Context, tenderId model.TenderId, params model.EditTenderParams) error {
	err := s.storage.UpdateTender(tenderId, params)
	if err != nil {
		s.logger.Error("EditTender error: ", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to edit tender")
	}
	return ctx.NoContent(http.StatusNoContent)
}

func (s *Server) RollbackTender(ctx echo.Context, tenderId model.TenderId, version int32, params model.RollbackTenderParams) error {
	err := s.storage.RollbackTender(ctx, tenderId, version, params)
	if err != nil {
		s.logger.Error("RollbackTender error: ", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to rollback tender")
	}
	return ctx.NoContent(http.StatusNoContent)
}

func (s *Server) GetTenderStatus(ctx echo.Context, tenderId model.TenderId, params model.GetTenderStatusParams) error {
	// Получаем стандартный context.Context из echo.Context
	stdCtx := ctx.Request().Context()

	// Вызываем метод хранения, передавая стандартный Context
	status, err := s.storage.GetTenderStatus(stdCtx, tenderId, params)
	if err != nil {
		s.logger.Error("GetTenderStatus error: ", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get tender status")
	}
	return ctx.JSON(http.StatusOK, status)
}

package api

import (
	"fmt"
	"go-tenders/model"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Получение списка ваших предложений
	// (GET /bids/my)
	GetUserBids(ctx echo.Context, params model.GetUserBidsParams) error
	// Создание нового предложения
	// (POST /bids/new)
	CreateBid(ctx echo.Context, bid model.Bid) error
	// Редактирование параметров предложения
	// (PATCH /bids/{bidId}/edit)
	EditBid(ctx echo.Context, bidId model.BidId, params model.EditBidParams) error
	// Отправка отзыва по предложению
	// (PUT /bids/{bidId}/feedback)
	SubmitBidFeedback(ctx echo.Context, bidId model.BidId, params model.SubmitBidFeedbackParams) error
	// Откат версии предложения
	// (PUT /bids/{bidId}/rollback/{version})
	RollbackBid(ctx echo.Context, bidId model.BidId, version int32, params model.RollbackBidParams) error
	// Получение текущего статуса предложения
	// (GET /bids/{bidId}/status)
	GetBidStatus(ctx echo.Context, bidId model.BidId, params model.GetBidStatusParams) error
	// Изменение статуса предложения
	// (PUT /bids/{bidId}/status)
	UpdateBidStatus(ctx echo.Context, bidId model.BidId, params model.UpdateBidStatusParams) error
	// Отправка решения по предложению
	// (PUT /bids/{bidId}/submit_decision)
	SubmitBidDecision(ctx echo.Context, bidId model.BidId, params model.SubmitBidDecisionParams) error
	// Получение списка предложений для тендера
	// (GET /bids/{tenderId}/list)
	GetBidsForTender(ctx echo.Context, tenderId model.TenderId, params model.GetBidsForTenderParams) error
	// Просмотр отзывов на прошлые предложения
	// (GET /bids/{tenderId}/reviews)
	GetBidReviews(ctx echo.Context, tenderId model.TenderId, params model.GetBidReviewsParams) error
	// Проверка доступности сервера
	// (GET /ping)
	CheckServer(ctx echo.Context) error
	// Получение списка тендеров
	// (GET /tenders)
	GetTenders(ctx echo.Context, params model.GetTendersParams) error
	// Получить тендеры пользователя
	// (GET /tenders/my)
	GetUserTenders(ctx echo.Context, params model.GetUserTendersParams) error
	// Создание нового тендера
	// (POST /tenders/new)
	CreateTender(ctx echo.Context) error
	// Редактирование тендера
	// (PATCH /tenders/{tenderId}/edit)
	EditTender(ctx echo.Context, tenderId model.TenderId, params model.EditTenderParams) error
	// Откат версии тендера
	// (PUT /tenders/{tenderId}/rollback/{version})
	RollbackTender(ctx echo.Context, tenderId model.TenderId, version int32, params model.RollbackTenderParams) error
	// Получение текущего статуса тендера
	// (GET /tenders/{tenderId}/status)
	GetTenderStatus(ctx echo.Context, tenderId model.TenderId, params model.GetTenderStatusParams) error
	// Изменение статуса тендера
	// (PUT /tenders/{tenderId}/status)
	UpdateTenderStatus(ctx echo.Context, tenderId model.TenderId, params model.UpdateTenderStatusParams) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetUserBids converts echo context to params.
func (w *ServerInterfaceWrapper) GetUserBids(ctx echo.Context) error {
	var err error

	ctx.Set(model.BearerAuthScopes, []string{})

	// Parameter object where we will unmarshal all parameters from the context
	var params model.GetUserBidsParams
	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", ctx.QueryParams(), &params.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter limit: %s", err))
	}

	// ------------- Optional query parameter "offset" -------------

	err = runtime.BindQueryParameter("form", true, false, "offset", ctx.QueryParams(), &params.Offset)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter offset: %s", err))
	}

	// ------------- Optional query parameter "username" -------------

	err = runtime.BindQueryParameter("form", true, false, "username", ctx.QueryParams(), &params.Username)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter username: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetUserBids(ctx, params)
	return err
}

// CreateBid converts echo context to params.
func (w *ServerInterfaceWrapper) CreateBid(ctx echo.Context) error {
	var err error
	var bid model.Bid
	ctx.Set(model.BearerAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.CreateBid(ctx, bid)
	return err
}

// EditBid converts echo context to params.
func (w *ServerInterfaceWrapper) EditBid(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "bidId" -------------
	var bidId model.BidId

	err = runtime.BindStyledParameterWithOptions("simple", "bidId", ctx.Param("bidId"), &bidId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter bidId: %s", err))
	}

	ctx.Set(model.BearerAuthScopes, []string{})

	// Parameter object where we will unmarshal all parameters from the context
	var params model.EditBidParams
	// ------------- Required query parameter "username" -------------

	err = runtime.BindQueryParameter("form", true, true, "username", ctx.QueryParams(), &params.Username)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter username: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.EditBid(ctx, bidId, params)
	return err
}

// SubmitBidFeedback converts echo context to params.
func (w *ServerInterfaceWrapper) SubmitBidFeedback(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "bidId" -------------
	var bidId model.BidId

	err = runtime.BindStyledParameterWithOptions("simple", "bidId", ctx.Param("bidId"), &bidId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter bidId: %s", err))
	}

	ctx.Set(model.BearerAuthScopes, []string{})

	// Parameter object where we will unmarshal all parameters from the context
	var params model.SubmitBidFeedbackParams
	// ------------- Required query parameter "bidFeedback" -------------

	err = runtime.BindQueryParameter("form", true, true, "bidFeedback", ctx.QueryParams(), &params.BidFeedback)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter bidFeedback: %s", err))
	}

	// ------------- Required query parameter "username" -------------

	err = runtime.BindQueryParameter("form", true, true, "username", ctx.QueryParams(), &params.Username)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter username: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.SubmitBidFeedback(ctx, bidId, params)
	return err
}

// RollbackBid converts echo context to params.
func (w *ServerInterfaceWrapper) RollbackBid(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "bidId" -------------
	var bidId model.BidId

	err = runtime.BindStyledParameterWithOptions("simple", "bidId", ctx.Param("bidId"), &bidId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter bidId: %s", err))
	}

	// ------------- Path parameter "version" -------------
	var version int32

	err = runtime.BindStyledParameterWithOptions("simple", "version", ctx.Param("version"), &version, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter version: %s", err))
	}

	ctx.Set(model.BearerAuthScopes, []string{})

	// Parameter object where we will unmarshal all parameters from the context
	var params model.RollbackBidParams
	// ------------- Required query parameter "username" -------------

	err = runtime.BindQueryParameter("form", true, true, "username", ctx.QueryParams(), &params.Username)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter username: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.RollbackBid(ctx, bidId, version, params)
	return err
}

// GetBidStatus converts echo context to params.
func (w *ServerInterfaceWrapper) GetBidStatus(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "bidId" -------------
	var bidId model.BidId

	err = runtime.BindStyledParameterWithOptions("simple", "bidId", ctx.Param("bidId"), &bidId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter bidId: %s", err))
	}

	ctx.Set(model.BearerAuthScopes, []string{})

	// Parameter object where we will unmarshal all parameters from the context
	var params model.GetBidStatusParams
	// ------------- Required query parameter "username" -------------

	err = runtime.BindQueryParameter("form", true, true, "username", ctx.QueryParams(), &params.Username)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter username: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetBidStatus(ctx, bidId, params)
	return err
}

// UpdateBidStatus converts echo context to params.
func (w *ServerInterfaceWrapper) UpdateBidStatus(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "bidId" -------------
	var bidId model.BidId

	err = runtime.BindStyledParameterWithOptions("simple", "bidId", ctx.Param("bidId"), &bidId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter bidId: %s", err))
	}

	ctx.Set(model.BearerAuthScopes, []string{})

	// Parameter object where we will unmarshal all parameters from the context
	var params model.UpdateBidStatusParams
	// ------------- Required query parameter "status" -------------

	err = runtime.BindQueryParameter("form", true, true, "status", ctx.QueryParams(), &params.Status)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter status: %s", err))
	}

	// ------------- Required query parameter "username" -------------

	err = runtime.BindQueryParameter("form", true, true, "username", ctx.QueryParams(), &params.Username)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter username: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.UpdateBidStatus(ctx, bidId, params)
	return err
}

// SubmitBidDecision converts echo context to params.
func (w *ServerInterfaceWrapper) SubmitBidDecision(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "bidId" -------------
	var bidId model.BidId

	err = runtime.BindStyledParameterWithOptions("simple", "bidId", ctx.Param("bidId"), &bidId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter bidId: %s", err))
	}

	ctx.Set(model.BearerAuthScopes, []string{})

	// Parameter object where we will unmarshal all parameters from the context
	var params model.SubmitBidDecisionParams
	// ------------- Required query parameter "decision" -------------

	err = runtime.BindQueryParameter("form", true, true, "decision", ctx.QueryParams(), &params.Decision)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter decision: %s", err))
	}

	// ------------- Required query parameter "username" -------------

	err = runtime.BindQueryParameter("form", true, true, "username", ctx.QueryParams(), &params.Username)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter username: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.SubmitBidDecision(ctx, bidId, params)
	return err
}

// GetBidsForTender converts echo context to params.
func (w *ServerInterfaceWrapper) GetBidsForTender(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "tenderId" -------------
	var tenderId model.TenderId

	err = runtime.BindStyledParameterWithOptions("simple", "tenderId", ctx.Param("tenderId"), &tenderId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter tenderId: %s", err))
	}

	ctx.Set(model.BearerAuthScopes, []string{})

	// Parameter object where we will unmarshal all parameters from the context
	var params model.GetBidsForTenderParams
	// ------------- Required query parameter "username" -------------

	err = runtime.BindQueryParameter("form", true, true, "username", ctx.QueryParams(), &params.Username)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter username: %s", err))
	}

	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", ctx.QueryParams(), &params.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter limit: %s", err))
	}

	// ------------- Optional query parameter "offset" -------------

	err = runtime.BindQueryParameter("form", true, false, "offset", ctx.QueryParams(), &params.Offset)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter offset: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetBidsForTender(ctx, tenderId, params)
	return err
}

// GetBidReviews converts echo context to params.
func (w *ServerInterfaceWrapper) GetBidReviews(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "tenderId" -------------
	var tenderId model.TenderId

	err = runtime.BindStyledParameterWithOptions("simple", "tenderId", ctx.Param("tenderId"), &tenderId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter tenderId: %s", err))
	}

	ctx.Set(model.BearerAuthScopes, []string{})

	// Parameter object where we will unmarshal all parameters from the context
	var params model.GetBidReviewsParams
	// ------------- Required query parameter "authorUsername" -------------

	err = runtime.BindQueryParameter("form", true, true, "authorUsername", ctx.QueryParams(), &params.AuthorUsername)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter authorUsername: %s", err))
	}

	// ------------- Required query parameter "requesterUsername" -------------

	err = runtime.BindQueryParameter("form", true, true, "requesterUsername", ctx.QueryParams(), &params.RequesterUsername)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter requesterUsername: %s", err))
	}

	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", ctx.QueryParams(), &params.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter limit: %s", err))
	}

	// ------------- Optional query parameter "offset" -------------

	err = runtime.BindQueryParameter("form", true, false, "offset", ctx.QueryParams(), &params.Offset)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter offset: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetBidReviews(ctx, tenderId, params)
	return err
}

// CheckServer converts echo context to params.
func (w *ServerInterfaceWrapper) CheckServer(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.CheckServer(ctx)
	return err
}

// GetTenders converts echo context to params.
func (w *ServerInterfaceWrapper) GetTenders(ctx echo.Context) error {
	var err error

	ctx.Set(model.BearerAuthScopes, []string{})

	// Parameter object where we will unmarshal all parameters from the context
	var params model.GetTendersParams
	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", ctx.QueryParams(), &params.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter limit: %s", err))
	}

	// ------------- Optional query parameter "offset" -------------

	err = runtime.BindQueryParameter("form", true, false, "offset", ctx.QueryParams(), &params.Offset)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter offset: %s", err))
	}

	// ------------- Optional query parameter "service_type" -------------

	err = runtime.BindQueryParameter("form", true, false, "service_type", ctx.QueryParams(), &params.ServiceType)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter service_type: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetTenders(ctx, params)
	return err
}

// GetUserTenders converts echo context to params.
func (w *ServerInterfaceWrapper) GetUserTenders(ctx echo.Context) error {
	var err error

	ctx.Set(model.BearerAuthScopes, []string{})

	// Parameter object where we will unmarshal all parameters from the context
	var params model.GetUserTendersParams
	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", ctx.QueryParams(), &params.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter limit: %s", err))
	}

	// ------------- Optional query parameter "offset" -------------

	err = runtime.BindQueryParameter("form", true, false, "offset", ctx.QueryParams(), &params.Offset)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter offset: %s", err))
	}

	// ------------- Optional query parameter "username" -------------

	err = runtime.BindQueryParameter("form", true, false, "username", ctx.QueryParams(), &params.Username)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter username: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetUserTenders(ctx, params)
	return err
}

// CreateTender converts echo context to params.
func (w *ServerInterfaceWrapper) CreateTender(ctx echo.Context) error {
	var err error

	ctx.Set(model.BearerAuthScopes, []string{})

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.CreateTender(ctx)
	return err
}

// EditTender converts echo context to params.
func (w *ServerInterfaceWrapper) EditTender(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "tenderId" -------------
	var tenderId model.TenderId

	err = runtime.BindStyledParameterWithOptions("simple", "tenderId", ctx.Param("tenderId"), &tenderId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter tenderId: %s", err))
	}

	ctx.Set(model.BearerAuthScopes, []string{})

	// Parameter object where we will unmarshal all parameters from the context
	var params model.EditTenderParams
	// ------------- Required query parameter "username" -------------

	err = runtime.BindQueryParameter("form", true, true, "username", ctx.QueryParams(), &params.Username)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter username: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.EditTender(ctx, tenderId, params)
	return err
}

// RollbackTender converts echo context to params.
func (w *ServerInterfaceWrapper) RollbackTender(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "tenderId" -------------
	var tenderId model.TenderId

	err = runtime.BindStyledParameterWithOptions("simple", "tenderId", ctx.Param("tenderId"), &tenderId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter tenderId: %s", err))
	}

	// ------------- Path parameter "version" -------------
	var version int32

	err = runtime.BindStyledParameterWithOptions("simple", "version", ctx.Param("version"), &version, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter version: %s", err))
	}

	ctx.Set(model.BearerAuthScopes, []string{})

	// Parameter object where we will unmarshal all parameters from the context
	var params model.RollbackTenderParams
	// ------------- Required query parameter "username" -------------

	err = runtime.BindQueryParameter("form", true, true, "username", ctx.QueryParams(), &params.Username)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter username: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.RollbackTender(ctx, tenderId, version, params)
	return err
}

// GetTenderStatus converts echo context to params.
func (w *ServerInterfaceWrapper) GetTenderStatus(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "tenderId" -------------
	var tenderId model.TenderId

	err = runtime.BindStyledParameterWithOptions("simple", "tenderId", ctx.Param("tenderId"), &tenderId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter tenderId: %s", err))
	}

	ctx.Set(model.BearerAuthScopes, []string{})

	// Parameter object where we will unmarshal all parameters from the context
	var params model.GetTenderStatusParams
	// ------------- Optional query parameter "username" -------------

	err = runtime.BindQueryParameter("form", true, false, "username", ctx.QueryParams(), &params.Username)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter username: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetTenderStatus(ctx, tenderId, params)
	return err
}

// UpdateTenderStatus converts echo context to params.
func (w *ServerInterfaceWrapper) UpdateTenderStatus(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "tenderId" -------------
	var tenderId model.TenderId

	err = runtime.BindStyledParameterWithOptions("simple", "tenderId", ctx.Param("tenderId"), &tenderId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter tenderId: %s", err))
	}

	ctx.Set(model.BearerAuthScopes, []string{})

	// Parameter object where we will unmarshal all parameters from the context
	var params model.UpdateTenderStatusParams
	// ------------- Required query parameter "status" -------------

	err = runtime.BindQueryParameter("form", true, true, "status", ctx.QueryParams(), &params.Status)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter status: %s", err))
	}

	// ------------- Required query parameter "username" -------------

	err = runtime.BindQueryParameter("form", true, true, "username", ctx.QueryParams(), &params.Username)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter username: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.UpdateTenderStatus(ctx, tenderId, params)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/bids/my", wrapper.GetUserBids)
	router.POST(baseURL+"/bids/new", wrapper.CreateBid)
	router.PATCH(baseURL+"/bids/:bidId/edit", wrapper.EditBid)
	router.PUT(baseURL+"/bids/:bidId/feedback", wrapper.SubmitBidFeedback)
	router.PUT(baseURL+"/bids/:bidId/rollback/:version", wrapper.RollbackBid)
	router.GET(baseURL+"/bids/:bidId/status", wrapper.GetBidStatus)
	router.PUT(baseURL+"/bids/:bidId/status", wrapper.UpdateBidStatus)
	router.PUT(baseURL+"/bids/:bidId/submit_decision", wrapper.SubmitBidDecision)
	router.GET(baseURL+"/bids/:tenderId/list", wrapper.GetBidsForTender)
	router.GET(baseURL+"/bids/:tenderId/reviews", wrapper.GetBidReviews)
	router.GET(baseURL+"/ping", wrapper.CheckServer)
	router.GET(baseURL+"/tenders", wrapper.GetTenders)
	router.GET(baseURL+"/tenders/my", wrapper.GetUserTenders)
	router.POST(baseURL+"/tenders/new", wrapper.CreateTender)
	router.PATCH(baseURL+"/tenders/:tenderId/edit", wrapper.EditTender)
	router.PUT(baseURL+"/tenders/:tenderId/rollback/:version", wrapper.RollbackTender)
	router.GET(baseURL+"/tenders/:tenderId/status", wrapper.GetTenderStatus)
	router.PUT(baseURL+"/tenders/:tenderId/status", wrapper.UpdateTenderStatus)

}

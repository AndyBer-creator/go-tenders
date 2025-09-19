package model

const (
	BearerAuthScopes = "bearerAuth.Scopes"
)

// Defines values for BidAuthorType.
const (
	Organization BidAuthorType = "Organization"
	User         BidAuthorType = "User"
)

// Defines values for BidDecision.
const (
	BidDecisionApproved BidDecision = "Approved"
	BidDecisionRejected BidDecision = "Rejected"
)

// Defines values for BidStatus.
const (
	BidStatusApproved  BidStatus = "Approved"
	BidStatusCanceled  BidStatus = "Canceled"
	BidStatusCreated   BidStatus = "Created"
	BidStatusPublished BidStatus = "Published"
	BidStatusRejected  BidStatus = "Rejected"
)

// Defines values for TenderServiceType.
const (
	Construction TenderServiceType = "Construction"
	Delivery     TenderServiceType = "Delivery"
	Manufacture  TenderServiceType = "Manufacture"
)

// Defines values for TenderStatus.
const (
	Closed    TenderStatus = "Closed"
	Created   TenderStatus = "Created"
	Published TenderStatus = "Published"
)

// Bid Информация о предложении
type Bid struct {
	// AuthorId Уникальный идентификатор автора предложения, присвоенный сервером.
	AuthorId BidAuthorId `json:"authorId"`

	// AuthorType Тип автора
	AuthorType BidAuthorType `json:"authorType"`

	// CreatedAt Серверная дата и время в момент, когда пользователь отправил предложение на создание.
	// Передается в формате RFC3339.
	CreatedAt string `json:"createdAt"`

	// Description Описание предложения
	Description BidDescription `json:"description"`

	// Id Уникальный идентификатор предложения, присвоенный сервером.
	Id BidId `json:"id"`

	// Name Полное название предложения
	Name BidName `json:"name"`

	// Status Статус предложения
	Status BidStatus `json:"status"`

	// TenderId Уникальный идентификатор тендера, присвоенный сервером.
	TenderId TenderId `json:"tenderId"`

	// Version Номер версии посел правок
	Version BidVersion `json:"version"`
}

// BidAuthorId Уникальный идентификатор автора предложения, присвоенный сервером.
type BidAuthorId = string

// BidAuthorType Тип автора
type BidAuthorType string

// BidDecision Решение по предложению
type BidDecision string

// BidDescription Описание предложения
type BidDescription = string

// BidFeedback Отзыв на предложение
type BidFeedback = string

// BidId Уникальный идентификатор предложения, присвоенный сервером.
type BidId = string

// BidIdEditBody defines model for bidId_edit_body.
type BidIdEditBody struct {
	// Description Описание предложения
	Description *BidDescription `json:"description,omitempty"`

	// Name Полное название предложения
	Name *BidName `json:"name,omitempty"`
}

// BidName Полное название предложения
type BidName = string

// BidReview Отзыв о предложении
type BidReview struct {
	// CreatedAt Серверная дата и время в момент, когда пользователь отправил отзыв на предложение.
	// Передается в формате RFC3339.
	CreatedAt string `json:"createdAt"`

	// Description Описание предложения
	Description BidReviewDescription `json:"description"`

	// Id Уникальный идентификатор отзыва, присвоенный сервером.
	Id BidReviewId `json:"id"`
}

// BidReviewDescription Описание предложения
type BidReviewDescription = string

// BidReviewId Уникальный идентификатор отзыва, присвоенный сервером.
type BidReviewId = string

// BidStatus Статус предложения
type BidStatus string

// BidVersion Номер версии посел правок
type BidVersion = int32

// BidsNewBody defines model for bids_new_body.
type BidsNewBody struct {
	// CreatorUsername Уникальный slug пользователя.
	CreatorUsername Username `json:"creatorUsername"`

	// Description Описание предложения
	Description BidDescription `json:"description"`

	// Name Полное название предложения
	Name BidName `json:"name"`

	// OrganizationId Уникальный идентификатор организации, присвоенный сервером.
	OrganizationId OrganizationId `json:"organizationId"`

	// Status Статус предложения
	Status BidStatus `json:"status"`

	// TenderId Уникальный идентификатор тендера, присвоенный сервером.
	TenderId TenderId `json:"tenderId"`
}

// ErrorResponse Используется для возвращения ошибки пользователю
type ErrorResponse struct {
	// Reason Описание ошибки в свободной форме
	Reason string `json:"reason"`
}

// OrganizationId Уникальный идентификатор организации, присвоенный сервером.
type OrganizationId = string

// Tender Информация о тендере
type Tender struct {
	// CreatedAt Серверная дата и время в момент, когда пользователь отправил тендер на создание.
	// Передается в формате RFC3339.
	CreatedAt string `json:"createdAt"`

	// Description Описание тендера
	Description TenderDescription `json:"description"`

	// Id Уникальный идентификатор тендера, присвоенный сервером.
	Id TenderId `json:"id"`

	// Name Полное название тендера
	Name TenderName `json:"name"`

	// OrganizationId Уникальный идентификатор организации, присвоенный сервером.
	OrganizationId OrganizationId `json:"organizationId"`

	// ServiceType Вид услуги, к которой относиться тендер
	ServiceType TenderServiceType `json:"serviceType"`

	// Status Статус тендер
	Status TenderStatus `json:"status"`

	// Version Номер версии посел правок
	Version TenderVersion `json:"version"`
}

// TenderDescription Описание тендера
type TenderDescription = string

// TenderId Уникальный идентификатор тендера, присвоенный сервером.
type TenderId = string

// TenderIdEditBody defines model for tenderId_edit_body.
type TenderIdEditBody struct {
	// Description Описание тендера
	Description *TenderDescription `json:"description,omitempty"`

	// Name Полное название тендера
	Name *TenderName `json:"name,omitempty"`

	// ServiceType Вид услуги, к которой относиться тендер
	ServiceType *TenderServiceType `json:"serviceType,omitempty"`
}

// TenderName Полное название тендера
type TenderName = string

// TenderServiceType Вид услуги, к которой относиться тендер
type TenderServiceType string

// TenderStatus Статус тендер
type TenderStatus string

// TenderVersion Номер версии посел правок
type TenderVersion = int32

// TendersNewBody defines model for tenders_new_body.
type TendersNewBody struct {
	// CreatorUsername Уникальный slug пользователя.
	CreatorUsername Username `json:"creatorUsername"`

	// Description Описание тендера
	Description TenderDescription `json:"description"`

	// Name Полное название тендера
	Name TenderName `json:"name"`

	// OrganizationId Уникальный идентификатор организации, присвоенный сервером.
	OrganizationId OrganizationId `json:"organizationId"`

	// ServiceType Вид услуги, к которой относиться тендер
	ServiceType TenderServiceType `json:"serviceType"`

	// Status Статус тендер
	Status TenderStatus `json:"status"`
}

// Username Уникальный slug пользователя.
type Username = string

// GetUserBidsParams defines parameters for GetUserBids.
type GetUserBidsParams struct {
	// Limit Максимальное число возвращаемых объектов. Используется для запросов с пагинацией.
	//
	// Сервер должен возвращать максимальное допустимое число объектов.
	Limit *int32 `form:"limit,omitempty" json:"limit,omitempty"`

	// Offset Какое количество объектов должно быть пропущено с начала. Используется для запросов с пагинацией.
	Offset   *int32    `form:"offset,omitempty" json:"offset,omitempty"`
	Username *Username `form:"username,omitempty" json:"username,omitempty"`
}

// EditBidParams defines parameters for EditBid.
type EditBidParams struct {
	Username Username `form:"username" json:"username"`
}

// SubmitBidFeedbackParams defines parameters for SubmitBidFeedback.
type SubmitBidFeedbackParams struct {
	BidFeedback BidFeedback `form:"bidFeedback" json:"bidFeedback"`
	Username    Username    `form:"username" json:"username"`
}

// RollbackBidParams defines parameters for RollbackBid.
type RollbackBidParams struct {
	Username Username `form:"username" json:"username"`
}

// GetBidStatusParams defines parameters for GetBidStatus.
type GetBidStatusParams struct {
	Username Username `form:"username" json:"username"`
}

// UpdateBidStatusParams defines parameters for UpdateBidStatus.
type UpdateBidStatusParams struct {
	Status   BidStatus `form:"status" json:"status"`
	Username Username  `form:"username" json:"username"`
}

// SubmitBidDecisionParams defines parameters for SubmitBidDecision.
type SubmitBidDecisionParams struct {
	Decision BidDecision `form:"decision" json:"decision"`
	Username Username    `form:"username" json:"username"`
}

// GetBidsForTenderParams defines parameters for GetBidsForTender.
type GetBidsForTenderParams struct {
	Username Username `form:"username" json:"username"`

	// Limit Максимальное число возвращаемых объектов. Используется для запросов с пагинацией.
	//
	// Сервер должен возвращать максимальное допустимое число объектов.
	Limit *int32 `form:"limit,omitempty" json:"limit,omitempty"`

	// Offset Какое количество объектов должно быть пропущено с начала. Используется для запросов с пагинацией.
	Offset *int32 `form:"offset,omitempty" json:"offset,omitempty"`
}

// GetBidReviewsParams defines parameters for GetBidReviews.
type GetBidReviewsParams struct {
	// AuthorUsername Имя пользователя автора предложений, отзывы на которые нужно просмотреть.
	AuthorUsername Username `form:"authorUsername" json:"authorUsername"`

	// RequesterUsername Имя пользователя, который запрашивает отзывы.
	RequesterUsername Username `form:"requesterUsername" json:"requesterUsername"`

	// Limit Максимальное число возвращаемых объектов. Используется для запросов с пагинацией.
	//
	// Сервер должен возвращать максимальное допустимое число объектов.
	Limit *int32 `form:"limit,omitempty" json:"limit,omitempty"`

	// Offset Какое количество объектов должно быть пропущено с начала. Используется для запросов с пагинацией.
	Offset *int32 `form:"offset,omitempty" json:"offset,omitempty"`
}

// GetTendersParams defines parameters for GetTenders.
type GetTendersParams struct {
	// Limit Максимальное число возвращаемых объектов. Используется для запросов с пагинацией.
	//
	// Сервер должен возвращать максимальное допустимое число объектов.
	Limit *int32 `form:"limit,omitempty" json:"limit,omitempty"`

	// Offset Какое количество объектов должно быть пропущено с начала. Используется для запросов с пагинацией.
	Offset *int32 `form:"offset,omitempty" json:"offset,omitempty"`

	// ServiceType Возвращенные тендеры должны соответствовать указанным видам услуг.
	//
	// Если список пустой, фильтры не применяются.
	ServiceType *[]TenderServiceType `form:"service_type,omitempty" json:"service_type,omitempty"`
}

// GetUserTendersParams defines parameters for GetUserTenders.
type GetUserTendersParams struct {
	// Limit Максимальное число возвращаемых объектов. Используется для запросов с пагинацией.
	//
	// Сервер должен возвращать максимальное допустимое число объектов.
	Limit *int32 `form:"limit,omitempty" json:"limit,omitempty"`

	// Offset Какое количество объектов должно быть пропущено с начала. Используется для запросов с пагинацией.
	Offset   *int32    `form:"offset,omitempty" json:"offset,omitempty"`
	Username *Username `form:"username,omitempty" json:"username,omitempty"`
}

// EditTenderParams defines parameters for EditTender.
type EditTenderParams struct {
	Username    Username `form:"username" json:"username"`
	Description *string  `form:"description,omitempty" json:"description,omitempty"`
}

// RollbackTenderParams defines parameters for RollbackTender.
type RollbackTenderParams struct {
	Username Username `form:"username" json:"username"`
}

// GetTenderStatusParams defines parameters for GetTenderStatus.
type GetTenderStatusParams struct {
	Username *Username    `form:"username,omitempty" json:"username,omitempty"`
	Status   TenderStatus `form:"status" json:"status"`
}

// UpdateTenderStatusParams defines parameters for UpdateTenderStatus.
type UpdateTenderStatusParams struct {
	Status   TenderStatus `form:"status" json:"status"`
	Username Username     `form:"username" json:"username"`
}

// CreateBidJSONRequestBody defines body for CreateBid for application/json ContentType.
type CreateBidJSONRequestBody = BidsNewBody

// EditBidJSONRequestBody defines body for EditBid for application/json ContentType.
type EditBidJSONRequestBody = BidIdEditBody

// CreateTenderJSONRequestBody defines body for CreateTender for application/json ContentType.
type CreateTenderJSONRequestBody = TendersNewBody

// EditTenderJSONRequestBody defines body for EditTender for application/json ContentType.
type EditTenderJSONRequestBody = TenderIdEditBody

package models

import "errors"

// maybe better place for it?
var (
	ErrNoRecord           = errors.New("models: no matching record found")
	ErrInvalidCredentials = errors.New("models: invalid credentials")
	ErrDuplicateEmail     = errors.New("models: duplicate email")
	ErrDuplicateUsername  = errors.New("models: duplicate username")
	ErrTooManySpaces      = errors.New("inupt data: too many spaces in field")
	ErrUserNotActivated   = errors.New("models: user is not activated")
	ErrUnauthorized       = errors.New("models: user has no rights to enter this page")

	ErrInternalServerError = errors.New("server: internal server error")
)

type User struct {
	Id             int    `json:"id,omitempty"`
	FirstName      string `json:"firstname,omitempty"`
	LastName       string `json:"lastname,omitempty"`
	Email          string `json:"email,omitempty"`
	Password       string `json:"password,omitempty"`
	HashedPassword []byte `json:"-"`
	Activated      *int   `json:"activated,omitempty"`
	// IsAdmin        int    `json:"is_admin,omitempty"`
	// UserPercent *float64 `json:"user_percent,omitempty"`
	DateCreated int `json:"date_created,omitempty"`

	// date of creation do we need ?
}

type Session struct {
	Id   string
	User *User
}

// third party API response struct
type ApiResponse struct {
	Prices []*OnePrice
}
type OnePrice struct {
	Price            float64 `json:"price,omitempty"`
	Article          string  `json:"article,omitempty"`
	Supplier         string  `json:"supplier,omitempty"`
	SupplierPriceNum float64 `json:"supplier_price_num,omitempty"`
	Brand            string  `json:"brand,omitempty"`
	Currency         string  `json:"currency,omitempty"`
	CurrencyRate     string  `json:"currency_rate,omitempty"`
	Delivery         string  `json:"delivery,omitempty"`
	Weight           float64 `json:"weight,omitempty"`
}

func ChangePrice(op *OnePrice, percent float64) {
	op.Price *= percent
}

// shopping cart and products in there
type Product struct {
	PositionId           string  `json:"position_id,omitempty"`
	Price                float64 `json:"price,omitempty"`
	Article              string  `json:"article,omitempty"`
	Supplier             string  `json:"supplier,omitempty"`
	SupplierPriceNum     float64 `json:"supplier_price_num,omitempty"`
	Brand                string  `json:"brand,omitempty"`
	Currency             string  `json:"currency,omitempty"`
	CurrencyRate         string  `json:"currency_rate,omitempty"`
	Delivery             string  `json:"delivery,omitempty"`
	Weight               float64 `json:"weight,omitempty"`
	Quantity             int     `json:"quantity"` // should not be empty; quantity should = 1 by default
	ProductQuantityPrice float64 `json:"product_quantity_price,omitempty"`
}

type ShoppingCart struct {
	OrderId    string     `json:"order_id"`
	TotalPrice float64    `json:"total_price"`
	Products   []*Product `json:"products"`
}

func CreateShoppingCart(sc *ShoppingCart, orderId string) {
	sc.OrderId = orderId
	for _, prod := range sc.Products {
		sc.TotalPrice += prod.ProductQuantityPrice
	}
}

// POST /search
type Article struct {
	Article string `json:"article"`
}

// POST /cart/removeitem
type PositionId struct {
	PositionId string `json:"position_id"`
}

// POST /cart/remove
type OrderId struct {
	OrderId string `json:"order_id"`
}

// GET /myorders
type UserOrders struct {
	Orders []string `json:"orders"`
}

// POST /order
type UserOrder struct {
	OrderId     string     `json:"order_id"`
	DateCreated int        `json:"date_created"`
	Positions   []*Product `json:"positions"`
}

func CollectUserOrder(order *UserOrder, orderId string) {
	order.OrderId = orderId
}

package models

type Employee struct {
	ID       int    `json:"id"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password"`
	Balance  int    `json:"balance"`
}

type Merch struct {
	ID    int    `json:"id"`
	Type  string `json:"type"`
	Price int    `json:"price"`
}

type AuthRequest struct {
	Username string `json:"username" validate:"required,min=3,max=20"`
	Password string `json:"password" validate:"required,min=6"`
}

type AuthResponse struct {
	Token string `json:"token"`
}

type ErrorResponse struct {
	Errors string `json:"errors"`
}

type SentCoinRequest struct {
	ToUser string `json:"toUser" validate:"required"`
	Amount int    `json:"amount" validate:"required"`
}

type InfoResponse struct {
	Coins       int         `json:"coins"`
	Inventory   []Item      `json:"item"`
	CoinHistory CoinHistory `json:"coinHistory"`
}

type Item struct {
	Type     string `json:"type"`
	Quantity int    `json:"quantity"`
}

type CoinHistory struct {
	Received []Received `json:"received"`
	Sent     []Sent     `json:"sent"`
}

type Received struct {
	FromUser string `json:"fromUser"`
	Amount   int    `json:"amount"`
}

type Sent struct {
	ToUser string `json:"toUser"`
	Amount int    `json:"amount"`
}

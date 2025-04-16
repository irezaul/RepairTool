package main

import "time"

type Client struct {
    ID        int       `json:"id"`
    Name      string    `json:"name"`
    Email     string    `json:"email"`
    Phone     string    `json:"phone"`
    Address   string    `json:"address"`
    CreatedAt time.Time `json:"createdAt"`
    UpdatedAt time.Time `json:"updatedAt"`
}

type Product struct {
    ID          int       `json:"id"`
    ClientID    int       `json:"clientId"`
    Name        string    `json:"name"`
    Serial      string    `json:"serial"`
    Description string    `json:"description"`
    CreatedAt   time.Time `json:"createdAt"`
    UpdatedAt   time.Time `json:"updatedAt"`
}

type Payment struct {
    ID          int       `json:"id"`
    ClientID    int       `json:"clientId"`
    Amount      float64   `json:"amount"`
    DueAmount   float64   `json:"dueAmount"`
    PaymentDate time.Time `json:"paymentDate"`
    Notes       string    `json:"notes"`
    CreatedAt   time.Time `json:"createdAt"`
    UpdatedAt   time.Time `json:"updatedAt"`
}
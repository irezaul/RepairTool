package main

import (
    "context"
    
)

type App struct {
    ctx context.Context
    db  *DB
}

func NewApp() *App {
    return &App{}
}

func (a *App) startup(ctx context.Context) {
    a.ctx = ctx
    db, err := NewDB("clientmanager.db")
    if err != nil {
        panic(err)
    }
    a.db = db
}

// Client methods
func (a *App) AddClient(name, email, phone, address string) (int64, error) {
    client := Client{
        Name:    name,
        Email:   email,
        Phone:   phone,
        Address: address,
    }
    return a.db.AddClient(client)
}

// Add similar methods for products and payments
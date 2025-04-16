package main

import (
    "database/sql"
    "time"
    _ "modernc.org/sqlite"
)

type DB struct {
    db *sql.DB
}

func NewDB(path string) (*DB, error) {
    db, err := sql.Open("sqlite", path)
    if err != nil {
        return nil, err
    }
    
    // Create tables if they don't exist
    queries := []string{
        `CREATE TABLE IF NOT EXISTS clients (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            name TEXT NOT NULL,
            email TEXT,
            phone TEXT,
            address TEXT,
            createdAt DATETIME DEFAULT CURRENT_TIMESTAMP,
            updatedAt DATETIME DEFAULT CURRENT_TIMESTAMP
        )`,
        `CREATE TABLE IF NOT EXISTS products (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            clientId INTEGER NOT NULL,
            name TEXT NOT NULL,
            serial TEXT NOT NULL,
            description TEXT,
            createdAt DATETIME DEFAULT CURRENT_TIMESTAMP,
            updatedAt DATETIME DEFAULT CURRENT_TIMESTAMP,
            FOREIGN KEY(clientId) REFERENCES clients(id)
        )`,
        `CREATE TABLE IF NOT EXISTS payments (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            clientId INTEGER NOT NULL,
            amount REAL NOT NULL,
            dueAmount REAL NOT NULL,
            paymentDate DATETIME NOT NULL,
            notes TEXT,
            createdAt DATETIME DEFAULT CURRENT_TIMESTAMP,
            updatedAt DATETIME DEFAULT CURRENT_TIMESTAMP,
            FOREIGN KEY(clientId) REFERENCES clients(id)
        )`,
    }
    
    for _, query := range queries {
        _, err = db.Exec(query)
        if err != nil {
            return nil, err
        }
    }
    
    return &DB{db: db}, nil
}

// Client operations
func (d *DB) AddClient(client Client) (int64, error) {
    res, err := d.db.Exec(
        "INSERT INTO clients (name, email, phone, address, createdAt, updatedAt) VALUES (?, ?, ?, ?, ?, ?)",
        client.Name, client.Email, client.Phone, client.Address, time.Now(), time.Now(),
    )
    if err != nil {
        return 0, err
    }
    return res.LastInsertId()
}

// Add similar methods for other operations (products, payments, etc.)
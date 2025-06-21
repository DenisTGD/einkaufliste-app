package database

import (
    "database/sql"
    "errors"
)

type User struct {
    UserID       string `json:"user_id"`
    Vorname      string `json:"vorname"`
    Nachname     string `json:"nachname"`
    Email        string `json:"email"`
    PasswordHash string `json:"password_hash"`
}


type UserRepository struct {
    DB *sql.DB
}

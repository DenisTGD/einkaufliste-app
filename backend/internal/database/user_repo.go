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
func (r *UserRepository) FindByEmail(email string) (*User, error) {
    query := `SELECT user_id, email, password_hash, role FROM users WHERE email = $1`

    var user User
    row := r.DB.QueryRow(query, email)
    err := row.Scan(&user.UserID, &user.Email, &user.PasswordHash, &user.Role)

    if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return nil, nil // Kein User gefunden
        }
        return nil, err // Technischer Fehler
    }

    return &user, nil
}

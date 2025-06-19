package dbHelper

import (
	"Todo/database"
	"Todo/models"
	"time"
)

func IsUserExists(email string) (bool, error) {
	SQL := `SELECT count(id) > 0 as is_exist
            FROM users
            WHERE email = TRIM($1)
              AND archived_at IS NULL`

	var check bool
	chkErr := database.Todo.Get(&check, SQL, email)
	return check, chkErr
}

func CreateUser(name, email, password string) error {
	SQL := `INSERT INTO users (name, email, password)
            VALUES (TRIM($1), TRIM($2), $3)`

	_, crtErr := database.Todo.Exec(SQL, name, email, password)
	return crtErr
}

func FindUserByEmail(email string, user *models.User) error {
	err := database.Todo.QueryRow("SELECT * FROM users WHERE email = $1", email).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.ArchivedAt,
	)
	return err
}

func CreateUserSession(userID string) (string, error) {
	var sessionID string
	SQL := `INSERT INTO user_session(user_id, archived_at)
             VALUES ($1,$2) RETURNING id`
	archive_at := time.Now().Add(7 * 24 * time.Hour)
	crtErr := database.Todo.Get(&sessionID, SQL, userID, archive_at)
	return sessionID, crtErr
}

func IsValidSession(token string) (bool, error) {
	var exists bool
	query := `SELECT EXISTS (SELECT 1 FROM user_session WHERE id = $1 AND (archived_at) > $2 )`
	err := database.Todo.QueryRow(query, token, time.Now()).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func LogoutSession(token string) error {
	_, err := database.Todo.Query("UPDATE user_session SET archived_at=$1 WHERE  id=$2", time.Now(), token)
	return err
}

func GetProfileDetails(token string, profile *models.ProfileRequest) error {
	query := `
		SELECT u.id, u.name, u.email
		FROM user_session s
		JOIN users u ON s.user_id = u.id
		WHERE s.id = $1 
		LIMIT 1;
	`
	err := database.Todo.QueryRow(query, token).Scan(
		&profile.ID, &profile.Name, &profile.Email,
	)
	return err
}

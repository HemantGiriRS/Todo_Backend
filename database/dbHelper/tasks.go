package dbHelper

import (
	"Todo/database"
	"Todo/models"
	"github.com/sirupsen/logrus"
	"net/http"
)

func GetUserIdBySession(token string) (string, error) {
	query := `SELECT  user_id FROM user_session WHERE id =$1`
	var userId string
	if err := database.Todo.QueryRow(query, token).Scan(&userId); err != nil {
		logrus.Errorf("Failed in getting user id from session", err, http.StatusInternalServerError)
		return "", err
	}
	return userId, nil
}

func AddTaskByUser(id string, name string, des string) error {
	query := `INSERT INTO todo (user_id, name,description) VALUES ($1, $2,$3)`
	if _, err := database.Todo.Query(query, id, name, des); err != nil {
		logrus.Errorf("Failed in adding task", err, http.StatusInternalServerError)
		return err
	}
	return nil
}

func GetTaskByUserId(userId string, tasks *[]models.GetTask) error {
	query := `SELECT id ,name,description ,is_completed,created_at FROM todo WHERE user_id=$1`
	rows, err := database.Todo.Query(query, userId)
	if err != nil {
		logrus.Errorf("Failed in getting task by user", err, http.StatusInternalServerError)
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var task models.GetTask
		if err := rows.Scan(&task.ID, &task.Name, &task.Description, &task.IsCompleted, &task.CreatedAt); err != nil {
			return err
		}
		*tasks = append(*tasks, task)
	}
	if err := rows.Err(); err != nil {
		return err
	}
	return nil
}

func UpdateTaskById(body models.UpdateTask) error {
	query := `UPDATE todo SET name=$1, description=$2, is_completed=$3 WHERE id=$4`
	if _, err := database.Todo.Exec(query, body.Name, body.Description, body.IsCompleted, body.ID); err != nil {
		logrus.Errorf("Failed in updating data %v", err, http.StatusInternalServerError)
		return err
	}
	return nil
}

func UpdateStatus(taskId string, status bool) error {
	query := `UPDATE todo SET is_completed = $1 WHERE id = $2;`
	if _, err := database.Todo.Exec(query, status, taskId); err != nil {
		logrus.Errorf("Failed in updating task %v", err, http.StatusInternalServerError)
		return err
	}
	return nil
}

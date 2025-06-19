package handlers

import (
	"Todo/database/dbHelper"
	"Todo/models"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
)

func CreateTask(w http.ResponseWriter, r *http.Request) {
	//get session token
	token := r.Header.Get("Session-Token")

	var body models.AddTask
	//assign json
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		logrus.Errorf("Failed in decode json body", err)
		return
	}

	//check token
	if token == "" {
		logrus.Errorf("Failed to get token")
		return
	}
	//check session
	var isExists bool
	isExists, err := dbHelper.IsValidSession(token)
	if err != nil {
		logrus.Println("Failed in validation of session ", err, http.StatusInternalServerError)
		return
	}
	if !isExists {
		logrus.Println("Session does not exist")
		return
	}
	//find user id
	var usrId string
	usrId, _ = dbHelper.GetUserIdBySession(token)
	//create task by user id
	if err := dbHelper.AddTaskByUser(usrId, body.Name, body.Description); err != nil {
		logrus.Println("Failed in adding task", err, http.StatusInternalServerError)
	}

	//response
	json.NewEncoder(w).Encode(map[string]string{"Message": "Task Added Successfully..."})
}

func GetTasks(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Session-Token")
	//check token
	if token == "" {
		logrus.Errorf("Failed to get token")
		return
	}
	//check session
	var isExists bool
	isExists, err := dbHelper.IsValidSession(token)
	if err != nil {
		logrus.Println("Failed in validation of session ", err, http.StatusInternalServerError)
		return
	}
	if !isExists {
		logrus.Println("Session does not exist")
		return
	}
	//find user id
	var usrId string
	usrId, _ = dbHelper.GetUserIdBySession(token)
	//get task by user id
	var tasks []models.GetTask
	if err := dbHelper.GetTaskByUserId(usrId, &tasks); err != nil {
		logrus.Errorf("Failed to fetch tasks", err, http.StatusInternalServerError)
		return
	}
	//response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	//get session token
	token := r.Header.Get("Session-Token")

	var body models.UpdateTask
	//assign json
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		logrus.Errorf("Failed in decode json body", err)
		return
	}
	//check token
	if token == "" {
		logrus.Errorf("Failed to get token")
		return
	}
	//check session
	var isExists bool
	isExists, err := dbHelper.IsValidSession(token)
	if err != nil {
		logrus.Println("Failed in validation of session ", err, http.StatusInternalServerError)
		return
	}
	if !isExists {
		logrus.Println("Session does not exist")
		return
	}
	if err := dbHelper.UpdateTaskById(body); err != nil {
		logrus.Println("Failed in adding task", err, http.StatusInternalServerError)
	}

	//response
	json.NewEncoder(w).Encode(map[string]string{"Message": "Task Updated Successfully..."})
}
func UpdateTaskStatus(w http.ResponseWriter, r *http.Request) {
	//get session token
	token := r.Header.Get("Session-Token")

	var body models.UpdateStatus
	//assign json
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		logrus.Errorf("Failed in decode json body", err)
		return
	}

	//check token
	if token == "" {
		logrus.Errorf("Failed to get token")
		return
	}
	//check session
	var isExists bool
	isExists, err := dbHelper.IsValidSession(token)
	if err != nil {
		logrus.Println("Failed in validation of session ", err, http.StatusInternalServerError)
		return
	}
	if !isExists {
		logrus.Println("Session does not exist")
		return
	}
	if err := dbHelper.UpdateStatus(body.Id, body.IsCompleted); err != nil {
		logrus.Println("Failed in adding task", err, http.StatusInternalServerError)
	}

	//response
	json.NewEncoder(w).Encode(map[string]string{"Message": "Task Status Updated Successfully..."})
}

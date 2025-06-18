package handlers

import (
	"Todo/database/dbHelper"
	"Todo/models"
	"Todo/utils"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var body models.RegisterRequest

	//assign json
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		logrus.Errorf("Failed in decode json body", err)
		return
	}
	//check user's existance
	exists, existsErr := dbHelper.IsUserExists(body.Email)
	if existsErr != nil {
		logrus.Errorf("Failed to check user existence %v %v", existsErr, http.StatusInternalServerError)
		return
	}
	if exists {
		logrus.Errorf("User already exists %v %v", existsErr, http.StatusBadRequest)
		return
	}
	//genrate hash pass
	hashedPassword, hasErr := utils.HashPassword(body.Password)
	if hasErr != nil {
		logrus.Println("Failed to genrate hash password %v %v", hasErr, http.StatusInternalServerError)
		return
	}
	//insert in db
	if saveErr := dbHelper.CreateUser(body.Username, body.Email, hashedPassword); saveErr != nil {
		logrus.Println("Failed to create user %v %v", saveErr, http.StatusInternalServerError)
		return
	}
	//response
	json.NewEncoder(w).Encode(map[string]string{"Message": "User Registered Successfully..."})
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	var body models.LoginRequest
	//assign json body
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		logrus.Errorf("Failed in decode json body", err)
		return
	}
	//find user in db
	var usr models.User
	var err error
	if err := dbHelper.FindUserByEmail(body.Email, &usr); err != nil {
		logrus.Errorf("Failed to find user by email %v %v", err, usr)
		return
	}
	if usr.ID == "" {
		logrus.Println("User does not exist")
		return
	}
	//compare pass
	if err != nil || bcrypt.CompareHashAndPassword([]byte(usr.Password), []byte(body.Password)) != nil {
		logrus.Errorf("Invaid Credentails...", http.StatusUnauthorized)
		return
	}
	//create session
	var sid string
	if sid, err = dbHelper.CreateUserSession(usr.ID); err != nil {
		logrus.Println("Failed to create session", err, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"Message": "User Login Successfully...", "SessionId": sid})
}

func LogoutUser(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Session-Token")
	//check token
	if token == "" {
		logrus.Errorf("Failed to get token")
		return
	}
	//if err:=dbHelper.IsValidSession()
	//log out
	if err := dbHelper.LogoutSession(token); err != nil {
		logrus.Println("Failed to logout session")
		return
	}

	//response
	json.NewEncoder(w).Encode(map[string]string{"Message": "User Logout Successfully..."})
}

func Profile(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Session-Token")
	//check token
	if token == "" {
		logrus.Errorf("Failed to get token")
		return
	}

}

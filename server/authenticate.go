package server

import (
	"chat-application/models"
	"chat-application/utils"
	"encoding/json"
	"errors"
	"net/http"
)


type SignupRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Confirm  string `json:"confirm"`
}

var ( 
	ErrEmailAndPassReq=errors.New("Email and password are Required")
	UserAlreadyExists=errors.New("User already Exists!!!")
	UserNotFound=errors.New("User not found!!")
)

var (
	UserCreated="User Created !!"
	SignInSuccessful="Sign in Successful"
)

func (s *Server) Signup(w http.ResponseWriter,r *http.Request){

	var srb *models.SignUpRequestBody

	// decode the request body into signup struct
	err:= json.NewDecoder(r.Body).Decode(&srb)
	if err!=nil{
		utils.CreateErrorResponse(ErrEmailAndPassReq,http.StatusInternalServerError,w)
		return
	}

	userFound,err:=s.store.GetUserByEmail(srb.Email)

	if(userFound!=nil){
		utils.CreateErrorResponse(UserAlreadyExists,http.StatusOK,w)
		return
	}

	user,err:=s.store.AddUser(srb)

	utils.CreateSuccessResponse(UserCreated,user,http.StatusOK,w)

}

func (s *Server) Signin(w http.ResponseWriter,r *http.Request){

	var srb *models.SignInRequestBody

	err:= json.NewDecoder(r.Body).Decode(&srb)
	if err!=nil{
		utils.CreateErrorResponse(ErrEmailAndPassReq,http.StatusInternalServerError,w)
		return
	}

	userFound,err:=s.store.GetUserByEmail(srb.Email)

	if(userFound==nil){
		utils.CreateErrorResponse(UserNotFound,http.StatusOK,w)
		return
	}else if(err!=nil){
		utils.CreateErrorResponse(err,http.StatusInternalServerError,w)
		return 
	}

	utils.CreateSuccessResponse(SignInSuccessful,userFound,http.StatusOK,w)
	
}
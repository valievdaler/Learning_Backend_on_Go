package main

import (
	//"log"
	"net/http"
	"encoding/json"
	"errors"
)


type api struct{
	addr string
}
var users =[]User{}

func (a *api) getUsersHandler(w http.ResponseWriter, r *http.Request){
	
	w.Header().Set("Content-Type","application/json")
	err:=json.NewEncoder(w).Encode(users)

	if err!=nil{
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return
	}
}




func (a *api) createUsersHandler(w http.ResponseWriter, r *http.Request){


	
	var payroad User
	err:=json.NewDecoder(r.Body).Decode(&payroad)
	if err!=nil{
		http.Error(w,err.Error(),http.StatusBadRequest)
		return
	}

	u:=User{
		FirstName: payroad.FirstName,
		LastName:payroad.LastName,
	}

	users=append(users,u)

	w.WriteHeader(http.StatusCreated)
}

func (a*api) InsertUser(u User) error{
	if u.FirstName==""{
		return errors.New("First Name is required")
	}
	if u.LastName==""{
		return errors.New("Last Name is required")
	}

	for _,user:=range users{
		if user.FirstName==u.FirstName && user.LastName==u.LastName{
			return errors.New("User already exists")
		}
	}
	users=append(users, u)
	return nil
}

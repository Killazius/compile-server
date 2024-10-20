package handlers

import (
	"compile-server/internal/compilation"
	"compile-server/internal/models"
	"encoding/json"
	"log"
	"net/http"
)

func CodeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("new request")
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var userCode models.Code
	err := json.NewDecoder(r.Body).Decode(&userCode)

	if err != nil {
		return
	}

	err = compilation.MakeFile(userCode.Link, userCode.Lang)
	if err != nil {
		return
	}

	switch userCode.Lang {
	case "cpp":
		{
			compilation.MakeCPPfile(userCode.TaskName, "user.cpp")
		}
	case "py":
		{
			compilation.MakePYfile(userCode.TaskName, "user.py")
		}
	}

}

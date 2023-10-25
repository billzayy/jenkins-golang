package controllers

import (
	"database/sql"
	"os"

	"github.com/joho/godotenv"
)

func GetUserController() (string, error) {
	godotenv.Load()

	dbSrc, err := sql.Open("mysql", os.Getenv("DBPATH"))

	if err != nil {
		return "", nil
	}

	dbPrep, err := dbSrc.Prepare("INSERT INTO User (UserName, Password, Email) VALUES (?,?,?)")

	if err != nil {
		return "", nil
	}

	_, err = dbPrep.Exec()

	if err != nil {
		return "", nil
	}

	return "Sign Up Successful", nil
}

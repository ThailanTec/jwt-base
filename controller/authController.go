package controller

import "errors"

func AuthenticateUser(username, password string) (bool, error) {

	if username == "admin" && password == "admin123" {
		return true, nil
	}

	return false, errors.New("usuario errado irmão")
}

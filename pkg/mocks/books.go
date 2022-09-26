package mocks

import "example/restapi/pkg/models"

var Books = []models.Books{{
	ID:     "1",
	Isbn:   "123456",
	Title:  "The last women",
	Author: &models.Author{FullName: "Milad Mizani"},
}}

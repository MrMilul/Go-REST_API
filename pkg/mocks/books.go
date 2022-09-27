package mocks

import (
	"example/restapi/pkg/models"
	"math/rand"
	"strconv"
)

var Books = []models.Books{{
	ID:     strconv.Itoa(rand.Intn(10000)),
	Isbn:   "123456",
	Title:  "The last women",
	Author: &models.Author{FullName: "Milad Mizani"},
}}

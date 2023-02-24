package test

import (
	"fmt"
	"gin-icqqg/utils/jwt"
	"testing"
)

func TestGenerateToken(t *testing.T) {

	token, err := jwt.GenerateToken("xfdada", "xfdadad")
	if err != nil {
		fmt.Println("err")
	}
	fmt.Println(token)
}

func TestParseToken(t *testing.T) {
	_, err := jwt.ParseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhcHBfa2V5IjoiMmRlMzAxODgzZjU3ZTBlZTQ2YWU3NDc1MDJkNjgwNGQiLCJhcHBfc2VjcmV0IjoiMzYxOTBlNjU1ZDAwYzgxN2MwMDMwOGUwZWQzYmQ5ZDUiLCJleHAiOjE2NzcwNDUwNTksImlzcyI6InhmIn0.eieGyhx3zQgNpEpxc_kTwAMbYQ0eQWYVbc-NDVPK0eo")

	if err != nil {
		fmt.Println("err")
	}
	fmt.Println("pass")
}

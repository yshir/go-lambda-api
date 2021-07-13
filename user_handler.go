package go_lambda_api

import (
	"fmt"
	"go_lambda_api/gen/models"
	"go_lambda_api/gen/restapi/operations"

	"github.com/go-openapi/runtime/middleware"
)

func GetUsers(p operations.GetUsersParams) middleware.Responder {
	ctx := p.HTTPRequest.Context()
	users, err := scanUsers(ctx)
	if err != nil {
		return operations.NewGetUsersInternalServerError().WithPayload(&models.Error{
			Message: fmt.Sprintf("scan users error: %v", err),
		})
	}

	var res models.Users
	for _, user := range users {
		user := user
		res = append(res, &models.User{
			UserID: &user.UserID,
			Name:   &user.UserName,
		})
	}
	return operations.NewGetUsersOK().WithPayload(res)
}

package handlers

import (
	"context"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
	"taskServer/internal/models"
	"taskServer/internal/userService"
	"taskServer/internal/web/users"
)

type UserHandler struct {
	service userService.UserService
}

func NewUserHandler(service userService.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) GetUsers(ctx context.Context, _ users.GetUsersRequestObject) (users.GetUsersResponseObject, error) {
	userList, err := h.service.GetAllUsers()
	if err != nil {
		return nil, err
	}
	var response users.GetUsers200JSONResponse
	for _, u := range userList {
		response = append(response, users.User{
			Id:        u.ID,
			Email:     u.Email,
			Password:  u.Password,
			CreatedAt: &u.CreatedAt,
			UpdatedAt: &u.UpdatedAt,
		})
	}
	return response, nil
}

func (h *UserHandler) PostUsers(ctx context.Context, req users.PostUsersRequestObject) (users.PostUsersResponseObject, error) {
	if req.Body == nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Empty Body")
	}

	newUser := &models.User{
		ID:       uuid.New().String(),
		Email:    req.Body.Email,
		Password: req.Body.Password,
	}

	err := h.service.CreateUser(newUser)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Failed to save user")
	}

	response := users.PostUsers201JSONResponse{
		Id:        newUser.ID,
		Email:     newUser.Email,
		Password:  newUser.Password,
		CreatedAt: &newUser.CreatedAt,
		UpdatedAt: &newUser.UpdatedAt,
	}
	return response, nil

}

func (h *UserHandler) PatchUsersId(ctx context.Context, req users.PatchUsersIdRequestObject) (users.PatchUsersIdResponseObject, error) {
	user, err := h.service.GetUserByID(req.Id)
	if err != nil {
		return users.PatchUsersId404Response{}, nil
	}

	if req.Body.Email != nil {
		user.Email = *req.Body.Email
	}

	if req.Body.Password != nil {
		user.Password = *req.Body.Password
	}

	err = h.service.UpdateUser(user)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Failed to update user")
	}

	response := users.PatchUsersId200JSONResponse{
		Id:        user.ID,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: &user.CreatedAt,
		UpdatedAt: &user.UpdatedAt,
	}
	return response, nil
}

func (h *UserHandler) DeleteUsersId(ctx context.Context, req users.DeleteUsersIdRequestObject) (users.DeleteUsersIdResponseObject, error) {
	_, err := h.service.GetUserByID(req.Id)
	if err != nil {
		return users.DeleteUsersId404Response{}, nil
	}

	err = h.service.DeleteUserByID(req.Id)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Failed to delete user")
	}
	return users.DeleteUsersId204Response{}, nil
}

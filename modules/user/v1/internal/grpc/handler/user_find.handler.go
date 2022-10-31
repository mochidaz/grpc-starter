package handler

import (
	"context"
	"github.com/google/uuid"
	"google.golang.org/grpc/status"
	userv1 "grpc-starter/api/user/v1"
	"grpc-starter/common/constant"
	"grpc-starter/common/errors"
	"net/http"
)

func (ah *UserHandler) GetProfile(ctx context.Context, request *userv1.GetProfileRequest) (*userv1.GetProfileResponse, error) {

	id, err := uuid.Parse(request.GetId())

	if err != nil {
		return nil, status.Errorf(
			http.StatusBadRequest,
			"Invalid ID",
		)
	}

	user, err := ah.userFinderSvc.FindByID(ctx, id)

	if err != nil {
		parseError := errors.ParseError(err)
		return nil, status.Errorf(
			parseError.Code,
			parseError.Message,
		)
	}

	return &userv1.GetProfileResponse{
		Code:    http.StatusOK,
		Message: constant.SuccessMessage,
		Data: &userv1.GetProfileResponse_Profile{
			Username:    user.Username.String,
			Email:       user.Email,
			PhoneNumber: user.PhoneNumber.String,
		},
	}, nil
}

func (ah *UserHandler) GetAllUsers(ctx context.Context, request *userv1.GetAllUsersRequest) (*userv1.GetAllUserResponse, error) {
	users, err := ah.userFinderSvc.FindAllUsers(ctx)

	if err != nil {
		parseError := errors.ParseError(err)
		return nil, status.Errorf(
			parseError.Code,
			parseError.Message,
		)
	}

	var response []*userv1.GetAllUserResponse_Profile

	for _, user := range users {
		response = append(response, &userv1.GetAllUserResponse_Profile{
			UserId:      user.ID.String(),
			Username:    user.Username.String,
			Email:       user.Email,
			PhoneNumber: user.PhoneNumber.String,
		})
	}

	return &userv1.GetAllUserResponse{
		Code:    http.StatusOK,
		Message: constant.SuccessMessage,
		Data:    response,
	}, nil
}

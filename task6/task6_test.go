package task6

import (
	"database/sql"
	"errors"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"go_testing/task6/dto"
	mock "go_testing/task6/repo/mock"
	"testing"
)

func TestUserService_FindByID(t *testing.T) {
	type mockBehaviour func(s *mock.MockUserRepository, id int)

	testsCases := []struct {
		Name             string
		InputID          int
		MockInputID      int
		mockBehaviour    mockBehaviour
		ExpectOutputUser dto.User
		ExpectError      error
	}{
		{
			Name:        "Valid_input_exist_user",
			InputID:     1,
			MockInputID: 1,
			mockBehaviour: func(s *mock.MockUserRepository, id int) {
				s.EXPECT().FindByID(id).Return(dto.User{
					ID:   1,
					Name: "Rustam",
				}, nil)
			},
			ExpectOutputUser: dto.User{
				ID:   1,
				Name: "Rustam",
			},
			ExpectError: nil,
		},
		{
			Name:        "Valid_input_no_rows",
			InputID:     10,
			MockInputID: 10,
			mockBehaviour: func(s *mock.MockUserRepository, id int) {
				s.EXPECT().FindByID(id).Return(dto.User{}, sql.ErrNoRows)
			},
			ExpectOutputUser: dto.User{},
			ExpectError:      sql.ErrNoRows,
		},
		{
			Name:        "Valid_input_unknown_err",
			InputID:     10,
			MockInputID: 10,
			mockBehaviour: func(s *mock.MockUserRepository, id int) {
				s.EXPECT().FindByID(id).Return(dto.User{}, errors.New("unknown error"))
			},
			ExpectOutputUser: dto.User{},
			ExpectError:      errors.New("unknown error"),
		},
	}

	for _, tt := range testsCases {
		t.Run(tt.Name, func(t *testing.T) {

			//Init Deps
			c := gomock.NewController(t)
			defer c.Finish()
			userRepo := mock.NewMockUserRepository(c)
			tt.mockBehaviour(userRepo, tt.InputID)

			userService := NewUserService(userRepo)

			// Testing
			readUser, err := userService.FindByID(tt.InputID)
			assert.Equal(t, tt.ExpectOutputUser, readUser)
			assert.Equal(t, tt.ExpectError, err)
		})
	}
}

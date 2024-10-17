package user

import (
	"context"
	"database/sql"
	"launchpad-go-rest/internal/lib/errors"
	mock_utils "launchpad-go-rest/internal/lib/utils/mock"
	"launchpad-go-rest/internal/repository/mock/mock_cache"
	"launchpad-go-rest/internal/repository/mock/mock_user"
	user_types "launchpad-go-rest/pkg/types/user"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"go.uber.org/mock/gomock"
)

func Test_User_UpdateByID(t *testing.T) {
	Convey("Test_User_UpdateByID", t, func() {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		deps := struct {
			mockUser  *mock_user.MockRepository
			mockUtils *mock_utils.MockUtils
			mockCache *mock_cache.MockRepository
		}{
			mockUser:  mock_user.NewMockRepository(ctrl),
			mockUtils: mock_utils.NewMockUtils(ctrl),
			mockCache: mock_cache.NewMockRepository(ctrl),
		}

		svc := New(deps.mockUser, deps.mockUtils, deps.mockCache)

		errMock := errors.New("mock error")

		testCases := []struct {
			desc        string
			req         user_types.UpdateByIDRequest
			mockCalls   func()
			errExpected bool
		}{
			{
				desc:        "Error: Validation error",
				req:         user_types.UpdateByIDRequest{},
				mockCalls:   func() {},
				errExpected: true,
			},
			{
				desc: "Error: Password confirmation mismatch",
				req: user_types.UpdateByIDRequest{
					ID:                   "id",
					Email:                "test@email.com",
					OldPassword:          "OldPassword1",
					Password:             "Password1",
					PasswordConfirmation: "Password2",
				},
				mockCalls:   func() {},
				errExpected: true,
			},
			{
				desc: "Error: User not found",
				req: user_types.UpdateByIDRequest{
					ID:                   "id",
					Email:                "test@email.com",
					OldPassword:          "OldPassword1",
					Password:             "Password1",
					PasswordConfirmation: "Password1",
				},
				mockCalls: func() {
					deps.mockUser.EXPECT().FindByID(gomock.Any(), gomock.Any()).Return(user_types.User{}, sql.ErrNoRows)
				},
				errExpected: true,
			},
			{
				desc: "Error: Find by id",
				req: user_types.UpdateByIDRequest{
					ID:                   "id",
					Email:                "test@email.com",
					OldPassword:          "OldPassword1",
					Password:             "Password1",
					PasswordConfirmation: "Password1",
				},
				mockCalls: func() {
					deps.mockUser.EXPECT().FindByID(gomock.Any(), gomock.Any()).Return(user_types.User{}, errMock)
				},
				errExpected: true,
			},
			{
				desc: "Error: Old password not match",
				req: user_types.UpdateByIDRequest{
					ID:                   "id",
					Email:                "test@email.com",
					OldPassword:          "OldPassword1",
					Password:             "Password1",
					PasswordConfirmation: "Password1",
				},
				mockCalls: func() {
					deps.mockUser.EXPECT().FindByID(gomock.Any(), gomock.Any()).Return(user_types.User{}, nil)
					deps.mockUtils.EXPECT().ComparePassword(gomock.Any(), gomock.Any()).Return(false)
				},
				errExpected: true,
			},
			{
				desc: "Error: Hash password",
				req: user_types.UpdateByIDRequest{
					ID:                   "id",
					Email:                "test@email.com",
					OldPassword:          "OldPassword1",
					Password:             "Password1",
					PasswordConfirmation: "Password1",
				},
				mockCalls: func() {
					deps.mockUser.EXPECT().FindByID(gomock.Any(), gomock.Any()).Return(user_types.User{}, nil)
					deps.mockUtils.EXPECT().ComparePassword(gomock.Any(), gomock.Any()).Return(true)
					deps.mockUtils.EXPECT().HashPassword(gomock.Any()).Return("", errMock)
				},
				errExpected: true,
			},
			{
				desc: "Error: Update user",
				req: user_types.UpdateByIDRequest{
					ID:                   "id",
					Email:                "test@email.com",
					OldPassword:          "OldPassword1",
					Password:             "Password1",
					PasswordConfirmation: "Password1",
				},
				mockCalls: func() {
					deps.mockUser.EXPECT().FindByID(gomock.Any(), gomock.Any()).Return(user_types.User{}, nil)
					deps.mockUtils.EXPECT().ComparePassword(gomock.Any(), gomock.Any()).Return(true)
					deps.mockUtils.EXPECT().HashPassword(gomock.Any()).Return("PasswordHash1", nil)
					deps.mockUser.EXPECT().UpdateByID(gomock.Any(), gomock.Any()).Return(errMock)
				},
				errExpected: true,
			},
			{
				desc: "Success",
				req: user_types.UpdateByIDRequest{
					ID:                   "id",
					Email:                "test@email.com",
					OldPassword:          "OldPassword1",
					Password:             "Password1",
					PasswordConfirmation: "Password1",
				},
				mockCalls: func() {
					deps.mockUser.EXPECT().FindByID(gomock.Any(), gomock.Any()).Return(user_types.User{}, nil)
					deps.mockUtils.EXPECT().ComparePassword(gomock.Any(), gomock.Any()).Return(true)
					deps.mockUtils.EXPECT().HashPassword(gomock.Any()).Return("PasswordHash1", nil)
					deps.mockUser.EXPECT().UpdateByID(gomock.Any(), gomock.Any()).Return(nil)
				},
				errExpected: false,
			},
		}

		for _, testCase := range testCases {
			Convey(testCase.desc, func() {
				testCase.mockCalls()

				err := svc.UpdateByID(context.Background(), testCase.req)
				if testCase.errExpected {
					So(err, ShouldNotBeNil)
				} else {
					So(err, ShouldBeNil)
				}
			})
		}
	})
}

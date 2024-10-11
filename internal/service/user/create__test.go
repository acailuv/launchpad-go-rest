package user

import (
	"context"
	"launchpad-go-rest/internal/lib/errors"
	mock_utils "launchpad-go-rest/internal/lib/utils/mock"
	"launchpad-go-rest/internal/repository/mock/mock_user"
	user_types "launchpad-go-rest/pkg/types/user"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"go.uber.org/mock/gomock"
)

func Test_User_Create(t *testing.T) {
	Convey("Test_User_Create", t, func() {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		deps := struct {
			mockUser  *mock_user.MockRepository
			mockUtils *mock_utils.MockUtils
		}{
			mockUser:  mock_user.NewMockRepository(ctrl),
			mockUtils: mock_utils.NewMockUtils(ctrl),
		}

		svc := New(deps.mockUser, deps.mockUtils)

		errMock := errors.New("mock error")

		testCases := []struct {
			desc        string
			req         user_types.CreateRequest
			mockCalls   func()
			errExpected bool
		}{
			{
				desc:        "Error: Validation error",
				req:         user_types.CreateRequest{},
				mockCalls:   func() {},
				errExpected: true,
			},
			{
				desc: "Error: Password confirmation mismatch",
				req: user_types.CreateRequest{
					Email:                "test@email.com",
					Password:             "Password1",
					PasswordConfirmation: "Password2",
				},
				mockCalls:   func() {},
				errExpected: true,
			},
			{
				desc: "Error: Hash password",
				req: user_types.CreateRequest{
					Email:                "test@email.com",
					Password:             "Password1",
					PasswordConfirmation: "Password1",
				},
				mockCalls: func() {
					deps.mockUtils.EXPECT().HashPassword(gomock.Any()).Return("", errMock)
				},
				errExpected: true,
			},
			{
				desc: "Error: Find duplicate user",
				req: user_types.CreateRequest{
					Email:                "test@email.com",
					Password:             "Password1",
					PasswordConfirmation: "Password1",
				},
				mockCalls: func() {
					deps.mockUtils.EXPECT().HashPassword(gomock.Any()).Return("", nil)
					deps.mockUser.EXPECT().FindByEmail(gomock.Any(), gomock.Any()).Return(user_types.User{}, nil)
				},
				errExpected: true,
			},
			{
				desc: "Error: Create user",
				req: user_types.CreateRequest{
					Email:                "test@email.com",
					Password:             "Password1",
					PasswordConfirmation: "Password1",
				},
				mockCalls: func() {
					deps.mockUtils.EXPECT().HashPassword(gomock.Any()).Return("PasswordHash1", nil)
					deps.mockUser.EXPECT().FindByEmail(gomock.Any(), gomock.Any()).Return(user_types.User{}, errMock)
					deps.mockUser.EXPECT().Create(gomock.Any(), gomock.Any()).Return(errMock)
				},
				errExpected: true,
			},
			{
				desc: "Success",
				req: user_types.CreateRequest{
					Email:                "test@email.com",
					Password:             "Password1",
					PasswordConfirmation: "Password1",
				},
				mockCalls: func() {
					deps.mockUtils.EXPECT().HashPassword(gomock.Any()).Return("PasswordHash1", nil)
					deps.mockUser.EXPECT().FindByEmail(gomock.Any(), gomock.Any()).Return(user_types.User{}, errMock)
					deps.mockUser.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil)
				},
				errExpected: false,
			},
		}

		for _, testCase := range testCases {
			Convey(testCase.desc, func() {
				testCase.mockCalls()

				err := svc.Create(context.Background(), testCase.req)
				if testCase.errExpected {
					So(err, ShouldNotBeNil)
				} else {
					So(err, ShouldBeNil)
				}
			})
		}
	})
}

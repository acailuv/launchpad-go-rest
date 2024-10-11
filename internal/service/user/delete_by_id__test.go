package user

import (
	"context"
	"launchpad-go-rest/internal/lib/errors"
	mock_utils "launchpad-go-rest/internal/lib/utils/mock"
	"launchpad-go-rest/internal/repository/mock/mock_user"
	"launchpad-go-rest/pkg/types/user"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"go.uber.org/mock/gomock"
)

func Test_User_DeleteByID(t *testing.T) {
	Convey("Test_User_DeleteByID", t, func() {
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
			req         user.DeleteByIDRequest
			mockCalls   func()
			errExpected bool
		}{
			{
				desc:        "Error: Validation error",
				req:         user.DeleteByIDRequest{},
				mockCalls:   func() {},
				errExpected: true,
			},
			{
				desc: "Error: Delete user",
				req: user.DeleteByIDRequest{
					ID: "id",
				},
				mockCalls: func() {
					deps.mockUser.EXPECT().DeleteByID(gomock.Any(), gomock.Any()).Return(errMock)
				},
				errExpected: true,
			},
			{
				desc: "Success",
				req: user.DeleteByIDRequest{
					ID: "id",
				},
				mockCalls: func() {
					deps.mockUser.EXPECT().DeleteByID(gomock.Any(), gomock.Any()).Return(nil)
				},
				errExpected: false,
			},
		}

		for _, testCase := range testCases {
			Convey(testCase.desc, func() {
				testCase.mockCalls()

				err := svc.DeleteByID(context.Background(), testCase.req)
				if testCase.errExpected {
					So(err, ShouldNotBeNil)
				} else {
					So(err, ShouldBeNil)
				}
			})
		}
	})
}

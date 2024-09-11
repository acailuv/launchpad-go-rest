package user

import (
	"context"
	"launchpad-go-rest/pkg/types/user"
)

func (s service) Find(ctx context.Context) ([]user.FindResponse, error) {
	var response []user.FindResponse = make([]user.FindResponse, 0)

	res, err := s.user.Find(ctx)
	if err != nil {
		return response, err
	}

	for _, user := range res {
		response = append(response, user.ToFindResponse())
	}

	return response, nil
}

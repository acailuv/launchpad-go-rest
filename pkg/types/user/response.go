package user

import "encoding/json"

type FindByIDResponse struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

func (x FindByIDResponse) MarshalBinary() ([]byte, error) {
	return json.Marshal(x)
}

func (x *FindByIDResponse) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, x)
}

type FindResponse struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

type FindResponseList []FindResponse

func (x FindResponseList) MarshalBinary() ([]byte, error) {
	return json.Marshal(x)
}

func (x *FindResponseList) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, x)
}

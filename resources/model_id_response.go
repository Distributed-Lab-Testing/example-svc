package resources

type IdResponse struct {
	Data Key `json:"data"`
}

func NewIdResponse(key Key) IdResponse {
	return IdResponse{
		Data: key,
	}
}

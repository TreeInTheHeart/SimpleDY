package pojo

import "encoding/json"

func UnmarshalFeedResponse(data []byte) (FeedResponse, error) {
	var r FeedResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FeedResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

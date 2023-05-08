package storage

import (
	"encoding/json"
	"fmt"
)

type SaiStorageGetRequest struct {
	Collection string      `json:"collection"`
	Select     interface{} `json:"select"`
	Config     interface{} `json:"options"`
}

type SaiStorageGetResponse struct {
	Result []interface{} `json:"result"`
}

func (saiStorage *SaiStorage) Get(request SaiStorageGetRequest) (*SaiStorageGetResponse, error) {

	// Define the request body
	requestBody, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %v", err)
	}

	response, err := saiStorage.makeRequest("get", requestBody)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %v", err)
	}
	defer response.Body.Close()

	// Parse the response body into the struct
	var result SaiStorageGetResponse
	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		return nil, fmt.Errorf("failed to parse response body: %v", err)
	}

	// Return the parsed results
	return &result, nil
}

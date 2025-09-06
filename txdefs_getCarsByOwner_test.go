package main

import (
	"encoding/json"
	"testing"

	"github.com/hyperledger-labs/cc-tools/mock"
)

func TestGetCarsByOwner(t *testing.T) {
	stub := mock.NewMockStub("org1MSP", new(CCDemo))

	// Create a person first
	person := map[string]interface{}{
		"@assetType": "person",
		"id":         "318.207.920-48",
		"name":       "Maria",
	}
	personJSON, _ := json.Marshal(person)

	// Create person
	res := stub.MockInvoke("createAsset", [][]byte{
		[]byte("createAsset"),
		personJSON,
	})
	if res.GetStatus() != 200 {
		t.Errorf("Expected status 200, got %d. Message: %s", res.GetStatus(), res.GetMessage())
	}

	// Create first car with dateTransfered
	car1 := map[string]interface{}{
		"@assetType":       "car",
		"id":               "car-001",
		"make":             "Toyota",
		"model":            "Corolla",
		"colour":           "blue",
		"dateTransfered":   "2023-01-15T10:30:00Z",
		"lastTransferTxId": "tx-001",
		"owner": map[string]interface{}{
			"@assetType": "person",
			"@key":       "318.207.920-48",
		},
	}
	car1JSON, _ := json.Marshal(car1)

	res = stub.MockInvoke("createAsset", [][]byte{
		[]byte("createAsset"),
		car1JSON,
	})
	if res.GetStatus() != 200 {
		t.Errorf("Expected status 200, got %d. Message: %s", res.GetStatus(), res.GetMessage())
	}

	// Create second car with different dateTransfered
	car2 := map[string]interface{}{
		"@assetType":       "car",
		"id":               "car-002",
		"make":             "Honda",
		"model":            "Civic",
		"colour":           "red",
		"dateTransfered":   "2023-06-20T14:45:00Z",
		"lastTransferTxId": "tx-002",
		"owner": map[string]interface{}{
			"@assetType": "person",
			"@key":       "318.207.920-48",
		},
	}
	car2JSON, _ := json.Marshal(car2)

	res = stub.MockInvoke("createAsset", [][]byte{
		[]byte("createAsset"),
		car2JSON,
	})
	if res.GetStatus() != 200 {
		t.Errorf("Expected status 200, got %d. Message: %s", res.GetStatus(), res.GetMessage())
	}

	// Create third car with most recent dateTransfered
	car3 := map[string]interface{}{
		"@assetType":       "car",
		"id":               "car-003",
		"make":             "Ford",
		"model":            "Focus",
		"colour":           "white",
		"dateTransfered":   "2023-12-01T09:15:00Z",
		"lastTransferTxId": "tx-003",
		"owner": map[string]interface{}{
			"@assetType": "person",
			"@key":       "318.207.920-48",
		},
	}
	car3JSON, _ := json.Marshal(car3)

	res = stub.MockInvoke("createAsset", [][]byte{
		[]byte("createAsset"),
		car3JSON,
	})
	if res.GetStatus() != 200 {
		t.Errorf("Expected status 200, got %d. Message: %s", res.GetStatus(), res.GetMessage())
	}

	// Test getCarsByOwner transaction
	ownerRef := map[string]interface{}{
		"@assetType": "person",
		"@key":       "318.207.920-48",
	}
	ownerJSON, _ := json.Marshal(ownerRef)

	res = stub.MockInvoke("getCarsByOwner", [][]byte{
		[]byte("getCarsByOwner"),
		[]byte(`{"owner":` + string(ownerJSON) + `}`),
	})

	if res.GetStatus() != 200 {
		t.Errorf("Expected status 200, got %d. Message: %s", res.GetStatus(), res.GetMessage())
	}

	// Verify response contains cars ordered by dateTransfered desc
	var response []map[string]interface{}
	err := json.Unmarshal(res.GetPayload(), &response)
	if err != nil {
		t.Errorf("Failed to unmarshal response: %v", err)
	}

	if len(response) != 3 {
		t.Errorf("Expected 3 cars, got %d", len(response))
	}

	// Verify order: car3 (2023-12-01), car2 (2023-06-20), car1 (2023-01-15)
	if response[0]["id"] != "car-003" {
		t.Errorf("Expected first car to be car-003, got %s", response[0]["id"])
	}
	if response[1]["id"] != "car-002" {
		t.Errorf("Expected second car to be car-002, got %s", response[1]["id"])
	}
	if response[2]["id"] != "car-001" {
		t.Errorf("Expected third car to be car-001, got %s", response[2]["id"])
	}
}

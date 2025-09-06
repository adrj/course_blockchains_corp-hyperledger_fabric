package txdefs

// ...existing code...

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger-labs/cc-tools/assets"
	"github.com/hyperledger-labs/cc-tools/errors"
	"github.com/hyperledger-labs/cc-tools/events"
	sw "github.com/hyperledger-labs/cc-tools/stubwrapper"
	tx "github.com/hyperledger-labs/cc-tools/transactions"
)

var RegisterCar = tx.Transaction{
	Tag:         "registerCar",
	Label:       "Register Car",
	Description: "Registers a new car asset using transaction ID as the car ID.",
	Method:      "POST",
	Args: []tx.Argument{
		{
			Tag:         "make",
			Label:       "Car Make",
			Description: "Car make",
			DataType:    "string",
			Required:    true,
		},
		{
			Tag:         "model",
			Label:       "Car Model",
			Description: "Car model",
			DataType:    "string",
			Required:    true,
		},
		{
			Tag:         "colour",
			Label:       "Car Colour",
			Description: "Car colour",
			DataType:    "string",
			Required:    true,
		},
		{
			Tag:         "owner",
			Label:       "Car Owner",
			Description: "Car owner",
			DataType:    "->person",
			Required:    true,
		},
		{
			Tag:         "dateTransfered",
			Label:       "Date Transferred",
			Description: "Date transferred",
			DataType:    "datetime",
			Required:    true,
		},
	},
	Routine: func(stub *sw.StubWrapper, req map[string]interface{}) ([]byte, errors.ICCError) {
		carID := stub.Stub.GetTxID()
		carMap := make(map[string]interface{})
		carMap["@assetType"] = "car"
		carMap["id"] = carID
		carMap["make"] = req["make"]
		carMap["model"] = req["model"]
		carMap["colour"] = req["colour"]
		carMap["owner"] = req["owner"]
		carMap["dateTransfered"] = req["dateTransfered"]

		carAsset, err := assets.NewAsset(carMap)
		if err != nil {
			return nil, errors.WrapError(err, "Failed to create a new car asset")
		}

		_, err = carAsset.PutNew(stub)
		if err != nil {
			return nil, errors.WrapErrorWithStatus(err, "Error saving car asset on blockchain", err.Status())
		}

		carJSON, nerr := json.Marshal(carAsset)
		if nerr != nil {
			return nil, errors.WrapError(nil, "failed to encode car asset to JSON format")
		}

		logMsg, ok := json.Marshal(fmt.Sprintf("New car registered: %s", carID))
		if ok != nil {
			return nil, errors.WrapError(nil, "failed to encode log message to JSON format")
		}

		events.CallEvent(stub, "registerCarLog", logMsg)

		return carJSON, nil
	},
}

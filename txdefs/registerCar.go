package txdefs

import (
	"github.com/hyperledger-labs/cc-tools/chaincode"
	"github.com/hyperledger-labs/cc-tools/assets"
	"github.com/hyperledger-labs/cc-tools-demo/chaincode/assettypes"
)

var RegisterCarTx = chaincode.Transaction{
	Tag:         "registerCar",
	Label:       "Register Car",
	Description: "Registers a new car asset using transaction ID as the car ID.",
	Params: []chaincode.Param{
		{
			Tag:      "make",
			Label:    "Car Make",
			DataType: "string",
			Required: true,
		},
		{
			Tag:      "model",
			Label:    "Car Model",
			DataType: "string",
			Required: true,
		},
		{
			Tag:      "colour",
			Label:    "Car Colour",
			DataType: "string",
			Required: true,
		},
		{
			Tag:      "owner",
			Label:    "Car Owner",
			DataType: "->person",
			Required: true,
		},
		{
			Tag:      "dateTransfered",
			Label:    "Date Transferred",
			DataType: "datetime",
			Required: true,
		},
	},
	Invoke: func(ctx chaincode.Context, params map[string]interface{}) (interface{}, error) {
		carID := ctx.Stub.GetTxID()
		car := map[string]interface{}{
			"id":             carID,
			"make":           params["make"],
			"model":          params["model"],
			"colour":         params["colour"],
			"owner":          params["owner"],
			"dateTransfered": params["dateTransfered"],
		}
		return assets.CreateAsset(ctx, assettypes.Car, car)
	},
}

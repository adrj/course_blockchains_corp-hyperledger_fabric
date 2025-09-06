package txdefs

import (
	"encoding/json"
	"github.com/hyperledger-labs/cc-tools/assets"
	"github.com/hyperledger-labs/cc-tools/errors"
	sw "github.com/hyperledger-labs/cc-tools/stubwrapper"
	tx "github.com/hyperledger-labs/cc-tools/transactions"
	"time"
)

var TransferCar = tx.Transaction{
	Tag:         "transferCar",
	Label:       "Transfer Car",
	Description: "Transfers a car to a new owner, only allowed once every 30 days.",
	Method:      "POST",
	Args: []tx.Argument{
		{
			Tag:         "carId",
			Label:       "Car ID",
			Description: "ID of the car to transfer",
			DataType:    "string",
			Required:    true,
		},
		{
			Tag:         "newOwner",
			Label:       "New Owner",
			Description: "Person asset reference for new owner",
			DataType:    "->person",
			Required:    true,
		},
	},
	Routine: func(stub *sw.StubWrapper, args map[string]interface{}) ([]byte, errors.ICCError) {
		carId := args["carId"].(string)
		newOwner := args["newOwner"]

		// Obtém o ID da transação
		txID := stub.Stub.GetTxID()

		// Cria chave para buscar o asset Car
		carKey, err := assets.NewKey(map[string]interface{}{
			"@assetType": "car",
			"id":         carId,
		})
		if err != nil {
			return nil, errors.WrapError(err, "failed to create car key")
		}

		carAsset, err := carKey.Get(stub)
		if err != nil {
			return nil, errors.WrapError(err, "car not found")
		}

		// Obtém o timestamp da transação
		ts, tsErr := stub.Stub.GetTxTimestamp()
		if tsErr != nil {
			return nil, errors.WrapError(tsErr, "failed to get transaction timestamp")
		}
		txTime := time.Unix(ts.Seconds, int64(ts.Nanos))

		carMap := (map[string]interface{})(*carAsset)

		var lastTransfer time.Time
		if dateVal, ok := carMap["dateTransfered"]; ok && dateVal != nil {
			switch v := dateVal.(type) {
			case time.Time:
				lastTransfer = v
			case string:
				t, err := time.Parse(time.RFC3339, v)
				if err == nil {
					lastTransfer = t
				}
			}
		}

		if !lastTransfer.IsZero() {
			if txTime.Sub(lastTransfer) < 30*24*time.Hour {
				return nil, errors.NewCCError("car can only be transferred once every 30 days", 400)
			}
		}

		// Atualiza owner e dateTransfered usando o timestamp da transação
		// Também registra o ID da transação que fez a transferência
		carMap["owner"] = newOwner
		carMap["dateTransfered"] = txTime.Format(time.RFC3339)
		carMap["lastTransferTxId"] = txID // Registra o ID da transação

		updatedCarAsset, err := assets.NewAsset(carMap)
		if err != nil {
			return nil, errors.WrapError(err, "failed to create updated car asset")
		}

		_, err = updatedCarAsset.Put(stub)
		if err != nil {
			return nil, errors.WrapError(err, "failed to transfer car")
		}

		carJSON, nerr := json.Marshal(updatedCarAsset)
		if nerr != nil {
			return nil, errors.WrapError(nil, "failed to encode car asset to JSON format")
		}

		return carJSON, nil
	},
}

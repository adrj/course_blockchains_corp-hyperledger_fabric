package txdefs

import (
	"encoding/json"

	"github.com/hyperledger-labs/cc-tools/accesscontrol"
	"github.com/hyperledger-labs/cc-tools/assets"
	"github.com/hyperledger-labs/cc-tools/errors"
	sw "github.com/hyperledger-labs/cc-tools/stubwrapper"
	tx "github.com/hyperledger-labs/cc-tools/transactions"
)

// GetCarsByOwner returns all cars owned by a specific person, ordered by transfer date (descending)
var GetCarsByOwner = tx.Transaction{
	Tag:         "getCarsByOwner",
	Label:       "Get Cars By Owner",
	Description: "Returns all cars owned by a specific person, ordered by transfer date in descending order",
	Method:      "GET",
	Callers: []accesscontrol.Caller{ // Any org can call this transaction
		{MSP: "org1MSP"},
		{MSP: "org2MSP"},
		{MSP: "org3MSP"},
	},

	Args: []tx.Argument{
		{
			Tag:         "owner",
			Label:       "Owner",
			Description: "The person who owns the cars",
			DataType:    "->person",
			Required:    true,
		},
	},
	Routine: func(stub *sw.StubWrapper, req map[string]interface{}) ([]byte, errors.ICCError) {
		// Get the owner asset key
		ownerKey, ok := req["owner"].(assets.Key)
		if !ok {
			return nil, errors.NewCCError("invalid owner format", 400)
		}

		// Get the owner ID from the key
		ownerID := ownerKey.Key()

		// Build the query selector for CouchDB
		query := map[string]interface{}{
			"selector": map[string]interface{}{
				"@assetType": "car",
				"owner": map[string]interface{}{
					"@key": ownerID,
				},
			},
			"sort": []map[string]interface{}{
				{
					"dateTransfered": "desc",
				},
			},
		}

		// Execute the rich query using assets.Search
		response, ccErr := assets.Search(stub, query, "", true)
		if ccErr != nil {
			return nil, ccErr
		}

		// Convert response to JSON
		responseBytes, err := json.Marshal(response)
		if err != nil {
			return nil, errors.WrapError(err, "failed to marshal response")
		}

		return responseBytes, nil
	},
}

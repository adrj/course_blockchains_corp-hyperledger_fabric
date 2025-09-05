package assettypes

import (
	"fmt"

	"github.com/hyperledger-labs/cc-tools/assets"
)

var Car = assets.AssetType{
	Tag:         "car",
	Label:       "Car",
	Description: "Car asset with make, model, colour and owner",

	Props: []assets.AssetProp{
		{
			// Primary key - id
			Required: true,
			IsKey:    true,
			Tag:      "id",
			Label:    "Car ID",
			DataType: "string",
			Validate: func(id interface{}) error {
				idStr := id.(string)
				if idStr == "" {
					return fmt.Errorf("id must be non-empty")
				}
				return nil
			},
		},
		{
			// Required property - colour (only org1 and org2 can modify)
			Required: true,
			Tag:      "colour",
			Label:    "Car Colour",
			DataType: "string",
			Writers:  []string{`org1MSP`, `org2MSP`},
			Validate: func(colour interface{}) error {
				colourStr := colour.(string)
				if colourStr == "" {
					return fmt.Errorf("colour must be non-empty")
				}
				return nil
			},
		},
		{
			// Date the car was transferred to the current owner
			Tag:      "dateTransfered",
			Label:    "Date Transferred",
			DataType: "datetime",
		},
		{
			// Reference to Person asset - owner
			Tag:      "owner",
			Label:    "Car Owner",
			DataType: "->person",
		},
	},
}

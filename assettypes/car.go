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
			// Primary key - make
			Required: true,
			IsKey:    true,
			Tag:      "make",
			Label:    "Car Make",
			DataType: "string",
			Validate: func(make interface{}) error {
				makeStr := make.(string)
				if makeStr == "" {
					return fmt.Errorf("make must be non-empty")
				}
				return nil
			},
		},
		{
			// Primary key - model
			Required: true,
			IsKey:    true,
			Tag:      "model",
			Label:    "Car Model",
			DataType: "string",
			Validate: func(model interface{}) error {
				modelStr := model.(string)
				if modelStr == "" {
					return fmt.Errorf("model must be non-empty")
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
			// Reference to Person asset - owner
			Tag:      "owner",
			Label:    "Car Owner",
			DataType: "->person",
		},
	},
}

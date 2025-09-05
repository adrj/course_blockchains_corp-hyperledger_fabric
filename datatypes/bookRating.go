package datatypes

import (
	"github.com/hyperledger-labs/cc-tools/assets"
	"github.com/hyperledger-labs/cc-tools/errors"
)

var bookRating = assets.DataType{
	AcceptedFormats: []string{"number"},
	Parse: func(data interface{}) (string, interface{}, errors.ICCError) {
		rating, ok := data.(float64)
		if !ok {
			return "", nil, errors.NewCCError("property must be a number", 400)
		}

		if rating < 1.0 {
			return "", nil, errors.NewCCError("rating must be greater than or equal to 1.0", 400)
		}

		if rating > 10.0 {
			return "", nil, errors.NewCCError("rating must be less than or equal to 10.0", 400)
		}

		return "number", rating, nil
	},
}

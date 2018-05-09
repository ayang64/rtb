package rtb

import (
	"encoding/json"
)

type Deal struct {
	ID            string      `json:"id"`       // id string; required A unique identifier for the deal.
	Quantity      int         `json:"qty"`      // qty integer Number of instances of the item to which the deal applies. Default is the full quantity specified in the “item.qty” attribute.
	Floor         float32     `json:"flr"`      // flr float Minimum deal price for this item expressed in CPM.
	FloorCurrency string      `json:"flrcur"`   // flrcur string; default “USD” Currency of the “flr” attribute specified using ISO-4217 alpha codes.
	Auctiontype   int         `json:"at"`       //  integer Optional override of the overall auction type of the request, where 1 = First Price, 2 = Second Price Plus, 3 = the value passed in “flr” is the agreed upon deal price. Additional auction types can be defined by the exchange.
	Seat          []string    `json:"seat"`     // seat string array Whitelist of buyer seats allowed to bid on this deal. IDs of seats and the buyer’s customers to which they refer must be coordinated between bidders and the exchange a priori. Omission implies no restrictions.
	AuthDomains   []string    `json:"wadomain"` //  string array Array of advertiser domains (e.g., advertiser.com) allowed to bid on this deal. Omission implies no restrictions.
	Ext           interface{} `json:"ext"`      //  object Optional exchange-specific extensions.
}

func (D *Deal) UnmarshalJSON(d []byte) error {
	type alias Deal

	a := alias{
		FloorCurrency: "USD",
	}

	if err := json.Unmarshal(d, &a); err != nil {
		return err
	}

	*D = Deal(a)

	return nil
}

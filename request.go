package rtb

import (
	"encoding/json"
)

type unk struct{}

type Request struct {
	Id          string      `json:"id"`             // required -  Unique ID of the bid request; provided by the exchange.
	Test        int         `json:"test,omitempty"` // default 0 -  Indicator of test mode in which auctions are not billable, where 0 = live mode, 1 = test mode.
	TMax        *int        `json:"tmax,omitempty"` // Maximum time in milliseconds the exchange allows for bids to be receivedincluding Internet latency to avoid timeout. This value supersedes any a priori guidance from the exchange.
	AuctionType int         `json:"at"`             // default 2 - Auction type, where 1 = First Price, 2 = Second Price Plus. Values greater than 500 can be used for exchange-specific auction types.
	Currencies  []string    `json:"curs"`           // default [“USD”] - Array of currencies for bids on this bid request using ISO-4217 alpha codes. Recommended if the exchange accepts multiple currencies. If omitted, the single currency of “USD” is assumed.
	Wcurs       int         `json:"wcurs"`          // wcurs integer; default 0 Flag that determines the restriction interpretation of the “curs” array, where 0 = block list, 1 = whitelist.  seats string array Restriction list of buyer seats for bidding on this item. Knowledge of buyer’s customers and their seat IDs must be coordinated between parties a priori. Omission implies no restrictions.
	Wseats      int         `json:"wseats"`         // wseats integer; default 0 Flag that determines the restriction interpretation of the “seats” array, where 0 = block list, 1 = whitelist.
	Source      Source      `json:"source"`         // source object A “Source” object that provides data about the inventory source and which entity makes the final decision.
	Offer       Offer       `json:"offer"`          // offer object; required An “Offer” object that conveys the item(s) being offered for sale.
	Domain      unk         `json:"domain"`         // domain object; recommended Layer-4 domain object structure that provides context for the items being offered (e.g., user, device, site or app, etc.) conforming to the specification and version referenced in “openrtb.domainspec” and “openrtb.domainver”.
	Ext         interface{} `json:"ext"`            // ext object Optional exchange-specific extensions.
}

func (r *Request) UnmarshalJSON(d []byte) error {
	type alias Request

	a := alias{
		Test:        0,
		AuctionType: 2,
		Currencies:  []string{"USD"},
	}

	if err := json.Unmarshal(d, &a); err != nil {
		return err
	}

	*r = Request(a)

	return nil
}

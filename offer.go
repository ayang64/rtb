package rtb

type Offer struct {
	Item    []Item      `json:"item"`    // object array; required Array of “Item” objects (at least one) that constitute the set of goods being offered for sale.
	Package int         `json:"package"` // integer; default 0 Flag to indicate if the Exchange can verify that the items offered represent all of the items available in context (e.g., all impressions on a web page, all video spots such as pre/mid/post roll) to support roadblocking, where 0 = no or unknown, 1 = yes.
	DBurl   string      `json:"dburl"`   // dburl string Billing notice URL called by the demand partner when a winning bid becomes billable based on partner-specific business policy. Substitution macros may be included. One of burl in response or dburl in request is required (exception for VAST).
	Ext     interface{} `json:"ext"`     // ext object Optional exchange-specific extensions.
}

func (o *Offer) UnmarshalJSON(d []byte) error {
	type alias Offer

	a := alias{
		Package: 0,
	}

	if err := json.Unmarshal(d, &a); err != nil {
		return err
	}

	*br = Offer(a)

	return nil
}

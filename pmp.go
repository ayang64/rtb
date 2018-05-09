package rtb

import (
	"encoding/json"
)

type Pmp struct {
	Private int         `json:"private"` // private integer; default 0 Indicator of auction eligibility to seats named in “Deal” objects, where 0 = all bids are accepted, 1 = bids are restricted to the deals specified and the terms thereof.
	Deal    *unk        `json:"deal"`    // deal object array Array of “Deal” objects that convey special terms applicable to this item.
	Ext     interface{} `json:"ext"`     // ext object Optional exchange-specific extensions.
}

func (p *Pmp) UnmarshalJSON(d []byte) error {
	type alias Pmp

	a := alias{
		Private: 0,
	}

	if err := json.Unmarshal(d, &a); err != nil {
		return err
	}

	*p = Pmp(a)

	return nil
}

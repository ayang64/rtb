package rtb

type Item struct {
	Id            string      `json:"id"`     // id string; required A unique identifier for this item within the context of the offer (typically starts with 1 and increments).
	Quantity      int         `json:"qty"`    // qty integer; default 1 The number of instances of this item being offered (e.g., multiple identical impressions in a digital out-of-home scenario).
	Floor         float32     `json:"flr"`    // flr float Minimum bid price for this item expressed in CPM.
	FloorCurrency string      `json:"flrcur"` // flrcur string; default “USD” Currency of the “flr” attribute specified using ISO-4217 alpha codes.
	Sequence      int         `json:"seq"`    // seq integer If multiple items are offered in the same bid request, the sequence number allows for the coordinated delivery.
	Pmp           unk         `json:"pmp"`    //  object A “Pmp” object containing any private marketplace deals for this item.
	Domain        unk         `json:"domain"` // domain object; required Layer-4 domain object structure that specifies the item being offered (e.g., impression) conforming to the specification and version referenced in “openrtb.domainspec” and “openrtb.domainver”.
	Ext           interface{} `json:"ext"`    // object Optional exchange-specific extensions.
}

func (i *Item) UnmarshalJSON(d []byte) error {
	type alias Item

	a := alias{
		Quantity:      1,
		FloorCurrency: "USD",
	}

	if err := json.Unmarshal(d, &a); err != nil {
		return err
	}

	*br = Item(a)

	return nil
}

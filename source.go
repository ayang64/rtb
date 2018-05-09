package rtb

type Source struct {
	FinalSaleDecision int         `json:"fd"`     // fd integer; recommended Entity responsible for the final sale decision, where 0 = exchange, 1 = upstream source.
	TransactionID     string      `json:"tid"`    // string; recommended Transaction ID that must be common across all participants in this bid request (e.g., potentially multiple exchanges).
	PaymentChain      string      `json:"pchain"` // recommended Payment ID chain string containing embedded syntax described in the TAG Payment ID Protocol.
	Ext               interface{} `json:"ext"`    // ext object Optional exchange-specific extensions.
}

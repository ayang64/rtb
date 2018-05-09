package rtb

type Response struct {
	ID       string      `json:"id"`      // id string; required ID of the bid request to which this is a response; must match the “request.id” attribute.
	BidID    string      `json:"bidid"`   // bidid string Bidder generated response ID to assist with logging/tracking.  nbr integer Reason for not bidding if applicable (see Enumerations).
	Currency string      `json:"cur"`     // cur string Bid currency using ISO-4217 alpha codes.  OpenRTB 3.0 Framework for Public Comment IAB Tech Lab https://iabtechlab.com/openrtb 18
	CDATA    string      `json:"cdata"`   // Allows bidder to set data in the exchange’s cookie if supported by the exchange. The string must be in base85 cookie-safe characters. JSON encoding must be used to include “escaped” quotation marks.
	SeatBid  *unk        `json:"seatbid"` // object array Array of “Seatbid” objects; 1+ required if a bid is to be made.
	Ext      interface{} `json:"ext"`     // object Optional demand source specific extensions.
}

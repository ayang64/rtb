package rtb

import (
	"encoding/json"
	"os"
	"testing"
)

func TestUnmarshalBidRequest(t *testing.T) {

	inf, err := os.Open("bid-request.json")

	if err != nil {
		t.Fatalf("%v", err)
		t.FailNow()
	}

	d := json.NewDecoder(inf)

	req := BidRequest{}
	if err := d.Decode(&req); err != nil {
		t.Fatalf("%v", err)
		t.FailNow()
	}

	t.Logf("%#v", req)
}

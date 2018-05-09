package rtb

import (
	"encoding/json"
	"os"
	"testing"
)

func TestUnmarshalRequest(t *testing.T) {
	inf, err := os.Open("test-fixtures/bid-request.json")

	if err != nil {
		t.Fatalf("%v", err)
		t.FailNow()
	}

	d := json.NewDecoder(inf)

	req := Request{}
	if err := d.Decode(&req); err != nil {
		t.Fatalf("%v", err)
		t.FailNow()
	}

	t.Logf("%#v", req)
}

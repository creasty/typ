package typ

import (
	"fmt"
	"testing"
)

func Test_URL(t *testing.T) {
	u := NewURL()
	rawData := "http://example.com:1234/path"
	data := []byte(fmt.Sprintf("%q", rawData))

	{
		got, err := u.MarshalJSON()
		if err != nil {
			t.Error("error", err)
			return
		}

		if string(got[:]) != "null" {
			t.Error("error", string(got[:]))
			return
		}
	}

	if err := u.UnmarshalJSON(data); err != nil {
		t.Error("should not fail", err)
		return
	}

	if u.URL == nil {
		t.Error("should unmarshal")
		return
	}

	if u.URL.String() != rawData {
		t.Error("error", u.URL)
		return
	}

	{
		got, err := u.MarshalJSON()
		if err != nil {
			t.Error("error", err)
			return
		}

		if string(got[:]) != string(data[:]) {
			t.Error("error", string(got[:]), string(data[:]))
		}
	}
}

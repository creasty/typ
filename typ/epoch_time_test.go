package typ

import (
	"testing"
)

func Test_EpochTime(t *testing.T) {
	et := NewEpochTime()
	data := []byte("1480411283")

	{
		got, err := et.MarshalJSON()
		if err != nil {
			t.Error("error", err)
			return
		}

		if string(got[:]) != "null" {
			t.Error("error", string(got[:]))
			return
		}
	}

	if err := et.UnmarshalJSON(data); err != nil {
		t.Error("should not fail", err)
		return
	}

	if et.Time == nil {
		t.Error("should unmarshal")
		return
	}

	if et.Time.UTC().Format("2006-01-02 15:04:05") != "2016-11-29 09:21:23" {
		t.Error("error", et.Time.UTC())
		return
	}

	{
		got, err := et.MarshalJSON()
		if err != nil {
			t.Error("error", err)
			return
		}

		if string(got[:]) != string(data[:]) {
			t.Error("error", string(got[:]), string(data[:]))
		}
	}
}

func Test_EpochMsec(t *testing.T) {
	et := NewEpochMsec()
	data := []byte("1480411283963")

	{
		got, err := et.MarshalJSON()
		if err != nil {
			t.Error("error", err)
			return
		}

		if string(got[:]) != "null" {
			t.Error("error", string(got[:]))
			return
		}
	}

	if err := et.UnmarshalJSON(data); err != nil {
		t.Error("should not fail", err)
		return
	}

	if et.Time == nil {
		t.Error("should unmarshal")
		return
	}

	if et.Time.UTC().Format("2006-01-02 15:04:05") != "2016-11-29 09:21:23" {
		t.Error("error", et.Time.UTC())
		return
	}

	if et.Time.Nanosecond() != 963*1e6 {
		t.Error("error", et.Time.Nanosecond())
		return
	}

	{
		got, err := et.MarshalJSON()
		if err != nil {
			t.Error("error", err)
			return
		}

		if string(got[:]) != string(data[:]) {
			t.Error("error", string(got[:]), string(data[:]))
		}
	}
}

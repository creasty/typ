package typ

import (
	"encoding/json"
	"strconv"
	"time"
)

/*  sec
-----------------------------------------------*/
type EpochTime struct {
	*time.Time
}

func NewEpochTime() *EpochTime {
	return &EpochTime{}
}

func NewEpochTimeFromTime(t time.Time) *EpochTime {
	return &EpochTime{Time: &t}
}

func NewEpochTimeFromTimePtr(t *time.Time) *EpochTime {
	return &EpochTime{Time: t}
}

func (self *EpochTime) MarshalJSON() ([]byte, error) {
	if self.Time == nil {
		return []byte("null"), nil
	} else {
		return []byte(strconv.FormatInt(self.Time.Unix(), 10)), nil
	}
}

func (self *EpochTime) UnmarshalJSON(data []byte) error {
	var et int64
	if err := json.Unmarshal(data, &et); err != nil {
		return err
	}

	t := time.Unix(et, 0)
	self.Time = &t

	return nil
}

/*  msec
-----------------------------------------------*/
type EpochMsec struct {
	*time.Time
}

func NewEpochMsec() *EpochMsec {
	return &EpochMsec{}
}

func NewEpochMsecFromTime(t time.Time) *EpochMsec {
	return &EpochMsec{Time: &t}
}

func NewEpochMsecFromTimePtr(t *time.Time) *EpochMsec {
	return &EpochMsec{Time: t}
}

func (self *EpochMsec) MarshalJSON() ([]byte, error) {
	if self.Time == nil {
		return []byte("null"), nil
	} else {
		return []byte(strconv.FormatInt(self.Time.UnixNano()/1e6, 10)), nil
	}
}

func (self *EpochMsec) UnmarshalJSON(data []byte) error {
	var et int64
	if err := json.Unmarshal(data, &et); err != nil {
		return err
	}

	t := time.Unix(et/1e3, (et%1e3)*1e6)
	self.Time = &t

	return nil
}

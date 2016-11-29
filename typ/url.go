package typ

import (
	"encoding/json"
	"fmt"
	"net/url"
	// "github.com/ghodss/yaml"
)

type URL struct {
	*url.URL
}

func NewURL() *URL {
	return &URL{}
}

func NewURLFromURL(u url.URL) *URL {
	return &URL{URL: &u}
}

func NewURLFromURLPtr(u *url.URL) *URL {
	return &URL{URL: u}
}

func (self *URL) UnmarshalJSON(b []byte) error {
	var str string
	if err := json.Unmarshal(b, &str); err != nil {
		return err
	}

	u, err := url.Parse(str)
	self.URL = u
	return err
}

func (self *URL) MarshalJSON() ([]byte, error) {
	if self.URL == nil {
		return []byte("null"), nil
	} else {
		return []byte(fmt.Sprintf("%q", self.URL.String())), nil
	}
}

/*
func (self *URL) UnmarshalYAML(b []byte) error {
	var str string
	if err := yaml.Unmarshal(b, &str); err != nil {
		return err
	}

	u, err := url.Parse(str)
	self.URL = u
	return err
}
*/

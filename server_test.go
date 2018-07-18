package puppetdb

import (
	"testing"
	"time"
)

func TestSetHTTPTimeout(t *testing.T) {
	s := NewServer("http://localhost:8080")
	expect := time.Second * 15
	s.SetHTTPTimeout(expect)
	if s.HTTPTimeout != expect {
		t.Errorf("Expected HTTPTimeout set to 15 got %s", s.HTTPTimeout)
	}
}

package puppetdb

import (
	"testing"
	"time"
)

func TestSetHTTPTimeout(t *testing.T) {
	expect := time.Second * 15
	s := NewServer("http://localhost:8080", WithTimeout(expect))
	if s.client.Timeout != expect {
		t.Errorf("Expected HTTPTimeout set to %s got %s", expect, s.client.Timeout)
	}
}

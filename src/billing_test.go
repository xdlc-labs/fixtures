package main

import (
	"net/http"
	"testing"
	"time"
)

func TestBillingUpstream(t *testing.T) {
	c := &http.Client{Timeout: 2 * time.Second}
	resp, err := c.Get("http://billing.internal.example:9/ready")
	if err != nil {
		t.Fatalf("billing upstream: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("billing status %d", resp.StatusCode)
	}
}

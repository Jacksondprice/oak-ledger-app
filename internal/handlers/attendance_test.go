package handlers

import (
	"net/http/httptest"
	"testing"
)

func TestListAttendance(t *testing.T) {
	rec := httptest.NewRecorder()
	ListAttendance(rec, httptest.NewRequest("GET", "/api/attendance", nil))
	if rec.Code != 200 {
		t.Fatalf("status %d", rec.Code)
	}
}

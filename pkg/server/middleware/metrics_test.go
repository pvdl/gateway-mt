// Copyright (C) 2021 Storj Labs, Inc.
// See LICENSE for copying information.

package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/spacemonkeygo/monkit/v3"
	"github.com/stretchr/testify/assert"
)

func TestMetrics(t *testing.T) {
	status := func(code int) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(code)
		})
	}

	req, err := http.NewRequest("GET", "", nil)
	assert.NoError(t, err)
	rr := httptest.NewRecorder()

	Metrics(status(200)).ServeHTTP(rr, req) // 1 200
	Metrics(status(400)).ServeHTTP(rr, req) // 2 400s
	Metrics(status(400)).ServeHTTP(rr, req)
	Metrics(status(500)).ServeHTTP(rr, req) // 3 500s
	Metrics(status(500)).ServeHTTP(rr, req)
	Metrics(status(500)).ServeHTTP(rr, req)

	c := monkit.Collect(monkit.ScopeNamed("storj.io/gateway-mt/pkg/server/middleware"))
	assert.Equal(t, 1.0, c["request_times,status_code=200 count"])
	assert.Equal(t, 2.0, c["request_times,status_code=400 count"])
	assert.Equal(t, 3.0, c["request_times,status_code=500 count"])

}
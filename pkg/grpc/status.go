// Copyright 2024 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package grpc

import (
	"net/http"

	"github.com/FishGoddess/errors"
	"github.com/FishGoddess/errors/status"
	grpccodes "google.golang.org/grpc/codes"
	grpcstatus "google.golang.org/grpc/status"
)

func init() {
	status.RegisterStatuses(
		status.New(http.StatusBadRequest, "Bad Request", errors.IsBadRequest),
		status.New(http.StatusForbidden, "Forbidden", errors.IsForbidden),
		status.New(http.StatusNotFound, "Not Found", errors.IsNotFound),
		status.New(http.StatusRequestTimeout, "Request Timeout", errors.IsRequestTimeout),
		status.New(http.StatusInternalServerError, "Internal Server Error", errors.IsInternalServerError),
	)
}

func wrapStatus(err error) error {
	code, msg := status.Parse(err)
	gCode := grpccodes.Code(code)

	return grpcstatus.New(gCode, msg).Err()
}

// Copyright (C) 2021 Storj Labs, Inc.
// See LICENSE for copying information.

package minio

import (
	"context"
	"net/http"
	"net/url"
	_ "unsafe" // for go:linkname

	"github.com/gorilla/mux"

	"storj.io/minio/cmd"
	"storj.io/minio/cmd/config/policy/opa"
	"storj.io/minio/pkg/auth"
	"storj.io/minio/pkg/bucket/policy"
)

// CollectAPIStats exposes minio's cmd.collectAPIStats.
//
//nolint:golint
//go:linkname CollectAPIStats storj.io/minio/cmd.collectAPIStats
func CollectAPIStats(api string, f http.HandlerFunc) http.HandlerFunc

// GlobalBucketQuotaSys exposes minio's cmd.globalBucketQuotaSys.
//
//nolint:golint
//go:linkname GlobalBucketQuotaSys storj.io/minio/cmd.globalBucketQuotaSys
var GlobalBucketQuotaSys *cmd.BucketQuotaSys

// GlobalCLIContext exposes minio's cmd.globalCLIContext.
//
//nolint:golint
//go:linkname GlobalCLIContext storj.io/minio/cmd.globalCLIContext
var GlobalCLIContext = struct {
	JSON, Quiet    bool
	Anonymous      bool
	Addr           string
	StrictS3Compat bool
}{}

// GlobalHandlers exposes minio's cmd.globalHandlers.
//
//nolint:golint
//go:linkname GlobalHandlers storj.io/minio/cmd.globalHandlers
var GlobalHandlers []mux.MiddlewareFunc

// GlobalIAMSys exposes minio's cmd.globalIAMSys.
//
//nolint:golint
//go:linkname GlobalIAMSys storj.io/minio/cmd.globalIAMSys
var GlobalIAMSys *cmd.IAMSys

// GlobalIsGateway exposes minio's cmd.globalIsGateway.
//
//nolint:golint
//go:linkname GlobalIsGateway storj.io/minio/cmd.globalIsGateway
var GlobalIsGateway bool

// GlobalNotificationSys exposes minio's cmd.globalNotificationSys.
//
//nolint:golint
//go:linkname GlobalNotificationSys storj.io/minio/cmd.globalNotificationSys
var GlobalNotificationSys *cmd.NotificationSys

// GlobalPolicyOPA exposes minio's cmd.globalPolicyOPA.
//
//nolint:golint
//go:linkname GlobalPolicyOPA storj.io/minio/cmd.globalPolicyOPA
var GlobalPolicyOPA *opa.Opa

// GlobalIsSSL exposes minio's cmd.globalIsSSL
//
//nolint:golint
//go:linkname GlobalIsSSL storj.io/minio/cmd.globalIsSSL
var GlobalIsSSL bool

// GetAPIError exposes minio's cmd.getAPIError.
//
//nolint:golint
//go:linkname GetAPIError storj.io/minio/cmd.getAPIError
func GetAPIError(code cmd.APIErrorCode) cmd.APIError

// HandleCommonEnvVars exposes minio's cmd.handleCommonEnvVars.
//
//nolint:golint
//go:linkname HandleCommonEnvVars storj.io/minio/cmd.handleCommonEnvVars
func HandleCommonEnvVars()

// HTTPTraceAll exposes minio's cmd.httpTraceAll.
//
//nolint:golint
//go:linkname HTTPTraceAll storj.io/minio/cmd.httpTraceAll
func HTTPTraceAll(f http.HandlerFunc) http.HandlerFunc

// HTTPTraceHdrs exposes minio's cmd.httpTraceHdrs.
//
//nolint:golint
//go:linkname HTTPTraceHdrs storj.io/minio/cmd.httpTraceHdrs
func HTTPTraceHdrs(f http.HandlerFunc) http.HandlerFunc

// MaxClients exposes minio's cmd.maxClients.
//
//nolint:golint
//go:linkname MaxClients storj.io/minio/cmd.maxClients
func MaxClients(f http.HandlerFunc) http.HandlerFunc

// MethodNotAllowedHandler exposes minio's cmd.methodNotAllowedHandler.
//
//nolint:golint
//go:linkname MethodNotAllowedHandler storj.io/minio/cmd.methodNotAllowedHandler
func MethodNotAllowedHandler(api string) http.HandlerFunc

// ErrorResponseHandler exposes minio's cmd.errorResponseHandler.
//
//nolint:golint
//go:linkname ErrorResponseHandler storj.io/minio/cmd.errorResponseHandler
func ErrorResponseHandler(w http.ResponseWriter, r *http.Request)

// SetObjectLayer exposes minio's cmd.setObjectLayer.
//
//nolint:golint
//go:linkname SetObjectLayer storj.io/minio/cmd.setObjectLayer
func SetObjectLayer(o cmd.ObjectLayer)

// writeErrorResponse exposes minio's cmd.writeErrorResponse.
//
//nolint:golint
//go:linkname writeErrorResponse storj.io/minio/cmd.writeErrorResponse
func writeErrorResponse(ctx context.Context, w http.ResponseWriter, err cmd.APIError, reqURL *url.URL, browser bool)

// writeSuccessResponseXML exposes minio's cmd.writeSuccessResponseXML.
//
//nolint:golint
//go:linkname writeSuccessResponseXML storj.io/minio/cmd.writeSuccessResponseXML
func writeSuccessResponseXML(w http.ResponseWriter, response []byte)

//
// Exposes below are needed for ListBucketsWithAttributionHandler
//

// checkRequestAuthTypeCredential exposes minio's cmd.checkRequestAuthTypeCredential.
//
//nolint:golint
//go:linkname checkRequestAuthTypeCredential storj.io/minio/cmd.checkRequestAuthTypeCredential
func checkRequestAuthTypeCredential(ctx context.Context, r *http.Request, action policy.Action, bucketName, objectName string) (cred auth.Credentials, owner bool, s3Err cmd.APIErrorCode)

// encodeResponse exposes minio's cmd.encodeResponse.
//
//nolint:golint
//go:linkname encodeResponse storj.io/minio/cmd.encodeResponse
func encodeResponse(response interface{}) []byte

// globalMinioDefaultOwnerID exposes minio's cmd.globalMinioDefaultOwnerID.
//
//nolint:golint
//go:linkname globalMinioDefaultOwnerID storj.io/minio/cmd.globalMinioDefaultOwnerID
var globalMinioDefaultOwnerID string

// newContext exposes minio's cmd.newContext.
//
//nolint:golint
//go:linkname newContext storj.io/minio/cmd.newContext
func newContext(r *http.Request, w http.ResponseWriter, api string) context.Context

// toAPIError exposes minio's cmd.toAPIError.
//
//nolint:golint
//go:linkname toAPIError storj.io/minio/cmd.toAPIError
func toAPIError(ctx context.Context, err error) cmd.APIError

// Copyright 2019 Google LLC.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated file. DO NOT EDIT.

// Package playcustomapp provides access to the Google Play Custom App Publishing API.
//
// For product documentation, see: https://developers.google.com/android/work/play/custom-app-api
//
// Creating a client
//
// Usage example:
//
//   import "google.golang.org/api/playcustomapp/v1"
//   ...
//   ctx := context.Background()
//   playcustomappService, err := playcustomapp.NewService(ctx)
//
// In this example, Google Application Default Credentials are used for authentication.
//
// For information on how to create and obtain Application Default Credentials, see https://developers.google.com/identity/protocols/application-default-credentials.
//
// Other authentication options
//
// To use an API key for authentication (note: some APIs do not support API keys), use option.WithAPIKey:
//
//   playcustomappService, err := playcustomapp.NewService(ctx, option.WithAPIKey("AIza..."))
//
// To use an OAuth token (e.g., a user token obtained via a three-legged OAuth flow), use option.WithTokenSource:
//
//   config := &oauth2.Config{...}
//   // ...
//   token, err := config.Exchange(ctx, ...)
//   playcustomappService, err := playcustomapp.NewService(ctx, option.WithTokenSource(config.TokenSource(ctx, token)))
//
// See https://godoc.org/google.golang.org/api/option/ for details on options.
package playcustomapp // import "google.golang.org/api/playcustomapp/v1"

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	googleapi "google.golang.org/api/googleapi"
	gensupport "google.golang.org/api/internal/gensupport"
	option "google.golang.org/api/option"
	htransport "google.golang.org/api/transport/http"
)

// Always reference these packages, just in case the auto-generated code
// below doesn't.
var _ = bytes.NewBuffer
var _ = strconv.Itoa
var _ = fmt.Sprintf
var _ = json.NewDecoder
var _ = io.Copy
var _ = url.Parse
var _ = gensupport.MarshalJSON
var _ = googleapi.Version
var _ = errors.New
var _ = strings.Replace
var _ = context.Canceled

const apiId = "playcustomapp:v1"
const apiName = "playcustomapp"
const apiVersion = "v1"
const basePath = "https://www.googleapis.com/playcustomapp/v1/accounts/"

// OAuth2 scopes used by this API.
const (
	// View and manage your Google Play Developer account
	AndroidpublisherScope = "https://www.googleapis.com/auth/androidpublisher"
)

// NewService creates a new Service.
func NewService(ctx context.Context, opts ...option.ClientOption) (*Service, error) {
	scopesOption := option.WithScopes(
		"https://www.googleapis.com/auth/androidpublisher",
	)
	// NOTE: prepend, so we don't override user-specified scopes.
	opts = append([]option.ClientOption{scopesOption}, opts...)
	client, endpoint, err := htransport.NewClient(ctx, opts...)
	if err != nil {
		return nil, err
	}
	s, err := New(client)
	if err != nil {
		return nil, err
	}
	if endpoint != "" {
		s.BasePath = endpoint
	}
	return s, nil
}

// New creates a new Service. It uses the provided http.Client for requests.
//
// Deprecated: please use NewService instead.
// To provide a custom HTTP client, use option.WithHTTPClient.
// If you are using google.golang.org/api/googleapis/transport.APIKey, use option.WithAPIKey with NewService instead.
func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.Accounts = NewAccountsService(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	Accounts *AccountsService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewAccountsService(s *Service) *AccountsService {
	rs := &AccountsService{s: s}
	rs.CustomApps = NewAccountsCustomAppsService(s)
	return rs
}

type AccountsService struct {
	s *Service

	CustomApps *AccountsCustomAppsService
}

func NewAccountsCustomAppsService(s *Service) *AccountsCustomAppsService {
	rs := &AccountsCustomAppsService{s: s}
	return rs
}

type AccountsCustomAppsService struct {
	s *Service
}

// CustomApp: This resource represents a custom app.
type CustomApp struct {
	// LanguageCode: Default listing language in BCP 47 format.
	LanguageCode string `json:"languageCode,omitempty"`

	// Title: Title for the Android app.
	Title string `json:"title,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "LanguageCode") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "LanguageCode") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *CustomApp) MarshalJSON() ([]byte, error) {
	type NoMethod CustomApp
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// method id "playcustomapp.accounts.customApps.create":

type AccountsCustomAppsCreateCall struct {
	s          *Service
	account    int64
	customapp  *CustomApp
	urlParams_ gensupport.URLParams
	mediaInfo_ *gensupport.MediaInfo
	ctx_       context.Context
	header_    http.Header
}

// Create: Create and publish a new custom app.
func (r *AccountsCustomAppsService) Create(account int64, customapp *CustomApp) *AccountsCustomAppsCreateCall {
	c := &AccountsCustomAppsCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.account = account
	c.customapp = customapp
	return c
}

// Media specifies the media to upload in one or more chunks. The chunk
// size may be controlled by supplying a MediaOption generated by
// googleapi.ChunkSize. The chunk size defaults to
// googleapi.DefaultUploadChunkSize.The Content-Type header used in the
// upload request will be determined by sniffing the contents of r,
// unless a MediaOption generated by googleapi.ContentType is
// supplied.
// At most one of Media and ResumableMedia may be set.
func (c *AccountsCustomAppsCreateCall) Media(r io.Reader, options ...googleapi.MediaOption) *AccountsCustomAppsCreateCall {
	c.mediaInfo_ = gensupport.NewInfoFromMedia(r, options)
	return c
}

// ResumableMedia specifies the media to upload in chunks and can be
// canceled with ctx.
//
// Deprecated: use Media instead.
//
// At most one of Media and ResumableMedia may be set. mediaType
// identifies the MIME media type of the upload, such as "image/png". If
// mediaType is "", it will be auto-detected. The provided ctx will
// supersede any context previously provided to the Context method.
func (c *AccountsCustomAppsCreateCall) ResumableMedia(ctx context.Context, r io.ReaderAt, size int64, mediaType string) *AccountsCustomAppsCreateCall {
	c.ctx_ = ctx
	c.mediaInfo_ = gensupport.NewInfoFromResumableMedia(r, size, mediaType)
	return c
}

// ProgressUpdater provides a callback function that will be called
// after every chunk. It should be a low-latency function in order to
// not slow down the upload operation. This should only be called when
// using ResumableMedia (as opposed to Media).
func (c *AccountsCustomAppsCreateCall) ProgressUpdater(pu googleapi.ProgressUpdater) *AccountsCustomAppsCreateCall {
	c.mediaInfo_.SetProgressUpdater(pu)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsCustomAppsCreateCall) Fields(s ...googleapi.Field) *AccountsCustomAppsCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
// This context will supersede any context previously provided to the
// ResumableMedia method.
func (c *AccountsCustomAppsCreateCall) Context(ctx context.Context) *AccountsCustomAppsCreateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *AccountsCustomAppsCreateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *AccountsCustomAppsCreateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("x-goog-api-client", "gl-go/1.11.0 gdcl/20191124")
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.customapp)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "{account}/customApps")
	if c.mediaInfo_ != nil {
		urls = googleapi.ResolveRelative(c.s.BasePath, "/upload/playcustomapp/v1/accounts/{account}/customApps")
		c.urlParams_.Set("uploadType", c.mediaInfo_.UploadType())
	}
	if body == nil {
		body = new(bytes.Buffer)
		reqHeaders.Set("Content-Type", "application/json")
	}
	body, getBody, cleanup := c.mediaInfo_.UploadRequest(reqHeaders, body)
	defer cleanup()
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	req.GetBody = getBody
	googleapi.Expand(req.URL, map[string]string{
		"account": strconv.FormatInt(c.account, 10),
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "playcustomapp.accounts.customApps.create" call.
// Exactly one of *CustomApp or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *CustomApp.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *AccountsCustomAppsCreateCall) Do(opts ...googleapi.CallOption) (*CustomApp, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	rx := c.mediaInfo_.ResumableUpload(res.Header.Get("Location"))
	if rx != nil {
		rx.Client = c.s.client
		rx.UserAgent = c.s.userAgent()
		ctx := c.ctx_
		if ctx == nil {
			ctx = context.TODO()
		}
		res, err = rx.Upload(ctx)
		if err != nil {
			return nil, err
		}
		defer res.Body.Close()
		if err := googleapi.CheckResponse(res); err != nil {
			return nil, err
		}
	}
	ret := &CustomApp{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Create and publish a new custom app.",
	//   "httpMethod": "POST",
	//   "id": "playcustomapp.accounts.customApps.create",
	//   "mediaUpload": {
	//     "accept": [
	//       "*/*"
	//     ],
	//     "maxSize": "100MB",
	//     "protocols": {
	//       "resumable": {
	//         "multipart": true,
	//         "path": "/resumable/upload/playcustomapp/v1/accounts/{account}/customApps"
	//       },
	//       "simple": {
	//         "multipart": true,
	//         "path": "/upload/playcustomapp/v1/accounts/{account}/customApps"
	//       }
	//     }
	//   },
	//   "parameterOrder": [
	//     "account"
	//   ],
	//   "parameters": {
	//     "account": {
	//       "description": "Developer account ID.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{account}/customApps",
	//   "request": {
	//     "$ref": "CustomApp"
	//   },
	//   "response": {
	//     "$ref": "CustomApp"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/androidpublisher"
	//   ],
	//   "supportsMediaUpload": true
	// }

}

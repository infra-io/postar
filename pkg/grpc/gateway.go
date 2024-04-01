// Copyright 2024 FishGoddess. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package grpc

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/infra-io/postar/pkg/grpc/contextutil"
	"google.golang.org/grpc/credentials"
	"google.golang.org/protobuf/encoding/protojson"
)

const (
	headerKeySpaceID    = "X-Postar-Space-Id"
	headerKeySpaceToken = "X-Postar-Space-Token"
	headerKeyTraceID    = "X-Postar-Trace-Id"
)

func matchRequestHeader(key string) (string, bool) {
	if key == headerKeySpaceID {
		return contextutil.MetadataKeySpaceID, true
	}

	if key == headerKeySpaceToken {
		return contextutil.MetadataKeySpaceToken, true
	}

	return key, false
}

func matchResponseHeader(key string) (string, bool) {
	if key == contextutil.MetadataKeyTraceID {
		return headerKeyTraceID, true
	}

	return key, false
}

func NewGatewayMux() *runtime.ServeMux {
	jpb := &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseProtoNames:     true,
			UseEnumNumbers:    true,
			EmitUnpopulated:   true,
			EmitDefaultValues: true,
		},
	}

	mux := runtime.NewServeMux(
		runtime.WithIncomingHeaderMatcher(matchRequestHeader),
		runtime.WithOutgoingHeaderMatcher(matchResponseHeader),
		runtime.WithMarshalerOption(runtime.MIMEWildcard, jpb),
	)

	return mux
}

func jsonifyProto(v any) string {
	jpb := &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseProtoNames:     true,
			EmitUnpopulated:   true,
			EmitDefaultValues: true,
		},
	}

	marshaled, err := jpb.Marshal(v)
	if err != nil {
		return fmt.Sprintf("%+v", v)
	}

	return string(marshaled)
}

func dnsNamesFromCert(certFile string) ([]string, error) {
	certFileBytes, err := os.ReadFile(certFile)
	if err != nil {
		return nil, err
	}

	certPem, _ := pem.Decode(certFileBytes)

	cert, err := x509.ParseCertificate(certPem.Bytes)
	if err != nil {
		return nil, err
	}

	return cert.DNSNames, nil
}

func NewClientTLSFromCert(certFile string) (credentials.TransportCredentials, error) {
	dnsNames, err := dnsNamesFromCert(certFile)
	if err != nil {
		return nil, err
	}

	serverName := ""
	if len(dnsNames) > 0 {
		serverName = dnsNames[0]
	}

	return credentials.NewClientTLSFromFile(certFile, serverName)
}

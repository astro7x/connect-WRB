/*
Copyright (c) Facebook, Inc. and its affiliates.
All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.
*/

package servicers_test

import (
	"crypto/x509"
	"testing"
	"time"

	"magma/orc8r/cloud/go/protos"
	"magma/orc8r/cloud/go/services/certifier/servicers"
	certifier_test_utils "magma/orc8r/cloud/go/services/certifier/test_utils"
	"magma/orc8r/cloud/go/test_utils"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
)

func TestCertifier(t *testing.T) {
	ds := test_utils.NewMockDatastore()
	ctx := context.Background()

	caCert, caKey, err := certifier_test_utils.CreateSignedCertAndPrivKey(
		time.Duration(time.Hour * 24 * 10))
	assert.NoError(t, err)

	// just test with default
	caMap := map[protos.CertType]*servicers.CAInfo{
		protos.CertType_DEFAULT: {caCert, caKey},
	}
	srv, err := servicers.NewCertifierServer(ds, caMap)
	assert.NoError(t, err)

	// sign and add
	csrMsg, err := certifier_test_utils.CreateCSR(time.Duration(time.Hour*24*10), "cn", "cn")
	assert.NoError(t, err)
	certMsg, err := srv.SignAddCertificate(ctx, csrMsg)
	assert.NoError(t, err)

	// get
	certInfoMsg, err := srv.GetIdentity(ctx, certMsg.Sn)
	assert.NoError(t, err)
	assert.True(t, proto.Equal(certInfoMsg.Id, csrMsg.Id))

	// do the same with CSN containing leading zeros (ngnix encoding)
	for i := 0; i < 3; i++ {
		certMsg.Sn.Sn = "0" + certMsg.Sn.Sn
		certInfoMsg, err = srv.GetIdentity(ctx, certMsg.Sn)
		assert.NoError(t, err)
		assert.True(t, proto.Equal(certInfoMsg.Id, csrMsg.Id))
	}

	// revoke
	_, err = srv.RevokeCertificate(ctx, certMsg.Sn)
	assert.NoError(t, err)

	// get should return not found error
	certInfoMsg, err = srv.GetIdentity(ctx, certMsg.Sn)
	assert.Error(t, err)

	// test expiration
	csrMsg, err = certifier_test_utils.CreateCSR(0, "cn", "cn")
	assert.NoError(t, err)
	certMsg, err = srv.SignAddCertificate(ctx, csrMsg)
	assert.NoError(t, err)
	certInfoMsg, err = srv.GetIdentity(ctx, certMsg.Sn)
	assert.Error(t, err)
	_, err = srv.RevokeCertificate(ctx, certMsg.Sn)
	assert.NoError(t, err)

	// test garbage collection
	servicers.CollectGarbageAfter = time.Duration(0)

	for i := 0; i < 3; i++ {
		csrMsg, err = certifier_test_utils.CreateCSR(0, "cn", "cn")
		assert.NoError(t, err)
		_, err = srv.SignAddCertificate(ctx, csrMsg)
		assert.NoError(t, err)
	}
	allSns, _ := ds.ListKeys(servicers.CERTIFICATE_INFO_TABLE)
	assert.Equal(t, 3, len(allSns))
	srv.CollectGarbage(ctx, nil)
	allSns, _ = ds.ListKeys(servicers.CERTIFICATE_INFO_TABLE)
	assert.Equal(t, 0, len(allSns))

	// test csr longer than cert
	csrMsg, err = certifier_test_utils.CreateCSR(time.Duration(time.Hour*24*100), "cn", "cn")
	assert.NoError(t, err)
	certMsg, err = srv.SignAddCertificate(ctx, csrMsg)
	assert.NoError(t, err)
	certInfoMsg, err = srv.GetIdentity(ctx, certMsg.Sn)
	assert.NoError(t, err)
	notAfter, _ := ptypes.Timestamp(certInfoMsg.NotAfter)
	assert.True(t, notAfter.Equal(caCert.NotAfter))

	// test CN mismatch
	csrMsg, err = certifier_test_utils.CreateCSR(time.Duration(time.Hour*1), "cn", "nc")
	assert.NoError(t, err)
	certMsg, err = srv.SignAddCertificate(ctx, csrMsg)
	assert.Error(t, err)

	// test CN onverwrite
	csrMsg, err = certifier_test_utils.CreateCSR(time.Duration(time.Hour*1), "", "cn")
	assert.NoError(t, err)
	certMsg, err = srv.SignAddCertificate(ctx, csrMsg)
	assert.NoError(t, err)
	cert, err := x509.ParseCertificate(certMsg.CertDer)
	assert.NoError(t, err)
	assert.Equal(t, cert.Subject.CommonName, *csrMsg.Id.ToCommonName())
}

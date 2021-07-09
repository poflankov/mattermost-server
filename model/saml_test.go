// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package model

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSamlCertificateStatusJSON(t *testing.T) {
	status := &SamlCertificateStatus{IDpCertificateFile: true, PrivateKeyFile: true, PublicCertificateFile: true}
	json := status.ToJSON()
	rstatus := SamlCertificateStatusFromJSON(strings.NewReader(json))

	require.Equal(t, status.IDpCertificateFile, rstatus.IDpCertificateFile, "IdpCertificateFile do not match")

	rstatus = SamlCertificateStatusFromJSON(strings.NewReader("junk"))
	require.Nil(t, rstatus, "should be nil")
}

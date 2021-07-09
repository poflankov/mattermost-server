// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package model

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTermsOfServiceIsValid(t *testing.T) {
	s := TermsOfService{}

	assert.NotNil(t, s.IsValid(), "should be invalid")

	s.ID = NewID()
	assert.NotNil(t, s.IsValid(), "should be invalid")

	s.CreateAt = GetMillis()
	assert.NotNil(t, s.IsValid(), "should be invalid")

	s.UserID = NewID()
	assert.Nil(t, s.IsValid(), "should be valid")

	s.Text = strings.Repeat("0", PostMessageMaxRunesV2+1)
	assert.NotNil(t, s.IsValid(), "should be invalid")

	s.Text = strings.Repeat("0", PostMessageMaxRunesV2)
	assert.Nil(t, s.IsValid(), "should be valid")

	s.Text = "test"
	assert.Nil(t, s.IsValid(), "should be valid")
}

func TestTermsOfServiceJSON(t *testing.T) {
	o := TermsOfService{
		ID:       NewID(),
		Text:     NewID(),
		CreateAt: GetMillis(),
		UserID:   NewID(),
	}
	j := o.ToJSON()
	ro := TermsOfServiceFromJSON(strings.NewReader(j))

	assert.NotNil(t, ro)
	assert.Equal(t, o, *ro)
}

// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package model

import (
	"encoding/json"
	"io"
)

type Preferences []Preference

func (o *Preferences) ToJSON() string {
	b, _ := json.Marshal(o)
	return string(b)
}

func PreferencesFromJSON(data io.Reader) (Preferences, error) {
	decoder := json.NewDecoder(data)
	var o Preferences
	err := decoder.Decode(&o)
	if err != nil {
		return nil, err
	}
	return o, nil
}

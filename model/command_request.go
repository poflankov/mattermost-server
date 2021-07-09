// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package model

import (
	"encoding/json"
	"io"
)

type CommandMoveRequest struct {
	TeamID string `json:"team_id"`
}

func CommandMoveRequestFromJSON(data io.Reader) (*CommandMoveRequest, error) {
	decoder := json.NewDecoder(data)
	var cmr CommandMoveRequest
	err := decoder.Decode(&cmr)
	if err != nil {
		return nil, err
	}
	return &cmr, nil
}

func (cmr *CommandMoveRequest) ToJSON() string {
	b, err := json.Marshal(cmr)
	if err != nil {
		return ""
	}
	return string(b)
}

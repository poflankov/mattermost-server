// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package model

import (
	"encoding/json"
	"io"
)

type ChannelList []*Channel

func (o *ChannelList) ToJSON() string {
	b, err := json.Marshal(o)
	if err != nil {
		return "[]"
	}
	return string(b)
}

func (o *ChannelList) Etag() string {

	id := "0"
	var t int64 = 0
	var delta int64 = 0

	for _, v := range *o {
		if v.LastPostAt > t {
			t = v.LastPostAt
			id = v.ID
		}

		if v.UpdateAt > t {
			t = v.UpdateAt
			id = v.ID
		}

	}

	return Etag(id, t, delta, len(*o))
}

func ChannelListFromJSON(data io.Reader) *ChannelList {
	var o *ChannelList
	json.NewDecoder(data).Decode(&o)
	return o
}

func ChannelSliceFromJSON(data io.Reader) []*Channel {
	var o []*Channel
	json.NewDecoder(data).Decode(&o)
	return o
}

type ChannelListWithTeamData []*ChannelWithTeamData

func (o *ChannelListWithTeamData) ToJSON() string {
	b, err := json.Marshal(o)
	if err != nil {
		return "[]"
	}
	return string(b)
}

func (o *ChannelListWithTeamData) Etag() string {

	id := "0"
	var t int64 = 0
	var delta int64 = 0

	for _, v := range *o {
		if v.LastPostAt > t {
			t = v.LastPostAt
			id = v.ID
		}

		if v.UpdateAt > t {
			t = v.UpdateAt
			id = v.ID
		}

		if v.TeamUpdateAt > t {
			t = v.TeamUpdateAt
			id = v.ID
		}
	}

	return Etag(id, t, delta, len(*o))
}

func ChannelListWithTeamDataFromJSON(data io.Reader) *ChannelListWithTeamData {
	var o *ChannelListWithTeamData
	json.NewDecoder(data).Decode(&o)
	return o
}

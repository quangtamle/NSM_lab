//go:generate go run ../../../scripts/gendecoder.go -package github.com/skydive-project/skydive/topology/probes/nsm

/*
 * Copyright (C) 2019 Orange
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy ofthe License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specificlanguage governing permissions and
 * limitations under the License.
 *
 */

package nsm

import (
	json "encoding/json"
	"fmt"

	"github.com/skydive-project/skydive/common"
)

// BaseConnectionMetadata holds common connection attributes
// easyjson:json
// gendecoder
type BaseConnectionMetadata struct {
	MechanismType       string
	MechanismParameters map[string]string
	Labels              map[string]string
	Metrics             Metrics
}

//Metrics holds connection metric attributes
// easyjson:json
// gendecoder
type Metrics struct {
	Rx_Bytes   string
	Tx_Bytes   string
	Rx_Packets string
	Tx_Packets string
}

// LocalConnectionMetadata holds local connection attributes
// easyjson:json
// gendecoder
type LocalConnectionMetadata struct {
	BaseConnectionMetadata
	IP string
}

// RemoteConnectionMetadata holds remote connection attributes
// easyjson:json
// gendecoder
type RemoteConnectionMetadata struct {
	BaseConnectionMetadata
	SourceNSM              string
	DestinationNSM         string
	NetworkServiceEndpoint string
}

// BaseNSMMetadata holds common attributes for NSM object
// easyjson:json
// gendecoder
type BaseNSMMetadata struct {
	NetworkService string
	Payload        string
	Source         LocalConnectionMetadata
	Destination    LocalConnectionMetadata
}

// LocalNSMMetadata holds attributes for a local NSM object
// easyjson:json
// gendecoder
type LocalNSMMetadata struct {
	CrossConnectID string
}

// RemoteNSMMetadata holds attributes for a remote NSM object
// easyjson:json
// gendecoder
type RemoteNSMMetadata struct {
	SourceCrossConnectID      string
	DestinationCrossConnectID string
	Via                       RemoteConnectionMetadata
}

// EdgeMetadata describes an NSM edge metadata
// easyjson:json
// gendecoder
type EdgeMetadata struct {
	BaseNSMMetadata
	LocalNSMMetadata
	RemoteNSMMetadata
}

// MetadataDecoder implements a json message raw decoder
func MetadataDecoder(raw json.RawMessage) (common.Getter, error) {
	var m EdgeMetadata
	if err := json.Unmarshal(raw, &m); err != nil {
		return nil, fmt.Errorf("unable to unmarshal NSM object %s: %s", string(raw), err)
	}

	return &m, nil
}

// Code generated - DO NOT EDIT.

package nsm

import (
	"github.com/skydive-project/skydive/common"
	"strings"
)

func (obj *BaseConnectionMetadata) GetFieldBool(key string) (bool, error) {
	return false, common.ErrFieldNotFound
}

func (obj *BaseConnectionMetadata) GetFieldInt64(key string) (int64, error) {
	return 0, common.ErrFieldNotFound
}

func (obj *BaseConnectionMetadata) GetFieldString(key string) (string, error) {
	switch key {
	case "MechanismType":
		return string(obj.MechanismType), nil
	}
	return "", common.ErrFieldNotFound
}

func (obj *BaseConnectionMetadata) GetFieldKeys() []string {
	return []string{
		"MechanismType",
		"MechanismParameters",
		"Labels",
		"Metrics",
	}
}

func (obj *BaseConnectionMetadata) MatchBool(key string, predicate common.BoolPredicate) bool {
	return false
}

func (obj *BaseConnectionMetadata) MatchInt64(key string, predicate common.Int64Predicate) bool {
	return false
}

func (obj *BaseConnectionMetadata) MatchString(key string, predicate common.StringPredicate) bool {
	if b, err := obj.GetFieldString(key); err == nil {
		return predicate(b)
	}
	return false
}

func (obj *BaseConnectionMetadata) GetField(key string) (interface{}, error) {
	if s, err := obj.GetFieldString(key); err == nil {
		return s, nil
	}
	return nil, common.ErrFieldNotFound
}

func (obj *BaseNSMMetadata) GetFieldBool(key string) (bool, error) {
	return false, common.ErrFieldNotFound
}

func (obj *BaseNSMMetadata) GetFieldInt64(key string) (int64, error) {
	return 0, common.ErrFieldNotFound
}

func (obj *BaseNSMMetadata) GetFieldString(key string) (string, error) {
	switch key {
	case "NetworkService":
		return string(obj.NetworkService), nil
	case "Payload":
		return string(obj.Payload), nil
	}
	return "", common.ErrFieldNotFound
}

func (obj *BaseNSMMetadata) GetFieldKeys() []string {
	return []string{
		"NetworkService",
		"Payload",
		"Source",
		"Destination",
	}
}

func (obj *BaseNSMMetadata) MatchBool(key string, predicate common.BoolPredicate) bool {
	return false
}

func (obj *BaseNSMMetadata) MatchInt64(key string, predicate common.Int64Predicate) bool {
	return false
}

func (obj *BaseNSMMetadata) MatchString(key string, predicate common.StringPredicate) bool {
	if b, err := obj.GetFieldString(key); err == nil {
		return predicate(b)
	}
	return false
}

func (obj *BaseNSMMetadata) GetField(key string) (interface{}, error) {
	if s, err := obj.GetFieldString(key); err == nil {
		return s, nil
	}
	return nil, common.ErrFieldNotFound
}

func (obj *EdgeMetadata) GetFieldBool(key string) (bool, error) {
	return false, common.ErrFieldNotFound
}

func (obj *EdgeMetadata) GetFieldInt64(key string) (int64, error) {
	return 0, common.ErrFieldNotFound
}

func (obj *EdgeMetadata) GetFieldString(key string) (string, error) {
	switch key {
	case "NetworkService":
		return string(obj.NetworkService), nil
	case "Payload":
		return string(obj.Payload), nil
	case "CrossConnectID":
		return string(obj.CrossConnectID), nil
	case "SourceCrossConnectID":
		return string(obj.SourceCrossConnectID), nil
	case "DestinationCrossConnectID":
		return string(obj.DestinationCrossConnectID), nil
	}
	return "", common.ErrFieldNotFound
}

func (obj *EdgeMetadata) GetFieldKeys() []string {
	return []string{
		"NetworkService",
		"Payload",
		"Source",
		"Destination",
		"CrossConnectID",
		"SourceCrossConnectID",
		"DestinationCrossConnectID",
		"Via",
	}
}

func (obj *EdgeMetadata) MatchBool(key string, predicate common.BoolPredicate) bool {
	return false
}

func (obj *EdgeMetadata) MatchInt64(key string, predicate common.Int64Predicate) bool {
	return false
}

func (obj *EdgeMetadata) MatchString(key string, predicate common.StringPredicate) bool {
	if b, err := obj.GetFieldString(key); err == nil {
		return predicate(b)
	}
	return false
}

func (obj *EdgeMetadata) GetField(key string) (interface{}, error) {
	if s, err := obj.GetFieldString(key); err == nil {
		return s, nil
	}
	return nil, common.ErrFieldNotFound
}

func (obj *LocalConnectionMetadata) GetFieldBool(key string) (bool, error) {
	return false, common.ErrFieldNotFound
}

func (obj *LocalConnectionMetadata) GetFieldInt64(key string) (int64, error) {
	return 0, common.ErrFieldNotFound
}

func (obj *LocalConnectionMetadata) GetFieldString(key string) (string, error) {
	switch key {
	case "MechanismType":
		return string(obj.MechanismType), nil
	case "IP":
		return string(obj.IP), nil
	}
	return "", common.ErrFieldNotFound
}

func (obj *LocalConnectionMetadata) GetFieldKeys() []string {
	return []string{
		"MechanismType",
		"MechanismParameters",
		"Labels",
		"Metrics",
		"IP",
	}
}

func (obj *LocalConnectionMetadata) MatchBool(key string, predicate common.BoolPredicate) bool {
	return false
}

func (obj *LocalConnectionMetadata) MatchInt64(key string, predicate common.Int64Predicate) bool {
	return false
}

func (obj *LocalConnectionMetadata) MatchString(key string, predicate common.StringPredicate) bool {
	if b, err := obj.GetFieldString(key); err == nil {
		return predicate(b)
	}
	return false
}

func (obj *LocalConnectionMetadata) GetField(key string) (interface{}, error) {
	if s, err := obj.GetFieldString(key); err == nil {
		return s, nil
	}
	return nil, common.ErrFieldNotFound
}

func (obj *LocalNSMMetadata) GetFieldBool(key string) (bool, error) {
	return false, common.ErrFieldNotFound
}

func (obj *LocalNSMMetadata) GetFieldInt64(key string) (int64, error) {
	return 0, common.ErrFieldNotFound
}

func (obj *LocalNSMMetadata) GetFieldString(key string) (string, error) {
	switch key {
	case "CrossConnectID":
		return string(obj.CrossConnectID), nil
	}
	return "", common.ErrFieldNotFound
}

func (obj *LocalNSMMetadata) GetFieldKeys() []string {
	return []string{
		"CrossConnectID",
	}
}

func (obj *LocalNSMMetadata) MatchBool(key string, predicate common.BoolPredicate) bool {
	return false
}

func (obj *LocalNSMMetadata) MatchInt64(key string, predicate common.Int64Predicate) bool {
	return false
}

func (obj *LocalNSMMetadata) MatchString(key string, predicate common.StringPredicate) bool {
	if b, err := obj.GetFieldString(key); err == nil {
		return predicate(b)
	}
	return false
}

func (obj *LocalNSMMetadata) GetField(key string) (interface{}, error) {
	if s, err := obj.GetFieldString(key); err == nil {
		return s, nil
	}
	return nil, common.ErrFieldNotFound
}

func (obj *Metrics) GetFieldBool(key string) (bool, error) {
	return false, common.ErrFieldNotFound
}

func (obj *Metrics) GetFieldInt64(key string) (int64, error) {
	return 0, common.ErrFieldNotFound
}

func (obj *Metrics) GetFieldString(key string) (string, error) {
	switch key {
	case "Rx_Bytes":
		return string(obj.Rx_Bytes), nil
	case "Tx_Bytes":
		return string(obj.Tx_Bytes), nil
	case "Rx_Packets":
		return string(obj.Rx_Packets), nil
	case "Tx_Packets":
		return string(obj.Tx_Packets), nil
	}
	return "", common.ErrFieldNotFound
}

func (obj *Metrics) GetFieldKeys() []string {
	return []string{
		"Rx_Bytes",
		"Tx_Bytes",
		"Rx_Packets",
		"Tx_Packets",
	}
}

func (obj *Metrics) MatchBool(key string, predicate common.BoolPredicate) bool {
	return false
}

func (obj *Metrics) MatchInt64(key string, predicate common.Int64Predicate) bool {
	return false
}

func (obj *Metrics) MatchString(key string, predicate common.StringPredicate) bool {
	if b, err := obj.GetFieldString(key); err == nil {
		return predicate(b)
	}
	return false
}

func (obj *Metrics) GetField(key string) (interface{}, error) {
	if s, err := obj.GetFieldString(key); err == nil {
		return s, nil
	}
	return nil, common.ErrFieldNotFound
}

func (obj *RemoteConnectionMetadata) GetFieldBool(key string) (bool, error) {
	return false, common.ErrFieldNotFound
}

func (obj *RemoteConnectionMetadata) GetFieldInt64(key string) (int64, error) {
	return 0, common.ErrFieldNotFound
}

func (obj *RemoteConnectionMetadata) GetFieldString(key string) (string, error) {
	switch key {
	case "MechanismType":
		return string(obj.MechanismType), nil
	case "SourceNSM":
		return string(obj.SourceNSM), nil
	case "DestinationNSM":
		return string(obj.DestinationNSM), nil
	case "NetworkServiceEndpoint":
		return string(obj.NetworkServiceEndpoint), nil
	}
	return "", common.ErrFieldNotFound
}

func (obj *RemoteConnectionMetadata) GetFieldKeys() []string {
	return []string{
		"MechanismType",
		"MechanismParameters",
		"Labels",
		"Metrics",
		"SourceNSM",
		"DestinationNSM",
		"NetworkServiceEndpoint",
	}
}

func (obj *RemoteConnectionMetadata) MatchBool(key string, predicate common.BoolPredicate) bool {
	return false
}

func (obj *RemoteConnectionMetadata) MatchInt64(key string, predicate common.Int64Predicate) bool {
	return false
}

func (obj *RemoteConnectionMetadata) MatchString(key string, predicate common.StringPredicate) bool {
	if b, err := obj.GetFieldString(key); err == nil {
		return predicate(b)
	}
	return false
}

func (obj *RemoteConnectionMetadata) GetField(key string) (interface{}, error) {
	if s, err := obj.GetFieldString(key); err == nil {
		return s, nil
	}
	return nil, common.ErrFieldNotFound
}

func (obj *RemoteNSMMetadata) GetFieldBool(key string) (bool, error) {
	return false, common.ErrFieldNotFound
}

func (obj *RemoteNSMMetadata) GetFieldInt64(key string) (int64, error) {
	return 0, common.ErrFieldNotFound
}

func (obj *RemoteNSMMetadata) GetFieldString(key string) (string, error) {
	switch key {
	case "SourceCrossConnectID":
		return string(obj.SourceCrossConnectID), nil
	case "DestinationCrossConnectID":
		return string(obj.DestinationCrossConnectID), nil
	}
	return "", common.ErrFieldNotFound
}

func (obj *RemoteNSMMetadata) GetFieldKeys() []string {
	return []string{
		"SourceCrossConnectID",
		"DestinationCrossConnectID",
		"Via",
	}
}

func (obj *RemoteNSMMetadata) MatchBool(key string, predicate common.BoolPredicate) bool {
	return false
}

func (obj *RemoteNSMMetadata) MatchInt64(key string, predicate common.Int64Predicate) bool {
	return false
}

func (obj *RemoteNSMMetadata) MatchString(key string, predicate common.StringPredicate) bool {
	if b, err := obj.GetFieldString(key); err == nil {
		return predicate(b)
	}
	return false
}

func (obj *RemoteNSMMetadata) GetField(key string) (interface{}, error) {
	if s, err := obj.GetFieldString(key); err == nil {
		return s, nil
	}
	return nil, common.ErrFieldNotFound
}

func init() {
	strings.Index("", ".")
}

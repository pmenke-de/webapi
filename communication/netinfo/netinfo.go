// Code generated by webidl-bind. DO NOT EDIT.

// +build !js

package netinfo

import js "github.com/gowebapi/webapi/core/js"

import (
	"github.com/gowebapi/webapi/dom/domcore"
)

// using following types:
// domcore.Event
// domcore.EventHandler
// domcore.EventTarget

// source idl files:
// netinfo.idl

// transform files:
// netinfo.go.md

// workaround for compiler error
func unused(value interface{}) {
	// TODO remove this method
}

type Union struct {
	Value js.Value
}

func (u *Union) JSValue() js.Value {
	return u.Value
}

func UnionFromJS(value js.Value) *Union {
	return &Union{Value: value}
}

// enum: ConnectionType
type ConnectionType int

const (
	BluetoothConnectionType ConnectionType = iota
	CellularConnectionType
	EthernetConnectionType
	MixedConnectionType
	NoneConnectionType
	OtherConnectionType
	UnknownConnectionType
	WifiConnectionType
	WimaxConnectionType
)

var connectionTypeToWasmTable = []string{
	"bluetooth", "cellular", "ethernet", "mixed", "none", "other", "unknown", "wifi", "wimax",
}

var connectionTypeFromWasmTable = map[string]ConnectionType{
	"bluetooth": BluetoothConnectionType, "cellular": CellularConnectionType, "ethernet": EthernetConnectionType, "mixed": MixedConnectionType, "none": NoneConnectionType, "other": OtherConnectionType, "unknown": UnknownConnectionType, "wifi": WifiConnectionType, "wimax": WimaxConnectionType,
}

// JSValue is converting this enum into a javascript object
func (this *ConnectionType) JSValue() js.Value {
	return js.ValueOf(this.Value())
}

// Value is converting this into javascript defined
// string value
func (this ConnectionType) Value() string {
	idx := int(this)
	if idx >= 0 && idx < len(connectionTypeToWasmTable) {
		return connectionTypeToWasmTable[idx]
	}
	panic("unknown input value")
}

// ConnectionTypeFromJS is converting a javascript value into
// a ConnectionType enum value.
func ConnectionTypeFromJS(value js.Value) ConnectionType {
	key := value.String()
	conv, ok := connectionTypeFromWasmTable[key]
	if !ok {
		panic("unable to convert '" + key + "'")
	}
	return conv
}

// enum: EffectiveConnectionType
type EffectiveConnectionType int

const (
	_2gEffectiveConnectionType EffectiveConnectionType = iota
	_3gEffectiveConnectionType
	_4gEffectiveConnectionType
	Slow2gEffectiveConnectionType
)

var effectiveConnectionTypeToWasmTable = []string{
	"2g", "3g", "4g", "slow-2g",
}

var effectiveConnectionTypeFromWasmTable = map[string]EffectiveConnectionType{
	"2g": _2gEffectiveConnectionType, "3g": _3gEffectiveConnectionType, "4g": _4gEffectiveConnectionType, "slow-2g": Slow2gEffectiveConnectionType,
}

// JSValue is converting this enum into a javascript object
func (this *EffectiveConnectionType) JSValue() js.Value {
	return js.ValueOf(this.Value())
}

// Value is converting this into javascript defined
// string value
func (this EffectiveConnectionType) Value() string {
	idx := int(this)
	if idx >= 0 && idx < len(effectiveConnectionTypeToWasmTable) {
		return effectiveConnectionTypeToWasmTable[idx]
	}
	panic("unknown input value")
}

// EffectiveConnectionTypeFromJS is converting a javascript value into
// a EffectiveConnectionType enum value.
func EffectiveConnectionTypeFromJS(value js.Value) EffectiveConnectionType {
	key := value.String()
	conv, ok := effectiveConnectionTypeFromWasmTable[key]
	if !ok {
		panic("unable to convert '" + key + "'")
	}
	return conv
}

// class: NetworkInformation
type NetworkInformation struct {
	domcore.EventTarget
}

// NetworkInformationFromJS is casting a js.Wrapper into NetworkInformation.
func NetworkInformationFromJS(value js.Wrapper) *NetworkInformation {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &NetworkInformation{}
	ret.Value_JS = input
	return ret
}

// Type returning attribute 'type' with
// type ConnectionType (idl: ConnectionType).
func (_this *NetworkInformation) Type() ConnectionType {
	var ret ConnectionType
	value := _this.Value_JS.Get("type")
	ret = ConnectionTypeFromJS(value)
	return ret
}

// EffectiveType returning attribute 'effectiveType' with
// type EffectiveConnectionType (idl: EffectiveConnectionType).
func (_this *NetworkInformation) EffectiveType() EffectiveConnectionType {
	var ret EffectiveConnectionType
	value := _this.Value_JS.Get("effectiveType")
	ret = EffectiveConnectionTypeFromJS(value)
	return ret
}

// DownlinkMax returning attribute 'downlinkMax' with
// type float64 (idl: unrestricted double).
func (_this *NetworkInformation) DownlinkMax() float64 {
	var ret float64
	value := _this.Value_JS.Get("downlinkMax")
	ret = (value).Float()
	return ret
}

// Downlink returning attribute 'downlink' with
// type float64 (idl: unrestricted double).
func (_this *NetworkInformation) Downlink() float64 {
	var ret float64
	value := _this.Value_JS.Get("downlink")
	ret = (value).Float()
	return ret
}

// Rtt returning attribute 'rtt' with
// type int (idl: unsigned long long).
func (_this *NetworkInformation) Rtt() int {
	var ret int
	value := _this.Value_JS.Get("rtt")
	ret = (value).Int()
	return ret
}

// SaveData returning attribute 'saveData' with
// type bool (idl: boolean).
func (_this *NetworkInformation) SaveData() bool {
	var ret bool
	value := _this.Value_JS.Get("saveData")
	ret = (value).Bool()
	return ret
}

// OnChange returning attribute 'onchange' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *NetworkInformation) OnChange() domcore.EventHandlerFunc {
	var ret domcore.EventHandlerFunc
	value := _this.Value_JS.Get("onchange")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		ret = domcore.EventHandlerFromJS(value)
	}
	return ret
}

// event attribute: domcore.Event
func eventFuncNetworkInformation_domcore_Event(listener func(event *domcore.Event, target *NetworkInformation)) js.Func {
	fn := func(this js.Value, args []js.Value) interface{} {
		var ret *domcore.Event
		value := args[0]
		incoming := value.Get("target")
		ret = domcore.EventFromJS(value)
		src := NetworkInformationFromJS(incoming)
		listener(ret, src)
		return js.Undefined
	}
	return js.FuncOf(fn)
}

// AddChange is adding doing AddEventListener for 'Change' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *NetworkInformation) AddEventChange(listener func(event *domcore.Event, currentTarget *NetworkInformation)) js.Func {
	cb := eventFuncNetworkInformation_domcore_Event(listener)
	_this.Value_JS.Call("addEventListener", "change", cb)
	return cb
}

// SetOnChange is assigning a function to 'onchange'. This
// This method is returning allocated javascript function that need to be released.
func (_this *NetworkInformation) SetOnChange(listener func(event *domcore.Event, currentTarget *NetworkInformation)) js.Func {
	cb := eventFuncNetworkInformation_domcore_Event(listener)
	_this.Value_JS.Set("onchange", cb)
	return cb
}

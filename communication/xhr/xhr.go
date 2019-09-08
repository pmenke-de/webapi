// Code generated by webidl-bind. DO NOT EDIT.

// +build !js

package xhr

import js "github.com/gowebapi/webapi/core/js"

import (
	"github.com/gowebapi/webapi/dom/domcore"
	"github.com/gowebapi/webapi/patch"
)

// using following types:
// domcore.Event
// domcore.EventHandler
// domcore.EventTarget
// patch.ByteString

// source idl files:
// xhr.idl

// transform files:
// xhr.go.md

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

// enum: XMLHttpRequestResponseType
type XMLHttpRequestResponseType int

const (
	EmptyString0XMLHttpRequestResponseType XMLHttpRequestResponseType = iota
	ArraybufferXMLHttpRequestResponseType
	BlobXMLHttpRequestResponseType
	DocumentXMLHttpRequestResponseType
	JsonXMLHttpRequestResponseType
	TextXMLHttpRequestResponseType
)

var xMLHttpRequestResponseTypeToWasmTable = []string{
	"", "arraybuffer", "blob", "document", "json", "text",
}

var xMLHttpRequestResponseTypeFromWasmTable = map[string]XMLHttpRequestResponseType{
	"": EmptyString0XMLHttpRequestResponseType, "arraybuffer": ArraybufferXMLHttpRequestResponseType, "blob": BlobXMLHttpRequestResponseType, "document": DocumentXMLHttpRequestResponseType, "json": JsonXMLHttpRequestResponseType, "text": TextXMLHttpRequestResponseType,
}

// JSValue is converting this enum into a javascript object
func (this *XMLHttpRequestResponseType) JSValue() js.Value {
	return js.ValueOf(this.Value())
}

// Value is converting this into javascript defined
// string value
func (this XMLHttpRequestResponseType) Value() string {
	idx := int(this)
	if idx >= 0 && idx < len(xMLHttpRequestResponseTypeToWasmTable) {
		return xMLHttpRequestResponseTypeToWasmTable[idx]
	}
	panic("unknown input value")
}

// XMLHttpRequestResponseTypeFromJS is converting a javascript value into
// a XMLHttpRequestResponseType enum value.
func XMLHttpRequestResponseTypeFromJS(value js.Value) XMLHttpRequestResponseType {
	key := value.String()
	conv, ok := xMLHttpRequestResponseTypeFromWasmTable[key]
	if !ok {
		panic("unable to convert '" + key + "'")
	}
	return conv
}

// dictionary: ProgressEventInit
type ProgressEventInit struct {
	Bubbles          bool
	Cancelable       bool
	Composed         bool
	LengthComputable bool
	Loaded           int
	Total            int
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *ProgressEventInit) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Bubbles
	out.Set("bubbles", value0)
	value1 := _this.Cancelable
	out.Set("cancelable", value1)
	value2 := _this.Composed
	out.Set("composed", value2)
	value3 := _this.LengthComputable
	out.Set("lengthComputable", value3)
	value4 := _this.Loaded
	out.Set("loaded", value4)
	value5 := _this.Total
	out.Set("total", value5)
	return out
}

// ProgressEventInitFromJS is allocating a new
// ProgressEventInit object and copy all values from
// input javascript object
func ProgressEventInitFromJS(value js.Wrapper) *ProgressEventInit {
	input := value.JSValue()
	var out ProgressEventInit
	var (
		value0 bool // javascript: boolean {bubbles Bubbles bubbles}
		value1 bool // javascript: boolean {cancelable Cancelable cancelable}
		value2 bool // javascript: boolean {composed Composed composed}
		value3 bool // javascript: boolean {lengthComputable LengthComputable lengthComputable}
		value4 int  // javascript: unsigned long long {loaded Loaded loaded}
		value5 int  // javascript: unsigned long long {total Total total}
	)
	value0 = (input.Get("bubbles")).Bool()
	out.Bubbles = value0
	value1 = (input.Get("cancelable")).Bool()
	out.Cancelable = value1
	value2 = (input.Get("composed")).Bool()
	out.Composed = value2
	value3 = (input.Get("lengthComputable")).Bool()
	out.LengthComputable = value3
	value4 = (input.Get("loaded")).Int()
	out.Loaded = value4
	value5 = (input.Get("total")).Int()
	out.Total = value5
	return &out
}

// class: ProgressEvent
type ProgressEvent struct {
	domcore.Event
}

// ProgressEventFromJS is casting a js.Wrapper into ProgressEvent.
func ProgressEventFromJS(value js.Wrapper) *ProgressEvent {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &ProgressEvent{}
	ret.Value_JS = input
	return ret
}

func NewProgressEvent(_type string, eventInitDict *ProgressEventInit) (_result *ProgressEvent) {
	_klass := js.Global().Get("ProgressEvent")
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := _type
	_args[0] = _p0
	_end++
	if eventInitDict != nil {
		_p1 := eventInitDict.JSValue()
		_args[1] = _p1
		_end++
	}
	_returned := _klass.New(_args[0:_end]...)
	var (
		_converted *ProgressEvent // javascript: ProgressEvent _what_return_name
	)
	_converted = ProgressEventFromJS(_returned)
	_result = _converted
	return
}

// LengthComputable returning attribute 'lengthComputable' with
// type bool (idl: boolean).
func (_this *ProgressEvent) LengthComputable() bool {
	var ret bool
	value := _this.Value_JS.Get("lengthComputable")
	ret = (value).Bool()
	return ret
}

// Loaded returning attribute 'loaded' with
// type int (idl: unsigned long long).
func (_this *ProgressEvent) Loaded() int {
	var ret int
	value := _this.Value_JS.Get("loaded")
	ret = (value).Int()
	return ret
}

// Total returning attribute 'total' with
// type int (idl: unsigned long long).
func (_this *ProgressEvent) Total() int {
	var ret int
	value := _this.Value_JS.Get("total")
	ret = (value).Int()
	return ret
}

// class: XMLHttpRequest
type XMLHttpRequest struct {
	XMLHttpRequestEventTarget
}

// XMLHttpRequestFromJS is casting a js.Wrapper into XMLHttpRequest.
func XMLHttpRequestFromJS(value js.Wrapper) *XMLHttpRequest {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &XMLHttpRequest{}
	ret.Value_JS = input
	return ret
}

const (
	UNSENT           int = 0
	OPENED           int = 1
	HEADERS_RECEIVED int = 2
	LOADING          int = 3
	DONE             int = 4
)

func NewXMLHttpRequest() (_result *XMLHttpRequest) {
	_klass := js.Global().Get("XMLHttpRequest")
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _klass.New(_args[0:_end]...)
	var (
		_converted *XMLHttpRequest // javascript: XMLHttpRequest _what_return_name
	)
	_converted = XMLHttpRequestFromJS(_returned)
	_result = _converted
	return
}

// OnReadyStateChange returning attribute 'onreadystatechange' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *XMLHttpRequest) OnReadyStateChange() domcore.EventHandlerFunc {
	var ret domcore.EventHandlerFunc
	value := _this.Value_JS.Get("onreadystatechange")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		ret = domcore.EventHandlerFromJS(value)
	}
	return ret
}

// ReadyState returning attribute 'readyState' with
// type int (idl: unsigned short).
func (_this *XMLHttpRequest) ReadyState() int {
	var ret int
	value := _this.Value_JS.Get("readyState")
	ret = (value).Int()
	return ret
}

// Timeout returning attribute 'timeout' with
// type uint (idl: unsigned long).
func (_this *XMLHttpRequest) Timeout() uint {
	var ret uint
	value := _this.Value_JS.Get("timeout")
	ret = (uint)((value).Int())
	return ret
}

// SetTimeout setting attribute 'timeout' with
// type uint (idl: unsigned long).
func (_this *XMLHttpRequest) SetTimeout(value uint) {
	input := value
	_this.Value_JS.Set("timeout", input)
}

// WithCredentials returning attribute 'withCredentials' with
// type bool (idl: boolean).
func (_this *XMLHttpRequest) WithCredentials() bool {
	var ret bool
	value := _this.Value_JS.Get("withCredentials")
	ret = (value).Bool()
	return ret
}

// SetWithCredentials setting attribute 'withCredentials' with
// type bool (idl: boolean).
func (_this *XMLHttpRequest) SetWithCredentials(value bool) {
	input := value
	_this.Value_JS.Set("withCredentials", input)
}

// Upload returning attribute 'upload' with
// type XMLHttpRequestUpload (idl: XMLHttpRequestUpload).
func (_this *XMLHttpRequest) Upload() *XMLHttpRequestUpload {
	var ret *XMLHttpRequestUpload
	value := _this.Value_JS.Get("upload")
	ret = XMLHttpRequestUploadFromJS(value)
	return ret
}

// ResponseURL returning attribute 'responseURL' with
// type string (idl: USVString).
func (_this *XMLHttpRequest) ResponseURL() string {
	var ret string
	value := _this.Value_JS.Get("responseURL")
	ret = (value).String()
	return ret
}

// Status returning attribute 'status' with
// type int (idl: unsigned short).
func (_this *XMLHttpRequest) Status() int {
	var ret int
	value := _this.Value_JS.Get("status")
	ret = (value).Int()
	return ret
}

// StatusText returning attribute 'statusText' with
// type patch.ByteString (idl: ByteString).
func (_this *XMLHttpRequest) StatusText() *patch.ByteString {
	var ret *patch.ByteString
	value := _this.Value_JS.Get("statusText")
	ret = patch.ByteStringFromJS(value)
	return ret
}

// ResponseType returning attribute 'responseType' with
// type XMLHttpRequestResponseType (idl: XMLHttpRequestResponseType).
func (_this *XMLHttpRequest) ResponseType() XMLHttpRequestResponseType {
	var ret XMLHttpRequestResponseType
	value := _this.Value_JS.Get("responseType")
	ret = XMLHttpRequestResponseTypeFromJS(value)
	return ret
}

// SetResponseType setting attribute 'responseType' with
// type XMLHttpRequestResponseType (idl: XMLHttpRequestResponseType).
func (_this *XMLHttpRequest) SetResponseType(value XMLHttpRequestResponseType) {
	input := value.JSValue()
	_this.Value_JS.Set("responseType", input)
}

// Response returning attribute 'response' with
// type Any (idl: any).
func (_this *XMLHttpRequest) Response() js.Value {
	var ret js.Value
	value := _this.Value_JS.Get("response")
	ret = value
	return ret
}

// ResponseText returning attribute 'responseText' with
// type string (idl: USVString).
func (_this *XMLHttpRequest) ResponseText() string {
	var ret string
	value := _this.Value_JS.Get("responseText")
	ret = (value).String()
	return ret
}

// ResponseXML returning attribute 'responseXML' with
// type js.Value (idl: Document).
func (_this *XMLHttpRequest) ResponseXML() js.Value {
	var ret js.Value
	value := _this.Value_JS.Get("responseXML")
	ret = value
	return ret
}

// event attribute: domcore.Event
func eventFuncXMLHttpRequest_domcore_Event(listener func(event *domcore.Event, target *XMLHttpRequest)) js.Func {
	fn := func(this js.Value, args []js.Value) interface{} {
		var ret *domcore.Event
		value := args[0]
		incoming := value.Get("target")
		ret = domcore.EventFromJS(value)
		src := XMLHttpRequestFromJS(incoming)
		listener(ret, src)
		return js.Undefined
	}
	return js.FuncOf(fn)
}

// AddReadyStateChange is adding doing AddEventListener for 'ReadyStateChange' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *XMLHttpRequest) AddEventReadyStateChange(listener func(event *domcore.Event, currentTarget *XMLHttpRequest)) js.Func {
	cb := eventFuncXMLHttpRequest_domcore_Event(listener)
	_this.Value_JS.Call("addEventListener", "readystatechange", cb)
	return cb
}

// SetOnReadyStateChange is assigning a function to 'onreadystatechange'. This
// This method is returning allocated javascript function that need to be released.
func (_this *XMLHttpRequest) SetOnReadyStateChange(listener func(event *domcore.Event, currentTarget *XMLHttpRequest)) js.Func {
	cb := eventFuncXMLHttpRequest_domcore_Event(listener)
	_this.Value_JS.Set("onreadystatechange", cb)
	return cb
}

func (_this *XMLHttpRequest) Open(method *patch.ByteString, url string) {
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := method.JSValue()
	_args[0] = _p0
	_end++
	_p1 := url
	_args[1] = _p1
	_end++
	_this.Value_JS.Call("open", _args[0:_end]...)
	return
}

func (_this *XMLHttpRequest) Open2(method *patch.ByteString, url string, async bool, username *string, password *string) {
	var (
		_args [5]interface{}
		_end  int
	)
	_p0 := method.JSValue()
	_args[0] = _p0
	_end++
	_p1 := url
	_args[1] = _p1
	_end++
	_p2 := async
	_args[2] = _p2
	_end++
	if username != nil {
		_p3 := username
		_args[3] = _p3
		_end++
	}
	if password != nil {
		_p4 := password
		_args[4] = _p4
		_end++
	}
	_this.Value_JS.Call("open", _args[0:_end]...)
	return
}

func (_this *XMLHttpRequest) SetRequestHeader(name *patch.ByteString, value *patch.ByteString) {
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := name.JSValue()
	_args[0] = _p0
	_end++
	_p1 := value.JSValue()
	_args[1] = _p1
	_end++
	_this.Value_JS.Call("setRequestHeader", _args[0:_end]...)
	return
}

func (_this *XMLHttpRequest) Send(body *Union) {
	var (
		_args [1]interface{}
		_end  int
	)
	if body != nil {
		_p0 := body.JSValue()
		_args[0] = _p0
		_end++
	}
	_this.Value_JS.Call("send", _args[0:_end]...)
	return
}

func (_this *XMLHttpRequest) Abort() {
	var (
		_args [0]interface{}
		_end  int
	)
	_this.Value_JS.Call("abort", _args[0:_end]...)
	return
}

func (_this *XMLHttpRequest) GetResponseHeader(name *patch.ByteString) (_result *patch.ByteString) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := name.JSValue()
	_args[0] = _p0
	_end++
	_returned := _this.Value_JS.Call("getResponseHeader", _args[0:_end]...)
	var (
		_converted *patch.ByteString // javascript: ByteString _what_return_name
	)
	if _returned.Type() != js.TypeNull && _returned.Type() != js.TypeUndefined {
		_converted = patch.ByteStringFromJS(_returned)
	}
	_result = _converted
	return
}

func (_this *XMLHttpRequest) GetAllResponseHeaders() (_result *patch.ByteString) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.Value_JS.Call("getAllResponseHeaders", _args[0:_end]...)
	var (
		_converted *patch.ByteString // javascript: ByteString _what_return_name
	)
	_converted = patch.ByteStringFromJS(_returned)
	_result = _converted
	return
}

func (_this *XMLHttpRequest) OverrideMimeType(mime string) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := mime
	_args[0] = _p0
	_end++
	_this.Value_JS.Call("overrideMimeType", _args[0:_end]...)
	return
}

// class: XMLHttpRequestEventTarget
type XMLHttpRequestEventTarget struct {
	domcore.EventTarget
}

// XMLHttpRequestEventTargetFromJS is casting a js.Wrapper into XMLHttpRequestEventTarget.
func XMLHttpRequestEventTargetFromJS(value js.Wrapper) *XMLHttpRequestEventTarget {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &XMLHttpRequestEventTarget{}
	ret.Value_JS = input
	return ret
}

// OnLoadStart returning attribute 'onloadstart' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *XMLHttpRequestEventTarget) OnLoadStart() domcore.EventHandlerFunc {
	var ret domcore.EventHandlerFunc
	value := _this.Value_JS.Get("onloadstart")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		ret = domcore.EventHandlerFromJS(value)
	}
	return ret
}

// OnProgress returning attribute 'onprogress' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *XMLHttpRequestEventTarget) OnProgress() domcore.EventHandlerFunc {
	var ret domcore.EventHandlerFunc
	value := _this.Value_JS.Get("onprogress")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		ret = domcore.EventHandlerFromJS(value)
	}
	return ret
}

// OnAbort returning attribute 'onabort' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *XMLHttpRequestEventTarget) OnAbort() domcore.EventHandlerFunc {
	var ret domcore.EventHandlerFunc
	value := _this.Value_JS.Get("onabort")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		ret = domcore.EventHandlerFromJS(value)
	}
	return ret
}

// OnError returning attribute 'onerror' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *XMLHttpRequestEventTarget) OnError() domcore.EventHandlerFunc {
	var ret domcore.EventHandlerFunc
	value := _this.Value_JS.Get("onerror")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		ret = domcore.EventHandlerFromJS(value)
	}
	return ret
}

// OnLoad returning attribute 'onload' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *XMLHttpRequestEventTarget) OnLoad() domcore.EventHandlerFunc {
	var ret domcore.EventHandlerFunc
	value := _this.Value_JS.Get("onload")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		ret = domcore.EventHandlerFromJS(value)
	}
	return ret
}

// OnTimeOut returning attribute 'ontimeout' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *XMLHttpRequestEventTarget) OnTimeOut() domcore.EventHandlerFunc {
	var ret domcore.EventHandlerFunc
	value := _this.Value_JS.Get("ontimeout")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		ret = domcore.EventHandlerFromJS(value)
	}
	return ret
}

// OnLoadEnd returning attribute 'onloadend' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *XMLHttpRequestEventTarget) OnLoadEnd() domcore.EventHandlerFunc {
	var ret domcore.EventHandlerFunc
	value := _this.Value_JS.Get("onloadend")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		ret = domcore.EventHandlerFromJS(value)
	}
	return ret
}

// event attribute: ProgressEvent
func eventFuncXMLHttpRequestEventTarget_ProgressEvent(listener func(event *ProgressEvent, target *XMLHttpRequestEventTarget)) js.Func {
	fn := func(this js.Value, args []js.Value) interface{} {
		var ret *ProgressEvent
		value := args[0]
		incoming := value.Get("target")
		ret = ProgressEventFromJS(value)
		src := XMLHttpRequestEventTargetFromJS(incoming)
		listener(ret, src)
		return js.Undefined
	}
	return js.FuncOf(fn)
}

// AddAbort is adding doing AddEventListener for 'Abort' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *XMLHttpRequestEventTarget) AddEventAbort(listener func(event *ProgressEvent, currentTarget *XMLHttpRequestEventTarget)) js.Func {
	cb := eventFuncXMLHttpRequestEventTarget_ProgressEvent(listener)
	_this.Value_JS.Call("addEventListener", "abort", cb)
	return cb
}

// SetOnAbort is assigning a function to 'onabort'. This
// This method is returning allocated javascript function that need to be released.
func (_this *XMLHttpRequestEventTarget) SetOnAbort(listener func(event *ProgressEvent, currentTarget *XMLHttpRequestEventTarget)) js.Func {
	cb := eventFuncXMLHttpRequestEventTarget_ProgressEvent(listener)
	_this.Value_JS.Set("onabort", cb)
	return cb
}

// AddError is adding doing AddEventListener for 'Error' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *XMLHttpRequestEventTarget) AddEventError(listener func(event *ProgressEvent, currentTarget *XMLHttpRequestEventTarget)) js.Func {
	cb := eventFuncXMLHttpRequestEventTarget_ProgressEvent(listener)
	_this.Value_JS.Call("addEventListener", "error", cb)
	return cb
}

// SetOnError is assigning a function to 'onerror'. This
// This method is returning allocated javascript function that need to be released.
func (_this *XMLHttpRequestEventTarget) SetOnError(listener func(event *ProgressEvent, currentTarget *XMLHttpRequestEventTarget)) js.Func {
	cb := eventFuncXMLHttpRequestEventTarget_ProgressEvent(listener)
	_this.Value_JS.Set("onerror", cb)
	return cb
}

// AddLoad is adding doing AddEventListener for 'Load' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *XMLHttpRequestEventTarget) AddEventLoad(listener func(event *ProgressEvent, currentTarget *XMLHttpRequestEventTarget)) js.Func {
	cb := eventFuncXMLHttpRequestEventTarget_ProgressEvent(listener)
	_this.Value_JS.Call("addEventListener", "load", cb)
	return cb
}

// SetOnLoad is assigning a function to 'onload'. This
// This method is returning allocated javascript function that need to be released.
func (_this *XMLHttpRequestEventTarget) SetOnLoad(listener func(event *ProgressEvent, currentTarget *XMLHttpRequestEventTarget)) js.Func {
	cb := eventFuncXMLHttpRequestEventTarget_ProgressEvent(listener)
	_this.Value_JS.Set("onload", cb)
	return cb
}

// AddLoadEnd is adding doing AddEventListener for 'LoadEnd' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *XMLHttpRequestEventTarget) AddEventLoadEnd(listener func(event *ProgressEvent, currentTarget *XMLHttpRequestEventTarget)) js.Func {
	cb := eventFuncXMLHttpRequestEventTarget_ProgressEvent(listener)
	_this.Value_JS.Call("addEventListener", "loadend", cb)
	return cb
}

// SetOnLoadEnd is assigning a function to 'onloadend'. This
// This method is returning allocated javascript function that need to be released.
func (_this *XMLHttpRequestEventTarget) SetOnLoadEnd(listener func(event *ProgressEvent, currentTarget *XMLHttpRequestEventTarget)) js.Func {
	cb := eventFuncXMLHttpRequestEventTarget_ProgressEvent(listener)
	_this.Value_JS.Set("onloadend", cb)
	return cb
}

// AddLoadStart is adding doing AddEventListener for 'LoadStart' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *XMLHttpRequestEventTarget) AddEventLoadStart(listener func(event *ProgressEvent, currentTarget *XMLHttpRequestEventTarget)) js.Func {
	cb := eventFuncXMLHttpRequestEventTarget_ProgressEvent(listener)
	_this.Value_JS.Call("addEventListener", "loadstart", cb)
	return cb
}

// SetOnLoadStart is assigning a function to 'onloadstart'. This
// This method is returning allocated javascript function that need to be released.
func (_this *XMLHttpRequestEventTarget) SetOnLoadStart(listener func(event *ProgressEvent, currentTarget *XMLHttpRequestEventTarget)) js.Func {
	cb := eventFuncXMLHttpRequestEventTarget_ProgressEvent(listener)
	_this.Value_JS.Set("onloadstart", cb)
	return cb
}

// AddProgress is adding doing AddEventListener for 'Progress' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *XMLHttpRequestEventTarget) AddEventProgress(listener func(event *ProgressEvent, currentTarget *XMLHttpRequestEventTarget)) js.Func {
	cb := eventFuncXMLHttpRequestEventTarget_ProgressEvent(listener)
	_this.Value_JS.Call("addEventListener", "progress", cb)
	return cb
}

// SetOnProgress is assigning a function to 'onprogress'. This
// This method is returning allocated javascript function that need to be released.
func (_this *XMLHttpRequestEventTarget) SetOnProgress(listener func(event *ProgressEvent, currentTarget *XMLHttpRequestEventTarget)) js.Func {
	cb := eventFuncXMLHttpRequestEventTarget_ProgressEvent(listener)
	_this.Value_JS.Set("onprogress", cb)
	return cb
}

// AddTimeOut is adding doing AddEventListener for 'TimeOut' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *XMLHttpRequestEventTarget) AddEventTimeOut(listener func(event *ProgressEvent, currentTarget *XMLHttpRequestEventTarget)) js.Func {
	cb := eventFuncXMLHttpRequestEventTarget_ProgressEvent(listener)
	_this.Value_JS.Call("addEventListener", "timeout", cb)
	return cb
}

// SetOnTimeOut is assigning a function to 'ontimeout'. This
// This method is returning allocated javascript function that need to be released.
func (_this *XMLHttpRequestEventTarget) SetOnTimeOut(listener func(event *ProgressEvent, currentTarget *XMLHttpRequestEventTarget)) js.Func {
	cb := eventFuncXMLHttpRequestEventTarget_ProgressEvent(listener)
	_this.Value_JS.Set("ontimeout", cb)
	return cb
}

// class: XMLHttpRequestUpload
type XMLHttpRequestUpload struct {
	XMLHttpRequestEventTarget
}

// XMLHttpRequestUploadFromJS is casting a js.Wrapper into XMLHttpRequestUpload.
func XMLHttpRequestUploadFromJS(value js.Wrapper) *XMLHttpRequestUpload {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &XMLHttpRequestUpload{}
	ret.Value_JS = input
	return ret
}

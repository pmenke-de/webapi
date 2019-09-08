// Code generated by webidl-bind. DO NOT EDIT.

// +build !js

package file

import js "github.com/gowebapi/webapi/core/js"

import (
	"github.com/gowebapi/webapi/communication/xhr"
	"github.com/gowebapi/webapi/dom/domcore"
	"github.com/gowebapi/webapi/javascript"
)

// using following types:
// domcore.DOMException
// domcore.EventHandler
// domcore.EventTarget
// javascript.ArrayBuffer
// javascript.PromiseFinally
// xhr.ProgressEvent

// source idl files:
// fileapi.idl
// html.idl
// promises.idl

// transform files:
// fileapi.go.md
// html.go.md
// promises.go.md

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

// enum: EndingType
type EndingType int

const (
	TransparentEndingType EndingType = iota
	NativeEndingType
)

var endingTypeToWasmTable = []string{
	"transparent", "native",
}

var endingTypeFromWasmTable = map[string]EndingType{
	"transparent": TransparentEndingType, "native": NativeEndingType,
}

// JSValue is converting this enum into a javascript object
func (this *EndingType) JSValue() js.Value {
	return js.ValueOf(this.Value())
}

// Value is converting this into javascript defined
// string value
func (this EndingType) Value() string {
	idx := int(this)
	if idx >= 0 && idx < len(endingTypeToWasmTable) {
		return endingTypeToWasmTable[idx]
	}
	panic("unknown input value")
}

// EndingTypeFromJS is converting a javascript value into
// a EndingType enum value.
func EndingTypeFromJS(value js.Value) EndingType {
	key := value.String()
	conv, ok := endingTypeFromWasmTable[key]
	if !ok {
		panic("unable to convert '" + key + "'")
	}
	return conv
}

// callback: BlobCallback
type BlobCallbackFunc func(blob *Blob)

// BlobCallback is a javascript function type.
//
// Call Release() when done to release resouces
// allocated to this type.
type BlobCallback js.Func

func BlobCallbackToJS(callback BlobCallbackFunc) *BlobCallback {
	if callback == nil {
		return nil
	}
	ret := BlobCallback(js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		var (
			_p0 *Blob // javascript: Blob blob
		)
		if args[0].Type() != js.TypeNull && args[0].Type() != js.TypeUndefined {
			_p0 = BlobFromJS(args[0])
		}
		callback(_p0)

		// returning no return value
		return nil
	}))
	return &ret
}

func BlobCallbackFromJS(_value js.Value) BlobCallbackFunc {
	return func(blob *Blob) {
		var (
			_args [1]interface{}
			_end  int
		)
		_p0 := blob.JSValue()
		_args[0] = _p0
		_end++
		_value.Invoke(_args[0:_end]...)
		return
	}
}

// callback: PromiseTemplateOnFulfilled
type PromiseBlobOnFulfilledFunc func(value *Blob)

// PromiseBlobOnFulfilled is a javascript function type.
//
// Call Release() when done to release resouces
// allocated to this type.
type PromiseBlobOnFulfilled js.Func

func PromiseBlobOnFulfilledToJS(callback PromiseBlobOnFulfilledFunc) *PromiseBlobOnFulfilled {
	if callback == nil {
		return nil
	}
	ret := PromiseBlobOnFulfilled(js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		var (
			_p0 *Blob // javascript: Blob value
		)
		_p0 = BlobFromJS(args[0])
		callback(_p0)

		// returning no return value
		return nil
	}))
	return &ret
}

func PromiseBlobOnFulfilledFromJS(_value js.Value) PromiseBlobOnFulfilledFunc {
	return func(value *Blob) {
		var (
			_args [1]interface{}
			_end  int
		)
		_p0 := value.JSValue()
		_args[0] = _p0
		_end++
		_value.Invoke(_args[0:_end]...)
		return
	}
}

// callback: PromiseTemplateOnRejected
type PromiseBlobOnRejectedFunc func(reason js.Value)

// PromiseBlobOnRejected is a javascript function type.
//
// Call Release() when done to release resouces
// allocated to this type.
type PromiseBlobOnRejected js.Func

func PromiseBlobOnRejectedToJS(callback PromiseBlobOnRejectedFunc) *PromiseBlobOnRejected {
	if callback == nil {
		return nil
	}
	ret := PromiseBlobOnRejected(js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		var (
			_p0 js.Value // javascript: any reason
		)
		_p0 = args[0]
		callback(_p0)

		// returning no return value
		return nil
	}))
	return &ret
}

func PromiseBlobOnRejectedFromJS(_value js.Value) PromiseBlobOnRejectedFunc {
	return func(reason js.Value) {
		var (
			_args [1]interface{}
			_end  int
		)
		_p0 := reason
		_args[0] = _p0
		_end++
		_value.Invoke(_args[0:_end]...)
		return
	}
}

// dictionary: BlobPropertyBag
type BlobPropertyBag struct {
	Type    string
	Endings EndingType
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *BlobPropertyBag) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Type
	out.Set("type", value0)
	value1 := _this.Endings.JSValue()
	out.Set("endings", value1)
	return out
}

// BlobPropertyBagFromJS is allocating a new
// BlobPropertyBag object and copy all values from
// input javascript object
func BlobPropertyBagFromJS(value js.Wrapper) *BlobPropertyBag {
	input := value.JSValue()
	var out BlobPropertyBag
	var (
		value0 string     // javascript: DOMString {type Type _type}
		value1 EndingType // javascript: EndingType {endings Endings endings}
	)
	value0 = (input.Get("type")).String()
	out.Type = value0
	value1 = EndingTypeFromJS(input.Get("endings"))
	out.Endings = value1
	return &out
}

// dictionary: FilePropertyBag
type FilePropertyBag struct {
	Type         string
	Endings      EndingType
	LastModified int
}

// JSValue is allocating a new javasript object and copy
// all values
func (_this *FilePropertyBag) JSValue() js.Value {
	out := js.Global().Get("Object").New()
	value0 := _this.Type
	out.Set("type", value0)
	value1 := _this.Endings.JSValue()
	out.Set("endings", value1)
	value2 := _this.LastModified
	out.Set("lastModified", value2)
	return out
}

// FilePropertyBagFromJS is allocating a new
// FilePropertyBag object and copy all values from
// input javascript object
func FilePropertyBagFromJS(value js.Wrapper) *FilePropertyBag {
	input := value.JSValue()
	var out FilePropertyBag
	var (
		value0 string     // javascript: DOMString {type Type _type}
		value1 EndingType // javascript: EndingType {endings Endings endings}
		value2 int        // javascript: long long {lastModified LastModified lastModified}
	)
	value0 = (input.Get("type")).String()
	out.Type = value0
	value1 = EndingTypeFromJS(input.Get("endings"))
	out.Endings = value1
	value2 = (input.Get("lastModified")).Int()
	out.LastModified = value2
	return &out
}

// class: Blob
type Blob struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *Blob) JSValue() js.Value {
	return _this.Value_JS
}

// BlobFromJS is casting a js.Wrapper into Blob.
func BlobFromJS(value js.Wrapper) *Blob {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &Blob{}
	ret.Value_JS = input
	return ret
}

func NewBlob(blobParts []*Union, options *BlobPropertyBag) (_result *Blob) {
	_klass := js.Global().Get("Blob")
	var (
		_args [2]interface{}
		_end  int
	)
	if blobParts != nil {
		_p0 := js.Global().Get("Array").New(len(blobParts))
		for __idx0, __seq_in0 := range blobParts {
			__seq_out0 := __seq_in0.JSValue()
			_p0.SetIndex(__idx0, __seq_out0)
		}
		_args[0] = _p0
		_end++
	}
	if options != nil {
		_p1 := options.JSValue()
		_args[1] = _p1
		_end++
	}
	_returned := _klass.New(_args[0:_end]...)
	var (
		_converted *Blob // javascript: Blob _what_return_name
	)
	_converted = BlobFromJS(_returned)
	_result = _converted
	return
}

// Size returning attribute 'size' with
// type int (idl: unsigned long long).
func (_this *Blob) Size() int {
	var ret int
	value := _this.Value_JS.Get("size")
	ret = (value).Int()
	return ret
}

// Type returning attribute 'type' with
// type string (idl: DOMString).
func (_this *Blob) Type() string {
	var ret string
	value := _this.Value_JS.Get("type")
	ret = (value).String()
	return ret
}

func (_this *Blob) Slice(start *int, end *int, contentType *string) (_result *Blob) {
	var (
		_args [3]interface{}
		_end  int
	)
	if start != nil {
		_p0 := start
		_args[0] = _p0
		_end++
	}
	if end != nil {
		_p1 := end
		_args[1] = _p1
		_end++
	}
	if contentType != nil {
		_p2 := contentType
		_args[2] = _p2
		_end++
	}
	_returned := _this.Value_JS.Call("slice", _args[0:_end]...)
	var (
		_converted *Blob // javascript: Blob _what_return_name
	)
	_converted = BlobFromJS(_returned)
	_result = _converted
	return
}

// class: File
type File struct {
	Blob
}

// FileFromJS is casting a js.Wrapper into File.
func FileFromJS(value js.Wrapper) *File {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &File{}
	ret.Value_JS = input
	return ret
}

func NewFile(fileBits []*Union, fileName string, options *FilePropertyBag) (_result *File) {
	_klass := js.Global().Get("File")
	var (
		_args [3]interface{}
		_end  int
	)
	_p0 := js.Global().Get("Array").New(len(fileBits))
	for __idx0, __seq_in0 := range fileBits {
		__seq_out0 := __seq_in0.JSValue()
		_p0.SetIndex(__idx0, __seq_out0)
	}
	_args[0] = _p0
	_end++
	_p1 := fileName
	_args[1] = _p1
	_end++
	if options != nil {
		_p2 := options.JSValue()
		_args[2] = _p2
		_end++
	}
	_returned := _klass.New(_args[0:_end]...)
	var (
		_converted *File // javascript: File _what_return_name
	)
	_converted = FileFromJS(_returned)
	_result = _converted
	return
}

// Name returning attribute 'name' with
// type string (idl: DOMString).
func (_this *File) Name() string {
	var ret string
	value := _this.Value_JS.Get("name")
	ret = (value).String()
	return ret
}

// LastModified returning attribute 'lastModified' with
// type int (idl: long long).
func (_this *File) LastModified() int {
	var ret int
	value := _this.Value_JS.Get("lastModified")
	ret = (value).Int()
	return ret
}

// WebkitRelativePath returning attribute 'webkitRelativePath' with
// type string (idl: USVString).
func (_this *File) WebkitRelativePath() string {
	var ret string
	value := _this.Value_JS.Get("webkitRelativePath")
	ret = (value).String()
	return ret
}

// class: FileList
type FileList struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *FileList) JSValue() js.Value {
	return _this.Value_JS
}

// FileListFromJS is casting a js.Wrapper into FileList.
func FileListFromJS(value js.Wrapper) *FileList {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &FileList{}
	ret.Value_JS = input
	return ret
}

// Length returning attribute 'length' with
// type uint (idl: unsigned long).
func (_this *FileList) Length() uint {
	var ret uint
	value := _this.Value_JS.Get("length")
	ret = (uint)((value).Int())
	return ret
}

func (_this *FileList) Index(index uint) (_result *File) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := index
	_args[0] = _p0
	_end++
	_returned := _this.Value_JS.Call("item", _args[0:_end]...)
	var (
		_converted *File // javascript: File _what_return_name
	)
	if _returned.Type() != js.TypeNull && _returned.Type() != js.TypeUndefined {
		_converted = FileFromJS(_returned)
	}
	_result = _converted
	return
}

func (_this *FileList) Item(index uint) (_result *File) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := index
	_args[0] = _p0
	_end++
	_returned := _this.Value_JS.Call("item", _args[0:_end]...)
	var (
		_converted *File // javascript: File _what_return_name
	)
	if _returned.Type() != js.TypeNull && _returned.Type() != js.TypeUndefined {
		_converted = FileFromJS(_returned)
	}
	_result = _converted
	return
}

// class: FileReader
type FileReader struct {
	domcore.EventTarget
}

// FileReaderFromJS is casting a js.Wrapper into FileReader.
func FileReaderFromJS(value js.Wrapper) *FileReader {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &FileReader{}
	ret.Value_JS = input
	return ret
}

const (
	EMPTY   int = 0
	LOADING int = 1
	DONE    int = 2
)

func NewFileReader() (_result *FileReader) {
	_klass := js.Global().Get("FileReader")
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _klass.New(_args[0:_end]...)
	var (
		_converted *FileReader // javascript: FileReader _what_return_name
	)
	_converted = FileReaderFromJS(_returned)
	_result = _converted
	return
}

// ReadyState returning attribute 'readyState' with
// type int (idl: unsigned short).
func (_this *FileReader) ReadyState() int {
	var ret int
	value := _this.Value_JS.Get("readyState")
	ret = (value).Int()
	return ret
}

// Result returning attribute 'result' with
// type Union (idl: Union).
func (_this *FileReader) Result() *Union {
	var ret *Union
	value := _this.Value_JS.Get("result")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		ret = UnionFromJS(value)
	}
	return ret
}

// Error returning attribute 'error' with
// type domcore.DOMException (idl: DOMException).
func (_this *FileReader) Error() *domcore.DOMException {
	var ret *domcore.DOMException
	value := _this.Value_JS.Get("error")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		ret = domcore.DOMExceptionFromJS(value)
	}
	return ret
}

// OnLoadStart returning attribute 'onloadstart' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *FileReader) OnLoadStart() domcore.EventHandlerFunc {
	var ret domcore.EventHandlerFunc
	value := _this.Value_JS.Get("onloadstart")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		ret = domcore.EventHandlerFromJS(value)
	}
	return ret
}

// OnProgress returning attribute 'onprogress' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *FileReader) OnProgress() domcore.EventHandlerFunc {
	var ret domcore.EventHandlerFunc
	value := _this.Value_JS.Get("onprogress")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		ret = domcore.EventHandlerFromJS(value)
	}
	return ret
}

// OnLoad returning attribute 'onload' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *FileReader) OnLoad() domcore.EventHandlerFunc {
	var ret domcore.EventHandlerFunc
	value := _this.Value_JS.Get("onload")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		ret = domcore.EventHandlerFromJS(value)
	}
	return ret
}

// OnAbort returning attribute 'onabort' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *FileReader) OnAbort() domcore.EventHandlerFunc {
	var ret domcore.EventHandlerFunc
	value := _this.Value_JS.Get("onabort")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		ret = domcore.EventHandlerFromJS(value)
	}
	return ret
}

// OnError returning attribute 'onerror' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *FileReader) OnError() domcore.EventHandlerFunc {
	var ret domcore.EventHandlerFunc
	value := _this.Value_JS.Get("onerror")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		ret = domcore.EventHandlerFromJS(value)
	}
	return ret
}

// OnLoadEnd returning attribute 'onloadend' with
// type domcore.EventHandler (idl: EventHandlerNonNull).
func (_this *FileReader) OnLoadEnd() domcore.EventHandlerFunc {
	var ret domcore.EventHandlerFunc
	value := _this.Value_JS.Get("onloadend")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		ret = domcore.EventHandlerFromJS(value)
	}
	return ret
}

// event attribute: xhr.ProgressEvent
func eventFuncFileReader_xhr_ProgressEvent(listener func(event *xhr.ProgressEvent, target *FileReader)) js.Func {
	fn := func(this js.Value, args []js.Value) interface{} {
		var ret *xhr.ProgressEvent
		value := args[0]
		incoming := value.Get("target")
		ret = xhr.ProgressEventFromJS(value)
		src := FileReaderFromJS(incoming)
		listener(ret, src)
		return js.Undefined
	}
	return js.FuncOf(fn)
}

// AddAbort is adding doing AddEventListener for 'Abort' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *FileReader) AddEventAbort(listener func(event *xhr.ProgressEvent, currentTarget *FileReader)) js.Func {
	cb := eventFuncFileReader_xhr_ProgressEvent(listener)
	_this.Value_JS.Call("addEventListener", "abort", cb)
	return cb
}

// SetOnAbort is assigning a function to 'onabort'. This
// This method is returning allocated javascript function that need to be released.
func (_this *FileReader) SetOnAbort(listener func(event *xhr.ProgressEvent, currentTarget *FileReader)) js.Func {
	cb := eventFuncFileReader_xhr_ProgressEvent(listener)
	_this.Value_JS.Set("onabort", cb)
	return cb
}

// AddError is adding doing AddEventListener for 'Error' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *FileReader) AddEventError(listener func(event *xhr.ProgressEvent, currentTarget *FileReader)) js.Func {
	cb := eventFuncFileReader_xhr_ProgressEvent(listener)
	_this.Value_JS.Call("addEventListener", "error", cb)
	return cb
}

// SetOnError is assigning a function to 'onerror'. This
// This method is returning allocated javascript function that need to be released.
func (_this *FileReader) SetOnError(listener func(event *xhr.ProgressEvent, currentTarget *FileReader)) js.Func {
	cb := eventFuncFileReader_xhr_ProgressEvent(listener)
	_this.Value_JS.Set("onerror", cb)
	return cb
}

// AddLoad is adding doing AddEventListener for 'Load' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *FileReader) AddEventLoad(listener func(event *xhr.ProgressEvent, currentTarget *FileReader)) js.Func {
	cb := eventFuncFileReader_xhr_ProgressEvent(listener)
	_this.Value_JS.Call("addEventListener", "load", cb)
	return cb
}

// SetOnLoad is assigning a function to 'onload'. This
// This method is returning allocated javascript function that need to be released.
func (_this *FileReader) SetOnLoad(listener func(event *xhr.ProgressEvent, currentTarget *FileReader)) js.Func {
	cb := eventFuncFileReader_xhr_ProgressEvent(listener)
	_this.Value_JS.Set("onload", cb)
	return cb
}

// AddLoadEnd is adding doing AddEventListener for 'LoadEnd' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *FileReader) AddEventLoadEnd(listener func(event *xhr.ProgressEvent, currentTarget *FileReader)) js.Func {
	cb := eventFuncFileReader_xhr_ProgressEvent(listener)
	_this.Value_JS.Call("addEventListener", "loadend", cb)
	return cb
}

// SetOnLoadEnd is assigning a function to 'onloadend'. This
// This method is returning allocated javascript function that need to be released.
func (_this *FileReader) SetOnLoadEnd(listener func(event *xhr.ProgressEvent, currentTarget *FileReader)) js.Func {
	cb := eventFuncFileReader_xhr_ProgressEvent(listener)
	_this.Value_JS.Set("onloadend", cb)
	return cb
}

// AddLoadStart is adding doing AddEventListener for 'LoadStart' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *FileReader) AddEventLoadStart(listener func(event *xhr.ProgressEvent, currentTarget *FileReader)) js.Func {
	cb := eventFuncFileReader_xhr_ProgressEvent(listener)
	_this.Value_JS.Call("addEventListener", "loadstart", cb)
	return cb
}

// SetOnLoadStart is assigning a function to 'onloadstart'. This
// This method is returning allocated javascript function that need to be released.
func (_this *FileReader) SetOnLoadStart(listener func(event *xhr.ProgressEvent, currentTarget *FileReader)) js.Func {
	cb := eventFuncFileReader_xhr_ProgressEvent(listener)
	_this.Value_JS.Set("onloadstart", cb)
	return cb
}

// AddProgress is adding doing AddEventListener for 'Progress' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *FileReader) AddEventProgress(listener func(event *xhr.ProgressEvent, currentTarget *FileReader)) js.Func {
	cb := eventFuncFileReader_xhr_ProgressEvent(listener)
	_this.Value_JS.Call("addEventListener", "progress", cb)
	return cb
}

// SetOnProgress is assigning a function to 'onprogress'. This
// This method is returning allocated javascript function that need to be released.
func (_this *FileReader) SetOnProgress(listener func(event *xhr.ProgressEvent, currentTarget *FileReader)) js.Func {
	cb := eventFuncFileReader_xhr_ProgressEvent(listener)
	_this.Value_JS.Set("onprogress", cb)
	return cb
}

func (_this *FileReader) ReadAsArrayBuffer(blob *Blob) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := blob.JSValue()
	_args[0] = _p0
	_end++
	_this.Value_JS.Call("readAsArrayBuffer", _args[0:_end]...)
	return
}

func (_this *FileReader) ReadAsBinaryString(blob *Blob) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := blob.JSValue()
	_args[0] = _p0
	_end++
	_this.Value_JS.Call("readAsBinaryString", _args[0:_end]...)
	return
}

func (_this *FileReader) ReadAsText(blob *Blob, encoding *string) {
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := blob.JSValue()
	_args[0] = _p0
	_end++
	if encoding != nil {
		_p1 := encoding
		_args[1] = _p1
		_end++
	}
	_this.Value_JS.Call("readAsText", _args[0:_end]...)
	return
}

func (_this *FileReader) ReadAsDataURL(blob *Blob) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := blob.JSValue()
	_args[0] = _p0
	_end++
	_this.Value_JS.Call("readAsDataURL", _args[0:_end]...)
	return
}

func (_this *FileReader) Abort() {
	var (
		_args [0]interface{}
		_end  int
	)
	_this.Value_JS.Call("abort", _args[0:_end]...)
	return
}

// class: FileReaderSync
type FileReaderSync struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *FileReaderSync) JSValue() js.Value {
	return _this.Value_JS
}

// FileReaderSyncFromJS is casting a js.Wrapper into FileReaderSync.
func FileReaderSyncFromJS(value js.Wrapper) *FileReaderSync {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &FileReaderSync{}
	ret.Value_JS = input
	return ret
}

func NewFileReaderSync() (_result *FileReaderSync) {
	_klass := js.Global().Get("FileReaderSync")
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _klass.New(_args[0:_end]...)
	var (
		_converted *FileReaderSync // javascript: FileReaderSync _what_return_name
	)
	_converted = FileReaderSyncFromJS(_returned)
	_result = _converted
	return
}

func (_this *FileReaderSync) ReadAsArrayBuffer(blob *Blob) (_result *javascript.ArrayBuffer) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := blob.JSValue()
	_args[0] = _p0
	_end++
	_returned := _this.Value_JS.Call("readAsArrayBuffer", _args[0:_end]...)
	var (
		_converted *javascript.ArrayBuffer // javascript: ArrayBuffer _what_return_name
	)
	_converted = javascript.ArrayBufferFromJS(_returned)
	_result = _converted
	return
}

func (_this *FileReaderSync) ReadAsBinaryString(blob *Blob) (_result string) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := blob.JSValue()
	_args[0] = _p0
	_end++
	_returned := _this.Value_JS.Call("readAsBinaryString", _args[0:_end]...)
	var (
		_converted string // javascript: DOMString _what_return_name
	)
	_converted = (_returned).String()
	_result = _converted
	return
}

func (_this *FileReaderSync) ReadAsText(blob *Blob, encoding *string) (_result string) {
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := blob.JSValue()
	_args[0] = _p0
	_end++
	if encoding != nil {
		_p1 := encoding
		_args[1] = _p1
		_end++
	}
	_returned := _this.Value_JS.Call("readAsText", _args[0:_end]...)
	var (
		_converted string // javascript: DOMString _what_return_name
	)
	_converted = (_returned).String()
	_result = _converted
	return
}

func (_this *FileReaderSync) ReadAsDataURL(blob *Blob) (_result string) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := blob.JSValue()
	_args[0] = _p0
	_end++
	_returned := _this.Value_JS.Call("readAsDataURL", _args[0:_end]...)
	var (
		_converted string // javascript: DOMString _what_return_name
	)
	_converted = (_returned).String()
	_result = _converted
	return
}

// class: Promise
type PromiseBlob struct {
	// Value_JS holds a reference to a javascript value
	Value_JS js.Value
}

func (_this *PromiseBlob) JSValue() js.Value {
	return _this.Value_JS
}

// PromiseBlobFromJS is casting a js.Wrapper into PromiseBlob.
func PromiseBlobFromJS(value js.Wrapper) *PromiseBlob {
	input := value.JSValue()
	if input.Type() == js.TypeNull {
		return nil
	}
	ret := &PromiseBlob{}
	ret.Value_JS = input
	return ret
}

func (_this *PromiseBlob) Then(onFulfilled *PromiseBlobOnFulfilled, onRejected *PromiseBlobOnRejected) (_result *PromiseBlob) {
	var (
		_args [2]interface{}
		_end  int
	)

	var __callback0 js.Value
	if onFulfilled != nil {
		__callback0 = (*onFulfilled).Value
	} else {
		__callback0 = js.Null()
	}
	_p0 := __callback0
	_args[0] = _p0
	_end++
	if onRejected != nil {

		var __callback1 js.Value
		if onRejected != nil {
			__callback1 = (*onRejected).Value
		} else {
			__callback1 = js.Null()
		}
		_p1 := __callback1
		_args[1] = _p1
		_end++
	}
	_returned := _this.Value_JS.Call("then", _args[0:_end]...)
	var (
		_converted *PromiseBlob // javascript: Promise _what_return_name
	)
	_converted = PromiseBlobFromJS(_returned)
	_result = _converted
	return
}

func (_this *PromiseBlob) Catch(onRejected *PromiseBlobOnRejected) (_result *PromiseBlob) {
	var (
		_args [1]interface{}
		_end  int
	)

	var __callback0 js.Value
	if onRejected != nil {
		__callback0 = (*onRejected).Value
	} else {
		__callback0 = js.Null()
	}
	_p0 := __callback0
	_args[0] = _p0
	_end++
	_returned := _this.Value_JS.Call("catch", _args[0:_end]...)
	var (
		_converted *PromiseBlob // javascript: Promise _what_return_name
	)
	_converted = PromiseBlobFromJS(_returned)
	_result = _converted
	return
}

func (_this *PromiseBlob) Finally(onFinally *javascript.PromiseFinally) (_result *PromiseBlob) {
	var (
		_args [1]interface{}
		_end  int
	)

	var __callback0 js.Value
	if onFinally != nil {
		__callback0 = (*onFinally).Value
	} else {
		__callback0 = js.Null()
	}
	_p0 := __callback0
	_args[0] = _p0
	_end++
	_returned := _this.Value_JS.Call("finally", _args[0:_end]...)
	var (
		_converted *PromiseBlob // javascript: Promise _what_return_name
	)
	_converted = PromiseBlobFromJS(_returned)
	_result = _converted
	return
}

// callbacks
package host

type SurfaceCallbackFunc func(width, height int32)
type VoidCallbackFunc func()
type PositionCallbackFunc func(x, y int32)
type KeyCallbackFunc func(k uint16)

var InitFunc SurfaceCallbackFunc = nil
var DestroyFunc VoidCallbackFunc = nil
var ResizeFunc SurfaceCallbackFunc = nil
var DrawFunc SurfaceCallbackFunc = nil
var MouseLeftButtonDownFunc PositionCallbackFunc = nil
var MouseLeftButtonUpFunc PositionCallbackFunc = nil
var MouseRightButtonDownFunc PositionCallbackFunc = nil
var MouseRightButtonUpFunc PositionCallbackFunc = nil
var MouseMoveFunc PositionCallbackFunc = nil
var KeyDownFunc KeyCallbackFunc = nil
var KeyUpFunc KeyCallbackFunc = nil

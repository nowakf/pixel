package pixelgl

import (
	"github.com/faiface/mainthread"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/nowakf/pixel"
)

// Pressed returns whether the Key is currently pressed down.
type ModKey int

const (
	ModShift   = ModKey(glfw.ModShift)
	ModControl = ModKey(glfw.ModControl)
	ModAlt     = ModKey(glfw.ModAlt)
	ModSuper   = ModKey(glfw.ModSuper)
	ModNone    = ModKey(KeyNone)
)

// Key is a keyboard or mouse button. Why distinguish?
type Key int

// List of all mouse buttons.
const (
	MouseButton1      = Key(glfw.MouseButton1)
	MouseButton2      = Key(glfw.MouseButton2)
	MouseButton3      = Key(glfw.MouseButton3)
	MouseButton4      = Key(glfw.MouseButton4)
	MouseButton5      = Key(glfw.MouseButton5)
	MouseButton6      = Key(glfw.MouseButton6)
	MouseButton7      = Key(glfw.MouseButton7)
	MouseButton8      = Key(glfw.MouseButton8)
	MouseButtonLast   = Key(glfw.MouseButtonLast)
	MouseButtonLeft   = Key(glfw.MouseButtonLeft)
	MouseButtonRight  = Key(glfw.MouseButtonRight)
	MouseButtonMiddle = Key(glfw.MouseButtonMiddle)
)

// List of all keyboard buttons.
const (
	KeyUnknown      = Key(glfw.KeyUnknown)
	KeyRune         = Key(glfw.KeyF24)
	KeyWorld1       = Key(glfw.KeyWorld1)
	KeyWorld2       = Key(glfw.KeyWorld2)
	KeyEscape       = Key(glfw.KeyEscape)
	KeyEnter        = Key(glfw.KeyEnter)
	KeyTab          = Key(glfw.KeyTab)
	KeyBackspace    = Key(glfw.KeyBackspace)
	KeyInsert       = Key(glfw.KeyInsert)
	KeyDelete       = Key(glfw.KeyDelete)
	KeyRight        = Key(glfw.KeyRight)
	KeyLeft         = Key(glfw.KeyLeft)
	KeyDown         = Key(glfw.KeyDown)
	KeyUp           = Key(glfw.KeyUp)
	KeyPageUp       = Key(glfw.KeyPageUp)
	KeyPageDown     = Key(glfw.KeyPageDown)
	KeyHome         = Key(glfw.KeyHome)
	KeyEnd          = Key(glfw.KeyEnd)
	KeyCapsLock     = Key(glfw.KeyCapsLock)
	KeyScrollLock   = Key(glfw.KeyScrollLock)
	KeyNumLock      = Key(glfw.KeyNumLock)
	KeyPrintScreen  = Key(glfw.KeyPrintScreen)
	KeyPause        = Key(glfw.KeyPause)
	KeyF1           = Key(glfw.KeyF1)
	KeyF2           = Key(glfw.KeyF2)
	KeyF3           = Key(glfw.KeyF3)
	KeyF4           = Key(glfw.KeyF4)
	KeyF5           = Key(glfw.KeyF5)
	KeyF6           = Key(glfw.KeyF6)
	KeyF7           = Key(glfw.KeyF7)
	KeyF8           = Key(glfw.KeyF8)
	KeyF9           = Key(glfw.KeyF9)
	KeyF10          = Key(glfw.KeyF10)
	KeyF11          = Key(glfw.KeyF11)
	KeyF12          = Key(glfw.KeyF12)
	KeyF13          = Key(glfw.KeyF13)
	KeyF14          = Key(glfw.KeyF14)
	KeyF15          = Key(glfw.KeyF15)
	KeyF16          = Key(glfw.KeyF16)
	KeyF17          = Key(glfw.KeyF17)
	KeyF18          = Key(glfw.KeyF18)
	KeyF19          = Key(glfw.KeyF19)
	KeyF20          = Key(glfw.KeyF20)
	KeyF21          = Key(glfw.KeyF21)
	KeyF22          = Key(glfw.KeyF22)
	KeyKP0          = Key(glfw.KeyKP0)
	KeyKP1          = Key(glfw.KeyKP1)
	KeyKP2          = Key(glfw.KeyKP2)
	KeyKP3          = Key(glfw.KeyKP3)
	KeyKP4          = Key(glfw.KeyKP4)
	KeyKP5          = Key(glfw.KeyKP5)
	KeyKP6          = Key(glfw.KeyKP6)
	KeyKP7          = Key(glfw.KeyKP7)
	KeyKP8          = Key(glfw.KeyKP8)
	KeyKP9          = Key(glfw.KeyKP9)
	KeyKPDecimal    = Key(glfw.KeyKPDecimal)
	KeyKPDivide     = Key(glfw.KeyKPDivide)
	KeyKPMultiply   = Key(glfw.KeyKPMultiply)
	KeyKPSubtract   = Key(glfw.KeyKPSubtract)
	KeyKPAdd        = Key(glfw.KeyKPAdd)
	KeyKPEnter      = Key(glfw.KeyKPEnter)
	KeyKPEqual      = Key(glfw.KeyKPEqual)
	KeyLeftShift    = Key(glfw.KeyLeftShift)
	KeyLeftControl  = Key(glfw.KeyLeftControl)
	KeyLeftAlt      = Key(glfw.KeyLeftAlt)
	KeyLeftSuper    = Key(glfw.KeyLeftSuper)
	KeyRightShift   = Key(glfw.KeyRightShift)
	KeyRightControl = Key(glfw.KeyRightControl)
	KeyRightAlt     = Key(glfw.KeyRightAlt)
	KeyRightSuper   = Key(glfw.KeyRightSuper)
	KeyMenu         = Key(glfw.KeyMenu)
	KeyLast         = Key(glfw.KeyLast)

	KeyNone   = Key(glfw.KeyF25)
	KeyString = Key(glfw.KeyF24)
)

// String returns a human-readable string describing the Key.
func (b Key) String() string {
	name, ok := buttonNames[b]
	if !ok {
		return "Invalid"
	}
	return name
}

var buttonNames = map[Key]string{
	MouseButton4:      "MouseButton4",
	MouseButton5:      "MouseButton5",
	MouseButton6:      "MouseButton6",
	MouseButton7:      "MouseButton7",
	MouseButton8:      "MouseButton8",
	MouseButtonLeft:   "MouseButtonLeft",
	MouseButtonRight:  "MouseButtonRight",
	MouseButtonMiddle: "MouseButtonMiddle",
	KeyWorld1:         "World1",
	KeyWorld2:         "World2",
	KeyEscape:         "Escape",
	KeyEnter:          "Enter",
	KeyTab:            "Tab",
	KeyBackspace:      "Backspace",
	KeyInsert:         "Insert",
	KeyDelete:         "Delete",
	KeyRight:          "Right",
	KeyLeft:           "Left",
	KeyDown:           "Down",
	KeyUp:             "Up",
	KeyPageUp:         "PageUp",
	KeyPageDown:       "PageDown",
	KeyHome:           "Home",
	KeyEnd:            "End",
	KeyCapsLock:       "CapsLock",
	KeyScrollLock:     "ScrollLock",
	KeyNumLock:        "NumLock",
	KeyPrintScreen:    "PrintScreen",
	KeyPause:          "Pause",
	KeyF1:             "F1",
	KeyF2:             "F2",
	KeyF3:             "F3",
	KeyF4:             "F4",
	KeyF5:             "F5",
	KeyF6:             "F6",
	KeyF7:             "F7",
	KeyF8:             "F8",
	KeyF9:             "F9",
	KeyF10:            "F10",
	KeyF11:            "F11",
	KeyF12:            "F12",
	KeyF13:            "F13",
	KeyF14:            "F14",
	KeyF15:            "F15",
	KeyF16:            "F16",
	KeyF17:            "F17",
	KeyF18:            "F18",
	KeyF19:            "F19",
	KeyF20:            "F20",
	KeyF21:            "F21",
	KeyF22:            "F22",
	KeyKP0:            "KP0",
	KeyKP1:            "KP1",
	KeyKP2:            "KP2",
	KeyKP3:            "KP3",
	KeyKP4:            "KP4",
	KeyKP5:            "KP5",
	KeyKP6:            "KP6",
	KeyKP7:            "KP7",
	KeyKP8:            "KP8",
	KeyKP9:            "KP9",
	KeyKPDecimal:      "KPDecimal",
	KeyKPDivide:       "KPDivide",
	KeyKPMultiply:     "KPMultiply",
	KeyKPSubtract:     "KPSubtract",
	KeyKPAdd:          "KPAdd",
	KeyKPEnter:        "KPEnter",
	KeyKPEqual:        "KPEqual",
	KeyLeftShift:      "LeftShift",
	KeyLeftControl:    "LeftControl",
	KeyLeftAlt:        "LeftAlt",
	KeyLeftSuper:      "LeftSuper",
	KeyRightShift:     "RightShift",
	KeyRightControl:   "RightControl",
	KeyRightAlt:       "RightAlt",
	KeyRightSuper:     "RightSuper",
	KeyMenu:           "Menu",
}

func (w *Window) initInput() {
	mainthread.Call(func() {
		w.window.SetMouseButtonCallback(func(_ *glfw.Window, button glfw.MouseButton, action glfw.Action, mod glfw.ModifierKey) {
			w.tempInp.key = KeyEvent{
				Action(action),
				Key(button),
				ModKey(mod),
				rune(0),
			}
			w.tempInp.event = true

		})

		w.window.SetKeyCallback(func(_ *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
			if key == glfw.KeyUnknown {
				return
			}
			w.tempInp.key = KeyEvent{
				Action(action),
				Key(key),
				ModKey(mods),
				rune(0),
			}
			w.tempInp.event = true
		})

		w.window.SetCursorPosCallback(func(_ *glfw.Window, x, y float64) {
			w.tempInp.mouse = pixel.V(
				x+w.bounds.Min.X,
				(w.bounds.H()-y)+w.bounds.Min.Y,
			)
			w.tempInp.event = true
		})

		w.window.SetScrollCallback(func(_ *glfw.Window, xoff, yoff float64) {
			w.tempInp.scroll.X += xoff
			w.tempInp.scroll.Y += yoff
			w.tempInp.event = true
		})

		w.window.SetCharCallback(func(_ *glfw.Window, r rune) {
			w.tempInp.key = KeyEvent{
				PRESS,
				KeyRune,
				ModNone,
				r,
			}
			w.tempInp.typed += string(r)
			w.tempInp.event = true
		})
	})
}
func (w *Window) MousePosition() pixel.Vec {
	return w.currInp.mouse
}

// UpdateInput polls window events. Call this function to poll window events
// without swapping buffers. Note that the Update method invokes UpdateInput.
func (w *Window) UpdateInput() {

	mainthread.Call(func() {
		glfw.WaitEvents()
	})

	w.currInp = w.tempInp
	w.tempInp.event = false
	w.tempInp.typed = ""
	w.tempInp.scroll = pixel.ZV
}

func (w *Window) PostEmpty() {
	glfw.PostEmptyEvent()
}

func (w *Window) Key() *KeyEvent {
	if !w.currInp.event {
		return nil
	}
	return &w.currInp.key
}
func (w *Window) Typed() string {
	return w.currInp.typed
}

type KeyEvent struct {
	Action
	Key
	ModKey
	Rune rune
}

type CursorEvent pixel.Vec

type ScrollEvent pixel.Vec

type Action int

const (
	PRESS   = Action(glfw.Press)
	RELEASE = Action(glfw.Release)
	REPEAT  = Action(glfw.Repeat)
)

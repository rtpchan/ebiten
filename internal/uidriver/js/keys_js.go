// Copyright 2013 The Ebiten Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by genkeys.go using 'go generate'. DO NOT EDIT.

package js

import (
	"syscall/js"

	"github.com/hajimehoshi/ebiten/v2/internal/driver"
)

var driverKeyToJSKey = map[driver.Key]js.Value{
	driver.KeyA:              js.ValueOf("KeyA"),
	driver.KeyAltLeft:        js.ValueOf("AltLeft"),
	driver.KeyAltRight:       js.ValueOf("AltRight"),
	driver.KeyArrowDown:      js.ValueOf("ArrowDown"),
	driver.KeyArrowLeft:      js.ValueOf("ArrowLeft"),
	driver.KeyArrowRight:     js.ValueOf("ArrowRight"),
	driver.KeyArrowUp:        js.ValueOf("ArrowUp"),
	driver.KeyB:              js.ValueOf("KeyB"),
	driver.KeyBackquote:      js.ValueOf("Backquote"),
	driver.KeyBackslash:      js.ValueOf("Backslash"),
	driver.KeyBackspace:      js.ValueOf("Backspace"),
	driver.KeyBracketLeft:    js.ValueOf("BracketLeft"),
	driver.KeyBracketRight:   js.ValueOf("BracketRight"),
	driver.KeyC:              js.ValueOf("KeyC"),
	driver.KeyCapsLock:       js.ValueOf("CapsLock"),
	driver.KeyComma:          js.ValueOf("Comma"),
	driver.KeyContextMenu:    js.ValueOf("ContextMenu"),
	driver.KeyControlLeft:    js.ValueOf("ControlLeft"),
	driver.KeyControlRight:   js.ValueOf("ControlRight"),
	driver.KeyD:              js.ValueOf("KeyD"),
	driver.KeyDelete:         js.ValueOf("Delete"),
	driver.KeyDigit0:         js.ValueOf("Digit0"),
	driver.KeyDigit1:         js.ValueOf("Digit1"),
	driver.KeyDigit2:         js.ValueOf("Digit2"),
	driver.KeyDigit3:         js.ValueOf("Digit3"),
	driver.KeyDigit4:         js.ValueOf("Digit4"),
	driver.KeyDigit5:         js.ValueOf("Digit5"),
	driver.KeyDigit6:         js.ValueOf("Digit6"),
	driver.KeyDigit7:         js.ValueOf("Digit7"),
	driver.KeyDigit8:         js.ValueOf("Digit8"),
	driver.KeyDigit9:         js.ValueOf("Digit9"),
	driver.KeyE:              js.ValueOf("KeyE"),
	driver.KeyEnd:            js.ValueOf("End"),
	driver.KeyEnter:          js.ValueOf("Enter"),
	driver.KeyEqual:          js.ValueOf("Equal"),
	driver.KeyEscape:         js.ValueOf("Escape"),
	driver.KeyF:              js.ValueOf("KeyF"),
	driver.KeyF1:             js.ValueOf("F1"),
	driver.KeyF10:            js.ValueOf("F10"),
	driver.KeyF11:            js.ValueOf("F11"),
	driver.KeyF12:            js.ValueOf("F12"),
	driver.KeyF2:             js.ValueOf("F2"),
	driver.KeyF3:             js.ValueOf("F3"),
	driver.KeyF4:             js.ValueOf("F4"),
	driver.KeyF5:             js.ValueOf("F5"),
	driver.KeyF6:             js.ValueOf("F6"),
	driver.KeyF7:             js.ValueOf("F7"),
	driver.KeyF8:             js.ValueOf("F8"),
	driver.KeyF9:             js.ValueOf("F9"),
	driver.KeyG:              js.ValueOf("KeyG"),
	driver.KeyH:              js.ValueOf("KeyH"),
	driver.KeyHome:           js.ValueOf("Home"),
	driver.KeyI:              js.ValueOf("KeyI"),
	driver.KeyInsert:         js.ValueOf("Insert"),
	driver.KeyJ:              js.ValueOf("KeyJ"),
	driver.KeyK:              js.ValueOf("KeyK"),
	driver.KeyL:              js.ValueOf("KeyL"),
	driver.KeyM:              js.ValueOf("KeyM"),
	driver.KeyMetaLeft:       js.ValueOf("MetaLeft"),
	driver.KeyMetaRight:      js.ValueOf("MetaRight"),
	driver.KeyMinus:          js.ValueOf("Minus"),
	driver.KeyN:              js.ValueOf("KeyN"),
	driver.KeyNumLock:        js.ValueOf("NumLock"),
	driver.KeyNumpad0:        js.ValueOf("Numpad0"),
	driver.KeyNumpad1:        js.ValueOf("Numpad1"),
	driver.KeyNumpad2:        js.ValueOf("Numpad2"),
	driver.KeyNumpad3:        js.ValueOf("Numpad3"),
	driver.KeyNumpad4:        js.ValueOf("Numpad4"),
	driver.KeyNumpad5:        js.ValueOf("Numpad5"),
	driver.KeyNumpad6:        js.ValueOf("Numpad6"),
	driver.KeyNumpad7:        js.ValueOf("Numpad7"),
	driver.KeyNumpad8:        js.ValueOf("Numpad8"),
	driver.KeyNumpad9:        js.ValueOf("Numpad9"),
	driver.KeyNumpadAdd:      js.ValueOf("NumpadAdd"),
	driver.KeyNumpadDecimal:  js.ValueOf("NumpadDecimal"),
	driver.KeyNumpadDivide:   js.ValueOf("NumpadDivide"),
	driver.KeyNumpadEnter:    js.ValueOf("NumpadEnter"),
	driver.KeyNumpadEqual:    js.ValueOf("NumpadEqual"),
	driver.KeyNumpadMultiply: js.ValueOf("NumpadMultiply"),
	driver.KeyNumpadSubtract: js.ValueOf("NumpadSubtract"),
	driver.KeyO:              js.ValueOf("KeyO"),
	driver.KeyP:              js.ValueOf("KeyP"),
	driver.KeyPageDown:       js.ValueOf("PageDown"),
	driver.KeyPageUp:         js.ValueOf("PageUp"),
	driver.KeyPause:          js.ValueOf("Pause"),
	driver.KeyPeriod:         js.ValueOf("Period"),
	driver.KeyPrintScreen:    js.ValueOf("PrintScreen"),
	driver.KeyQ:              js.ValueOf("KeyQ"),
	driver.KeyQuote:          js.ValueOf("Quote"),
	driver.KeyR:              js.ValueOf("KeyR"),
	driver.KeyS:              js.ValueOf("KeyS"),
	driver.KeyScrollLock:     js.ValueOf("ScrollLock"),
	driver.KeySemicolon:      js.ValueOf("Semicolon"),
	driver.KeyShiftLeft:      js.ValueOf("ShiftLeft"),
	driver.KeyShiftRight:     js.ValueOf("ShiftRight"),
	driver.KeySlash:          js.ValueOf("Slash"),
	driver.KeySpace:          js.ValueOf("Space"),
	driver.KeyT:              js.ValueOf("KeyT"),
	driver.KeyTab:            js.ValueOf("Tab"),
	driver.KeyU:              js.ValueOf("KeyU"),
	driver.KeyV:              js.ValueOf("KeyV"),
	driver.KeyW:              js.ValueOf("KeyW"),
	driver.KeyX:              js.ValueOf("KeyX"),
	driver.KeyY:              js.ValueOf("KeyY"),
	driver.KeyZ:              js.ValueOf("KeyZ"),
}

var edgeKeyCodeToDriverKey = map[int]driver.Key{
	8:   driver.KeyBackspace,
	9:   driver.KeyTab,
	13:  driver.KeyEnter,
	16:  driver.KeyShiftLeft,
	17:  driver.KeyControlLeft,
	18:  driver.KeyAltLeft,
	19:  driver.KeyPause,
	20:  driver.KeyCapsLock,
	27:  driver.KeyEscape,
	32:  driver.KeySpace,
	33:  driver.KeyPageUp,
	34:  driver.KeyPageDown,
	35:  driver.KeyEnd,
	36:  driver.KeyHome,
	37:  driver.KeyArrowLeft,
	38:  driver.KeyArrowUp,
	39:  driver.KeyArrowRight,
	40:  driver.KeyArrowDown,
	45:  driver.KeyInsert,
	46:  driver.KeyDelete,
	48:  driver.KeyDigit0,
	49:  driver.KeyDigit1,
	50:  driver.KeyDigit2,
	51:  driver.KeyDigit3,
	52:  driver.KeyDigit4,
	53:  driver.KeyDigit5,
	54:  driver.KeyDigit6,
	55:  driver.KeyDigit7,
	56:  driver.KeyDigit8,
	57:  driver.KeyDigit9,
	65:  driver.KeyA,
	66:  driver.KeyB,
	67:  driver.KeyC,
	68:  driver.KeyD,
	69:  driver.KeyE,
	70:  driver.KeyF,
	71:  driver.KeyG,
	72:  driver.KeyH,
	73:  driver.KeyI,
	74:  driver.KeyJ,
	75:  driver.KeyK,
	76:  driver.KeyL,
	77:  driver.KeyM,
	78:  driver.KeyN,
	79:  driver.KeyO,
	80:  driver.KeyP,
	81:  driver.KeyQ,
	82:  driver.KeyR,
	83:  driver.KeyS,
	84:  driver.KeyT,
	85:  driver.KeyU,
	86:  driver.KeyV,
	87:  driver.KeyW,
	88:  driver.KeyX,
	89:  driver.KeyY,
	90:  driver.KeyZ,
	91:  driver.KeyMetaLeft,
	92:  driver.KeyMetaRight,
	93:  driver.KeyContextMenu,
	96:  driver.KeyNumpad0,
	97:  driver.KeyNumpad1,
	98:  driver.KeyNumpad2,
	99:  driver.KeyNumpad3,
	100: driver.KeyNumpad4,
	101: driver.KeyNumpad5,
	102: driver.KeyNumpad6,
	103: driver.KeyNumpad7,
	104: driver.KeyNumpad8,
	105: driver.KeyNumpad9,
	106: driver.KeyNumpadMultiply,
	107: driver.KeyNumpadAdd,
	109: driver.KeyNumpadSubtract,
	110: driver.KeyNumpadDecimal,
	111: driver.KeyNumpadDivide,
	112: driver.KeyF1,
	113: driver.KeyF2,
	114: driver.KeyF3,
	115: driver.KeyF4,
	116: driver.KeyF5,
	117: driver.KeyF6,
	118: driver.KeyF7,
	119: driver.KeyF8,
	120: driver.KeyF9,
	121: driver.KeyF10,
	122: driver.KeyF11,
	123: driver.KeyF12,
	144: driver.KeyNumLock,
	145: driver.KeyScrollLock,
	186: driver.KeySemicolon,
	187: driver.KeyEqual,
	188: driver.KeyComma,
	189: driver.KeyMinus,
	190: driver.KeyPeriod,
	191: driver.KeySlash,
	192: driver.KeyBackquote,
	219: driver.KeyBracketLeft,
	220: driver.KeyBackslash,
	221: driver.KeyBracketRight,
	222: driver.KeyQuote,
}

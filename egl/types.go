// THE AUTOGENERATED LICENSE. ALL THE RIGHTS ARE RESERVED BY ROBOTS.

// WARNING: This file has automatically been generated on Tue, 10 Jan 2017 23:39:22 MSK.
// By https://git.io/cgogen. DO NOT EDIT.

package egl

/*
#cgo LDFLAGS: -lEGL
#include <EGL/egl.h>
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"
import "unsafe"

// Boolean type as declared in EGL/egl.h:44
type Boolean uint32

// Enum type as declared in EGL/egl.h:45
type Enum uint32

// Config type as declared in EGL/egl.h:46
type Config unsafe.Pointer

// Context type as declared in EGL/egl.h:47
type Context unsafe.Pointer

// Display type as declared in EGL/egl.h:48
type Display unsafe.Pointer

// Surface type as declared in EGL/egl.h:49
type Surface unsafe.Pointer

// ClientBuffer type as declared in EGL/egl.h:50
type ClientBuffer unsafe.Pointer

// NativeWindowType as declared in EGL/eglplatform.h:91
type NativeWindowType C.EGLNativeWindowType

// NativePixmapType as declared in EGL/eglplatform.h:92
type NativePixmapType C.EGLNativePixmapType

// NativeDisplayType type as declared in EGL/eglplatform.h:93
type NativeDisplayType unsafe.Pointer

// Int type as declared in EGL/eglplatform.h:122
type Int int32

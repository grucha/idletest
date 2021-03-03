package os_specific

// #cgo LDFLAGS: -framework CoreGraphics
// #include <CoreGraphics/CoreGraphics.h>
import "C"
import (
	"fmt"
	"math"
)

func ShowInactiveTime() {
	// Returns the number of seconds a osx user is idle!
	result := C.CGEventSourceSecondsSinceLastEventType(1, math.MaxInt32)
	fmt.Println(result)
}


// Package zvg contains go bindings for libzvg.so. See https://github.com/rhew/zvg-linux.
package zvg

// #cgo LDFLAGS: -lzvg
// #include </usr/local/include/zvg/zvgPort.h>
// #include </usr/local/include/zvg/zvgFrame.h>
import "C"


func ZvgBanner() {
    C.zvgBanner((&C.ZvgSpeeds[0]), &C.ZvgID)
}

func ZvgFrameOpen() int {
    return int(C.zvgFrameOpen())
}

func ZvgFrameClose() {
    C.zvgFrameClose()
}

func ZvgError(err int) {
    C.zvgError(C.uint(err))
}

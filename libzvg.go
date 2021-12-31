// Package zvg contains go bindings for libzvg.so. See https://github.com/rhew/zvg-linux.
package zvg

// #cgo LDFLAGS: -lzvg
// #include </usr/local/include/zvg/zvgPort.h>
// #include </usr/local/include/zvg/zvgFrame.h>
import "C"


func ZvgBanner() {
    C.zvgBanner((&C.ZvgSpeeds[0]), &C.ZvgID)
}

func ZvgFrameOpen() uint {
    return uint(C.zvgFrameOpen())
}

func ZvgFrameClose() {
    C.zvgFrameClose()
}

func ZvgError(err uint) {
    C.zvgError(C.uint(err))
}

func ZvgFrameVector(xStart, yStart, xEnd, yEnd uint) uint {
    return uint(C.zvgFrameVector(C.uint(xStart), C.uint(yStart), C.uint(xEnd), C.uint(yEnd)))
}

//   uint zvgFrameSend(void)
func ZvgFrameSend() uint {
    return uint(C.zvgFrameSend())
}

func ZvgFrameSetColor(color uint) {
    C.zvgEncSetColor(C.uint(color))
}

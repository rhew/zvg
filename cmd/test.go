package main

import (
    // "github.com/rhew/zvg/pkg/zvg"
    "github.com/rhew/zvg"
)

func main() {
    err := zvg.ZvgFrameOpen()
    if err != 0 {
        zvg.ZvgError(err)
        return
    }

    defer zvg.ZvgFrameClose()
    zvg.ZvgBanner()
}

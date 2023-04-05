//go:build darwin

package avcore

import "C"
import (
	"github.com/ebitengine/purego/objc"
	"github.com/go-musicfox/go-musicfox/pkg/macdriver"
	"github.com/go-musicfox/go-musicfox/pkg/macdriver/core"
)

func init() {
	importFramework()
	class_AVPlayer = objc.GetClass("AVPlayer")
}

var (
	class_AVPlayer objc.Class
)

var (
	sel_initWithPlayerItem = objc.RegisterName("initWithPlayerItem:")
	sel_setActionAtItemEnd = objc.RegisterName("setActionAtItemEnd:")
	sel_setVolume          = objc.RegisterName("setVolume:")
	sel_currentItem        = objc.RegisterName("currentItem")
	sel_currentTime        = objc.RegisterName("currentTime")
	sel_pause              = objc.RegisterName("pause")
	sel_play               = objc.RegisterName("play")
	sel_seekToTime         = objc.RegisterName("seekToTime")
)

type AVPlayer struct {
	core.NSObject
}

func AVPlayer_alloc() AVPlayer {
	return AVPlayer{
		NSObject: core.NSObject{
			ID: objc.ID(class_AVPlayer).Send(macdriver.SEL_alloc),
		},
	}
}

func (p AVPlayer) InitWithPlayerItem(item AVPlayerItem) AVPlayer {
	p.Send(sel_initWithPlayerItem, item.ID)
	return p
}

func (p AVPlayer) SetActionAtItemEnd(value core.NSInteger) {
	p.Send(sel_setActionAtItemEnd, value)
}

func (p AVPlayer) SetVolume(value float32) {
	p.Send(sel_setVolume, value)
}

func (p AVPlayer) CurrentItem() (item AVPlayerItem) {
	item.SetObjcID(p.Send(sel_currentItem))
	return
}

func (p AVPlayer) CurrentTime() objc.ID {
	return p.Send(sel_currentTime)
}

func (p AVPlayer) Pause() {
	p.Send(sel_pause)
}

func (p AVPlayer) Play() {
	p.Send(sel_play)
}

func (p AVPlayer) SeekToTime(time CMTime) {
	p.Send(sel_seekToTime, time)
}

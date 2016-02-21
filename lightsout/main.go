package main

import (
	"flag"
	"image"
	"image/color"
	"log"

	"golang.org/x/exp/shiny/driver"
	"golang.org/x/exp/shiny/screen"
	"golang.org/x/mobile/event/key"
	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/mouse"
	"golang.org/x/mobile/event/paint"
	"golang.org/x/mobile/event/size"
)

const maxLevel = 50

var level = flag.Int("level", 1, "starting level (1-50)")

var gray = color.RGBA{0xCC, 0xCC, 0xCC, 0xFF}

func main() {
	flag.Parse()

	//TODO support level==0 for random
	if *level < 1 || *level > maxLevel {
		log.Fatal("illegal level argument: ", *level)
	}

	bd := new(board)
	bd.load(*level)

	driver.Main(func(s screen.Screen) {
		w, err := s.NewWindow(nil)
		if err != nil {
			log.Fatal(err.Error())
		}
		defer w.Release()

		var b screen.Buffer
		defer func() {
			if b != nil {
				b.Release()
			}
		}()

		for {
			switch e := w.NextEvent().(type) {
			case lifecycle.Event:
				if e.To == lifecycle.StageDead {
					return
				}
			case key.Event:
				if e.Code == key.CodeEscape {
					return
				}
			case mouse.Event:
				if e.Direction == mouse.DirRelease && e.Button == mouse.ButtonLeft {
					bd.click(b.RGBA(), int(e.X), int(e.Y))
					if bd.allOff() {
						*level++
						if *level > maxLevel {
							*level = 1
						}
						bd.load(*level)
					}
					bd.draw(b.RGBA())
					w.Send(paint.Event{})
				}
			case paint.Event:
				// Gray background
				w.Fill(b.Bounds(), gray, screen.Src)
				// Board buffer
				w.Upload(image.Point{}, b, b.Bounds())
				w.Publish()
			case size.Event:
				if b != nil {
					b.Release()
				}
				b, err = s.NewBuffer(e.Size())
				if err != nil {
					log.Fatal(err)
				}
				// Overlay board
				bd.draw(b.RGBA())
			case error:
				log.Print(e)
			}
		}
	})
}

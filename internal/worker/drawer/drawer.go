package drawer

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/term"
)

type config struct {
	XPos int
	Xend int
	Pict [][][]string
}

type drawer struct {
	config
	context.Context
}

func NewDrawer(ctx context.Context) *drawer {
	termW, _, err := term.GetSize(int(os.Stdin.Fd()))
	if err != nil {
		log.Fatalln(err)
	}

	repeat := func(ll []string, num int) [][]string {
		r := make([][]string, 0)
		for i := 0; i < num; i++ {
			rl := append([]string{}, ll...) //deepcopy
			r = append(r, rl)
		}
		return r
	}
	appendSlices := func(ll ...[][]string) [][]string {
		r := [][]string{}
		for _, rl := range ll {
			r = append(r, rl...)
		}
		return r
	}

	pict := [][][]string{
		appendSlices(
			repeat([]string{
				"            Snake Is Sneaking",
			}, 20),
			repeat([]string{
				"                       まぁ気長にやろうや",
			}, 20),
			repeat([]string{
				"",
			}, 20),
		),
		{
			{
				"ヘビ",
			},
		},
		appendSlices(
			repeat([]string{
				"          _______  ",
				"         / ___   \\ ",
				"        |  *      \\ ",
				"     ===\\__       \\ ",
				"            |     | ",
				"             |    | ",
			}, 5),
			repeat([]string{
				"                   ",
				"          _______  ",
				"         / ___   \\ ",
				"        |  *      \\ ",
				"     ===\\__       \\ ",
				"            |     | ",
			}, 5),
		),
		appendSlices(
			//state1
			repeat([]string{
				"             |    | ",
				"             |    | ",
				"            |      \\   _/¯¯¯¯¯¯¯¯¯\\                 ",
				"           /        ¯¯¯             \\            __",
				"          |                 _____     \\_________/  / ",
				"          \\              _/      \\               /  ",
				"            ____________/         \\____________/    ",
			}, 2),
			//state2
			repeat([]string{
				"             |    | ",
				"             |    | ",
				"            |      \\   _____________     ",
				"            |        ¯¯¯             \\    ",
				"           |                            \\____________ ",
				"           \\              ________                   /  ",
				"             ____________/         \\______________/¯¯  ",
			}, 2),
			//state3
			repeat([]string{
				"             |    | ",
				"             |    | ",
				"            |      \\                   ",
				"            |        ¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯\\     ",
				"           |                            \\-----------___   ",
				"           \\                                          /  ",
				"             ____________/¯¯¯¯¯¯¯¯¯¯¯\\____________---¯¯ ",
			}, 2),
			//state4
			repeat([]string{
				"             |    | ",
				"             |    | ",
				"            |      \\        _________________  ",
				"            |        ¯¯¯¯¯¯¯¯               \\___    ",
				"             |                                     \\_____ ",
				"               \\              ___________                 \\",
				"                  ____________/            \\________________\\  ",
			}, 2),
			//state3
			repeat([]string{
				"             |    | ",
				"             |    | ",
				"            |      \\  ",
				"            |        ¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯¯\\   ",
				"           |                            \\-----------___ ",
				"           \\                                          / ",
				"             ____________/¯¯¯¯¯¯¯¯¯¯¯\\____________---¯  ",
			}, 2),
			//state2
			repeat([]string{
				"             |    | ",
				"             |    | ",
				"            |      \\   _____________  ",
				"            |        ¯¯¯             \\    ",
				"           |                            \\____________  ",
				"           \\              ________                   / ",
				"             ____________/         \\______________/¯¯  ",
			}, 2),
		),
	}

	for _, ll := range pict {
		for _, l := range ll {
			for n := range l {
				l[n] = strings.Repeat(" ", termW) + l[n]
			}
		}
	}

	max_width := func() int {
		w := 0
		for _, ll := range pict {
			for _, l := range ll {
				if len(l) > w {
					w = len(l)
				}
			}
		}
		return w
	}()

	return NewDrawerWithConfig(ctx, config{
		XPos: 0,
		Xend: termW + max_width + 100,
		Pict: pict,
	})
}

func NewDrawerWithConfig(ctx context.Context, config config) *drawer {
	return &drawer{
		config,
		ctx,
	}
}

func (d *drawer) Clear() {
	// https://en.wikipedia.org/wiki/ANSI_escape_code
	// \033: clear console
	// [2J: unix系っぽい?
	fmt.Print("\033[H\033[2J")
}

func (d *drawer) DrawNext() (isEnd bool) {
	d.Clear()
	for _, ll := range d.Pict {
		for _, l := range ll[d.XPos%len(ll)] {
			line := string(l)
			if d.XPos < len(line) {
				line = line[d.XPos:]
			}

			termW, _, err := term.GetSize(int(os.Stdin.Fd()))
			if err != nil {
				log.Fatalln(err)
			}
			if len(line) > termW {
				line = line[:termW]
			}
			fmt.Println(line)
		}
	}

	d.XPos++
	isEnd = d.XPos > d.Xend
	return
}

func (d *drawer) Draw() {
loop:
	for {
		select {
		case <-d.Context.Done():
			break loop
		default:
			d.DrawNext()
		}
	}
}

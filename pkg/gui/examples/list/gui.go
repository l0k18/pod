package main

import (
	"github.com/l0k18/pod/pkg/gui"
	qu "github.com/l0k18/pod/pkg/util/quit"
	
	l "gioui.org/layout"
	
	"github.com/l0k18/pod/pkg/gui/fonts/p9fonts"
)

type Model struct {
	th *gui.Theme
}

func main() {
	quit := qu.T()
	th := gui.NewTheme(p9fonts.Collection(), quit)
	minerModel := Model{
		th: th,
	}
	go func() {
		if err := gui.NewWindow(th).
			Size(64, 32).
			Title("example").
			Open().
			Run(
				minerModel.mainWidget,
				func(l.Context) {},
				func() {
					quit.Q()
				}, quit,
			); Check(err) {
		}
	}()
	<-quit
}

func (m *Model) mainWidget(gtx l.Context) l.Dimensions {
	return m.th.Flex().Flexed(1,
		m.th.Fill("red", m.th.Flex().AlignMiddle().SpaceAround().
			Flexed(1,
				m.th.H6("example").Fn,
			).Fn, l.Center).Fn,
	).Fn(gtx)
}

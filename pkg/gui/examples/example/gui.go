package main

import (
	"fmt"
	
	"github.com/l0k18/pod/pkg/gui"
	qu "github.com/l0k18/pod/pkg/util/quit"
	
	l "gioui.org/layout"
	"gioui.org/text"
	mico "golang.org/x/exp/shiny/materialdesign/icons"
	
	"github.com/l0k18/pod/pkg/gui/fonts/p9fonts"
	icons "github.com/l0k18/pod/pkg/gui/ico/svg"
)

type MinerModel struct {
	th                                                 *gui.Theme
	button0, button1, button2, iconbutton, iconbutton1 *gui.Clickable
	boolButton1, boolButton2                           *gui.Bool
	quit                                               qu.C
	progress                                           int
	slider                                             *gui.Float
	lineEditor, areaEditor                             *gui.Editor
	radio                                              *gui.Enum
}

func main() {
	quit := qu.T()
	th := gui.NewTheme(p9fonts.Collection(), quit)
	minerModel := MinerModel{
		th:      th,
		button0: th.Clickable(),
		button1: th.Clickable(),
		button2: th.Clickable().SetClick(
			func() {
				Info("clicked default style button")
			},
		),
		boolButton1: th.Bool(false),
		boolButton2: th.Bool(false),
		iconbutton:  th.Clickable(),
		iconbutton1: th.Clickable(),
		quit:        qu.T(),
		progress:    0,
		slider: th.Float().SetHook(
			func(fl float32) {
				Debug("float now at value", fl)
			},
		),
		lineEditor: th.Editor().SingleLine().Submit(true),
		areaEditor: th.Editor().SingleLine().Submit(false),
		radio: th.Enum().SetOnChange(
			func(value string) {
				Debug("changed radio button to", value)
			},
		),
	}
	go func() {
		if err := gui.NewWindow(th).
			Size(64, 32).
			Title("example").
			Open().
			Run(
				minerModel.testLabels, func(l.Context) {}, func() {
					quit.Q()
					// os.Exit(0)
				}, quit,
		); Check(err) {
		}
	}()
	<-quit
	// app.Main()
}

func (m *MinerModel) testLabels(gtx l.Context) l.Dimensions {
	m.progress++
	if m.progress == 100 {
		m.progress = 0
	}
	th := m.th
	return th.Flex().Flexed(
		1,
		th.Flex().Rigid(
			th.Flex().Flexed(
				0.5,
				th.Fill("PanelBg", th.Inset(
					0.25,
					m.blocks(),
				).Fn, l.Center).Fn,
			).Flexed(
				0.5,
				th.Fill("DocBg", th.Inset(
					0.25,
					m.buttons(),
				).Fn, l.Center).Fn,
			).Fn,
		).Fn,
	).Fn(gtx)
}

func (m *MinerModel) blocks() l.Widget {
	th := m.th
	return th.Flex().Vertical().Rigid(
		th.Inset(
			0.25,
			th.Flex().Rigid(
				th.H1("this is a H1").
					Color("PanelText").
					Fn,
			).Fn,
		).Fn,
	).Rigid(
		th.Inset(
			0.25,
			th.H2("this is a H2").
				Font("bariol regular").
				Color("PanelText").Fn,
		).Fn,
	).Rigid(
		th.Inset(
			0.25,
			th.H3("this is a H3").
				Alignment(text.End).
				Color("PanelText").Fn,
		).Fn,
	).Rigid(
		th.Fill("DocBg", th.Inset(
			0.25,
			th.H4("this is a H4").
				Alignment(text.Middle).
				Color("DocText").Fn,
		).Fn, l.Center).Fn,
	).Rigid(
		th.Fill("PanelBg", th.Inset(
			0.25,
			th.H5("this is a H5").
				Color("PanelText").
				Fn,
		).Fn, l.Center).Fn,
	).Rigid(
		th.Inset(
			0.25,
			th.H6("this is a H6").
				Color("PanelText").Fn,
		).Fn,
	).Rigid(
		th.Inset(
			0.25,
			th.Body1("this is a Body1").Color("PanelText").Fn,
		).Fn,
	).Rigid(
		th.Inset(
			0.25,
			th.Body2("this is a Body2").Color("PanelText").Fn,
		).Fn,
	).Rigid(
		th.Inset(
			0.25,
			th.Caption("this is a Caption").Color("PanelText").Fn,
		).Fn,
	).Fn
}

func (m *MinerModel) buttons() l.Widget {
	th := m.th
	return th.Flex().Vertical().Rigid(
		th.Inset(
			0.25,
			th.Flex().Rigid(
				th.Button(
					m.button0.SetClick(
						func() {
							Info("clicked customised button")
						},
					),
				).
					CornerRadius(3).
					Background("Secondary").
					Color("Dark").
					Font("bariol bold").
					TextScale(2).
					Text("customised button").
					Inset(1.5).
					Fn,
			).Fn,
		).Fn,
	).Rigid(
		th.Flex().Rigid(
			th.Inset(
				0.25,
				th.Button(
					m.button2,
				).
					Text("default style").
					Fn,
			).Fn,
		).Rigid(
			th.Inset(
				0.25,
				th.Flex().Rigid(
					th.Indefinite().Color("Primary").Fn,
				).Fn,
			).Fn,
		).Rigid(
			th.Inset(
				0.25,
				th.IconButton(
					m.iconbutton.SetClick(
						func() {
							Debug("clicked parallelcoin button")
						},
					),
				).
					Icon(th.Icon().Src(&icons.ParallelCoin)).
					Fn,
			).Fn,
		).Rigid(
			th.Inset(
				0.25,
				th.IconButton(
					m.iconbutton1.SetClick(
						func() {
							Debug("clicked android button")
						},
					),
				).
					Scale(1).
					Background("Secondary").
					Icon(th.Icon().Src(&mico.ActionAndroid)).
					Fn,
			).Fn,
		).Fn,
	).Rigid(
		th.ProgressBar().Color("Primary").SetProgress(m.progress).Fn,
	).Rigid(
		th.ProgressBar().Color("Primary").SetProgress(int(m.slider.Value())).Fn,
	).Rigid(
		th.Flex().
			Flexed(
				1,
				th.Slider().
					Float(m.slider).
					Min(0).Max(100).
					Fn,
			).
			Rigid(
				th.Body1(fmt.Sprintf("%3v", int(m.slider.Value()))).
					Font("go regular").Color("DocText").
					Fn,
			).Fn,
	).Rigid(
		th.Flex().Rigid(
			th.Icon().Scale(2).Color("DocText").Src(&icons.ParallelCoinRound).Fn,
			// ).Rigid(
			// 	th.RadioButton(m.radio, "first", "first").Fn,
			// ).Rigid(
			// 	th.RadioButton(m.radio, "second", "second").Fn,
			// ).Rigid(
			// 	th.RadioButton(m.radio, "third", "third").Fn,
		).Rigid(
			th.Switch(
				m.boolButton2.SetOnChange(
					func(b bool) {
						Debug("switch state set to", b)
					},
				),
			).Fn,
		).Rigid(
			th.CheckBox(
				m.boolButton1.SetOnChange(
					func(b bool) {
						Debug("change state to", b)
					},
				),
			).
				IconColor("Primary").
				TextColor("DocText").
				// IconScale(0.1).
				Text("checkbox").
				Fn,
		).Fn,
	).Rigid(
		th.Inset(
			0.25,
			th.Border().Embed(
				th.Inset(
					0.25,
					th.TextInput(
						m.lineEditor.
							SetChange(
								func(txt string) {
									Debug("lineEditor changed to:\n" + txt)
								},
						).
							SetFocus(
								func(is bool) {
									Debug("lineEditor is focused", is)
								},
							).
							SetSubmit(
								func(txt string) {
									Debug("lineEditor submitted with text:\n" + txt)
								},
							), "hint",
					).Fn,
				).Fn,
			).Fn,
		).Fn,
	).Flexed(
		1,
		th.Inset(
			0.25,
			th.Border().Embed(
				th.Inset(
					0.25,
					th.TextInput(
						m.areaEditor.
							SetChange(
								func(txt string) {
									Debug("areaEditor changed to:\n" + txt)
								},
						).
							SetFocus(
								func(is bool) {
									Debug("areaEditor is focused", is)
								},
							).
							SetSubmit(
								func(txt string) {
									Debug("areaEditor submitted with text:\n" + txt)
								},
							), "hint",
					).Fn,
				).Fn,
			).Fn,
		).Fn,
	).Fn
}

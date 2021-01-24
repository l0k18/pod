package gui

import (
	l "gioui.org/layout"
	"gioui.org/text"
	
	"github.com/l0k18/pod/pkg/gui"
	p9icons "github.com/l0k18/pod/pkg/gui/ico/svg"
	"github.com/l0k18/pod/version"
)

func (wg *WalletGUI) HelpPage() func(gtx l.Context) l.Dimensions {
	return func(gtx l.Context) l.Dimensions {
		return wg.VFlex().AlignMiddle().
			Flexed(0.5, gui.EmptyMaxWidth()).
			Rigid(
				wg.H5("ParallelCoin Pod Gio Wallet").Alignment(text.Middle).Fn,
			).
			Rigid(
				wg.Fill("DocBg", l.Center, wg.TextSize.V, 0, wg.Inset(0.5,
					wg.VFlex().
						AlignMiddle().
						Rigid(
							
							wg.VFlex().AlignMiddle().
								Rigid(
									wg.Inset(0.25,
										wg.Caption("Built from git repository:").
											Font("bariol bold").Fn,
									).Fn,
								).
								Rigid(
									wg.Caption(version.URL).Fn,
								).
								Fn,
						
						).
						Rigid(
							
							wg.VFlex().AlignMiddle().
								Rigid(
									wg.Inset(0.25,
										wg.Caption("GitRef:").
											Font("bariol bold").Fn,
									).Fn,
								).
								Rigid(
									wg.Caption(version.GitRef).Fn,
								).
								Fn,
						
						).
						Rigid(
							
							wg.VFlex().AlignMiddle().
								Rigid(
									wg.Inset(0.25,
										wg.Caption("GitCommit:").
											Font("bariol bold").Fn,
									).Fn,
								).
								Rigid(
									wg.Caption(version.GitCommit).Fn,
								).
								Fn,
						
						).
						Rigid(
							
							wg.VFlex().AlignMiddle().
								Rigid(
									wg.Inset(0.25,
										wg.Caption("BuildTime:").
											Font("bariol bold").Fn,
									).Fn,
								).
								Rigid(
									wg.Caption(version.BuildTime).Fn,
								).
								Fn,
						
						).
						Rigid(
							
							wg.VFlex().AlignMiddle().
								Rigid(
									wg.Inset(0.25,
										wg.Caption("Tag:").
											Font("bariol bold").Fn,
									).Fn,
								).
								Rigid(
									wg.Caption(version.Tag).Fn,
								).
								Fn,
						
						).
						Rigid(
							wg.Icon().Scale(gui.Scales["H6"]).
								Color("DocText").
								Src(&p9icons.Gio).
								Fn,
						).
						Rigid(
							wg.Caption("powered by Gio").Fn,
						).
						Fn,
				).Fn).Fn,
			).
			Flexed(0.5, gui.EmptyMaxWidth()).
			Fn(gtx)
	}
}

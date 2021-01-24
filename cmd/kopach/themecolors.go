package kopach

import "github.com/l0k18/pod/app/save"

func (m *MinerModel) FlipTheme() {
	m.DarkTheme = !m.DarkTheme
	Debug("dark theme:", m.DarkTheme)
	m.SetTheme(m.DarkTheme)
}

func (m *MinerModel) SetTheme(dark bool) {
	m.Theme.Colors.SetTheme(dark)
	*m.Cx.Config.DarkTheme = dark
	save.Pod(m.Cx.Config)
}

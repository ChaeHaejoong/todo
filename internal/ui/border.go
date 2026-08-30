package ui

import (
	"strings"

	"charm.land/lipgloss/v2"
)

var titledBorderStyle = lipgloss.NewStyle().
	Border(lipgloss.RoundedBorder()).
	Padding(0, 1)

func TitledBorder(width, height int, title, content string, focused bool) string {
	if width < 2 || height < 2 {
		return content
	}

	style := titledBorderStyle
	if focused {
		style = style.BorderForeground(ColorFocused)
	} else {
		style = style.BorderForeground(ColorNormal)
	}

	border, _, _, _, _ := style.GetBorder()
	innerWidth := width - 2
	label := " " + title + " "
	labelWidth := lipgloss.Width(label)

	if labelWidth+1 > innerWidth {
		label = lipgloss.NewStyle().MaxWidth(max(0, innerWidth-1)).Render(label)
		labelWidth = lipgloss.Width(label)
	}

	topBorderWidth := max(0, innerWidth-labelWidth-1)

	topBorderStyle := lipgloss.NewStyle().Foreground(style.GetBorderTopForeground())
	titleStyle := lipgloss.NewStyle().Foreground(style.GetBorderTopForeground())
	top := topBorderStyle.Render(border.TopLeft+border.Top) +
		titleStyle.Render(label) +
		topBorderStyle.Render(strings.Repeat(border.Top, topBorderWidth)+border.TopRight)

	bodyStyle := style.BorderTop(false).Width(width).Height(height - 1)
	return lipgloss.JoinVertical(lipgloss.Left, top, bodyStyle.Render(content))
}

func TitledBorderContentSize(width, height int) (int, int) {
	contentWidth := width - titledBorderStyle.GetBorderLeftSize() - titledBorderStyle.GetBorderRightSize()
	contentWidth -= titledBorderStyle.GetPaddingLeft() + titledBorderStyle.GetPaddingRight()

	contentHeight := height - titledBorderStyle.GetBorderTopSize() - titledBorderStyle.GetBorderBottomSize()
	contentHeight -= titledBorderStyle.GetPaddingTop() + titledBorderStyle.GetPaddingBottom()

	return max(0, contentWidth), max(0, contentHeight)
}

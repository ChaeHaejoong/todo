package appointment

func (m Model) View(width, height int, focused bool) string {
	return render(width, height, "", focused)
}

package appointment

func (m Model) View(width, height int, focused bool) string {
	if m.loading {
		return render(width, height, "fetching data...", focused)
	}
	if m.err != nil {
		return render(width, height, "Appointment를 불러오지 못했습니다: "+m.err.Error(), focused)
	}
	return render(width, height, renderAppointments(m.appointments, m.cursor), focused)
}

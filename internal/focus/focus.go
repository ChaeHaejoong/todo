package focus

type Target int

const (
	TodoList Target = iota
	TodoModal
	Calendar
)

type Manager struct {
	current Target
}

func New() Manager {
	return Manager{current: TodoList}
}

func (m *Manager) Set(target Target) {
	m.current = target
}

func (m Manager) Current() Target {
	return m.current
}

func (m Manager) Is(target Target) bool {
	return m.current == target
}

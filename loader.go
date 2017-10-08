package cors

import . "github.com/goline/lapi"

func NewCorsLoader(policy Policy, priority int) Loader {
	l := &corsLoader{Policy: policy}
	l.WithPriority(priority)
	return l
}

type corsLoader struct {
	Policy Policy
	PriorityAware
}

func (l *corsLoader) Load(app App) {
	app.Router().Options("/.*", &corsHandler{l.Policy})
	app.Router().WithHook(&corsHook{l.Policy})
}

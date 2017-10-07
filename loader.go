package cors

import . "github.com/goline/lapi"

func NewCorsLoader(policy Policy) Loader {
	return &corsLoader{policy}
}

type corsLoader struct {
	Policy Policy
}

func (l *corsLoader) Load(app App) {
	app.Router().Options("/.*", &corsHandler{l.Policy})
	app.Router().WithHook(&corsHook{l.Policy})
}

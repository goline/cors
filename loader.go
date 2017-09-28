package cors

import . "github.com/goline/lapi"

func NewLoader(policy Policy) Loader {
	return &corsLoader{policy}
}

type corsLoader struct {
	Policy Policy
}

func (l *corsLoader) Load(app App) {
	app.Router().Any("/.*", NewCorsHandler(l.Policy))
}

package ds

type Run struct {
	f func(...interface{})
	args []interface{}
}

func NewRun(f func(...interface{}), scenarios ...interface{}) *Run {
	return &Run{f: f, args: scenarios}
}

func (r *Run) Run() {
	for _, scenario := range r.args {
		r.f(scenario)
	}
}
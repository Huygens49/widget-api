package working

type Widget struct {
	Value int
}

func (w *Widget) Work() {
	w.Value++
}

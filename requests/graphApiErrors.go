package requests

type InvalidGraphApiIntervals struct{}

func (m *InvalidGraphApiIntervals) Error() string {
	return "graph api returned invalid intervals - from and to were the same"
}

type GraphNotVisible struct{}

func (m *GraphNotVisible) Error() string {
	return "graph was not visible, color was set to transparent"
}

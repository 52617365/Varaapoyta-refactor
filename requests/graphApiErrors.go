package requests

type InvalidGraphApiIntervals struct{}

func (m *InvalidGraphApiIntervals) Error() string {
	return "graph api returned invalid intervals - from and to were the same"
}

type GraphNotVisible struct{}

func (m *GraphNotVisible) Error() string {
	return "graph was not visible, color was set to transparent"
}

type UrlShouldBeSkipped struct{}

func (m *UrlShouldBeSkipped) Error() string {
	return "url should be skipped."
}

type RequiredInfoDoesNotExist struct{}

func (m *RequiredInfoDoesNotExist) Error() string {
	return "required info doesn't exist"
}
package responseStructures

type GraphApiResponse []struct {
	Name      string `json:"name"`
	Intervals []struct {
		From  int64  `json:"from"`
		To    int64  `json:"to"`
		Color string `json:"color"`
	} `json:"intervals"`
	ID int `json:"id"`
}
type RelevantIndex = struct {
	Name      string `json:"name"`
	Intervals []struct {
		From  int64  `json:"from"`
		To    int64  `json:"to"`
		Color string `json:"color"`
	} `json:"intervals"`
	ID int `json:"id"`
}

type Intervals = []struct {
	From  int64  `json:"from"`
	To    int64  `json:"to"`
	Color string `json:"color"`
}

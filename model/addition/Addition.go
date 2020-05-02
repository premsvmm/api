package addition


type Add struct {
	Value1 int64 `json:"value_1"`
	Value2 int64 `json:"value_2"`
}

type Result struct {
	Value int64 `json:"value"`
}
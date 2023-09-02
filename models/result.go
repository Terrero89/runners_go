package result

type Result struct {
	ID         string `json:"id"`
	RunnersID  string `json:"runners_id"`
	RaceResult string `json:"race_result"`
	Location   string `json:"location"`
	Position   int    `json:"position,omitempty"`
	Year       int    `json:"year"`
}

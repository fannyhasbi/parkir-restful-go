package realtime

// Realtime struct represents single data of realtime scan data
type Realtime struct {
	ID    int    `json:"id"`
	Waktu string `json:"waktu"`
	Merk  string `json:"merk"`
	Tipe  string `json:"tipe"`
	Nama  string `json:"nama"`
}

// ResponseRealtime struct represents realtime data followed by HTTP status
type ResponseRealtime struct {
	Status  int        `json:"status"`
	Message string     `json:"message"`
	Data    []Realtime `json:"data"`
}

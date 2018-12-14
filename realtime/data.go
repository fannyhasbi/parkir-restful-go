package realtime

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/fannyhasbi/parkir-restful-go/database"
)

func ReturnRealtime(w http.ResponseWriter, r *http.Request) {
	var realtime Realtime
	var arrRealtimes []Realtime
	var response ResponseRealtime

	db := database.Connect()
	defer db.Close()

	rows, err := db.Query(`
		SELECT s.id, s.waktu, v.merk, v.tipe, o.nama FROM scan s
		INNER JOIN vehicle v
			ON s.id_vehicle = v.id
		INNER JOIN owner o
			ON v.id_owner = o.id
	`)

	if err != nil {
		log.Println(err)
	}

	for rows.Next() {
		if err := rows.Scan(&realtime.ID, &realtime.Waktu, &realtime.Merk, &realtime.Tipe, &realtime.Nama); err != nil {
			log.Fatal(err.Error())
		} else {
			arrRealtimes = append(arrRealtimes, realtime)
		}
	}

	response.Status = 200
	response.Message = "OK"
	response.Data = arrRealtimes

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

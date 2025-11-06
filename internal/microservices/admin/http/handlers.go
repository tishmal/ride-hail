package admin

import (
	"encoding/json"
	"net/http"
)

func OverviewHandler(w http.ResponseWriter, r *http.Request) {
	resp := map[string]any{
		"metrics": map[string]int{
			"active_rides":      2,
			"available_drivers": 5,
		},
	}
	_ = json.NewEncoder(w).Encode(resp)
}

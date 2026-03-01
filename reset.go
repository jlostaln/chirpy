package main

import (
	"net/http"
)

func (cfg *apiConfig) handlerReset(w http.ResponseWriter, r *http.Request) {
	if cfg.platform != "dev" {
		respondWithError(
			w,
			http.StatusForbidden,
			"Reset is only allowed in dev environment.",
			nil,
		)
		return
	}
	err := cfg.db.ResetDB(r.Context())
	if err != nil {
		respondWithError(
			w,
			http.StatusInternalServerError,
			"Couldn't delete users"+err.Error(),
			nil,
		)
		return
	}
	cfg.fileserverHits.Store(0)
	respondWithJSON(w, http.StatusOK, map[string]string{
		"success": "Hits reset to 0 and database reset to initial state.",
	})
}

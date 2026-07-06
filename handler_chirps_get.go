package main

import (
	"net/http"
	"sort"

	"github.com/google/uuid"
	"github.com/jlostaln/chirpy/internal/database"
)

func (cfg *apiConfig) handlerChirpsRetrieve(w http.ResponseWriter, r *http.Request) {
	author := r.URL.Query().Get("author_id")
	sortFlag := r.URL.Query().Get("sort")

	var dbChirps []database.Chirp
	var err error

	if author == "" {
		dbChirps, err = cfg.db.GetAllChirps(r.Context())
	} else {
		var authorID uuid.UUID
		authorID, err = uuid.Parse(author)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Couldn't parse User ID", err)
			return
		}
		dbChirps, err = cfg.db.GetChirpByUserId(r.Context(), authorID)
	}

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't get all chirps from db", err)
		return
	}

	if sortFlag == "desc" {
		sort.Slice(dbChirps, func(i, j int) bool { return dbChirps[i].CreatedAt.After(dbChirps[j].CreatedAt) })
	} else {
		sort.Slice(dbChirps, func(i, j int) bool { return dbChirps[i].CreatedAt.Before(dbChirps[j].CreatedAt) })
	}

	chirps := []Chirp{}
	for _, dbChirp := range dbChirps {
		chirps = append(chirps, Chirp{
			ID:        dbChirp.ID,
			CreatedAt: dbChirp.CreatedAt,
			UpdatedAt: dbChirp.UpdatedAt,
			Body:      dbChirp.Body,
			UserID:    dbChirp.UserID,
		})
	}

	respondWithJSON(w, http.StatusOK, chirps)
}

func (cfg *apiConfig) handlerChirpsGetByID(w http.ResponseWriter, r *http.Request) {
	chirpID, err := uuid.Parse(r.PathValue("chirpID"))
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error(), err)
		return
	}
	dbChirp, err := cfg.db.GetChirpById(r.Context(), chirpID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Couldn't get chirp by id", err)
	}
	respondWithJSON(w, http.StatusOK, Chirp{
		ID:        dbChirp.ID,
		CreatedAt: dbChirp.CreatedAt,
		UpdatedAt: dbChirp.UpdatedAt,
		Body:      dbChirp.Body,
		UserID:    dbChirp.UserID,
	})
}

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Index page!")
}

func ArtistIndex(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Contet-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(artists); err != nil {
		panic(err)
	}
}

func ArtistShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	artistId := vars["artistId"]
	fmt.Fprintln(w, "Artist: ", artistId)
}

func ArtistCreate(w http.ResponseWriter, r *http.Request) {
	var artist Artist
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &artist); err != nil {
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.WriteHeader(422) // unprocessable
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	t := RepoCreateArtist(artist)
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}

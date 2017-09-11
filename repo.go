package main

import "fmt"

var currentId int

var artists Artists

// Seed some data.
func init() {
	RepoCreateArtist(Artist{Name: "Paramore"})
	RepoCreateArtist(Artist{Name: "John Mayer"})
}

func RepoFindArtist(id int) Artist {
	for _, t := range artists {
		if t.Id == id {
			return t
		}
	}
	// Returns an empty Artist if not found.
	return Artist{}
}

func RepoCreateArtist(t Artist) Artist {
	currentId += 1
	t.Id = currentId
	artists = append(artists, t)
	return t
}

func RepoDestroyArtist(id int) error {
	for i, t := range artists {
		if t.Id == id {
			artists = append(artists[:i], artists[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Artist with id of %d to delete", id)
}

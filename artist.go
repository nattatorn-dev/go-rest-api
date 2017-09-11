package main

type Artist struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Genre string `json:"genre"`
	Bio   string `json:"bio"`
}

type Artists []Artist

package main

import (
	commit "flutter/assessment/domain/commit"
	"flutter/assessment/helpers"
)

func main() {
	repositories, err := commit.NewCommitsData()
	helpers.ErrHandler(err)

	repositories.CountLinesPerRepository()

	repositories.ShowRaking()
}

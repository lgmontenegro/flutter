package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"
)

// RepositoryData is where all the data regarding the repository stored in the CSV file
type RepositoryData struct {
	// CVSRecords handle all the CSV lines in a slice of string slices
	CVSRecords [][]string
	// RepositoryNames is a slice of strings
	// with the names of the repository retrieved from the CVS file
	RepositoryNames []string
	// TotalOfActivity handles the sum of the lines in a slice of integer
	// where each line represents the sum of a repository line.
	// With this I garantee a fast access to the total of lines and I can order
	// in ascending or descending direction
	TotalOfLinesPerRepository []int
	// TotalOfLinesGroupByRepository has the total number of lines as index and
	// a slice of string as value with a list of the names of the repositories
	// that have that number of lines/activities in the CVS file
	TotalOfLinesGroupByRepository map[int][]string
}

func main() {
	repositories, err := NewRepositoryData()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	repositories.CountRepoInterations()
}

// NewRepositoryData starts a RepositoryData CVSRecords
// with data from commits.csv
func NewRepositoryData() (*RepositoryData, error) {
	file, err := os.Open("commits.csv")
	if err != nil {
		return nil, err
	}
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()

	if err != nil {
		return nil, err
	}

	return &RepositoryData{
		CVSRecords: records,
	}, nil
}

func (r *RepositoryData) CountRepoInterations() {
	var repositoryCount = make(map[string]int)
	var totalsOrdered = make(map[int]any)
	r.TotalOfLinesGroupByRepository = make(map[int][]string)

	//count each repo
	for _, line := range r.CVSRecords {
		repositoryCount[line[02]]++
		r.RepositoryNames = append(r.RepositoryNames, line[2])
	}

	//grouping repositories by total of commits
	for repositoryName, totalOfCommits := range repositoryCount {
		r.TotalOfLinesGroupByRepository[totalOfCommits] = append(
			r.TotalOfLinesGroupByRepository[totalOfCommits],
			repositoryName,
		)

		totalsOrdered[totalOfCommits] = nil
	}
	//creating ordered index to get top 10
	uniqueTotals := make([]int, 0, len(totalsOrdered))
	for unique := range totalsOrdered {
		uniqueTotals = append(uniqueTotals, unique)
	}

	sort.SliceStable(uniqueTotals, func(i, j int) bool {
		return uniqueTotals[i] < uniqueTotals[j]
	})

	r.TotalOfLinesPerRepository = uniqueTotals
}

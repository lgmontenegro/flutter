package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"
	"strconv"
)

// CommitsData is where all the data regarding the Commits stored in the CSV file
// is read and ranked
type CommitsData struct {
	// CSVRecords handle all the CSV lines in a slice of string slices.
	// Each string slice follows the column order: timestamp, username, repository, files, additions, deletions
	CSVRecords [][]string
	// RepositoryNames is a slice of strings
	// with the names of the repository retrieved from the CVS file
	RepositoryNames []string
	// TotalOfLinesPerRepositoryIndex handles the sum of the lines in a slice of integer
	// where each line represents the sum of a repository line.
	// With this I garantee a fast access to the total of lines and I can order
	// in ascending or descending direction
	TotalOfLinesPerRepositoryIndex []int
	// TotalOfLinesPerRepository is a list of the repositories and its respective sum os lines
	TotalOfLinesPerRepository map[string]int
	// TotalOfLinesGroupByRepository has the total number of lines as index and
	// a slice of string as value with a list of the names of the repositories
	// that have that number of lines/activities in the CVS file
	TotalOfLinesGroupByRepository map[int][]string

	TotalOfFilesPerRepository      map[string]int // TotalOfFilesPerRepository
	TotalOfAddictionsPerRepository map[string]int // TotalOfAddictionsPerRepository
	TotalOfDeletionPerRepository   map[string]int // TotalOfDeletionPerRepository
}

func main() {
	repositories, err := NewCommitsData()
	errHandler(err)

	repositories.CountLinesPerRepository()

	for _, rep := range repositories.RepositoryNames {
		println(repositories.TotalOfFilesPerRepository[rep])
	}

}

// NewCommitsData starts a CommitsData struct
// attribute  with data from commits.csv
// into CSVRecords
func NewCommitsData() (*CommitsData, error) {
	file, err := os.Open("commits.csv")
	if err != nil {
		return nil, err
	}
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()

	if err != nil {
		return nil, err
	}

	return &CommitsData{
		CSVRecords: records,
	}, nil
}

// CountLinesPerRepository count the lines associates
// for each Repository in the CSV file
func (c *CommitsData) CountLinesPerRepository() {
	c.TotalOfLinesPerRepository = make(map[string]int)
	c.TotalOfAddictionsPerRepository = make(map[string]int)
	c.TotalOfDeletionPerRepository = make(map[string]int)
	c.TotalOfFilesPerRepository = make(map[string]int)
	c.TotalOfLinesGroupByRepository = make(map[int][]string)

	totalsOrdered := make(map[int]any)

	//count each repo
	for ln, line := range c.CSVRecords {
		if ln != 0 {
			c.TotalOfLinesPerRepository[line[02]]++
			files, err := strconv.Atoi(line[3])
			errHandler(err)

			addictions, err := strconv.Atoi(line[4])
			errHandler(err)

			deletions, err := strconv.Atoi(line[5])
			errHandler(err)

			c.TotalOfFilesPerRepository[line[02]] = c.TotalOfLinesPerRepository[line[02]] + files
			c.TotalOfAddictionsPerRepository[line[02]] = c.TotalOfAddictionsPerRepository[line[02]] + addictions
			c.TotalOfDeletionPerRepository[line[02]] = c.TotalOfDeletionPerRepository[line[02]] + deletions
			c.RepositoryNames = append(c.RepositoryNames, line[2])
		}
	}

	//grouping repositories by total of commits
	for repositoryName, totalOfCommits := range c.TotalOfLinesPerRepository {
		c.TotalOfLinesGroupByRepository[totalOfCommits] = append(
			c.TotalOfLinesGroupByRepository[totalOfCommits],
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

	c.TotalOfLinesPerRepositoryIndex = uniqueTotals
}

// TopTen return the first ten lines from TotalOfLinesPerRepository
// when ordered descending
func (c *CommitsData) TopTen() []int {
	topTen := c.TotalOfLinesPerRepositoryIndex[len(c.TotalOfLinesPerRepositoryIndex)-10:]

	sort.SliceStable(topTen, func(i, j int) bool {
		return topTen[i] > topTen[j]
	})

	return topTen
}

func errHandler(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

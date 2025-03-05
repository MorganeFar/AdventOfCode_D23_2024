package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	sets := lanParty("../test_input_day23_2024.txt")
	fmt.Println("** Part 1 **")
	fmt.Printf("Number of sets of 3 interconnected computers : %d\n", len(sets))
}

func lanParty(filename string) [][]string {
	allSets := [][]string{}
	lanRelations := readFile(filename)
	for computer, relations := range lanRelations {
		// Look only for computers with t
		match, _ := regexp.MatchString("^t.*$", computer)
		if match {
			// Create sets 3 interconnected computers
			set := createSet(computer, relations, lanRelations)
			for i := range set {
				if !arrExists(allSets, set[i]) {
					allSets = append(allSets, set[i])
				}
			}
		}
	}
	return allSets
}

// Check if 2 arrays contain precisely the same elements (no order)
func isSameArray[N int | string](arr1, arr2 []N) bool {
	same := len(arr1) == len(arr2)
	ind := 0
	for same && ind < len(arr1) {
		same = isInTab(arr1, arr2[ind])
		ind++
	}
	return same
}

// Test if a set is already in the sets array (no element order)
func arrExists(sets [][]string, set []string) bool {
	exists := false
	for i := range sets {
		if exists {
			break
		}
		exists = isSameArray(sets[i], set)
	}
	return exists
}

func createSet(computer string, relations []string, lanRelations map[string][]string) [][]string {
	set := [][]string{}
	for r := range relations {
		computer2 := relations[r]
		relations2 := lanRelations[computer2]
		for i := range relations2 {
			// find a third computer
			if isInTab(relations, relations2[i]) {
				set = append(set, []string{computer, computer2, relations2[i]})
			}
		}
	}
	return set
}

// Read file and create dico of LAN relation
func readFile(filename string) map[string][]string {
	f, _ := os.Open(filename)
	sc := bufio.NewScanner(f)
	dicoRelations := make(map[string][]string)
	tmp := []string{}

	for sc.Scan() {
		line := sc.Text()
		tmp = strings.Split(line, "-")
		// Insert only once each element in both (cf. no direction)
		if !isInTab(dicoRelations[tmp[0]], tmp[1]) {
			dicoRelations[tmp[0]] = append(dicoRelations[tmp[0]], tmp[1])
		}
		if !isInTab(dicoRelations[tmp[1]], tmp[0]) {
			dicoRelations[tmp[1]] = append(dicoRelations[tmp[1]], tmp[0])
		}
	}
	return dicoRelations
}

// Check if an element is in the given array
func isInTab[N int | string](arr []N, elt N) bool {
	isIn := false
	for i := range len(arr) {
		if arr[i] == elt {
			isIn = true
			break
		}
	}
	return isIn
}

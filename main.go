package main

/*
One solution to day 23 of 2024 advent of code problem.
*/

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

func main() {

	argsFile := os.Args
	var filepath string

	if len(argsFile) > 1 {
		filepath = argsFile[1]
	} else {
		// default value
		filepath = "./datas_tests/example_input_day23_2024.txt"
	}

	if fileExists(filepath) {

		var part string

		if len(argsFile) < 3 {
			part = "3"
		} else {
			part = argsFile[2]
		}

		// Part of the problem to execute
		switch part {
		case "1":
			part1(filepath)

		case "2":
			part2(filepath)
		default:
			part1(filepath)
			part2(filepath)
		}

	} else {
		fmt.Println("File not found")
	}
}

func part1(filename string) {
	sets := lanParty(filename)
	fmt.Println("** Part 1 **")
	fmt.Printf("Number of sets of 3 interconnected computers : %d\n", len(sets))
}

func part2(filename string) {
	code := lanPartyP2(filename)
	fmt.Println("** Part 2 **")
	fmt.Printf("Code : %s\n", code)
}

/*
Naive algorithm to solve part 1 :
Find number of 3 interconnected computers sets that contains at least one computer with a name that starts with t.
*/
func lanParty(filename string) [][]string {
	allSets := [][]string{}
	lanRelations, _ := readFile(filename)

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

// Create sets of 3 interconnected vertexes
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

/*
Bron Kerbosch algorithm to solve part 2
*/

func lanPartyP2(filename string) string {
	graph, err := readFile(filename)
	if err == nil {
		/*
			R : clique
			P : potential vertices
			X : excluded vertices
		*/
		R := make(map[string][]string)
		P := make(map[string][]string)
		X := make(map[string][]string)

		// Init map P with all nodes
		for node := range graph {
			P[node] = []string{}
		}

		// Slice with all cliques
		cliques := [][]string{}

		bronKerbosch(R, P, X, graph, &cliques)

		//Find maximal cliques
		maxi := cliques[0]
		for _, clique := range cliques {
			if len(clique) > len(maxi) {
				maxi = clique
			}
		}
		// Sort alphabetic order
		sort.Strings(maxi)
		// Transform to string
		return setToString(maxi)
	}
	return ""
}

func bronKerbosch(R, P, X map[string][]string, graph map[string][]string, cliques *[][]string) {
	// If P and X are both empty, R is a maximal clique
	if len(P) == 0 && len(X) == 0 {
		// Convert R (set) to slice and add to the result
		// length = 0 and capacity = len(R)
		clique := make([]string, 0, len(R))
		for node := range R {
			clique = append(clique, node)
		}
		*cliques = append(*cliques, clique)
		return
	}

	for node := range P {
		neighbors := graph[node]

		PNew := intersection(P, neighbors)
		XNew := intersection(X, neighbors)

		bronKerbosch(union(R, node), PNew, XNew, graph, cliques)

		// Update P and X (backtracking)
		P = remove(P, node)
		X = add(X, node)
	}
}

// Find the common elements bewteen 2 sets
func intersection(a map[string][]string, b []string) map[string][]string {
	result := make(map[string][]string)
	for key := range a {
		if isInTab(b, key) {
			result[key] = []string{}
		}
	}
	return result
}

// Copy and add an element to a set
func union(set map[string][]string, element string) map[string][]string {
	result := make(map[string][]string)
	for key := range set {
		result[key] = []string{}
	}
	result[element] = []string{}
	return result
}

// Add an element to a set
func add(set map[string][]string, element string) map[string][]string {
	set[element] = []string{}
	return set
}

// Remove an element from a set
func remove(set map[string][]string, element string) map[string][]string {
	delete(set, element)
	return set
}

func setToString(biggestSet []string) string {
	var buffer bytes.Buffer
	size := len(biggestSet)
	for i := 0; i < size; i++ {
		buffer.WriteString(biggestSet[i])
		if i < size-1 {
			buffer.WriteString(",")
		}
	}
	return buffer.String()
}

// Read file and create map of LAN relation
/*
Ex :
{
	vertex1 : [ all neighbors of vertex1]
	vertex2 : [ all neighbors of vertex2]
	...
}
*/
func readFile(filename string) (map[string][]string, error) {
	f, err := os.Open(filename)
	dicoRelations := make(map[string][]string)
	if err == nil {
		sc := bufio.NewScanner(f)
		for sc.Scan() {
			line := sc.Text()
			tmp := strings.Split(line, "-")
			// Insert only once each element in both (cf. no direction)
			if !isInTab(dicoRelations[tmp[0]], tmp[1]) {
				dicoRelations[tmp[0]] = append(dicoRelations[tmp[0]], tmp[1])
			}
			if !isInTab(dicoRelations[tmp[1]], tmp[0]) {
				dicoRelations[tmp[1]] = append(dicoRelations[tmp[1]], tmp[0])
			}
		}

		return dicoRelations, nil

	} else {
		fmt.Println("File not found")
		return dicoRelations, err
	}

}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

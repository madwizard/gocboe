// Package map provides functions operating on map files:
// reading, parsing, saving.
package _map

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// separators is a string containing all characters that can be
// field separators in map field.
var separators string = "#*<v>^:!@$&hHbB"

type Tile struct {
	Terrain int
	Separator string
	Special int
}

type Row []Tile

type Map struct {
	town bool
	mapHeight int
	mapLength int
	Map []Row
}

// IsTown returns information if map is a town
// Returns boolean true if current map is a town
// and a boolean false if outside town
func (m Map) IsTown() bool {
	if m.town == true {
		return true
	}
	return false
}

func (m Map) Height() int {
	return m.mapHeight
}

func (m Map) Length() int {
	return m.mapLength
}

// Read reads up map from .map file. It doesn't parse the file
// nor does it read map .xml or scenarios.
func Read(path string) (Map, error) {
	mapFile, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer mapFile.Close()

	var ParsedMap Map
	if strings.Contains(path, "town") {
		ParsedMap.town = true
	}

	scanner := bufio.NewScanner(mapFile)
	i := 0
	for scanner.Scan() {
		var ParsedRow Row
		ParsedRow.parse(scanner.Text())
		ParsedMap.mapLength = len(ParsedRow)
		ParsedMap.appendRow(ParsedRow)
		i++
	}
	ParsedMap.mapHeight = i
	return ParsedMap, scanner.Err()
}

// parseTile is private function which reads a field string from parsed map
// and loads it into Tile structure
func parseTile(input string) (Tile, error) {
	var tile Tile
	if strings.ContainsAny(input, separators) {
		b := []byte(separators)
		for i, _ := range b {
			s := string(b[i])
			if strings.Contains(input, s) {
				tmp := strings.Split(input, s)
				tile.Separator = s
				tile.Terrain, _ = strconv.Atoi(tmp[0])
				tile.Special, _ = strconv.Atoi(tmp[1])
			}
		}
	} else {
		tile.Separator = "-"
		tile.Special = -1
		tile.Terrain, _ = strconv.Atoi(input)
	}
	return tile, nil
}

// parse accepts a row of map data and sends it to Tile parser
// filling up Tile structure.
func (r* Row) parse(mapLine string) error {
		mapLine = strings.TrimSpace(mapLine)
		fields := strings.Split(mapLine, ",")
		for _, field := range fields {
			tile, _ := parseTile(field)
			*r = append(*r, tile)
		}
	return nil
}

// appendRow adds row passed as a function argument
// to slice of rows in Map structure
func (m* Map) appendRow(r Row) error {
	(*m).Map = append((*m).Map, r)
	return nil
}

func (m Map) Print() {
	for _, row := range m.Map {
		row.print()
	}
}

func (r Row) print() {
	for _, tile := range r {
		tile.print()
		fmt.Print(",")
	}
	fmt.Println()
}

func (t Tile) print() {
	fmt.Print(t.Terrain)
	if t.Separator != "-" {
		fmt.Print(t.Separator)
	}
	if t.Special != -1 {
		fmt.Print(t.Special)
	}
}
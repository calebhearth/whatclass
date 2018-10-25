package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
)

type jsonError struct {
	Error string `json:"error"`
}

func main() {
	args := os.Args[1:]
	if len(args) > 0 && args[0] == "serve" {
		port := os.Getenv("PORT")
		if os.Getenv("PORT") == "" {
			port = "3000"
		}
		http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
			statString := request.FormValue("text")
			statSlice := strings.Split(statString, " ")
			if len(statSlice) != 6 {
				err := jsonError{fmt.Sprintf("Usage: %s [Str] [Dex] [Con] [Int] [Wis] [Cha]", request.FormValue("command"))}
				body, e := json.Marshal(err)
				if e != nil {
					log.Fatalf("%s", e)
				}
				response.Write(body)
				return
			}

			statline, err := parseStats(statSlice)
			if err != nil {
				log.Fatal(err)
			}
			statline.score()
			body, e := json.Marshal(classes)
			if e != nil {
				log.Fatalf("%s", e)
			}
			response.Write(body)
		})
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
	} else if len(args) == 6 {
		statline, err := parseStats(args)
		if err != nil {
			log.Fatal(err)
		}
		statline.score()
		statline.print()
	} else {
		fmt.Printf("Usage: %s [Str] [Dex] [Con] [Int] [Wis] [Cha]\n", os.Args[0])
		os.Exit(0)
	}
}

func parseStats(args []string) (stats, error) {
	values := []float64{}
	for _, s := range args {
		result, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return nil, fmt.Errorf("Arguments must be numeric: %s", err)
		}
		values = append(values, result)
	}
	return stats{
		Str: values[0],
		Dex: values[1],
		Con: values[2],
		Int: values[3],
		Wis: values[4],
		Cha: values[5],
	}, nil
}

func (statline stats) score() {
	for i, c := range classes {
		classes[i].Score = score(statline, c)
	}
	sort.Sort(sort.Reverse(classes))
}

func (statline stats) print() {
	fmt.Printf("Stats: %2.0f STR  %2.0f DEX  %2.0f CON  %2.0f INT  %2.0f WIS  %2.0f CHA\n", statline[Str], statline[Dex], statline[Con], statline[Int], statline[Wis], statline[Cha])
	fmt.Printf("\n%-12s %s  %-6s  %-6s  %-6s  %-6s\n", "Class", "Match", "1st", "2nd", "Dump", "Alt Dump")
	for _, bc := range classes {
		fmt.Printf("%-12s %4.0f%%  %2.0f %s  %2.0f %s  %2.0f %s", bc.Name, bc.Score*100, statline[bc.Primary], bc.Primary, statline[bc.Secondary], bc.Secondary, statline[bc.Dump], bc.Dump)
		if bc.Dump2 != None {
			fmt.Printf("  %.0f %s", statline[bc.Dump2], bc.Dump2)
		}
		fmt.Println()
	}
}

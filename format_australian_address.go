package main

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type formatted_address struct {
	StreetAddress string `json:"street_address"`
	Suburb        string `json:"suburb"`
	State         string `json:"state"`
	Postcode      int    `json:"postcode"`
}

func format_address(addr string) formatted_address {
	var (
		street_address string
		suburb         string
		state          string
		postcode       int
	)

	// postcode
	pc_pattern := `\b\d{3,4}\b`
	pc_regex, err := regexp.Compile(pc_pattern)

	matched_postcodes := pc_regex.FindAllString(addr, -1)
	if len(matched_postcodes) > 0 {
		postcode, err = strconv.Atoi(matched_postcodes[len(matched_postcodes)-1]) // last match
		if err != nil {
			panic(err)
		}
	}

	// state
	state_pattern := `(?i)\b(NT|NSW|VIC|QLD|WA|SA|ACT|TAS)\b`
	state_regex, err := regexp.Compile(state_pattern)
	if err != nil {
		panic(err)
	}
	state = state_regex.FindString(addr)

	// suburb and street address
	split_address := strings.SplitN(addr, state, 2)
	cleaned_address := strings.Trim(strings.TrimSpace(split_address[0]), ",")
	streetaddr_suburb := strings.Split(cleaned_address, ",")
	if len(streetaddr_suburb) > 0 {
		last_idx := len(streetaddr_suburb) - 1
		suburb = strings.TrimSpace(streetaddr_suburb[last_idx])
		street_address = strings.TrimSpace(strings.Join(streetaddr_suburb[:last_idx], ""))
	}

	return formatted_address{
		StreetAddress: street_address,
		Suburb:        suburb,
		State:         state,
		Postcode:      postcode,
	}
}

func main() {
	start_time := time.Now()

	addresses := []string{
		"9/49 The Causeway, Maroubra NSW 2035",
		"8 Horrocks Street, Torrens ACT 2607",
		"61 Mary Gillespie Ave, Gungahlin ACT 2912",
		"269 Charleston Road, Deanside VIC 3336",
		"218/2 Kerridge Street, Kingston ACT 2604",
		"2/38 Villiers Street, New Farm QLD 4005",
		"11/11 Belmont Avenue, Wollstonecraft NSW 2065",
		"58/73-87 Caboolture River Road, Morayfield QLD 4506",
		"2/38 Villiers Street, New Farm QLD 4005",
	}

	for _, address := range addresses {
		fa := format_address(address)

		data, err := json.Marshal(fa)
		if err != nil {
			panic(err)
		}

		fmt.Println(address, "-", string(data))
	}

	elapsed_time := time.Since(start_time)
	fmt.Println("total time taken :", elapsed_time.Seconds(), "seconds")
}

package sample

// go mod init "github.com/mitchellh/mapstructure"

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

// Record is a weather record
type Record struct {
	Time        time.Time
	Station     string
	Temperature float64 //`json: "temperature"` //celsius
	Rain        float64 //millimeter
}

func readRecord(r io.Reader) (Record, error) {
	var rec Record
	dec := json.NewDecoder(r)
	if err := dec.Decode(&rec); err != nil {
		return Record{}, err
	}
	return rec, nil
}

func main() {

	file, err := os.Open("record.json")
	fmt.Printf("file: %v\n", file)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	rec, err := readRecord(file)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", rec)
}

func parseTime(ts string) (time.Time, error) {
	return time.Parse("2006-01-02 15:04:05 PM", ts)
}

// laggingStations return stations that are lagging in their check time
func laggingStations(r io.Reader, timeout time.Duration) ([]string, error) {

	var reply struct {
		LastCheckTime string
		Stations      []struct {
			Name      string
			Status    string
			LastCheck struct {
				Time string
			}
		}
	}
	dec := json.NewDecoder(r)
	if err := dec.Decode(&reply); err != nil {
		return nil, err
	}
	checkTime, err := parseTime(reply.LastCheckTime)
	if err != nil {
		return nil, err
	}
	var lagging []string
	for _, station := range reply.Stations {
		if station.Status != "Active" {
			continue
		}
		lastCheck, err := parseTime(station.
			LastCheck.Time)
		if err != nil {
			return nil, err
		}
		if checkTime.Sub(lastCheck) > timeout {
			lagging = append(lagging, station.Name)
		}
	}
	return lagging, nil
}

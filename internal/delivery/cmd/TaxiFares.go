package cmd

import (
	constant "TaxiFares/model/constant"
	"TaxiFares/model/entity"
	"errors"
	"log"

	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func (d *CmdDelivery) TaxiFares() {

	for {
		scanner := bufio.NewScanner(os.Stdin)

		var records []entity.Record
		var prevTime time.Time
		var prevDistance float64
		var err error

		for scanner.Scan() {
			line := scanner.Text()
			if line == "" {
				break
			}

			//validate input
			prevTime, prevDistance, _, err = d.validateInput(line, prevTime, prevDistance, records)
			if err != nil {
				fmt.Println("error validating input: ", err)
				log.Printf("error validating input: %v", err)
				continue
			}

		}

		if errScanner := scanner.Err(); errScanner != nil || err != nil {
			return
		}

		err = d.processInput(records)
		if err != nil {
			fmt.Println("error processing input: ", err)
			log.Printf("error processing input: %v", err)
		}

	}
}

func (d *CmdDelivery) validateInput(input string, prevTime time.Time, prevDistance float64, records []entity.Record) (timeVal time.Time, distance, traveledDistance float64, err error) {

	//validate the part have time and mileage
	parts := strings.Split(input, " ")
	if len(parts) != 2 {
		err = errors.New("invalid input format")
		return
	}

	//split into 2 part
	timeStr := parts[0]
	distanceStr := parts[1]

	//parse elapsed time
	timeVal, err = time.Parse("15:04:05.000", timeStr)
	if err != nil {
		err = errors.New("error parsing time")
		return
	}

	//parse mileage
	distance, err = strconv.ParseFloat(distanceStr, 64)
	if err != nil {
		err = errors.New("error parsing distance")
		return
	}

	//compare current time with prev time
	elapsedTime := timeVal.Sub(prevTime).Minutes()

	// check past time
	if elapsedTime <= 0 && !prevTime.IsZero() {
		err = fmt.Errorf("invalid time interval: %v minutes", elapsedTime)
		return
	}

	// check interval more than 5 minutes apart
	if elapsedTime > constant.MaxInterval {
		err = fmt.Errorf("time interval exceeds maximum allowed: %v minutes", elapsedTime)
		return
	}

	//get travel distance
	traveledDistance = distance - prevDistance

	//validate the travel distance not in decrease value
	if traveledDistance < 0 {
		err = fmt.Errorf("invalid distance: %v meters", traveledDistance)
		return
	}

	record := entity.Record{
		Time:     timeVal,
		Distance: distance,
		Diff:     traveledDistance,
	}
	records = append(records, record)

	return
}

func (d *CmdDelivery) processInput(records []entity.Record) (err error) {
	//validate must have 2 lines data
	if len(records) < 2 {
		err = errors.New("need at least two records")
		fmt.Println(err)
		return
	}

	if records[len(records)-1].Distance == 0 {
		err = errors.New("total mileage is 0.0m")
		fmt.Println(err)
		return
	}

	fare := d.taxiFaresUc.CalculateFare(records[len(records)-1].Distance)
	fmt.Println("fare : ", fare)

	sort.Slice(records, func(i, j int) bool {
		return records[i].Diff > records[j].Diff
	})

	for _, record := range records {
		fmt.Printf("%s %.1f %.1f\n", record.Time.Format("15:04:05.000"), record.Distance, record.Diff)
	}

	return
}

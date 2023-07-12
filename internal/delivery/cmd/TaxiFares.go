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
		var (
			records      []entity.Record
			prevTime     time.Time
			prevDistance float64
			err          error
		)

		for scanner.Scan() {
			line := scanner.Text()
			if line == "" {
				break
			}

			//validate input
			timeVal, distance, traveledDistance, err := d.validateInput(line, prevTime, prevDistance)
			if err != nil {
				fmt.Println("error validating input: ", err)
				log.Printf("error validating input: %v", err)
				continue
			}

			record := entity.Record{
				Time:     timeVal,
				Distance: distance,
				Diff:     traveledDistance,
			}
			records = append(records, record)

			prevTime = timeVal
			prevDistance = distance

		}

		if errScanner := scanner.Err(); errScanner != nil {
			continue
		}

		if err != nil {
			continue
		}

		//validate must have 2 lines data
		if len(records) < 2 {
			fmt.Println("need at least two records")
			log.Print("need at least two records")
			continue
		}

		if records[len(records)-1].Distance == 0 {
			fmt.Println("total mileage is 0.0m")
			log.Print("total mileage is 0.0m")
			continue
		}

		fare := d.taxiFaresUc.CalculateFare(records[len(records)-1].Distance)
		fmt.Println(fare)

		sort.Slice(records, func(i, j int) bool {
			return records[i].Diff > records[j].Diff
		})

		for _, record := range records {
			fmt.Printf("%s %.1f %.1f\n", record.Time.Format("15:04:05.000"), record.Distance, record.Diff)
		}

	}
}

func (d *CmdDelivery) validateInput(input string, prevTime time.Time, prevDistance float64) (timeVal time.Time, distance, traveledDistance float64, err error) {
	parts := strings.Split(input, " ")
	if len(parts) != 2 {
		err = errors.New("invalid input format")
		return
	}

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

	traveledDistance = distance - prevDistance

	if traveledDistance < 0 {
		err = fmt.Errorf("invalid distance: %v meters", traveledDistance)
		return
	}

	return
}

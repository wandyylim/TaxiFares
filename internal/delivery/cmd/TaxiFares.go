package cmd

import (
	taxifares "TaxiFares/internal/usecase/TaxiFares"
	constant "TaxiFares/model/constant"
	"TaxiFares/model/entity"

	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func TaxiFares(taxiFaresUC *taxifares.TaxiFaresUsecase) {

	scanner := bufio.NewScanner(os.Stdin)
	var records []entity.Record
	var prevTime time.Time
	var prevDistance float64

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		parts := strings.Split(line, " ")
		if len(parts) != 2 {
			fmt.Fprintln(os.Stderr, "Invalid input format:", line)
			os.Exit(1)
		}

		timeStr := parts[0]
		distanceStr := parts[1]

		timeVal, err := time.Parse("15:04:05.000", timeStr)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error parsing time:", err)
			os.Exit(1)
		}

		distance, err := strconv.ParseFloat(distanceStr, 64)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error parsing distance:", err)
			os.Exit(1)
			continue
		}

		elapsedTime := timeVal.Sub(prevTime).Minutes()

		if elapsedTime <= 0 {
			fmt.Fprintln(os.Stderr, "Invalid time interval:", elapsedTime, "minutes")
			os.Exit(1)
		}

		if elapsedTime > constant.MaxInterval {
			fmt.Fprintln(os.Stderr, "Time interval exceeds maximum allowed:", elapsedTime, "minutes")
			os.Exit(1)
		}

		traveledDistance := distance - prevDistance

		if traveledDistance < 0 {
			fmt.Fprintln(os.Stderr, "Invalid distance:", traveledDistance, "meters")
			os.Exit(1)
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

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading standard input:", err)
		os.Exit(1)
	}

	if len(records) < 2 {
		fmt.Fprintln(os.Stderr, "Insufficient data. Need at least two records.")
		os.Exit(1)
	}

	if records[0].Distance == 0 {
		fmt.Fprintln(os.Stderr, "Total mileage is 0.0m")
		os.Exit(1)
	}

	fare := taxiFaresUC.CalculateFare(records[len(records)-1].Time.Sub(records[0].Time).Minutes(), records[len(records)-1].Distance)
	fmt.Println(fare)

	sort.Slice(records, func(i, j int) bool {
		return records[i].Diff > records[j].Diff
	})

	for _, record := range records {
		fmt.Printf("%s %.1f %.1f\n", record.Time.Format("15:04:05.000"), record.Distance, record.Diff)
	}

	return
}

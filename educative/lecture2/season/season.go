package season

import "fmt"

func season(month int) string {
	var seasonName string
	switch month {
	case 12, 1, 2:
		seasonName = "Winter"
	case 3, 4, 5:
		seasonName = "Spring"
	case 6, 7, 8:
		seasonName = "Summer"
	case 9, 10, 11:
		seasonName = "Autumn"
	default:
		seasonName = "Invalid month"
	}

	return seasonName
}

// Run the above function
func Run() {
	month := 12
	fmt.Printf("month %d is season of %s \n\n", month, season(month))
}

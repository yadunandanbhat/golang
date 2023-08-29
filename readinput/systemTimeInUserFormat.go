package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	fmt.Println("----Time Formatter----")
	fmt.Println("1. Tue Oct 4 18:54:28 2022")
	fmt.Println("2. Tue Oct 4 18:54:28 +07 2022")
	fmt.Println("3. Tue Oct 04 18:54:28 +0700 2022")
	fmt.Println("4. 04 Oct 22 18:54 +07")
	fmt.Println("5. 04 Oct 22 18:56 +0700")
	fmt.Println("6. Tuesday, 04-Oct-22 18:54:28 +07")
	fmt.Println("7. Tue, 04 Oct 2022 18:54:28 +07")
	fmt.Println("8. Tue, 04 Oct 2022 18:54:28 +0700")
	fmt.Println("9. 2022-10-04T18:54:28+07:00")
	fmt.Println("10. 2022-10-04T18:54:28.0977814+07:00")
	fmt.Println("11. 04/10 06:54:28PM '22 +0700")
	fmt.Println("12. 6:54PM")
	fmt.Println("Please input the number of any standard time format you want the current time in: ")
	var format string
	currentTime := time.Now()
	n, err := fmt.Scanln(&format)
	if err != nil || n <= 0 {
		log.Fatal(err)
	}
	switch format {
	case "1":
		fmt.Println(currentTime.Format(time.ANSIC))
	case "2":
		fmt.Println(currentTime.Format(time.UnixDate))
	case "3":
		fmt.Println(currentTime.Format(time.RubyDate))
	case "4":
		fmt.Println(currentTime.Format(time.RFC822))
	case "5":
		fmt.Println(currentTime.Format(time.RFC822Z))
	case "6":
		fmt.Println(currentTime.Format(time.RFC850))
	case "7":
		fmt.Println(currentTime.Format(time.RFC1123))
	case "8":
		fmt.Println(currentTime.Format(time.RFC1123Z))
	case "9":
		fmt.Println(currentTime.Format(time.RFC3339))
	case "10":
		fmt.Println(currentTime.Format(time.RFC3339Nano))
	case "11":
		fmt.Println(currentTime.Format(time.Layout))
	case "12":
		fmt.Println(currentTime.Format(time.Kitchen))
	default:
		fmt.Println("You entered a wrong option, so the default output is " + currentTime.Format(time.ANSIC))
	}

}

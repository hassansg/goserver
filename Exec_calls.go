//Go Algorithm for sysadmin exec calls
//Implementation 3 for Network Dynamics

package main

import (
	"log"
	"os/exec"
)

func main() {

	outv1, err := exec.Command("ping", "-c", "2", "8.8.8.8").Output()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("The results are -->  %s\n", outv1)

	// run the date command and store the output
	outv2, err := exec.Command("date").Output()
	if err != nil {
		log.Print(err)
	}
	//comparison date and time for accuracy, mean is taken
	log.Printf("The date is --> %s ", outv2)

}

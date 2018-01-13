/*
Plugin check_go_fun
2018_01_13
studost

Copy of check-graylog2-steam from mariussturm
Just for fun

CHANGELOG:
2018_01_13 1.0.0 initial version

EXAMPLES:

*/

package main

// IMPORT
import (
	"flag"
	"fmt"
	"github.com/fractalcat/nagiosplugin"
	"os"
        "math/rand"
        "time"
)

//FLAGS
var warning *int
var critical *int
var message *string

// FLAG INIT
func init() {
	warning = flag.String("warning", "<INT>", "Threshold for WARNING (int)")
	critical = flag.String("critical", "<INT>", "Threshold for CRITICAL (int)")
	message = flag.String("message", "<message>", "A simple message")
}

func main() {
        rand.Seed(time.Now().UTC().UnixNano())
	flag.Parse()
	checkArguments()

	check := nagiosplugin.NewCheck()
	defer check.Finish()

	//var random_result *integer
	// TODO:
	// generate random result
	random_result := random(1, 100)
	fmt.Println(random_result)

	if random_result > critical {
		nagiosplugin.Exit(nagiosplugin.CRITICAL, fmt.Sprintf("%g", random_result)+" is greater than "+*critical)
	}
	if random_result > warning {
		nagiosplugin.Exit(nagiosplugin.WARNING, fmt.Sprintf("%g", random_result)+" is greater than "+*warning)
	} else {
		nagiosplugin.Exit(nagiosplugin.OK, fmt.Sprintf("%g", random_result)+" is below all thresholds")
	}

	// check.AddResult(nagiosplugin.OK, "No stream alerts triggered for stream: " + stream_title)
}

func checkArguments() {
	if *critical == "<INT>" {
		fmt.Println("Usage:")
		flag.PrintDefaults()
		os.Exit(1)
	}
}

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

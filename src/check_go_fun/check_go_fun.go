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
	"math/rand"
	"time"
)

//FLAGS
var warning int
var critical int

//var message string

// FLAG INIT
func init() {
	flag.IntVar(&warning, "warning", 80, "Threshold for WARNING (int)")
	flag.IntVar(&critical, "critical", 90, "Threshold for CRITICAL (int)")
	//message = flag.String("message", "<message>", "A simple message")
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	flag.Parse()
	//checkArguments()

	check := nagiosplugin.NewCheck()
	defer check.Finish()

	//var random_result *integer
	// TODO:
	// generate random result
	random_result := random(1, 100)
	fmt.Println(random_result)

	if random_result > critical {
		nagiosplugin.Exit(nagiosplugin.CRITICAL, fmt.Sprintf("%d", random_result)+" is by way greater than "+fmt.Sprintf("%d", critical))
	} else if random_result > warning {
		nagiosplugin.Exit(nagiosplugin.WARNING, fmt.Sprintf("%d", random_result)+" is a little bit greater than "+fmt.Sprintf("%d", warning))
	} else {
		nagiosplugin.Exit(nagiosplugin.OK, fmt.Sprintf("%d", random_result)+" is really below all thresholds")
	}

	// check.AddResult(nagiosplugin.OK, "No stream alerts triggered for stream: " + stream_title)
}

/*
func checkArguments() {
	if critical == "<INT>" {
		fmt.Println("Usage:")
		flag.PrintDefaults()
		os.Exit(1)
	}
}
*/

/*///////////////////////////////////////////////
FUNC RANDOM
arg1 min int
arg2 max int
return result int
/*/ //////////////////////////////////////////////
func random(min, max int) int {
	return rand.Intn(max-min) + min
}

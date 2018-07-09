// designate destination
package jobs

// import lib
import (
	"fmt"

	"github.com/revel/modules/jobs/app/jobs"
	"github.com/revel/revel"
)

type LogRequestor struct{}

// for implementing Cron.Job

func (r LogRequestor) Run() {
	fmt.Println("Running")
	// 1. Request URL (using local static asset)
	// 2. Save the CSV File ()
	// 3. Parse the CSV File ()
	// 4. Push the Data on to the Elastic Search ()
	// 5. Write the self Log ()

}

// for scheduling Job

func init() {
	revel.OnAppStart(func() {
		jobs.Schedule("@every 1m", LogRequestor{})
	})
}

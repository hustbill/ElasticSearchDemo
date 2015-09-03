package main

import (
	/*
    "flag"
	"math"
	"math/rand"
    "time"
    */
    
    "fmt"    
    "log"
    "net/http"
    "os"
    
    "./routers"
    "github.com/prometheus/client_golang/prometheus"
    
    "github.com/golang/protobuf/proto"
    dto "github.com/prometheus/client_model/go"
    "math"
    
   
   _ "net/http/pprof"
)

func main() {
    
    router := routers.NewRouter()
    var port string
    fmt.Printf("len(os.Args) = %d ", len(os.Args))
   
    if ( len(os.Args) <= 1) {
        fmt.Println("Please speficy the port or use default port 8090")
        port = "8090"
    } else {
         port = os.Args[1]
    }
    fmt.Printf("port = %s", port)
    
    
    // histogram 
	temps := prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "pond_temperature_celsius",
		Help:    "The temperature of the frog pond.", // Sorry, we can't measure how badly it smells.
		Buckets: prometheus.LinearBuckets(20, 5, 5),  // 5 buckets, each 5 centigrade wide.
	})

	// Simulate some observations.
	for i := 0; i < 1000; i++ {
		temps.Observe(30 + math.Floor(120*math.Sin(float64(i)*0.1))/10)
	}

	// Just for demonstration, let's check the state of the histogram by
	// (ab)using its Write method (which is usually only used by Prometheus
	// internally).
	metric := &dto.Metric{}
	temps.Write(metric)
	fmt.Println(proto.MarshalTextString(metric))

    router.Handle("/metrics", prometheus.Handler())
    
    //log.Fatal(http.ListenAndServe(":8090", router))    
    log.Fatal(http.ListenAndServe(":" + port, router))
}



/*
 // backup code for duration , rpc

var (
	addr              = flag.String("listen-address", ":8080", "The address to listen on for HTTP requests.")
	uniformDomain     = flag.Float64("uniform.domain", 200, "The domain for the uniform distribution.")
	normDomain        = flag.Float64("normal.domain", 200, "The domain for the normal distribution.")
	normMean          = flag.Float64("normal.mean", 10, "The mean for the normal distribution.")
	oscillationPeriod = flag.Duration("oscillation-period", 10*time.Minute, "The duration of the rate oscillation period.")
)


var (
	// Create a summary to track fictional interservice RPC latencies for three
	// distinct services with different latency distributions. These services are
	// differentiated via a "service" label.
	rpcDurations = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name: "rpc_durations_microseconds",
			Help: "RPC latency distributions.",
		},
		[]string{"service"},
	)
	// The same as above, but now as a histogram, and only for the normal
	// distribution. The buckets are targeted to the parameters of the
	// normal distribution, with 20 buckets centered on the mean, each
	// half-sigma wide.
	rpcDurationsHistogram = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "rpc_durations_histogram_microseconds",
		Help:    "RPC latency distributions.",
		Buckets: prometheus.LinearBuckets(*normMean-5**normDomain, .5**normDomain, 20),
	})
)







	flag.Parse()

	start := time.Now()

	oscillationFactor := func() float64 {
		return 2 + math.Sin(math.Sin(2*math.Pi*float64(time.Since(start))/float64(*oscillationPeriod)))
	}

	// Periodically record some sample latencies for the three services.
	go func() {
		for {
			v := rand.Float64() * *uniformDomain
			rpcDurations.WithLabelValues("uniform").Observe(v)
			time.Sleep(time.Duration(100*oscillationFactor()) * time.Millisecond)
		}
	}()

	go func() {
		for {
			v := (rand.NormFloat64() * *normDomain) + *normMean
			rpcDurations.WithLabelValues("normal").Observe(v)
			rpcDurationsHistogram.Observe(v)
			time.Sleep(time.Duration(75*oscillationFactor()) * time.Millisecond)
		}
	}()

	go func() {
		for {
			v := rand.ExpFloat64()
			rpcDurations.WithLabelValues("exponential").Observe(v)
			time.Sleep(time.Duration(50*oscillationFactor()) * time.Millisecond)
		}
	}()
    

*/

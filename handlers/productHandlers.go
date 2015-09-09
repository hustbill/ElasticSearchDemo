package handlers

import (
	"encoding/json"
    
    "flag"

    "time"
    "io"
    "io/ioutil"
    
	"fmt"
	"net/http"
   "../models"
   "../daos"
   
	"github.com/gorilla/mux"
    "github.com/prometheus/client_golang/prometheus"
)

var products models.Products

var (
    postReqCounter = prometheus.NewCounter(prometheus.CounterOpts{
        Namespace: "BEB",
        Subsystem: "ReviewSevice",
        Name:      "review_post_request",
        Help:      "The number of post request",
    })
)

var (
    getReqCounter = prometheus.NewCounter(prometheus.CounterOpts{
        Namespace: "BEB",
        Subsystem: "ReviewSevice",
        Name:      "review_get_request",
        Help:      "The number of get request",
    })
)

var (
	addr              = flag.String("listen-address", ":8080", "The address to listen on for HTTP requests.")
	uniformDomain     = flag.Float64("uniform.domain", 200, "The domain for the uniform distribution.")
	normDomain        = flag.Float64("normal.domain", 200, "The domain for the normal distribution.")
	normMean          = flag.Float64("normal.mean", 10, "The mean for the normal distribution.")
	oscillationPeriod = flag.Duration("oscillation-period", 10*time.Minute, "The duration of the rate oscillation period.")
)

var (
    // Create a summary to track fictional interservice RPC latencies for three
    // distinct services. These services are differentiated via a "service" label.
    rpcDurations = prometheus.NewSummaryVec(
        prometheus.SummaryOpts{
            Name: "rpc_durations_microseconds",
            Help: "RPC latencey distributions.",
        },
        []string{"service"},
    )
    
    // Create a histogram for track fictional interservice RPC latencies for three
    // distinct services. 
    rpcDurationsHistogram = prometheus.NewHistogram(prometheus.HistogramOpts {
        Name: "rpc_durations_histogram_microseconds",
        Help: "RPC latency distibutions.",
        Buckets: prometheus.LinearBuckets(*normMean - 5**normDomain, .5**normDomain, 20),
        
    })
)



func init() {
	// Register the post and get review request counter
    prometheus.MustRegister(postReqCounter)
    prometheus.MustRegister(getReqCounter)
    
    // Register the summary and the histogram with Prometheus's default registry
    prometheus.MustRegister(rpcDurations)
    prometheus.MustRegister(rpcDurationsHistogram)
}




func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "welcome!")
}


func ProductIndex(w http.ResponseWriter, r *http.Request) {
     
    w.Header().Set("Context-Type", "application/json;charset = UTF-8")
    w.WriteHeader(http.StatusOK)


    if err := json.NewEncoder(w).Encode(products); err != nil {
        panic(err)
    }
}

func ProductShow(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    var name string
    name = vars["name"]
    fmt.Println(name)
    product := daos.RepoFindProduct(name)
    fmt.Println(product)
    fmt.Println(product.Name)

    if len(product.Name) >0 {
        w.Header().Set("Context-Type", "application/json; charset=UTF-8")
        w.WriteHeader(http.StatusOK)
        if err := json.NewEncoder(w).Encode(product); err!= nil {
            panic(err)
        }
        return
    }
    
	// If we didn't find it, 404
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
    // if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
   //      panic(err)
   //  }
}


func ProductCreate(w http.ResponseWriter, r *http.Request) {
    var product models.Product
    body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048567))
    
    if err != nil {
        panic(err)
    }
    
    if err := r.Body.Close(); err != nil {
        panic(err)
    }
    
    if err := json.Unmarshal(body, &product); err != nil {
        w.Header().Set("Context-Type", "application/json;charset=UTF-8")
        w.WriteHeader(422)
        if err := json.NewEncoder(w).Encode(err); err!= nil {
            panic(err)
        }
    }
    
    p := daos.RepoCreateProduct(product)
    w.Header().Set("Context-Type", "application/json;charset=UTF-8")
    w.WriteHeader(http.StatusCreated)
    if err := json.NewEncoder(w).Encode(p); err != nil {
        panic(err)
    }
}

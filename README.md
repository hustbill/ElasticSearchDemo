# Review Service API in Golang 


├── conf  
│   └── app.conf  
├── controllers  
│   └── reviewController.go  
├── main.go  
├── models  
├── routers  
│   └── router.go  
├── static  
│   ├── css  
│   ├── img  
│   └── js  
├── tests  
│   └── default_test.go  
└── views  
    └── index.tpl  


## Use Prometheus

### Installing 
$ export GOPATH= {your home folder}/golang
$ cd $GOPATH
$ go get github.com/prometheus/client_golang/prometheus


### Run 

1. start Prometheus Server  
$ cd review-service  && mkdir prometheus && cd prometheus
 Download the latest release of Prometheus for your platform, then extract and run it   
$ tar xvfz prometheus-*.tar.gz  
$ cp ../prometheus.yml.sample ./promethesual.yml  
$ ./prometheus prometheus.yml  

2. start Client with our APIs  
$ cd review-service/go  
$ go run *.go  
our review services will be started on port 8090  
  
### View by Web Browser
1. server side 
http://localhost:9090/

2. client side
http://localhost:8090/metrics

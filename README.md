# Product Service API in Golang 


├── conf  
│   └── app.conf  
├── controllers  
│   └── productController.go  
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


## Use elastic go lib

### Installing 
import "github.com/olivere/elastic"

Package elastic provides an interface to the Elasticsearch server 

### Run 
go run main.go

### View by Web Browser
1. Our golang API side  - (fetch from database)  
http://localhost:8090/products/6  

2. ElasticSearch server side (search from ElasticSearch server)  
http://127.0.0.1:9200/products  


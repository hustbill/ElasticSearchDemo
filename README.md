# Product Service API in Golang 


├── conf  
│   └── app.conf  
├── controllers  
│   └── productController.go  
├── main.go  
├── elasticClient.go  
├── models  
│   └── product.go  
├── kafka  
│   └── consumer.go  
│   └── producer.go  
├── routers  
│   └── router.go  
│   └── routes.go  
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

3. ElasticSearch client side (create index / query)  
go run elasticClient.go  
source code: https://github.com/hustbill/ElasticSearchDemo/blob/master/elasticClient.go  

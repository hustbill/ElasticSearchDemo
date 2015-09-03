# Review Service API in Golang 


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
1. client side
http://localhost:8090/products/6

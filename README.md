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

##Product Info Demo
  When we insert a new product into products table,  our client code will publish the database insert event to Kafka server,
  a consumer can listen on this event, and  then update the data into elastic search server via  elastic go library.
  
 (https://github.com/hustbill/ElasticSearchDemo/blob/master/ElasticSearch_Kafka_Postgres_in_order.png)

1. batch migrate data from database to ElasticSearch  
We use Postgres SQL database: medicus_dev, pull data from products table in medicus_dev database to ElasticSearch server

2. search product related data from ElasticSearch  
We start with products search. At first, pull one row from products table, and insert into ElasticSearch server

3. insert new data into ElasticSearch  
when a new product is added in database, we insert  one new data into ElasticSearch server
 
4. update existing ElasticSearch data  
After product information are changed, we also update existing ElasticSearch data
 


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

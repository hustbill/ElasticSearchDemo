package routers 

import (
        "net/http"
        "../handlers"

    )

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}


type Routes []Route

var routes = Routes{
    Route{
        "Index",
        "GET",
        "/",
        handlers.Index,
    },    
    Route{
        "ProductIndex",
        "GET",
        "/products",
        handlers.ProductIndex,
    }, 
    Route{
        "ProductShow",
        "GET",
        "/products/{name}",
        handlers.ProductShow,
    },
    Route{
        "ProductCreate",
        "POST",
        "/products",
        handlers.ProductCreate,
    },
}


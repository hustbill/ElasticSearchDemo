package daos 

import ( 
    "database/sql"
    "fmt"
    "os"
    
    "../models"
    "../kafka"

    _ "github.com/lib/pq"
)


var products models.Products


func GetProductById (db *sql.DB, id string) (*models.Product, error) {
    const query = `SELECT id, name, description, permalink, tax_category_id, shipping_category_id,  is_featured from products
        where id =$1`
 
        var retval models.Product
        err := db.QueryRow(query, id).Scan(&retval.Id, &retval.Name, &retval.Description, &retval.Permalink, 
            &retval.TaxCategoryId, &retval.ShippingCategoryId,  &retval.IsFeatured)
       
        return &retval, err
}



func RepoFindProduct(id string) models.Product {
    dbInfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
    db, err := sql.Open("postgres", dbInfo)    
    product, err := GetProductById(db, id)
    if err != nil {
        fmt.Fprintf(os.Stdout, "Error:%s\n", err)
    } else {
        fmt.Fprintf(os.Stdout, "Product:%v\n", product)
        kafka.Producer(models.Product{Name: product.Name, 
                                        Description: product.Description, 
                                        Permalink: product.Permalink, 
                                        IsFeatured: product.IsFeatured, 
                                        MetaDescription: product.MetaDescription})
        return models.Product{Name: product.Name, Description: product.Description, Permalink: product.Permalink, IsFeatured: product.IsFeatured, MetaDescription: product.MetaDescription}
    }
    // return empty Product if not found
    return models.Product{}
}



// Reference :
// (1)  http://astaxie.gitbooks.io/build-web-application-with-golang/content/en/05.4.html
// (2) http://thenewstack.io/make-a-restful-json-api-go/

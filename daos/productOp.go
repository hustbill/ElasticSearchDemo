package daos 

import ( 
    "database/sql"
    "fmt"
    "os"
    "log"
    
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

func RepoCreateProduct(p models.Product)  models.Product {
    log.Println("# Insert a product")
    dbInfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME) 
    db, err := sql.Open("postgres", dbInfo)
    checkErr(err)
    defer db.Close()
    
    var lastInsertId int64
    err = db.QueryRow(`INSERT INTO products (
        name,
        description,
        permalink,
        tax_category_id,
        shipping_category_id,
        deleted_at,
        meta_description,
        meta_keywords,
        position,
        is_featured,
        can_discount,
        distributor_only_membership
        ) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
     returning id;`,  
         p.Name,                     
         p.Description,              
         p.Permalink,                
         p.TaxCategoryId,
         p.ShippingCategoryId,
         p.DeletedAt,      
         p.MetaDescription,
         p.MetaKeywords,         
         p.Position,            
         p.IsFeatured,     
         p.CanDiscount,              
         p.DistributorOnlyMembership).Scan(&lastInsertId)
    
    log.Println("last inserted id = ", lastInsertId)
    kafka.Producer(models.Product{Name: p.Name, 
                                    Description: p.Description, 
                                    Permalink: p.Permalink, 
                                    IsFeatured: p.IsFeatured, 
                                    MetaDescription: p.MetaDescription})
    
    p.Id = lastInsertId
    return p
}



// Reference :
// (1)  http://astaxie.gitbooks.io/build-web-application-with-golang/content/en/05.4.html
// (2) http://thenewstack.io/make-a-restful-json-api-go/

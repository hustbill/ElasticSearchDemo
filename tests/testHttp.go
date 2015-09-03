package testHttp

import (
    "database/sql"
    "fmt"
    "net/http"
)

func TestHelloHandler_ServeHTTP(t *testing.T) {
    // Open our connection and setup our handler.
    db, err := sql.Open("postgres", "user=sc_admin dbname=review_service sslmode=disable")
    if err != nil {
        panic(err)
    }
    defer db.Close()
    h := HelloHandler{db: db}

    // Execute our handler with a simple buffer.
    rec := httptest.NewRecorder()
    rec.Body = bytes.NewBuffer()
    h.ServeHTTP(rec, nil)
    if rec.Body.String() != "hi bob!\n" {
        t.Errorf("unexpected response: %s", rec.Body.String())
    }
}


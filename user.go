package main

import (
    "net/http"
    "encoding/json"
    "github.com/gorilla/context"
)


/*
 * device
 */
func User(w http.ResponseWriter, r *http.Request) {
    var user ModelUserAccount
    if _, err := Db.QueryOne(&user, `SELECT * FROM user_account WHERE email = ?`, context.Get(r, "email").(string)); err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    result, err := json.Marshal(user)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }
    
    w.Header().Set("Content-Type", "application/json")
    w.Write(result)
}

# urban-farming-backend
    Backend of the Urban Farming student project

#Language
    Go:
    https://golang.org/dl/

#Modules
(go get [module-src])
    - "github.com/go-pg/pg"
    - "github.com/go-pg/pg/orm"
    - "github.com/satori/go.uuid"
    - "github.com/gorilla/mux"
    - "github.com/gorilla/context"
    - "github.com/urfave/negroni"
    - "github.com/SermoDigital/jose/jws"
    - "github.com/SermoDigital/jose/crypto"
    - "golang.org/x/crypto/bcrypt"
    - "gopkg.in/validator.v2"

#Build
    go build server.go database.go database_model.go router.go middleware.go device.go authentication.go user.go


#Environment Variables
(Should be set on the system the backend is running)

    DbHostname=172.17.0.2       //the adress of the db
    DbPort=5432                 //the port where your db is running on
    DbUsername=urban_farming    //the database user
    DbPassword=urban_farming    //the database password
    DbDatabase=urban_farming    //the database name
    JwtSecret=secret
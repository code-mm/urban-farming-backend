FROM golang:1.11.1-stretch

WORKDIR /go/src
COPY *.go /go/src/

RUN go get -d github.com/SermoDigital/jose/crypto github.com/SermoDigital/jose/jws github.com/go-pg/pg github.com/go-pg/pg/orm github.com/gorilla/context github.com/gorilla/mux github.com/satori/go.uuid github.com/urfave/negroni golang.org/x/crypto/bcrypt gopkg.in/validator.v2

RUN go build -o urban-farming-backend server.go database.go database_model.go router.go middleware.go device.go user.go authentication.go

EXPOSE 8080
CMD ["./urban-farming-backend"]
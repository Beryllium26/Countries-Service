# Countries-Service
micro-service that return a list of valid countries that currency can be transferred from or transferred to.

This is a JSON 'REST' API written in go for the express purpose of an exercise I have been asked to do. I had never used 'go' before but I felt that picking what I felt was one of the best tool for the job, was worth the challenge.

It exposes a single location resource fro the api:  /api/v1/

An about page is available at /about and / 



Get it running

just grab the project, build, and you’re ready to go.

go get https://github.com/Beryllium26/Countries-Service

cd $GOPATH/src/github.com/Beryllium26/Countries-Service

go build

go test ./...

start the http server
./


Testing

curl -sv --header "Accept: application/json" http://127.0.0.1:8083/api/v1/countries?target=source

curl -sv --header "Accept: application/json" http://127.0.0.1:8083/api/v1/countries?target=destination



Install gocheck ( http://gopkg.in/check.v1 ) as the internal testing library is relatively slim, and we could do with some of the nice features that gocheck provides: http://labix.org/gocheck






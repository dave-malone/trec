## Run locally with in-memory repositories

`go run main.go` 

## Create a Test User
`curl -H "Content-Type: application/json" -X POST -d '{"first_name":"John","last_name":"Doe", "email":"johndoe@test.com", "password":"p@$$w3Rd"}' http://localhost:3000/user`

## Run with MySQL back end 

### Run MySQL in a Docker container
`docker run --net=host --name trec-mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=c3ESBVltKfas -e MYSQL_DATABASE=trec -e MYSQL_USER=trec -e MYSQL_PASSWORD=trec -d mysql`


### Run TREC and connect with the Running Docker MySQL Container
`PROFILE=MYSQL go run main.go -dburl="trec:trec@tcp(192.168.99.100:3306)/trec"` 

### Connect to the MySQL Instance running in Docker
`mysql -h 192.168.99.100 -u root -pc3ESBVltKfas`
[![wercker status](https://app.wercker.com/status/1caa098ae0d53cda8de56daacd96d3ad/m "wercker status")](https://app.wercker.com/project/bykey/1caa098ae0d53cda8de56daacd96d3ad)


## Run locally with defaults

`go run main.go`

Optionally, use [fresh](https://github.com/pilu/fresh) to auto-reload changes to speed up your dev cycles.

This will run your app using the defaults described under Runtime Configuration via Environment variables.


## Runtime Configuration via Environment variables

Environment variables are used to configure the way the system initializes.

### Persistence Profile

The PROFILE environment variable is used to change which backends the system persists to. The current supported value is "mysql"; e.g. PROFILE=mysql. If the PROFILE env var is not set, or is set to something other than "mysql", then the system will run using the default in-memory repository.

While using in-memory repositories, every time the app is reloaded, you will lose all
previously created data.

TODO - add the option to easily load test data when using in-memory repos.

### Email

Currently, there are two implementations for sending email. The default option is a noopEmailSender, which will do nothing with email notifications.

To use the Amazon SES Email Sender, you must set three environment variables:

1. AWS_ENDPOINT - the url for your  
1. AWS_ACCESS_KEY_ID
1. AWS_SECRET_ACCESS_KEY

To use the Smtp Email sender, you must set these environment variables:

1. SMTP_HOST
1. SMTP_PORT
1. SMTP_USERNAME
1. SMTP_PASSWORD

## Interact with the API

[API Docs](https://app.getpostman.com/dashboard/documentation/view?collection_id=bcfed53e-c8a2-bc46-4146-63e03c357840&owner=314441)

### Create a user
`curl -H "Content-Type: application/json" -X POST -d '{"first_name":"John","last_name":"Doe", "email":"johndoe@test.com", "password":"p@$$w3Rd"}' http://localhost:3000/user/`

### List all Users
`curl http://localhost:3000/user/`

## Run with MySQL back end

### Run MySQL in a Docker container
`docker run --net=host --name trec-mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=c3ESBVltKfas -e MYSQL_DATABASE=trec -e MYSQL_USER=trec -e MYSQL_PASSWORD=trec -d mysql`


### Run TREC and connect with the Running Docker MySQL Container
`PROFILE=MYSQL go run main.go -dburl="trec:trec@tcp(192.168.99.100:3306)/trec"`

### Connect to the MySQL Instance running in Docker
`mysql -h 192.168.99.100 -u root -pc3ESBVltKfas`

# multiverse
multiverse is a project that contains multiple technologies such as grpc, rabitmq, prometheus, graphana, docker, kubernetes etc. 
welcomer and calculator services are from [grpc course](https://www.udemy.com/course/grpc-golang/) + it's hands on.

## Build&Run

You only need doecker to be installed then you can run `make multiverse` to build and run all services for the first time, after that you can run by running `make multiverseRun`.

## Documentation
[swagger documentation](https://app.swaggerhub.com/apis-docs/Armingodiz/Multiverse/1.0.0)

## Services 
### mongoExpress
This service is a gui for mongo, check `localhost:8081`
### notifier 
This service sends a welcome email to email that you gave at signup.

If you want this service to work you need to create a sendgrid token and add .env file to root/notifier and put this lines in it
```json
MULTIVERSE_NOTIFIER_SENDER_EMAIL=your_email_in_sendgrid
MULTIVERSE_NOTIFIER_SENDGRID_TOKEN=sendgrid_token
```



## TODO 
- [ ] Add NATS or KAFKA as broker service to notifier
- [ ] Play with more complex concepts of mongo in core service 

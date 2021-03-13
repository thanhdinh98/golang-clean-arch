# Backend Template for Golang Project
!["golang"](https://miro.medium.com/max/3152/1*Ifpd_HtDiK9u6h68SZgNuA.png)

## How can I run the app ?
```
sudo docker-compose up
```
- API port `:8000`
- Database (PostgreSQL) port `:35432`

For test
```
go test -v ./test/...
```
### Software Architecture: 
___
I closely flow the `Clean Architecture` of Uncle Bob 
1. The structure (in my view):
```
.
├── api
│   └── http
│       ├── v1
├── controller
├── domain
├── entity
├── infra
│   ├── repository
├── schema
```
- `api` present for web server as a third party framework, ... My point to seperate this layer framework to the domain business code
- `controller` as the routing. It's responsiblity would be like the transport layer that comunicate between web framework and our application
- `domain` is our core business code, it just like a stand-alone module which isolate with every outer layer includes the database
- `repository` is our data layer that comunicate between our application and the database
- `enitty` is data structure for our model in database. It would only used in repository and called from domain
- `schema` is reuqest-response structure presentation. It's useful for checking validate the request before go through the application depper

Please feel free to check the source and improve the source code with feedbacks

**Diagram about the architecture would be upload soon**
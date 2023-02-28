- [Type Api Gateway in Microservice](#type_pai_gateway_in_microservice)
- [How to choose api gateway](#how_to_choose_api_gateway)

## Type Api Gateway in Microservice  <a name="type_pai_gateway_in_microservice"></a>

![many_entry_point.png](img%2Fmany_entry_point.png) </br>
![single_entry_point.png](img%2Fsingle_entry_point.png) </br>
There are two main types of api gateways: Single entry point gateway and multiple entry point gateway. </br>
In the single entry point model, all your entry points will be entered into a single api gateway, which will do all the
navigation of the service, authen, author. In the many entry point model, split into a separate service for special
needs: web, mobile, public..., the work of each point is similar to the single entry point model. Details
link: https://microservices.io/patterns/apigateway.html </br>

## How to choose api gateway  <a name="how_to_choose_api_gateway"></a>
A selected api gateway must ensure fast, lightweight, stable, secure, full authen, authorization.... All famous open
source like Kong, cloud like AWS ApiGateway fully support features. Basic web app. In cases where it is necessary to
deeply customize Api Gateway logic, high-performance languages such as go, rust, C++... are used to ensure flexibility.
An API gateway acts as a reverse proxy to accept all application programming interface (API) calls, aggregate the
various services required to fulfill them, and return the appropriate result. You write your own gateway api in this
case. </br>
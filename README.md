# TrustWallet Test App

To have a production-ready application, it will be required to:
## Code application

- minimize the external library and if they are a requirement use specific versions of the library that have been audited
- improve the code to have a stable connection to the Ethereum network with eventual Heath Check or Circuit Breaking, etc ...
- improve the code to provide application metrics for better monitoring
- improve the code to properly expose the logs of the application
- if the application doesn't need to expose an HTTP endpoint, implement an internal health check via check present of PID file or health file or execution of command line
- implement the code to allow the horizontal scalability of the application
- implement performance tests, to identify the behavior of the application, ( minimum resource requirement, max performance, ... )

## CI/CD
- perform some security checks, systematically during the CI process of the application; static code analysis, dependency check, container scanning, ...
- deploy the application to allow the export of the application logs and metrics to an observability stack for better analysis of the application's behavior
- deploy the application to allow the implementation of alerting rules due to proactively manage the application based on the information that has been collected
- if the application allows it, deploy it with an autoscaling mechanize

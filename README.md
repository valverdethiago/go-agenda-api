# go-agenda-api

Build an HTTP API that's responsible for handling a phone agenda (i know it's obvious, but the nuances on building the project are what really matters).
The http server should contain:
-An endpoint for pushing new contacts
-An endpoint for editing contact information
-An endpoint for deleting a contact
-An endpoint for searching a contact by it's id
-An endpoint for searching contacts by a part of their name
-An endpoint that lists all contacts
-The http service should be configurable through flags on startup (timeouts and port to use)
-Log messages should be written for each endpoint hit, with the response status code and the time it took to fulfill the request
If an error occurs, a log message should be written with the response status code and a helpful error message, to help an engineer troubleshoot the issue
-Service and host metrics should be collected. I suggest using Prometheus (https://prometheus.io/docs/guides/go-application/)
-The application should have a reasonable test coverage, preferably above 70%
-The application should have end-to-end tests (this is a good way to try out the http cl

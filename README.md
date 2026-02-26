# Assignment 1 PROG2005
Gabriel Brindle

A web service using REST API to fetch information about countries.

# Endpoints
Using specified endpoints:

REST Countries API
http://129.241.150.113:8080/v3.1/ 

Currency API
http://129.241.150.113:9090/currency/

The below endpoints fetches information based specifications. Endpoints marked with {} requires to use Alpha-2 country codes.

## /Status/
Status is used to retrieve the status of the APIs, uptime of the service and the version of the service.
## /Info/{country code}
Info is used to return various information about designated country code, such as continent, population, capital etc.

E.g. /info/cn would fetch information about China
## /Exchange/{country code}
Is used to fetch information about which currency designated country code uses, and what bordering countries exchange rates are.

E.g. /exchange/no would fetch the currency of Norway and their neighbouring countries.

# Render
The render deployment should be live on:

https://cloud-assignment-1-rfd0.onrender.com

use /countryinfo/v1 + endpoint to test the service's functionality

E.g. https://cloud-assignment-1-rfd0.onrender.com/countryinfo/v1/status


#!/bin/bash



go build 

# ./searchserver -addr="localhost:3003" -indexpath="dataindex-fosdem-ical" -indexname="events" -update="500s" -datasourcepath="datacal/fosdem.ical"


./searchserver -addr="localhost:3002" -indexpath="dataindex-sportscheck-product" -indexname="products" -update="500s" -datasourcepath="datasourcesportscheck/data/products-mock.json"








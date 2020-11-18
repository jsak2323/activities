# activities

this repository is use to report every activity users, 

1. IS THIS A CLEAN ARCHITECTURE ?
- yes, this apps use a clean architecture with Domain Driven Design, how can we know, this is clean architecture by identify every layer part of this apps; 

** [Delivery/Rest] At first this apps have layer rest, in this layer only use for make data interface that user want, like format json/xml data structure for front end its all in rest layer, all endpoints ,and all about input output of apps is in Delivery/Rest layer.
** [Service] service layer, dealing about every process, every business logic will be here, so it won't be able to disturb the other layers, except for the process data in apps.
**Entity-Layer** entity layer is used to connecting the service layer into Repository Layer, in this case the developer put the entity inside of Domain layer, but it depends of conditions.
** [Repository] the last layer is Repository layer, use to get data from, db or other repository, and convert from entity model into repository model, and vice versa, all htpp call to get data can be in this layer too, so it wont be distract another layer.

2. HOW TO SCALE UP THE APPLICATION AND WHEN IT NEEDS TO BE ?
- that is a use of Golang Language. Just like c/c++ golang language is from binary code, so it wont took too long to process data, services, etc. and the advantage of golang is the multi thread that not every programming language have, so it can be processing services in concurrency, just like paralellism, not the same but related, concurrency is about structure, and paralelism is about execution, and have significant lower development cost, serves more request and have a efficient code 

How to Run Apps:
command in activities folder: 
1. mkdir logs
2. go run ./main.go


* Get all activity every user that use in this app, 

use this curl:
- curl --location --request GET 'http://localhost:8000/activities/all'

response :
{
    "errors": {},
    "data": [
        {
            "id": 2,
            "activityType": "select-all-activities",
            "activityAttribute": {
                "name": "jason",
                "act": "berlari",
                "status": 1
            }
        },
        {
            "id": 3,
            "activityType": "select-all-activities",
            "activityAttribute": {
                "name": "jason",
                "act": "berlari",
                "status": 1
            }
        }
    ],
    "messages": [
        "success"
    ]
}

* Get all activity user by id, 
use this curl:

- curl --location --request GET 'http://localhost:8000/activities/2'

response :
{
    "errors": {},
    "data": {
        "id": 2,
        "activityType": "select-activity-by-id",
        "activityAttribute": {
            "name": "jason",
            "act": "berlari",
            "status": 1
        }
    },
    "messages": [
        "success"
    ]
}

* Insert new user activity,
use this curl:

curl --location --request POST 'http://localhost:8000/activities' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "angga",
    "act": "Running-from-tomorrow",
    "status": 1
}'

response:
{
    "errors": {},
    "data": {
        "id": 9,
        "activityType": "insert-activities",
        "activityAttribute": {
            "name": "angga",
            "act": "Running-from-tomorrow",
            "status": 1
        }
    },
    "messages": [
        "success"
    ]
}

* Update user activity by id,
use this curl:

curl --location --request PUT 'http://localhost:8000/activities' \
--header 'Content-Type: application/json' \
--data-raw '{
    "id": 9,
    "name": "jason-angga",
    "act": "Running-from-tomorrow",
    "status": 1
}'

response:
{
    "errors": {},
    "data": {
        "id": 9,
        "activityType": "update-activities",
        "activityAttribute": {
            "name": "jason-angga",
            "act": "Running-from-tomorrow",
            "status": 1
        }
    },
    "messages": [
        "success"
    ]
}

* Delete activity user by id,
use this curl:

curl --location --request DELETE 'http://localhost:8000/activities/9'

response:
{
    "errors": {},
    "data": "OK",
    "messages": [
        "success"
    ]
}

<!-- every error or mistakes, will be record in logs folder that create at first before run this apps -->
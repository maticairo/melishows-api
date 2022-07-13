
# Melishows REST API

This REST API is entirely built over Golang 1.16
Deployed over GCP AppEngine

## Run the app

`go run cmd/main.go`
## Run the tests

`go test melishows-api`
# REST API  Documentation

The REST API to the example app is described below.

## Get list of Shows

### Request

`GET /shows/all`

https://melishows.rj.r.appspot.com/allShows
### Response

Status: 200 OK  
Content-Type: application/json
 ```
 [
   {
      "id":"62963928-f501-4af3-bafd-0acce2321668",
      "name":"El Lago de los Cisnes",
      "functions":[
         {
            "id":"7c336060-02c7-4d30-aec0-507e7b4c4c40",
            "show_date": "2006-01-02T15:04:05Z",
            "duration":120,
            "theater_id":"1eed8488-bac2-466c-95a8-cc6c450082b5",
            "theater_room_id":"28dea577-40a7-49c6-b37e-c6760f68d49a",
            "pricing":[
               {
                  "id":"36657660-f602-4f7f-9ad5-d2be83941609",
                  "price":100,
                  "seats":[
                     {
                        "row_number":1,
                        "identifier":"A",
                        "booked":false
                     },
                     {
                        "row_number":1,
                        "identifier":"B",
                        "booked":false
                     },
                     {
                        "row_number":1,
                        "identifier":"C",
                        "booked":false
                     }
                  ]
               }
            ]
         },
         {
            "id":"d122ca4f-58d3-4b6f-8344-bcb8f915d655",
            "show_date": "2019-01-02T15:04:05Z",
            "duration":120,
            "theater_id":"1eed8488-bac2-466c-95a8-cc6c450082b5",
            "theater_room_id":"4e3ae664-a47c-48c9-b8c5-241902003c74",
            "pricing":[
               {
                  "id":"36657660-f602-4f7f-9ad5-d2be83941609",
                  "price":100,
                  "seats":[
                     {
                        "row_number":1,
                        "identifier":"A",
                        "booked":false
                     },
                     {
                        "row_number":1,
                        "identifier":"B",
                        "booked":false
                     },
                     {
                        "row_number":1,
                        "identifier":"C",
                        "booked":false
                     }
                  ]
               }
            ]
         }
      ]
   }
]
```  

## Get Available Seats

### Request

`GET /availableSeats`

https://melishows.rj.r.appspot.com/availableSeats?show_id=62963928-f501-4af3-bafd-0acce2321668&function_id=d122ca4f-58d3-4b6f-8344-bcb8f915d655
### Response

 ```
 {
   "show":"62963928-f501-4af3-bafd-0acce2321668",
   "function":"7c336060-02c7-4d30-aec0-507e7b4c4c40",
   "seats":[
      {
         "id":"36657660-f602-4f7f-9ad5-d2be83941609",
         "price":100,
         "seats":[
            {
               "row_number":1,
               "identifier":"A",
               "booked":false
            },
            {
               "row_number":1,
               "identifier":"B",
               "booked":false
            },
            {
               "row_number":1,
               "identifier":"C",
               "booked":false
            }
         ]
      }
   ]
}
 ```

## Search Shows

### Request

`GET /shows/search`

https://melishows.rj.r.appspot.com/shows/search/?date_from=2005-01-01T15:04:05Z&date_to=2015-01-01T15:04:05Z&price_from=0&price_to=1000&order_kind=DESC

### Response

  ```
   [
   {
      "id":"62963928-f501-4af3-bafd-0acce2321668",
      "name":"El Lago de los Cisnes",
      "functions":[
         {
            "id":"7c336060-02c7-4d30-aec0-507e7b4c4c40",
            "show_date": "2006-01-02T15:04:05Z",
            "duration":120,
            "theater_id":"1eed8488-bac2-466c-95a8-cc6c450082b5",
            "theater_room_id":"28dea577-40a7-49c6-b37e-c6760f68d49a",
            "pricing":[
               {
                  "id":"36657660-f602-4f7f-9ad5-d2be83941609",
                  "price":100,
                  "seats":[
                     {
                        "row_number":1,
                        "identifier":"A",
                        "booked":false
                     },
                     {
                        "row_number":1,
                        "identifier":"B",
                        "booked":false
                     },
                     {
                        "row_number":1,
                        "identifier":"C",
                        "booked":false
                     }
                  ]
               }
            ]
         }
      ]
   }
]
  ```

## Book a ticket

### Request

`POST /book`

Body:
```
{
    "show_id": "62963928-f501-4af3-bafd-0acce2321668",
    "function_id": "7c336060-02c7-4d30-aec0-507e7b4c4c40",
    "dni": 12345678,
    "name": "Matias",
    "seats": [
        "1-A",
        "1-B",
        "1-C"
	 ]
}
``` 

https://melishows.rj.r.appspot.com/book

### Response
  ```
  {
   "dni":12345678,
   "name":"Matias",
   "show_name":"",
   "theater_name":"Teatro Colon",
   "theater_room":2,
   "day":"",
   "show_date":"2006-01-02T15:04:05Z",
   "seats":[
      "1-A",
      "1-B",
      "1-C"
   ],
   "total_price":300
}
  ```

## Third party libraries

[gorilla/mux v1.8.0](https://github.com/gorilla/mux)

[karlseguin/ccache v2.0.8](https://github.com/karlseguin/ccache/v2)

## Some considerations
Cache implements an in-memory cache. Some considerations: I've implemented this type of cache to fulfill the
exercise requirements and taking into account that I've deployed the application in a single VM environment.
If we want to scale the application horizontally, we have to implement a distributed cache or a mechanism to
synchronize the VMs (like a cron job)

The repository layer loads information from two main json files and keeps it in memory. No external
library was used to handle an in-memory DB in order to code it quickly.

Tests are implemented over the main layers of this software (I've avoided testing models in order to deliver the
exercise quickly)

## Authors

Contributors names and contact info

ex. Matias Cairo
ex. [@maticairo](https://www.linkedin.com/in/matias-cairo-56b7a0b6/)
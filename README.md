# Go Events API
[![Deploy](https://www.herokucdn.com/deploy/button.png)](https://heroku.com/deploy)

A Go Events CRUD API using the native net/http library and [mgo](https://github.com/go-mgo/mgo). 

Feel free to use and/or make any improvements to it. 

The Go gopher was designed by Renee French. (http://reneefrench.blogspot.com/) The gopher vector data was made by Takuya Ueda (http://u.hinoichi.net). Licensed under the Creative Commons 3.0 Attributions license.

## Instructions

1. Install mgo dependency. 

   ```
	go get gopkg.in/mgo.v2
   ```

2. Start a MongoDB instance.

   ```
	mongod --dbpath D:/dbpath
   ```

3. Run my app!

   ```
	go run main.go -port=<port> -conn=<connectionString>
   ```
  
There are 2 parameters: the port represents the HTTP port to the API and the conn represents the MongoDB connection string where the event data will be persisted. The default values are "8080" and "mongodb://localhost:27017", respectively.

## Usage

There are 5 URLs that allow CRUD operations in a event resource:

- **<code>GET</code> events/** to get all available events.
- **<code>GET</code> events/:id** to get a specific event.
- **<code>POST</code> events/** to include a event.
- **<code>PUT</code> events/:id/** to update a event.
- **<code>DELETE</code> events/:id** to delete a event.

The event resource has the following format:

```json
{ 
"name": "Golang POA Meetup", 
"begin": "2020-04-15T13:00:15-07:00", 
"end": "2020-04-15T18:00:15-07:00", 
"location": 
 { 
  "latitude": 30.0331,
  "longitude": 51.2300, 
  "name": "Porto Alegre" 
  }
}
```


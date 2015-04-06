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

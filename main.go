package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

type Event struct {
	Id       bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Begin    time.Time
	End      time.Time
	Name     string
	Location struct {
		Latitude  float32
		Longitude float32
		Name      string
	}
}

var connectionString *string

func main() {
	connectionString = flag.String("conn", "mongodb://localhost:27017", "MongoDB connection string")
	port := flag.String("port", "8080", "API HTTP port")
	flag.Parse()

	http.HandleFunc("/events/", eventsHandler)

	fmt.Printf("API started on port %s\n", *port)
	http.ListenAndServe(":"+*port, nil)

}

func eventsHandler(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "POST":
		event, err := readBody(request.Body)
		if err != nil {
			badRequest(writer, err)
		}
		createEvent(event)
		break
	case "GET":
		if id := readId(request); id != "" {
			event, err := readEvent(id)
			if err != nil {
				badRequest(writer, err)
			} else {
				success(writer, event)
			}
		} else {
			success(writer, readEvents())
		}
		break
	case "PUT":
		if id := readId(request); id != "" {
			event, err := readBody(request.Body)
			if err != nil {
				badRequest(writer, err)
			}
			updateEvent(id, event)
		} else {
			badRequest(writer, errors.New("Please provide a event id."))
		}

		break
	case "DELETE":
		if id := readId(request); id != "" {
			removeEvent(id)
		} else {
			badRequest(writer, errors.New("Please provide a event id."))
		}
		break
	}
}

// CRUD operations

func createEvent(event *Event) {
	session := getMongoSession()
	defer session.Close()

	collection := session.DB("").C("events")
	err := collection.Insert(event)
	if err != nil {
		panic(err)
	}
}

func readEvent(id string) (*Event, error) {
	var objectId bson.ObjectId
	if bson.IsObjectIdHex(id) {
		objectId = bson.ObjectIdHex(id)
	} else {
		return nil, errors.New("Invalid id.")
	}

	session := getMongoSession()
	defer session.Close()

	event := Event{}
	collection := session.DB("").C("events")
	query := collection.Find(bson.M{"_id": objectId})
	err := query.One(&event)
	if err != nil {
		panic(err)
	}

	return &event, nil
}

func readEvents() []Event {
	session := getMongoSession()
	defer session.Close()

	events := []Event{}
	collection := session.DB("").C("events")
	query := collection.Find(nil)
	query.All(&events)

	return events
}

func updateEvent(id string, event *Event) error {
	var objectId bson.ObjectId
	if bson.IsObjectIdHex(id) {
		objectId = bson.ObjectIdHex(id)
	} else {
		return errors.New("Invalid id.")
	}

	session := getMongoSession()
	defer session.Close()

	collection := session.DB("").C("events")
	err := collection.UpdateId(objectId, event)
	if err != nil {
		panic(err)
	}

	return nil
}

func removeEvent(id string) error {
	var objectId bson.ObjectId
	if bson.IsObjectIdHex(id) {
		objectId = bson.ObjectIdHex(id)
	} else {
		return errors.New("Invalid id.")
	}

	session := getMongoSession()
	defer session.Close()

	collection := session.DB("").C("events")
	err := collection.Remove(bson.M{"_id": objectId})
	if err != nil {
		panic(err)
	}

	return nil
}

// Helper methods

func readBody(body io.ReadCloser) (*Event, error) {
	data, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, errors.New("Invalid request body")
	}
	event := &Event{}
	json.Unmarshal(data, &event)
	return event, nil
}

func getMongoSession() *mgo.Session {
	session, err := mgo.Dial(*connectionString)
	if err != nil {
		panic(err)
	}

	return session
}

func badRequest(writer http.ResponseWriter, err error) {
	writer.Header().Set("Content-Type", "text/plain")
	writer.WriteHeader(400)
	fmt.Fprintln(writer, err.Error())
}

func success(writer http.ResponseWriter, response interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	r, _ := json.Marshal(response)
	fmt.Fprintln(writer, string(r))
}

func readId(request *http.Request) string {
	return request.URL.Path[len("/events/"):]
}

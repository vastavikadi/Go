package main

import (
	"encoding/json"
	"fmt"
	"log" //log: logging utilities. log.Fatal(err) prints an error and exits (calls os.Exit(1)).
	"math/rand"
	"net/http"
	"strconv" //string conversion utilities, strconv.Itoa(int) converts int → string.
	"time"

	"github.com/gorilla/mux" //uses github.com/gorilla/mux for routing and encoding/json to read/write JSON.
)

// Model for course - file(goes to a different folder/model)
type Course struct { //using fake db for testing is called seeding
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"` //*Author is type of the Author (defining the types here) it is going to a pointer of type Author (struct) Using a pointer allows nil author (encoded as null in JSON) and avoids copying a possibly-large Author struct.
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake DB
var courses []Course //Declares a package-level variable courses of type []Course (slice of Course). It's initially nil (zero-value slice). The code uses this slice as an in-memory fake database. Every change is only in memory, so it is transient (lost on program restart).

// middleware, helper - file
func (c *Course) IsEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == "" //just to check the entry is empty or not
}

func main() {
	fmt.Println("API - LearnCodeOnline.in")
	r := mux.NewRouter() // r := mux.NewRouter() creates a new router from gorilla/mux. r is used to register handlers for routes.

	//seeding
	courses = append(courses, Course{CourseId: "2", CourseName: "ReactJS", CoursePrice: 299, Author: &Author{Fullname: "Hitesh Choudhary", Website: "lco.dev"}})
	courses = append(courses, Course{CourseId: "4", CourseName: "MERN Stack", CoursePrice: 199, Author: &Author{Fullname: "Hitesh Choudhary", Website: "go.dev"}}) //append adds values to the courses slice. Because courses was nil, append creates a new backing array and returns a non-nil slice assigned back to courses. The Author: &Author{...} syntax creates an Author struct and takes its address — storing a pointer. These two entries are the initial dataset.

	//routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE") //r.HandleFunc(path, handler).Methods("GET") registers an HTTP handler function for path and restricts it to the specified HTTP method(s). {id} in path is a path variable (mux will extract it). Each handler must have signature func(w http.ResponseWriter, r *http.Request).

	// listen to a port
	log.Fatal(http.ListenAndServe(":4000", r)) /// log.Fatal(...) logs the error returned by ListenAndServe and exits the program. If ListenAndServe returns nil (unlikely), log.Fatal would not run—ListenAndServe returns non-nil only on error.
}

//controllers - file

// serve home route

func serveHome(w http.ResponseWriter, r *http.Request) { //(res, req) and w http.ResponseWriter is the API used to write headers and body back to the client.
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)                                  //WriteHeader sends an HTTP response header with the provided status code.
	w.Write([]byte("<h1>Welcome to API by LearnCodeOnline</h1>")) //this one is the GET / (Home)
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses) //json.NewEncoder(w).Encode(courses) marshals the courses slice to JSON and writes it directly to the response body. If encoding fails it returns an error — but this code ignores that error.
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	// grab id from request
	params := mux.Vars(r) //params := mux.Vars(r) extracts path variables into map[string]string, so params["id"] is the path {id}.
	fmt.Printf("The type of the params: %T", params)
	fmt.Printf("The value of the params: %v", params)

	// loop through courses, find matching id and return the response
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound) //to return 404 and return a JSON object like {"error":"No course..."} by default it is always 200OK Once you call Write (like w.Write([]byte(...)) or json.NewEncoder(w).Encode(...)), if you haven’t already set a status, Go will automatically set it to 200 OK before writing the body.
	json.NewEncoder(w).Encode(map[string]string{"error": "No course found with the requested id"})
	return //this is the response: //this is the response:
	// 	HTTP/1.1 404 Not Found
	// Content-Type: application/json
	// Date: Sat, 16 Aug 2025 14:00:00 GMT
	// Content-Length: 32
}

// Header: tells the client how to interpret the body.
// WriteHeader: tells the client whether the request was successful or not (with the proper HTTP status code).

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")

	// what if: body is empty
	if r.Body == nil { //from router body is directly available
		json.NewEncoder(w).Encode("Please send some data")
		return
	}

	// what about - {}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)//parses the json from the requets body and adds it to the variable course
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside JSON")
		return
	}

	//TODO: check only if title is duplicate
	// loop, title matches with course.coursename, JSON

	// generate unique id, string
	// append course into courses

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return

}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "applicatioan/json")

	// first - grab id from req
	params := mux.Vars(r)

	// loop, id, remove, add with my ID

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	//TODO: send a response when id is not found
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course")
	w.Header().Set("Content-Type", "applicatioan/json")

	params := mux.Vars(r)

	//loop, id, remove (index, index+1)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			// TODO: send a confirm or deny response
			break
		}
	}
}

func getUsersBooks(w http.ResponseWriter, r *http.Request) {}

// GET / — home HTML
// GET /courses — list all courses
// GET /course/{id} — get one course by id
// POST /course — create a course
// PUT /course/{id} — update a course
// DELETE /course/{id} — delete a course

// IMPORTANT NOTE:
// Only exported fields (beginning with uppercase letter) are visible to encoding/json. If fields were courseId with lowercase, they would be ignored by json marshal/unmarshal.

// The tag names are lowercase; when you json.Marshal(course) you'll get keys like "courseid", "coursename", "price", "author".

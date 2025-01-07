package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("API - Authors and Courses")
	r := mux.NewRouter()

	// seeding the data
	courses = append(courses, Course{CourseId: "2", CourseName: "React JS", CoursePrice: 200, Author: &Author{Fullname: "Param Kansagra", Website: "paramkansagra.com"}}, Course{CourseId: "4", CourseName: "MERN", CoursePrice: 300, Author: &Author{Fullname: "Param Kansagra", Website: "paramkansagra.com"}})

	// routing the different endpoints to functions
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET") // idhar ham jo bhi likhenge like id,courseId etc vo same to same we need to write in params
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	// listen to a port and log
	log.Fatal(http.ListenAndServe(":4000", r))
}

// Model for course - file
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

// Model for the author - file
type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// lets make a fake db
var courses []Course

// middleware / helper functions - file
func (c *Course) IsEmpty() bool {
	return c.CourseName == ""
}

/*
	Controller are the functions who handel some routes in the API
*/

// serve home route
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API by ParamKansagra</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")

	// we will be setting the Content-type header as application/json
	w.Header().Set("Content-type", "application/json")

	// and as we are getting all the courses we will just say
	// encode all the "courses" in the http.ResponceWriter that we are getting in
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get only one course")
	w.Header().Set("Content-type", "application/json")

	// lets grab the id from the request
	params := mux.Vars(r)

	// loop thru the courses and grab the correct course
	for _, course := range courses {
		if course.CourseId == params["id"] {
			w.WriteHeader(http.StatusFound)   // we have found the course so set the headers
			json.NewEncoder(w).Encode(course) // now write the course as json
			return                            // and return
		}
	}

	w.WriteHeader(http.StatusNotFound)                         // after looping if we havent found the course then set the header as statusNotFound
	json.NewEncoder(w).Encode("No course found with given id") // and in the json write "no course found"
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-type", "application/json")

	// what if the body is empty
	if r.Body == nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Please send some data")
		return
	}

	// what about body is like -> {}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)

	if course.IsEmpty() {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("No data inside json sent")
		return
	}

	// Lets also check if there is any course with that same id

	for _, courseLoop := range courses {
		if course.CourseName == courseLoop.CourseName {

			// make a map having the course already present in the database along with a message about it
			response := map[string]interface{}{"course": courseLoop, "message": "Course with same name already exisiting in database"}

			// we have to return that course and all
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(response)

			return
		}
	}

	// generate a unique id and convert the id into string
	// append the new course into courses

	rand.Seed(time.Now().UnixNano())

	course.CourseId = strconv.Itoa(rand.Intn(1000))

	courses = append(courses, course)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(course)
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Updating one course")
	w.Header().Set("Content-type", "application/json")

	// first we will check if the body is empty or not
	params := mux.Vars(r)

	/*
		so we have to
			1. get the id of the course to delete
			2. delete the course
			3. add the new course at the same index with the same course id
	*/

	for index, course := range courses {
		if course.CourseId == params["id"] {
			// we go the course to delete and update
			courses = append(courses[:index], courses[index+1:]...)

			// make a new course and get the data from the body
			var newCourse Course
			_ = json.NewDecoder(r.Body).Decode(&newCourse)

			// the new course id would be equal to that of the previous one
			newCourse.CourseId = params["id"]

			// append this course to the array
			courses = append(courses, newCourse)

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(newCourse)
			return
		}
	}

	// in case we do not find the course we will return statusNotFound
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode("Course not found to update")
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course")
	w.Header().Set("Content-type", "application/json")

	// now first we will get the parameters
	params := mux.Vars(r)

	/*
		So we steps to be followed
			1. get the index of the course to be delete
			2. If found in the array
				2.1 return status ok and remove it from the array
				2.2 return the course that we are deleting if needed
			3. if not found in the array
				3.1 return status not found and return
	*/

	// in case we dont find any id in params
	if params["id"] == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("No ID found in params")
		return
	}

	for index, course := range courses {
		if course.CourseId == params["id"] {
			// we will remove the course from the array
			courses = append(courses[:index], courses[index+1:]...)

			// set the headers as status ok
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode("Course deleted successfully")
			return
		}
	}

	// if id is not found in the array
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode("Course with id not found")
}

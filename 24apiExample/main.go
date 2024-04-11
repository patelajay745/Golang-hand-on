package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Course struct {
	CourseId    string
	CourseName  string
	CoursePrice int
	Author      *Author
}

type Author struct {
	Fullname string
	Website  string
}

var Courses []Course

func (c *Course) IsEmpty() bool {
	return c.CourseName == ""
}
func main() {
	fmt.Println("welcome to example of API")

	r := mux.NewRouter()

	Courses = append(Courses, Course{
		CourseId: "1", CourseName: "DevOps", CoursePrice: 299, Author: &Author{Fullname: "Ajay", Website: "ajaypatel.live"}})

	Courses = append(Courses, Course{
		CourseId: "2", CourseName: "Go Lang", CoursePrice: 249, Author: &Author{Fullname: "Ajay Patel", Website: "ajaypatel.net"}})

	r.HandleFunc("/", ServerHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	http.ListenAndServe(":8001", r)

}

func ServerHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-type", "application/Json")
	json.NewEncoder(w).Encode(Courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/Json")
	params := mux.Vars(r)
	for _, course := range Courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found")

}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/Json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please Send some data")
	}

	var course Course

	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside json")
		return
	}

	// Create a new random number generator with a custom seed (e.g., current time)
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)

	// Generate a random number of minutes between 1 and 15
	course.CourseId = strconv.Itoa(rng.Intn(100))

	Courses = append(Courses, course)

	json.NewEncoder(w).Encode(fmt.Sprintf("Added course: %s", course.CourseName))

}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/Json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please Send some data")
	}

	params := mux.Vars(r)

	for index, course := range Courses {
		if course.CourseId == params["id"] {
			Courses = append(Courses[:index], Courses[index+1:]...)
			var course Course

			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			// If course with the specified ID is not found, append the new course
			Courses = append(Courses, course)
			json.NewEncoder(w).Encode(course)

			return
		}
	}
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/Json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please Send some data")
	}

	params := mux.Vars(r)

	id := params["id"]
	if id == "" {
		http.Error(w, "No ID is provided", http.StatusBadRequest)
		return
	}

	for index, course := range Courses {
		if course.CourseId == params["id"] {
			Courses = append(Courses[:index], Courses[index+1:]...)
			json.NewEncoder(w).Encode("Deleted")
			break
		}
	}

}

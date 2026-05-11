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

// model for courses - file

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"courseprice"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake DB
var courses []Course

// middleware , helper
func (c *Course) IsEmpty() bool {
	return c.CourseId == "" && c.CourseName == ""
}

func main() {

}

// controllers - file

// server Home route

func serveHome(w http.ResponseWriter, _ *http.Request) {
	var str string = `{"msg": "hello from learn code online"}`

	w.Write([]byte(str))
}

func getAllCourses(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get One Course")
	w.Header().Set("Content-Type", "application/json")

	// grab id from req

	params := mux.Vars(r)

	// loop through courses,

	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course Found with given id")

}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("please send some data")
	}
	var course Course

	_ = json.NewDecoder(r.Body).Decode(&course)

	if course.IsEmpty() {
		json.NewEncoder(w).Encode("please send some data")
		return
	}
	// generate unique ID
	rand.Seed(time.Now().UTC().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	// append
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)

}

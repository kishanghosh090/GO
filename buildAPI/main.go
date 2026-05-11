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
	r := mux.NewRouter()
	courses = append(courses, Course{
		CourseId:    "2",
		CourseName:  "React",
		CoursePrice: 299,
		Author: &Author{
			Fullname: "kishan",
			Website:  "kishanranaghosh.xyz",
		},
	})
	courses = append(courses, Course{
		CourseId:    "3",
		CourseName:  "React",
		CoursePrice: 299,
		Author: &Author{
			Fullname: "kishan",
			Website:  "kishanranaghosh.xyz",
		},
	})
	courses = append(courses, Course{
		CourseId:    "4",
		CourseName:  "React",
		CoursePrice: 299,
		Author: &Author{
			Fullname: "kishan",
			Website:  "kishanranaghosh.xyz",
		},
	})
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course", deleteOneCourse).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":4007", r))

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

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// first grab id
	params := mux.Vars(r)

	// loop id

	for idx, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:idx], courses[idx+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
		}
	}
	json.NewEncoder(w).Encode("no id found")

}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// first grab id
	params := mux.Vars(r)

	// loop id

	for idx, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:idx], courses[idx+1:]...)
			json.NewEncoder(w).Encode(course)
		}
	}
	json.NewEncoder(w).Encode("no id found")

}

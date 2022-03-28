package main

type Course struct {
  CourseNumber string
  CourseName string
  CourseDescription string
  Professor string
  StartTime string
  EndTime string
  Days[5] bool
}

type Student struct {
  FirstName string
  LastName string
  EnrolledCourses[] Course
}

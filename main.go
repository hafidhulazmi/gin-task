package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/gin-gonic/gin"
)

type (
	Siswa struct {
		NIS    int    `json:"nis"`
		Nama   string `json:"nama"`
		Kelas  int    `json:"kelas"`
		Gender string `json:"gender"`
	}
	Guru struct {
		NIG    int    `json:"nig"`
		Nama   string `json:"nama"`
		Gender string `json:"gender"`
	}

	MataPelajaran struct {
		Nama  string `json:"nama"`
		ID    int    `json:"id"`
		NIG   int    `json:"nig"`
		Kelas int    `json:"kelas"`
	}
)

var Students []Siswa = []Siswa{
	{NIS: 1, Nama: "Kekar", Kelas: 4, Gender: "Laki-laki"},
	{NIS: 2, Nama: "Kiti", Kelas: 5, Gender: "Perempuan"},
}
var Teachers []Guru = []Guru{
	{NIG: 1, Nama: "Kekar", Gender: "Laki-laki"},
	{NIG: 2, Nama: "Kiti", Gender: "Perempuan"},
}

var Lesson []MataPelajaran = []MataPelajaran{
	{ID: 1, Nama: "Matematika", NIG: 1, Kelas: 4},
	{ID: 2, Nama: "Fisika", NIG: 2, Kelas: 5},
	{ID: 2, Nama: "Kimia", NIG: 2, Kelas: 6},
	{ID: 2, Nama: "Biologi", NIG: 2, Kelas: 6},
}

func LessonsByNIS() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer c.Request.Body.Close()
		dataBody, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			panic(err)
		}
		nis := struct {
			NIS int `json:"nis"`
		}{}
		err = json.Unmarshal(dataBody, &nis)
		if err != nil {
			panic(err)
		}
		dataStudent := Siswa{}
		for _, v := range Students {
			if v.NIS == nis.NIS {
				dataStudent = v
			}
		}
		if dataStudent.NIS == 0 {
			c.Writer.WriteHeader(404)
			_, err := c.Writer.Write([]byte("data siswa tidak ditemukan"))
			if err != nil {
				panic(err)
			}
			return
		}
		dataLesson := []MataPelajaran{}
		for _, v := range Lesson {
			if v.Kelas == dataStudent.Kelas {
				dataLesson = append(dataLesson, v)
			}
		}
		dataLessonJson, err := json.Marshal(dataLesson)
		if err != nil {
			panic(err)
		}
		_, err = c.Writer.Write(dataLessonJson)
		if err != nil {
			panic(err)
		}
		return
	}
}

func StudentsGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.WriteHeader(200)
		dataSiswaJson, err := json.Marshal(Students)
		if err != nil {
			panic(err)
		}
		_, err = c.Writer.Write(dataSiswaJson)
		if err != nil {
			panic(err)
		}
		return
	}
}

func StudentsPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer c.Request.Body.Close()
		dataBody, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			panic(err)
		}
		nis := struct {
			NIS int `json:"nis"`
		}{}
		err = json.Unmarshal(dataBody, &nis)
		if err != nil {
			panic(err)
		}
		dataSiswa := Siswa{NIS: 0}
		for _, v := range Students {
			if v.NIS == nis.NIS {
				dataSiswa = v
			}
		}
		if dataSiswa.NIS == 0 {
			c.Writer.WriteHeader(404)
			_, err := c.Writer.Write([]byte("data siswa tidak ditemukan"))
			if err != nil {
				panic(err)
			}
			return
		}
		dataSiswaJson, err := json.Marshal(dataSiswa)
		if err != nil {
			panic(err)
		}
		_, err = c.Writer.Write(dataSiswaJson)
		if err != nil {
			panic(err)
		}
		return
	}
}

func TeachersGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.WriteHeader(200)
		dataGuruJson, err := json.Marshal(Teachers)
		if err != nil {
			panic(err)
		}
		_, err = c.Writer.Write(dataGuruJson)
		if err != nil {
			panic(err)
		}
		return
	}
}

func Teacherspost() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer c.Request.Body.Close()
		dataBody, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			panic(err)
		}
		nig := struct {
			NIG int `json:"nig"`
		}{}
		err = json.Unmarshal(dataBody, &nig)
		if err != nil {
			panic(err)
		}
		dataGuru := Guru{NIG: 0}
		for _, v := range Teachers {
			if v.NIG == nig.NIG {
				dataGuru = v
			}
		}
		if dataGuru.NIG == 0 {
			c.Writer.WriteHeader(404)
			_, err := c.Writer.Write([]byte("data guru tidak ditemukan"))
			if err != nil {
				panic(err)
			}
			return
		}
		dataGuruJson, err := json.Marshal(dataGuru)
		if err != nil {
			panic(err)
		}
		_, err = c.Writer.Write(dataGuruJson)
		if err != nil {
			panic(err)
		}
		return
	}
}

func Lessonpost() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer c.Request.Body.Close()
		dataBody, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			panic(err)
		}
		id := struct {
			ID int `json:"id"`
		}{}
		err = json.Unmarshal(dataBody, &id)
		if err != nil {
			panic(err)
		}
		dataLesson := MataPelajaran{ID: 0}
		for _, v := range Lesson {
			if v.ID == id.ID {
				dataLesson = v
			}
		}
		if dataLesson.ID == 0 {
			c.Writer.WriteHeader(404)
			_, err := c.Writer.Write([]byte("data mapellajaran tidak ditemukan"))
			if err != nil {
				panic(err)
			}
			return
		}
		dataLessonJson, err := json.Marshal(dataLesson)
		if err != nil {
			panic(err)
		}
		_, err = c.Writer.Write(dataLessonJson)
		if err != nil {
		}
		return
	}
}

func GetGin() *gin.Engine {
	r := gin.Default()
	r.Use(GlobalMiddleware())
	r.GET("/students", RouteLevelMiddleware(), StudentsGet())
	r.GET("/teachers", RouteLevelMiddleware(), TeachersGet())
	r.GET("/lesson", Lessonpost())
	r.POST("/students", StudentsPost())
	r.POST("/teachers", Teacherspost())
	r.GET("/lessons/nis", LessonsByNIS())
	return r
}

func GlobalMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("Global middleware")
		c.Next()
	}
}

func RouteLevelMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("Route level middleware")
		c.Next()
	}
}

func main() {
	if err := GetGin().Run("localhost:8080"); err != nil {
		log.Printf("Error running server: %v", err)
	}
}

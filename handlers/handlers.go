package handlers

import (
	"17-assignment/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func GetStudentFromId(w http.ResponseWriter, r *http.Request) {
	
	studentIdString := r.URL.Query().Get("student_id")
	studentId, err := strconv.ParseUint(studentIdString, 10, 64)

	if err != nil {
		fmt.Println("error", err)
		ErrInConversion := map[string]string{"Message": "Cannot convert string to int: not a numeric string"}

		ErrInConversionJson, err := json.Marshal(ErrInConversion)
		if err != nil {
			log.Println("error", err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, http.StatusText(http.StatusInternalServerError))
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(ErrInConversionJson))
		return

	}

	studentDatafromMap, err := models.FetchStudentData(studentId)
	if err != nil {
		ErrRecordNotFound := map[string]string{"Message": err.Error()}
		ErrRecordNotFoundJson, err := json.Marshal(ErrRecordNotFound)
		if err != nil {
			log.Println("error", err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, http.StatusText(http.StatusInternalServerError))
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(ErrRecordNotFoundJson))
		return

	}
	studentDatafromMaptoJson, err := json.Marshal(studentDatafromMap)
	if err != nil {
		log.Println("error", err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, http.StatusText(http.StatusInternalServerError))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(studentDatafromMaptoJson))

}

func GetAllStudentsData(w http.ResponseWriter, r *http.Request) {
	studentDataAsMap := models.FetchAllStudentData()

	studentDataAsMapToJson, err := json.Marshal(studentDataAsMap)
	if err != nil {
		log.Println("error", err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, http.StatusText(http.StatusInternalServerError))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(studentDataAsMapToJson))

}

package routers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-twitter/bd"
	"github.com/go-twitter/models"
	"github.com/go-twitter/utils"
)

func ViewProfile(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro ID", http.StatusBadRequest)
		return
	}
	clientRedis := utils.GetClientRedis()
	key := "profile_" + ID
	var profile models.User
	value, err := clientRedis.Get(key).Result()
	if err != nil {
		// no existe profile
		profile, err := bd.GetProfile(ID)
		if err != nil {
			http.Error(w, "Ocurrio un error al intentar el buscar el registro "+err.Error(), http.StatusBadRequest)
			return
		}

		jsonProfile, errJson := json.Marshal(profile)
		if errJson != nil {
			fmt.Println("err Json ", errJson.Error())
		}
		errRedis := clientRedis.Set(key, jsonProfile, 100*time.Second).Err()
		if errRedis != nil {
			fmt.Println(errRedis.Error())
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(profile)
	} else {
		fmt.Println("Se encontro")
		json.Unmarshal([]byte(value), &profile)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(profile)
	}

}

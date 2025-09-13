package json

import (
	"encoding/json"
	"log"
	"smartgopher/modules"
)

type Week struct {
	monday    string `json:"monday"`
	tuesday   string `json:"tuesday"`
	wednesday string `json:"wednesday"`
	thursday  string `json:"thursday"`
	friday    string `json:"friday"`
}

func (wee Week) toMap() map[string]string {
	return map[string]string{ // Надо дореализовать преобразование расписания в json файл так,что бы было некие ветвления, глянь это: https://excalidraw.com/
		"monday":    wee.monday,
		"tuesday":   wee.tuesday,
		"wednesday": wee.wednesday,
		"thursay":   wee.thursday,
		"friday":    wee.friday,
	}
}

func SaveMap_to_Jsonfile(data map[string]interface{}, filename string) {
	bytesArray, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Println("Cant marshal to byte array")
	}

	err = modules.Write(filename, string(bytesArray))
	if err != nil {
		log.Println("Cant marshal to byte array")
	}
}

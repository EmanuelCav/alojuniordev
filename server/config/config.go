package config

import (
	"log"

	"github.com/joho/godotenv"
)

func Config() map[string]string {

	envFile, err := godotenv.Read(".env")

	if err != nil {
		log.Fatal("Error loading env file")
	}

	variables := make(map[string]string)

	variables["port"] = envFile["PORT"]
	variables["uri"] = envFile["URI"]
	variables["database"] = envFile["DATABASE"]
	variables["jwt"] = envFile["JWT"]
	variables["userCollection"] = envFile["USER_COLLECTION"]
	variables["roleCollection"] = envFile["ROLE_COLLECTION"]
	variables["categoryCollection"] = envFile["CATEGORY_COLLECTION"]
	variables["toolCollection"] = envFile["TOOL_COLLECTION"]
	variables["userMainRole"] = envFile["USER_MAIN_ROLE"]
	variables["cloud_name"] = envFile["CLOUD_NAME"]
	variables["api_key"] = envFile["API_KEY"]
	variables["api_secret"] = envFile["API_SECRET"]
	variables["folder"] = envFile["FOLDER"]

	return variables

}

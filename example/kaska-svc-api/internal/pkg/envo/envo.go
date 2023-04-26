package envo

import (
	"os"
	//"github.com/joho/godotenv"
)

func EnvString(env, fallback string) string {

	e := os.Getenv(env)

	if e == "" {
		//log.Printf("using fallback env value for %s : %s", env, fallback)
		return fallback
	}

	//log.Printf("using config env value for %s : %s", env, e)
	return e
}

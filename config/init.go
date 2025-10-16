package config

import "os"

func init() {
	URIMONGODB = os.Getenv("URIMONGODB")
	DBNAME = os.Getenv("DBNAME")
}

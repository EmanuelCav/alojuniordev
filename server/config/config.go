package config

import "os"

func Port() string {
	return os.Getenv("PORT")
}

func Database() string {
	return os.Getenv("DATABASE")
}

func Uri() string {
	return os.Getenv("URI")
}

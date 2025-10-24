package config

import "github.com/joho/godotenv"

// Load loads environment variables from the specified .env file path.
func Load(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		return err
	}
	return nil
}

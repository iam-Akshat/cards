package testutils

import (
	"fmt"
	"os/exec"
	"time"

	db "github.com/iam-Akshat/cards/utils"
)

func ExecutePostgresDockerContainer(config *db.DatabaseConfig) error {
	port := config.Port
	password := config.Password
	host := config.Host
	dbName := config.DBName
	dbUser := config.User

	commandArgs := []string{
		"run",
		"--rm",
		"-d",
		"-p",
		fmt.Sprintf("%s:5432", port),
		"--name", "test_db_container",
		"-e",
		fmt.Sprintf("POSTGRES_PASSWORD=%s", password),
		"-e",
		fmt.Sprintf("POSTGRES_HOST=%s", host),
		"-e",
		fmt.Sprintf("POSTGRES_DB=%s", dbName),
		"-e",
		fmt.Sprintf("POSTGRES_USER=%s", dbUser),
		"-e",
		fmt.Sprintf("POSTGRES_PORT=%s", port),
		"postgres:13",
	}
	fmt.Println("Executing docker command: ", commandArgs)
	out, err := exec.Command("docker", commandArgs...).Output()
	if err != nil {
		if err.Error() == "exit status 125" {
			fmt.Println("Docker container already running. Skipping...")
			return nil
		}
		fmt.Println("Error executing docker command: ", err)
		return err
	}
	fmt.Println(string(out))
	fmt.Println("Waiting for 4 seconds for docker container to start...")
	time.Sleep(4 * time.Second)

	return nil
}

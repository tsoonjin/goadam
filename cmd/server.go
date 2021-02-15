package main

import (
    "log"
    "fmt"
    "flag"
    "os"
    "io/ioutil"
    "encoding/json"
    "net/http"
    "github.com/joho/godotenv"
    "github.com/tsoonjin/goadam/internal/service"
)

type CLIFlags struct {
    name string
}

type Config struct {
    Env string
    Port string
}

type AppConfig struct {
    Version string `json:"version"`
}

func loadAppConfig(configPath string) (*AppConfig, error) {
    appConfig := AppConfig{}
    raw, err := ioutil.ReadFile(configPath)
    if err != nil {
        log.Fatal("Error parsing app config")
        return nil, err
    }
    json.Unmarshal(raw, &appConfig)
    return &appConfig, nil
}

func loadConfig() Config {
    if err := godotenv.Load(); err != nil {
        log.Fatal("Failed to load environment variable...")
        panic(err)
    }
    config := Config {
        Env: "develop",
        Port: "3000",
    }
    env := os.Getenv("ENV")
    if env != "" {
        config.Env = env
    }
    port := os.Getenv("PORT")
    if port != "" {
        config.Port = port
    }
    return config
}

func loadFlags() CLIFlags {
    serviceName := flag.String("n", "goadam", "App name")
    return CLIFlags {
        name: *serviceName,
    }
}

func setupServer(port string) {
    http.HandleFunc("/", service.IndexHandler)
    http.HandleFunc("/request", service.JSON2CSVHandler)
    http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}

func main() {
    config := loadConfig()
    cliFlags := loadFlags()
    appConfig, _ := loadAppConfig("./config/app.json")
    log.Println(fmt.Sprintf("Starting app %s@%s at port %s", cliFlags.name, appConfig.Version, config.Port))
    setupServer(config.Port)
}

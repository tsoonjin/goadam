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
    config string
}

type Config struct {
    Env string
    Port string
}

type AppConfig struct {
    Version string `json:"version"`
}

func loadAppConfig(configPath string) (*AppConfig, error) {
    appConfig := AppConfig{
        Version: "1.0.0",
    }
    raw, err := ioutil.ReadFile(configPath)
    if err != nil {
        log.Println("Failed to load app config. Using default config")
    }
    json.Unmarshal(raw, &appConfig)
    return &appConfig, err
}

func loadConfig() Config {
    if os.Getenv("ENV") != "production" {
        if err := godotenv.Load(); err != nil {
            log.Println("Failed to load .env file")
        }
    }
    config := Config {
        Env: "develop",
        Port: "3000",
    }
    env := os.Getenv("ENV")
    if env != "" {
        config.Env = env
    }
    port := os.Getenv("APP_PORT")
    if port != "" {
        config.Port = port
    }
    return config
}

func loadFlags() CLIFlags {
    serviceName := flag.String("n", "goadam", "App Name")
    configFile := flag.String("c", "./config.json", "App config file")
    return CLIFlags {
        name: *serviceName,
        config: *configFile,
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
    appConfig, _ := loadAppConfig(cliFlags.config)
    log.Println(fmt.Sprintf("Starting app %s@%s at port %s", cliFlags.name, appConfig.Version, config.Port))
    setupServer(config.Port)
}

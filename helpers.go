package main

import (
    "os"
    "log"
    "strconv"
    "strings"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
    "./models"
)


// database connection settings
type DatabaseSettings struct {
    Hostname        string
    Port            int
    Username        string
    Password        string
    DatabaseName    string
}

func (setting *DatabaseSettings) parseDatabaseHostname() {
    hostname, set := os.LookupEnv("DatabaseHostname")
    if set {
        setting.Hostname = hostname
    } else {
        setting.Hostname = "localhost"
    }
}

func (setting *DatabaseSettings) parseDatabasePort() {
    var err error
    port, set := os.LookupEnv("DatabasePort")
    if set {
        setting.Port, err = strconv.Atoi(port)
        if err != nil {
            log.Fatal(err)
        }
    } else {
        setting.Port = 5432 // postgres default port
    }
}

func (setting *DatabaseSettings) parseDatabaseUsername() {
    username, set := os.LookupEnv("DatabaseUsername")
    if set {
        setting.Username = username
    } else {
        setting.Username = "postgres" // default postgres username
    }
}

func (setting *DatabaseSettings) parseDatabasePassword() {
    password, set := os.LookupEnv("DatabasePassword")
    if set {
        setting.Password = password
    } else {
        setting.Password = "postgres" // default postgres password
    } 
}

func (setting *DatabaseSettings) parseDatabaseDatabaseName() {
    databaseName, set := os.LookupEnv("DatabaseName")
    if set {
        setting.DatabaseName = databaseName
    } else {
        setting.DatabaseName = "urban_farming"
    }
}

func (setting *DatabaseSettings) GetDatabaseArguments() string {
    var arguments strings.Builder
    arguments.WriteString("host=")
    arguments.WriteString(setting.Hostname)
    arguments.WriteString(" ")
    arguments.WriteString("port=")
    arguments.WriteString(strconv.Itoa(setting.Port))
    arguments.WriteString(" ")
    arguments.WriteString("user=")
    arguments.WriteString(setting.Username)
    arguments.WriteString(" ")
    arguments.WriteString("password=")
    arguments.WriteString(setting.Password)
    arguments.WriteString(" ")
    arguments.WriteString("dbname=")
    arguments.WriteString(setting.DatabaseName)
    return arguments.String()
}

func DatabaseSettingsNew() *DatabaseSettings {
    var setting DatabaseSettings
    setting.parseDatabaseHostname()
    setting.parseDatabasePort()
    setting.parseDatabaseUsername()
    setting.parseDatabasePassword()
    setting.parseDatabaseDatabaseName()
    return &setting
}

// jwt settings
type JwtSettings struct {
    Secret       string
    ValidityFarm int
    ValidityUser int
}

func (setting *JwtSettings) parseJwtSecret() {
    secret, set := os.LookupEnv("JwtSecret")
    if set {
        setting.Secret = secret
    } else {
        setting.Secret = "secret"
    }
}

func (setting *JwtSettings) parseJwtValidityFarm() {
    var err error
    validity, set := os.LookupEnv("JwtValidityFarm")
    if set {
        setting.ValidityFarm, err = strconv.Atoi(validity)
        if err != nil {
            log.Fatal(err)
        }
    } else {
        setting.ValidityFarm = 3600
    }
}

func (setting *JwtSettings) parseJwtValidityUser() {
    var err error
    validity, set := os.LookupEnv("JwtValidityUser")
    if set {
        setting.ValidityUser, err = strconv.Atoi(validity)
        if err != nil {
            log.Fatal(err)
        }
    } else {
        setting.ValidityUser = 3600
    }
}

func JwtSettingsNew() *JwtSettings {
    var setting JwtSettings
    setting.parseJwtSecret()
    setting.parseJwtValidityFarm()
    setting.parseJwtValidityUser()
    return &setting
}

// database connection operations
func DbOpen(db **gorm.DB, settings *DatabaseSettings) {
    var err error
    *db, err = gorm.Open("postgres", settings.GetDatabaseArguments())
    if err != nil {
        log.Fatal(err)
    }
}

func DbClose(db *gorm.DB) {
    db.Close()
}

func DbInit(db *gorm.DB) {
    db.AutoMigrate(&models.User{}, &models.UserFarmPermission{}, &models.Farm{}, &models.DataPointPh{}, &models.DataPointOxygen{}, &models.DataPointTemperature{})
    db.Model(&models.UserFarmPermission{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
    db.Model(&models.UserFarmPermission{}).AddForeignKey("farm_id", "farms(id)", "CASCADE", "CASCADE")
    db.Model(&models.DataPointPh{}).AddForeignKey("farm_id", "farms(id)", "CASCADE", "CASCADE")
    db.Model(&models.DataPointOxygen{}).AddForeignKey("farm_id", "farms(id)", "CASCADE", "CASCADE")
    db.Model(&models.DataPointTemperature{}).AddForeignKey("farm_id", "farms(id)", "CASCADE", "CASCADE")
}

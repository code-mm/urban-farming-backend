package main

import (
    "os"
    "log"
    "strconv"
    "fmt"
    "crypto/rand"
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
        setting.Password = "urban_farming" // default application password
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

func (setting *DatabaseSettings) GetDatabaseDSN() string {
    dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
        setting.Hostname,
        strconv.Itoa(setting.Port),
        setting.Username,
        setting.Password,
        setting.DatabaseName
    )
    return dsn
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
        // check for a minimum of 32 byte secret
        if len(secret) < 32 {
            log.Fatal("Invalid value for JwtSecret: 32 characters are required")
        }
        setting.Secret = secret
    } else {
        // generate random secret
        secret := make([]byte, 32)
        _, err := rand.Read(secret)
        if err != nil {
            log.Fatal(err)
        }
        setting.Secret = string(secret)
    }
}

func (setting *JwtSettings) parseJwtValidityFarm() {
    var err error
    validity, set := os.LookupEnv("JwtValidityFarm")
    if set {
        setting.ValidityFarm, err = strconv.Atoi(validity)
        if err != nil {
            log.Fatal("Invalid value for JwtValidityFarm")
        }
    } else {
        setting.ValidityFarm = 60
    }
}

func (setting *JwtSettings) parseJwtValidityUser() {
    var err error
    validity, set := os.LookupEnv("JwtValidityUser")
    if set {
        setting.ValidityUser, err = strconv.Atoi(validity)
        if err != nil {
            log.Fatal("Invalid value for JwtValidityUser")
        }
    } else {
        setting.ValidityUser = 30
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
    *db, err = gorm.Open("postgres", settings.GetDatabaseDSN())
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

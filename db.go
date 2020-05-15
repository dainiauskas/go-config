package config

import (
	"fmt"
	"net/url"
	"time"

	"github.com/go-sql-driver/mysql"
)

var (
	defaultMaxIdleConns = 10
	defaultMaxOpenConns = 10
	defaultMaxIdleTime  = time.Minute * 60
	defaultMaxLifetime  = time.Minute * 480
)

// Database configuration for connecting to database
type Database struct {
	// Dialect: mysql or mssql
	Dialect string

	// Host values: localhost, server name or ip
	Host string

	// Instance used only for mssql.
	Instance string

	// Database port, this value is optional, not used in mssql
	Port int

	// Database User name
	User string

	// Database User password
	Pass string

	// Database name
	Name string

	// Additional params
	Params map[string]string

	// Location for database time
	Location string

	// Net Database connection protocol
	Net string

	// Database collation
	Collation string

	// Pool connection settings
	MaxIdleConns *int
	MaxOpenConns *int

	// Poll Time Minutes
	MaxIdleTime *time.Duration
	MaxLifetime *time.Duration
}

// setDefaults set default values if empty required data
func (d *Database) setDefaults() *Database {
	if d.Location == "" {
		d.Location = "Europe/Vilnius"
	}

	if d.Net == "" {
		d.Net = "tcp"
	}

	if d.Collation == "" {
		d.Collation = "cp1257_lithuanian_ci"
	}

	if d.MaxIdleConns == nil {
		d.MaxIdleConns = &defaultMaxIdleConns
	}

	if d.MaxOpenConns == nil {
		d.MaxOpenConns = &defaultMaxOpenConns
	}

	if d.MaxIdleTime != nil {
		defaultMaxIdleTime = *d.MaxIdleTime * time.Minute
	}
	d.MaxIdleTime = &defaultMaxIdleTime

	if d.MaxLifetime != nil {
		defaultMaxLifetime = *d.MaxLifetime * time.Minute
	}
	d.MaxLifetime = &defaultMaxLifetime

	return d
}

// FormatDSN formats the given Config into a DSN string which can be passed to the driver.
func (d *Database) FormatDSN() string {
	d.setDefaults()

	switch d.Dialect {
	case "mysql":
		return d.myToString()
	case "mssql":
		return d.msToString()
	}

	return ""
}

// getLocation loading location by config, on error return UTC location
func (d *Database) getLocation() *time.Location {
	loc, err := time.LoadLocation(d.Location)
	if err != nil {
		loc = time.UTC
	}

	return loc
}

// myToString generate and return MySql connection string
func (d *Database) myToString() string {
	host := d.Host
	if d.Port > 0 {
		host += fmt.Sprintf(":%d", d.Port)
	}

	config := &mysql.Config{
		User:                 d.User,
		Passwd:               d.Pass,
		Net:                  d.Net,
		Addr:                 host,
		DBName:               d.Name,
		Params:               d.Params,
		ParseTime:            true,
		AllowNativePasswords: true,
		// MaxAllowedPacket: 4 << 20,
		Loc:               d.getLocation(),
		Collation:         d.Collation,
		MultiStatements:   true,
		InterpolateParams: true,
	}

	return config.FormatDSN()
}

// msToString generate and return MsSql connection string
func (d *Database) msToString() string {
	query := url.Values{}
	query.Add("database", d.Name)
	query.Add("charset", d.Collation)

	u := &url.URL{
		Scheme:   "sqlserver",
		User:     url.UserPassword(d.User, d.Pass),
		Host:     d.Host,
		Path:     d.Instance,
		RawQuery: query.Encode(),
	}

	if d.Port > 0 && d.Instance == "" {
		u.Host = fmt.Sprintf("%s:%d", d.Host, d.Port)
	}

	return u.String()
}

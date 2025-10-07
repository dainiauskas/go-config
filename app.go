package config

// Secure used for TLS
type Secure struct {
	Cache string
	Cert  string
	Key   string
}

// App is Aplication Configuration
type App struct {
	ServiceName    string
	ServiceDisplay string

	// Verbose use for debug (default - false).
	// Posible values: true, false
	Verbose bool

	// Console use for output to screen
	// Posible values: true, false
	Console bool

	// Host URL for API
	Host string

	// Port used for API listen
	Port int

	// Echo server configuration
	Recover bool
	Gzip    bool

	// Updater
	AutoUpdate bool

	Secure *Secure

	// Schedule use string value, example: "* * * * * *". This value required!
	//
	// Field name   | Mandatory? | Allowed values  | Allowed spec. characters
	// ----------   | ---------- | --------------  | ------------------------
	// Seconds      | Yes        | 0-59            | * / , -
	// Minutes      | Yes        | 0-59            | * / , -
	// Hours        | Yes        | 0-23            | * / , -
	// Day of month | Yes        | 1-31            | * / , - ?
	// Month        | Yes        | 1-12 or JAN-DEC | * / , -
	// Day of week  | Yes        | 0-6 or SUN-SAT  | * / , - ?
	//
	// # Special Characters
	// Asterisk ( * ) The asterisk indicates that the cron expression will
	// match for all values of the field; e.g., using an asterisk in the
	// 5th field (month) would indicate every month.
	//
	// Slash ( / ) Slashes are used to describe increments of ranges. For
	// example 3-59/15 in the 1st field (minutes) would indicate the 3rd
	// minute of the hour and every 15 minutes thereafter. The form "*\/..."
	// is equivalent to the form "first-last/...", that is, an increment over
	// the largest possible range of the field. The form "N/..." is accepted
	// as meaning "N-MAX/...", that is, starting at N, use the increment until
	// the end of that specific range. It does not wrap around.
	//
	// Comma ( , ) Commas are used to separate items of a list. For example,
	// using "MON,WED,FRI" in the 5th field (day of week) would mean Mondays,
	// Wednesdays and Fridays.
	//
	// Hyphen ( - ) Hyphens are used to define ranges. For example, 9-17 would
	// indicate every hour between 9am and 5pm inclusive.
	//
	// Question mark ( ? ) Question mark may be used instead of '*' for
	// leaving either day-of-month or day-of-week blank.
	//
	// Predefined schedules
	// May use one of several pre-defined schedules in place of a cron
	// expression.
	//
	// Entry    | Description                                | Equivalent To
	// -----    | -----------                                | -------------
	// @yearly  | Run once a year, midnight, Jan. 1st        | 0 0 0 1 1 *
	// @monthly | Run once a month, midnight, first of month | 0 0 0 1 * *
	// @weekly  | Run once a week, midnight between Sat/Sun  | 0 0 0 * * 0
	// @daily   | Run once a day, midnight                   | 0 0 0 * * *
	// @hourly  | Run once an hour, beginning of hour        | 0 0 * * * *
	Schedule string
	// ExposeErrorStack controls whether internal stack traces are included in
	// HTTP JSON error responses. Default is false (do not expose).
	ExposeErrorStack bool `mapstructure:"expose_error_stack"`
}

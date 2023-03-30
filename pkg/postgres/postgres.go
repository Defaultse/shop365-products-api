// // Package postgres implements postgres connection.
package postgres

// import (
// 	"fmt"
// 	"time"

// 	_ "github.com/jackc/pgx/stdlib"
// 	"github.com/jmoiron/sqlx"
// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// )

// const (
// 	_defaultMaxPoolSize  = 1
// 	_defaultConnAttempts = 10
// 	_defaultConnTimeout  = time.Second
// )

// // Postgres -.
// type Postgres struct {
// 	maxPoolSize  int
// 	connAttempts int
// 	connTimeout  time.Duration
// 	conn         *sqlx.DB
// }

// // New -.
// func New(url string, opts ...Option) (*Postgres, error) {
// 	// conn, err := sqlx.Connect("pgx", url)
// 	// if err != nil {
// 	// 	fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
// 	// 	return nil, err
// 	// }

// 	// if err := conn.Ping(); err != nil {
// 	// 	panic(err)
// 	// }

// 	// if err != nil {
// 	// 	return nil, fmt.Errorf("postgres - NewPostgres - connAttempts == 0: %w", err)
// 	// }

// 	db, err := gorm.Open(postgres.New(postgres.Config{
// 		DSN:                  "user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai",
// 		PreferSimpleProtocol: true, // disables implicit prepared statement usage
// 	}), &gorm.Config{})

// 	if err != nil {
// 		return nil, fmt.Errorf("postgres - NewPostgres - connAttempts == 0: %w", err)
// 	}

// 	pg := &Postgres{
// 		maxPoolSize:  _defaultMaxPoolSize,
// 		connAttempts: _defaultConnAttempts,
// 		connTimeout:  _defaultConnTimeout,
// 		pg:           db,
// 	}

// 	// Custom options
// 	for _, opt := range opts {
// 		opt(pg)
// 	}

// 	return pg, nil
// }

// // Close -.
// func (p *Postgres) Close() {
// 	if p.conn != nil {
// 		p.conn.Close()
// 	}
// }

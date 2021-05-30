package main

import (
	"github.com/jackc/pgx"
	"log"
	"strconv"
)

func main() {
	config := pgx.ConnPoolConfig{
		ConnConfig:     pgx.ConnConfig{
			Host:                 "localhost",
			Port:                 5432,
			Database:             "high",
			User:                 "high",
			Password:             "high",
			TLSConfig:            nil,
			UseFallbackTLS:       false,
			FallbackTLSConfig:    nil,
			Logger:               nil,
			LogLevel:             0,
			Dial:                 nil,
			RuntimeParams:        nil,
			OnNotice:             nil,
			CustomConnInfo:       nil,
			CustomCancel:         nil,
			PreferSimpleProtocol: false,
			TargetSessionAttrs:   "",
		},
		MaxConnections: 100,
		AfterConnect:   nil,
		AcquireTimeout: 0,
	}
	db, err := pgx.NewConnPool(config)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	for i := 0; i <= 1000; i++ {
		//var lang, lat float64
		sqlState := "INSERT INTO location(uid, lat, lng, location)\nVALUES(" + strconv.Itoa(i*100) + ",20,40,ST_SetSRID(ST_MakePoint(20, 40), 4326))"
		for j := 1; j < 100; j++ {
			sqlState += ", (" + strconv.Itoa(i*100+j) + ",30,50,ST_SetSRID(ST_MakePoint(30, 50), 4326))"
		}
		//print(sqlState)
		sqlState += " ON CONFLICT (uid) \nDO \n   UPDATE SET lat = EXCLUDED.lat, lng = EXCLUDED.lng, location = EXCLUDED.location"
		_, err = db.Exec(sqlState)
		//print(i, "\n")
		if err != nil {
			log.Fatal(err)
		}
	}
}

package persistence

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/modl"
	"github.com/joho/godotenv"
	"github.com/zainafifaldi/digicafe/domain/entity"
)

func DBConn() (*sql.DB, error) {
	if _, err := os.Stat("./../../.env"); !os.IsNotExist(err) {
		err := godotenv.Load(os.ExpandEnv("./../../.env"))
		if err != nil {
			log.Fatalf("Error getting env %v\n", err)
		}
		return LocalDatabase()
	}
	return CIBuild()
}

//Circle CI DB
func CIBuild() (*sql.DB, error) {
	var err error
	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", "((CI_DB_HOST))", "((CI_DB_PORT))", "((CI_DB_USERNAME))", "((CI_DB_NAME))", "((CI_DB_PASSWORD))")
	db, err := sql.Open("mysql", DBURL)
	if err != nil {
		log.Fatal("This is the error:", err)
	}
	return db, nil
}

//Local DB
func LocalDatabase() (*sql.DB, error) {
	dbDriver := os.Getenv("DB_CONNECTION")
	dbHost := os.Getenv("DB_TEST_HOST")
	dbPort := os.Getenv("DB_TEST_PORT")
	dbUser := os.Getenv("DB_TEST_USERNAME")
	dbPassword := os.Getenv("DB_TEST_PASSWORD")
	dbName := os.Getenv("DB_TEST_NAME")
	parseTime := "true"

	dbUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=%s", dbUser, dbPassword, dbHost, dbPort, dbName, parseTime)
	db, err := sql.Open(dbDriver, dbUrl)
	if err != nil {
		return nil, err
	} else {
		log.Println("CONNECTED TO: ", dbDriver)
	}

	dropQuery := "TRUNCATE TABLE stores;"
	if _, err = db.Exec(dropQuery); err != nil {
		return nil, err
	}

	return db, nil
}

func seedStore(db *sql.DB) (*entity.Store, error) {
	store := &entity.Store{
		ID:                1,
		Type:              1,
		Name:              "Main Store",
		Description:       "This is main store",
		Address:           "Bandung, Jawa Barat",
		Active:            true,
		LocationLatitude:  1.23456,
		LocationLongitude: 2.34567,
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
	}

	dialect := modl.MySQLDialect{Engine: "InnoDB", Encoding: "UTF8"}
	dbMap := modl.NewDbMap(db, dialect)
	dbMap.AddTableWithName(entity.Store{}, "stores")

	err := dbMap.Insert(store)
	if err != nil {
		return nil, err
	}
	return store, nil
}

func seedStores(db *sql.DB) ([]entity.Store, error) {
	stores := []entity.Store{
		{
			Type:              1,
			Name:              "Main Store",
			Description:       "This is main store",
			Address:           "Kota Bandung, Jawa Barat",
			Active:            true,
			LocationLatitude:  1.23456,
			LocationLongitude: 2.34567,
			CreatedAt:         time.Now(),
			UpdatedAt:         time.Now(),
		},
		{
			Type:              1,
			Name:              "Branch Store",
			Description:       "This is branch store",
			Address:           "Kab. Bandung Barat, Jawa Barat",
			Active:            true,
			LocationLatitude:  1.23568,
			LocationLongitude: 3.76543,
			CreatedAt:         time.Now(),
			UpdatedAt:         time.Now(),
		},
	}

	dialect := modl.MySQLDialect{Engine: "InnoDB", Encoding: "UTF8"}
	dbMap := modl.NewDbMap(db, dialect)
	dbMap.AddTableWithName(entity.Store{}, "stores")

	for _, store := range stores {
		err := dbMap.Insert(&store)
		if err != nil {
			return nil, err
		}
	}
	return stores, nil
}

// func seedFood(db *sql.DB) (*entity.Food, error) {
// 	food := &entity.Food{
// 		ID:          1,
// 		Title:       "food title",
// 		Description: "food desc",
// 		UserID:      1,
// 	}
// 	// err := db.Create(&food).Error
// 	_, err := db.Query("Lalala")
// 	if err != nil {
// 		return nil, err
// 	}
// 	return food, nil
// }

// func seedFoods(db *sql.DB) ([]entity.Food, error) {
// 	foods := []entity.Food{
// 		{
// 			ID:          1,
// 			Title:       "first food",
// 			Description: "first desc",
// 			UserID:      1,
// 		},
// 		{
// 			ID:          2,
// 			Title:       "second food",
// 			Description: "second desc",
// 			UserID:      1,
// 		},
// 	}
// 	for _, v := range foods {
// 		// err := db.Create(&v).Error
// 		_, err := db.Query("Lalala %d", v)
// 		if err != nil {
// 			return nil, err
// 		}
// 	}
// 	return foods, nil
// }

// dsn := "user=postgres password=s197328645S! dbname=todo port=5432 sslmode=disable TimeZone=Europe/Warsaw"
// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

// if err != nil {
// 	panic("failed to connect database")
// }
// Migrate the schema
//db.Migrator().DropTable(&task{})
//db.AutoMigrate(&task{})

// Create
//db.Create(&task{Description: "Wash the dishes", Completed: true})

// Read
// var product task
// db.First(&product, 1)                        // find product with integer primary key
// db.First(&product, "Description = ?", "D42") // find product with code D42

// // Update - update product's price to 200
// db.Model(&product).Update("Completed", false)
// // Update - update multiple fields
// db.Model(&product).Updates(task{Completed: false, Description: "R2D2"}) // non-zero fields
// db.Model(&product).Updates(map[string]interface{}{"Completed": false, "Description": "R2D2"})

// //Delete - delete product
// db.Delete(&product, 1)
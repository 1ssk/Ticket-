package initializers

import "github.com/1ssk/admin/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.Ticket{})
}

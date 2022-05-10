package utils

import (
	"fmt"
	"gorm.io/gorm"
)

func PrintResult(tx *gorm.DB){
	fmt.Println("RowsAffected : ",tx.RowsAffected)
	fmt.Println("Error : ",tx.Error)
}

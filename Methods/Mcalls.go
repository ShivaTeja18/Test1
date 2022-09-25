package Methods

import (
	. "EMP_API/Details"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

var DB *gorm.DB
var R = gin.Default()

func Fetching(a *gin.Context) {
	var empobj []EMP
	DB.Find(&empobj)
	a.IndentedJSON(http.StatusOK, gin.H{
		"message": "SUCCESSFUL",
		"result":  &empobj,
	})
}

func Creating(c *gin.Context) {
	var empobj EMP

	if err := c.BindJSON(&empobj); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
	}
	if empobj.Name == "" || empobj.Password == "" || empobj.City == "" {
		c.JSON(http.StatusNoContent, &empobj)
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{
			"message": "SUCCESSFUL",
			"result":  &empobj,
		})
	}
	DB.Create(&empobj)
}

func Fbyid(a *gin.Context) {
	var empobj EMP

	id := a.Param("id")
	if err := DB.Where("id = ?", id).First(&empobj).Error; err != nil {
		a.JSON(http.StatusNotFound, &empobj)
		return
	}
	a.JSON(http.StatusOK, gin.H{
		"message": "SUCCESSFUL",
		"result":  empobj,
	})
}

func Updating(a *gin.Context) {
	var empobj EMP

	id := a.Param("id")

	if err := a.ShouldBindJSON(&empobj); err != nil {
		a.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	if err := DB.Where("id = ?", id).Updates(&empobj).Error; err != nil {
		a.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	//a.JSON(http.StatusBadRequest, &empobj)
	a.IndentedJSON(http.StatusOK, gin.H{
		"message": "SUCCESSFUL",
		"result":  empobj,
	})
}

func Deleting(c *gin.Context) {
	var empobj EMP

	id := c.Param("id")
	if err := DB.Where("id = ?", id).Unscoped().Delete(&empobj).Error; err != nil {
		c.JSON(http.StatusNotFound, &empobj)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "SUCCESSFUL",
		"result":  empobj,
	})
}

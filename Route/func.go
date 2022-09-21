package Route

import (
	"EMP_API/Details"
	"github.com/gin-gonic/gin"
	_ "golang.org/x/crypto/openpgp/errors"
	"gorm.io/gorm"
	"net/http"
)

var DB *gorm.DB
var R = gin.Default()

func Creating(c *gin.Context) {
	var empobj Details.EMP

	if err := c.BindJSON(&empobj); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
	}
	if empobj.Name == "" || empobj.Password == "" || empobj.City == "" {
		c.JSON(http.StatusNoContent, &empobj)
	} else {
		c.IndentedJSON(http.StatusOK, &empobj)
	}
	DB.Create(&empobj)
}

func Fetching(a *gin.Context) {
	var empobj []Details.EMP
	DB.Find(&empobj)
	a.IndentedJSON(http.StatusOK, &empobj)
}

func Fbyid(a *gin.Context) {
	var empobj Details.EMP

	id := a.Param("id")
	if err := DB.Where("id = ?", id).First(&empobj).Error; err != nil {
		a.JSON(http.StatusNotFound, &empobj)
		return
	}
	a.JSON(http.StatusOK, empobj)
}

func Updating(a *gin.Context) {
	var empobj Details.EMP

	id := a.Param("id")
	if err := DB.Model(&empobj).Select("id = ?", id).Updates(&empobj).Error; err != nil {
		a.JSON(http.StatusBadRequest, &empobj)
		return
	}
	//err := a.BindJSON(&empobj)
	//if err != nil {
	//	a.JSON(http.StatusNotAcceptable, &empobj)
	//} else {
	a.IndentedJSON(http.StatusOK, &empobj)
	DB.Save(&empobj)
}

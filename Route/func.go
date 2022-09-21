package Route

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

var DB *gorm.DB
var R = gin.Default()

type EMP struct {
	gorm.Model
	Name     string `json:"name"`
	Password string `json:"password"`
	City     string `json:"city"`
}

func Creating(c *gin.Context) {
	var empobj EMP

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
	var empobj []EMP
	DB.Find(&empobj)
	a.IndentedJSON(http.StatusOK, &empobj)
}

func Fbyid(a *gin.Context) {
	var empobj EMP

	id := a.Param("id")
	if err := DB.Where("id = ?", id).First(&empobj).Error; err != nil {
		a.JSON(http.StatusNotFound, &empobj)
		return
	}
	a.JSON(http.StatusOK, empobj)
}

//f unc updating()

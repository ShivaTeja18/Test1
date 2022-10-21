package Methods

import (
	"EMP_API/Details"
	"EMP_API/dbc"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type Handler struct {
	DB *gorm.DB
}

var DB *gorm.DB
var R = gin.Default()

func (h Handler) Fetching(a *gin.Context) {
	var empobj []Details.EMP
	h.DB.Find(&empobj)
	a.IndentedJSON(http.StatusOK, gin.H{
		"message": "SUCCESSFUL",
		"result":  &empobj,
	})
}

func (h Handler) Creating(c *gin.Context) {
	var empobj Details.EMP

	//c.PostForm(empobj.Name)
	//c.Request.FormValue(empobj.City)
	//c.Request.FormValue(empobj.Password)
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
	h.DB.Create(&empobj)
}

func Fbyid(a *gin.Context) {
	var empobj Details.EMP

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
	var empobj Details.EMP

	id := a.Param("id")

	if err := a.ShouldBindJSON(&empobj); err != nil {
		a.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	if err := DB.Where("id = ?", id).Updates(&empobj).Error; err != nil {
		a.JSON(http.StatusNotFound, gin.H{"err": err.Error()})
		return
	}
	//a.JSON(http.StatusBadRequest, &empobj)
	a.IndentedJSON(http.StatusOK, gin.H{
		"message": "SUCCESSFUL",
		"result":  empobj,
	})
}

func Deleting(c *gin.Context) {
	var empobj Details.EMP

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
func Dbinit(db *gorm.DB) Handler {
	return Handler{DB: dbc.Connect()}
}

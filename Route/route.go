package Route

import (
	. "EMP_API/Details"
	. "EMP_API/Methods"
	"fmt"
)

func Initialize() {

	R.GET("/fetch", Fetching)
	R.GET("/fetch/:id", Fbyid)
	R.POST("/create", Creating)
	R.PUT("/change/:id", Updating)
	R.DELETE("/delete/:id", Deleting)
	err := R.Run(":8000")
	if err != nil {
		fmt.Println(err.Error())
	}
}

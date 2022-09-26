package Route

import (
	"EMP_API/Methods"

	"fmt"
)

func Initialize() {

	Methods.R.GET("/fetch", Methods.Fetching)
	Methods.R.GET("/fetch/:id", Methods.Fbyid)
	Methods.R.POST("/create", Methods.Creating)
	Methods.R.PUT("/change/:id", Methods.Updating)
	Methods.R.DELETE("/delete/:id", Methods.Deleting)
	err := Methods.R.Run(":8000")
	if err != nil {
		fmt.Println(err.Error())
	}
}

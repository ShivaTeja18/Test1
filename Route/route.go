package Route

import (
	"EMP_API/Methods"
	"EMP_API/dbc"
	"fmt"
)

var h = Methods.Dbinit(dbc.Connect())

func Initialize() {

	Methods.R.GET("/fetch", h.Fetching)
	Methods.R.GET("/fetch/:id", Methods.Fbyid)
	Methods.R.POST("/create", h.Creating)
	Methods.R.PUT("/change/:id", Methods.Updating)
	Methods.R.DELETE("/delete/:id", Methods.Deleting)
	err := Methods.R.Run(":8000")
	if err != nil {
		fmt.Println(err.Error())
	}

}

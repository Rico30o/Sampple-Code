package webtool

type (
	WTUsers struct {
		// Id       int    `json:"id"`
		RoleId   int    `json:"roleId"`
		Username string `json:"username"`
		Password string `json:"password"`
	}

	RegisterUsers struct {
		RegisterUser string `json:"response"`
	}
)

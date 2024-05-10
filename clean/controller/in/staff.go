package in

type StaffRequest struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Address string `json:"address"`
	Salary  int    `json:"salary"`
}

type DeleteStaffRequest struct {
	ID int `json:"id"`
}

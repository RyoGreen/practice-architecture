package in

type StaffRequest struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type DeleteStaffRequest struct {
	ID int `json:"id"`
}

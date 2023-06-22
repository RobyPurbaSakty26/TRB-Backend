package response

/**
 * Created by Goland & VS Code.
 * User : 1. Roby Purba Sakty 			: obykao26@gmail.com
		  2. Muhammad Irfan 			: mhd.irfann00@gmail.com
   		  3. Andre Rizaldi Brillianto	: andrerizaldib@gmail.com
 * Date: Saturday, 12 Juni 2023
 * Time: 08.30 AM
 * Description: BRI-CMP-Service-Backend
 **/

type UserResponse struct {
	Status string       `json:"status"`
	Data   ItemResponse `json:"data"`
}
type ItemResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	IsActive bool   `json:"is_active"`
	Role     string `json:"role"`
	RoleId   uint   `json:"role_id"`
}

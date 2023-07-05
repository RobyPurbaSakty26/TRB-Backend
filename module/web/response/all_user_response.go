package response

/*
*
  - Created by Goland & VS Code.
  - User : 1. Roby Purba Sakty 			: obykao26@gmail.com
    2. Muhammad Irfan 			: mhd.irfann00@gmail.com
    3. Andre Rizaldi Brillianto	: andrerizaldib@gmail.com
  - Date: Saturday, 12 Juni 2023
  - Time: 08.30 AM
  - Description: BRI-CMP-Service-Backend
    *
*/
type PaginateUserResponse struct {
	Page       int            `json:"page"`
	Limit      int            `json:"limit"`
	Total      int            `json:"total"`
	TotalPages float64        `json:"total_pages"`
	Data       []ItemResponse `json:"data"`
}

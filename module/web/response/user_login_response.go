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

type LoginResponse struct {
	Status string             `json:"status"`
	Data   LoginItemsResponse `json:"data"`
}

type LoginItemsResponse struct {
	Token    string `json:"token"`
	IsActive bool   `json:"is_active"`
}

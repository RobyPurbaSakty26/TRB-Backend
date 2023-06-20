package entity

import "gorm.io/gorm"

/**
 * Created by Goland & VS Code.
 * User : 1. Roby Purba Sakty 			: obykao26@gmail.com
		  2. Muhammad Irfan 			: mhd.irfann00@gmail.com
   		  3. Andre Rizaldi Brillianto	: andrerizaldib@gmail.com
 * Date: Saturday, 12 Juni 2023
 * Time: 08.30 AM
 * Description: BRI-CMP-Service-Backend
 **/

type User struct {
	gorm.Model
	Fullname   string `json:"fullname"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	RoleId     uint   `json:"role_id"`
	Active     bool   `json:"active"`
	Role       Role
	InputFalse int `json:"input_false"`
}

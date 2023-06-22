package helpers

import (
	"golang.org/x/crypto/bcrypt"
)

/**
 * Created by Goland & VS Code.
 * User : 1. Roby Purba Sakty 			: obykao26@gmail.com
		  2. Muhammad Irfan 			: mhd.irfann00@gmail.com
   		  3. Andre Rizaldi Brillianto	: andrerizaldib@gmail.com
 * Date: Saturday, 12 Juni 2023
 * Time: 08.30 AM
 * Description: BRI-CMP-Service-Backend
 **/

func HashPass(p string) (string, error) {
	pass := []byte(p)
	hash, err := bcrypt.GenerateFromPassword(pass, bcrypt.DefaultCost)

	return string(hash), err
}

func ComparePass(h, p []byte) error {
	hash, pass := []byte(h), []byte(p)
	err := bcrypt.CompareHashAndPassword(hash, pass)

	if err != nil {
		return err
	}

	return nil
}

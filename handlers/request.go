package handlers

import "fmt"

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is require", name, typ)
}

//Create Openning/models

type CreateUserRequest struct {
	Name          string `json:"name"`
	LastName      string `json:"lastname"`
	Phone         string `json:"phone"`
	Email         string `json:"email"`
	F_inscripicon string `json:"fecha_inscripcion"`
	Titular       *bool  `json:"titular"`
	CI            string `json:"ci"`
}

func (r *CreateUserRequest) Validate() error {

	/*
		se hace con todos los atributos que posee r
		aqui solo se tiene uno porque el struct solo tiene role
	*/
	// if r.Name == "" && r.LastName == "" && r.Titular == nil && r.Email == "" && r.CI == "" && r.F_inscripicon == ""{
	// 	return fmt.Errorf("request body is empty")
	// }
	if r.Name == "" {
		return errParamIsRequired("Name", "string")
	}
	if r.LastName == "" {
		return errParamIsRequired("Lastname", "string")
	}
	if r.Email == "" {
		return errParamIsRequired("Email", "string")
	}
	if r.Titular == nil {
		return errParamIsRequired("Titular", "Boolean")
	}
	if r.Phone == "" {
		return errParamIsRequired("Phone", "string")
	}
	if r.CI == "" {
		return errParamIsRequired("CI", "string")
	}
	if r.F_inscripicon == "" {
		return errParamIsRequired("Fecha de inscripcion", "string")
	}
	return nil
}
type UpdateUserRequest struct {
	Name          string `json:"name"`
	LastName      string `json:"lastname"`
	Phone         string `json:"phone"`
	Email         string `json:"email"`
	F_inscripicon string `json:"fecha_inscripcion"`
	Titular       *bool  `json:"titular"`
	CI            string `json:"ci"`
}

func (r *UpdateUserRequest) Validate() error{

	if r.Name != "" ||
	r.LastName != "" ||
	r.Phone != "" || 
	r.Email != "" ||  
	r.F_inscripicon != "" ||
	r.Titular != nil || 
	r.CI != "" {
	
	 return nil 
 }
	return fmt.Errorf("at least one valid field must be provied")

}
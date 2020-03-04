package domain

// A Users belong to the domain layer.
type AuthUsers []Auth

// A User belong to the domain layer.
type Auth struct {
	UID  int    `json:"uid"`
	Name string `json:"name"`
}

package graphql

import "github.com/graph-gophers/graphql-go"

// model for enterprises
type EnterpriseGQL struct {
	ID        graphql.ID     `json:"id" bson:"_id,omitempty"`
	Name      string         `json:"name"`
	Logo      string         `json:"logo"`
	Address   *AddressGQL    `json:"address"`
	Employees *[]EmployeeGQL `json:"employees"`
	Media     *MediaGQL      `json:"media"`
}
type AddressGQL struct {
	Country string `json:"country"`
	Zipcode string `json:"zipcode"`
	City    string `json:"city"`
	Street  string `json:"street"`
}
type MediaGQL struct {
	Config  map[string]string `json:"config"`
	Posts   []graphql.ID      `json:"posts"`
	Stories []graphql.ID      `json:"stories"`
}

// model for employees
type EmployeeGQL struct {
	Ref         graphql.ID `json:"ref"`
	Permissions string     `json:"permissions"`
}

// model for users
type UserGQL struct {
	ID          graphql.ID    `json:"id" bson:"_id,omitempty"`
	Email       string        `json:"email"`
	Firstname   string        `json:"firstname"`
	Lastname    string        `json:"lastname"`
	Enterprises *[]graphql.ID `json:"enterprises"`
}

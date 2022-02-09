package request

type ActorCreateRequest struct {
	FirstName string `validation:"required,min=3," json:"first_name"`
	LastName  string `validation:"required,min=3," json:"last_name"`
}

type ActorUpdateRequest struct {
	ActorID   int64  `validation:"required,numeric" json:"actor_id"`
	FirstName string `validation:"required,min=3," json:"first_name"`
	LastName  string `validation:"required,min=3," json:"last_name"`
}

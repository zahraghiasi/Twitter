package requests

type RelationRequest struct {
	UserID uint
}

type PostRequest struct {
	ReTweetID int
	Message   string
}

type LoginRequest struct {
	Email    string
	Password string
}

type FollowRequest struct {
	UserID int
}

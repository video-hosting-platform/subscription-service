package repository

type Repository struct {
	SubscribeRepository
}

type SubscribeRepository interface {
	CreateSubscription(authorID, userID int) error
	IncrementSubscriptions(authorID int) error
	DecrementSubscriptions(authoID int) error
	GetSubscriptionStatus(authorID, userID int) (bool, error)
	GetSubscriptions(userID int) ([]int, error)
	GetSubscriptionCount(authorID int) (int, error)
}

package service

type Service struct {
	SubscriptionsService
}

type SubscriptionsService interface {
	Subscribe(authorID, userID int) error
	Unsubscribe(authorID, userID int) error

	GetSubscriptions(userID int) ([]int, error)
	GetSubscriptionStatus(authorID int) (bool, error)
	GetSubscriptionCount(authorID int) (int, error)
}

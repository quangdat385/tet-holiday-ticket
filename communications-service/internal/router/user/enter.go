package user

type UserRouterGroup struct {
	MessageRouter
	NotificationRouter
	ConversationRouter
	UserInformationRouter
}

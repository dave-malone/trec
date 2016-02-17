package user

import "github.com/xchapter7x/lo"

func newUserRegistrationEvent(user User) {
	lo.G.Debugf("new user registration event %v", user)

	// emailMessage := email.NewMessage(
	// 	"no-reply@therealestatecrm.com",
	// 	user.Email,
	// 	"TheRealEstateCRM.com Verification",
	// 	"{bodyText}",
	// 	"{bodyHtml}",
	// )

	//TODO - propagate event
}

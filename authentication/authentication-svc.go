package authentication

type service struct {
	userInfoRepo UserInfoRepo
	emailRepo    EmailRepo
	cache        Cache
}

func NewService(userInfoRepo UserInfoRepo, emailRepo EmailRepo, cache Cache) Service {
	return &service{
		userInfoRepo: userInfoRepo,
		emailRepo:    emailRepo,
		cache:        cache,
	}
}

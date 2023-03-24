package redisKey

const (
	TokenHeaderKey = "Authorization"

	// UserTokenKey 用户token
	UserTokenKey         = "login-token:%s"
	UserTokenExpireHours = 7 * 24
	UserTokenValue       = "userId"

	// CnyPriceKey 人民币价格每日缓存
	CnyPriceKey = "cny-price-cache:%s"
	// CnyPriceTodayKey 当前价格
	CnyPriceCurrentKey = "cny-price:current"
	CnyPricePrevKey    = "cny-price:prev"
)

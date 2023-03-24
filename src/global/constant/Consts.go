package constant

const (
	ENV_MODE = "CLOUDHERD_ENV"

	PageSize        = "pageSize"
	PageNum         = "pageNum"
	DefaultPageNum  = "1"
	DefaultPageSize = "10"
	TimeStart       = "2006-01-02T15:04:05Z"
	SimpleTimeStart = "2006-01-02"

	// StatusEnable 状态  0：禁用 1：启用
	StatusEnable  = 1
	StatusDisable = 0

	//Rpc
	RpcHeaderKey         = "Authorization"
	RpcHeaderTokenPrefix = "Bearer "

	RpcMethodWithdraw         = "16"
	RpcMethodSend             = "0"
	RpcMethodWithdrawCategory = 0
	RpcMethodSendCategory     = 1
)
const CALC_ACCURACY = 10e6
const SealedPlageAmt = 1 * 10e18

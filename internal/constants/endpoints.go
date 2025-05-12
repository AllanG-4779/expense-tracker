package constants

const (
	ApiEntryPoint      = "/api/v1"
	ApiEntryPointUsers = ApiEntryPoint + "/users"
	ApiLogin           = "/login"
	ApiRegister        = "/register"
	ApiGetUser         = "/me"
)

const (
	SetupEntryPoint       = ApiEntryPoint + "/setup"
	SetupCategoryEndpoint = "/category"
)

const (
	ActivateAccount = "/activate"
	AddTransaction  = "/add/transaction"
	CreateBudget    = "/create/budget"
	GetTransactions = "/get/transactions"
	GetAccounts     = "/get/accounts"
)

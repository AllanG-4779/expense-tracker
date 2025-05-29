package constants

const (
	ApiEntryPoint      = "/api/v1"
	ApiEntryPointUsers = ApiEntryPoint + "/users"
	ApiLogin           = "/login"
	ApiRegister        = "/register"
	ApiGetUser         = "/me"
)

const (
	SetupEntryPoint          = ApiEntryPoint + "/setup"
	SetupCategoryEndpoint    = "/category"
	SetupCategoryEndpointGET = "/category/get"
)

const (
	ActivateAccount   = "/activate"
	AddTransaction    = "/add/transaction"
	CreateBudget      = "/create/budget"
	GetTransactions   = "/get/transactions"
	GetAccounts       = "/get/accounts"
	UpdateTransaction = "/update/transaction"
	FilterTransaction = "/filter/transaction"
	DeleteTransaction = "/delete/transaction"
	DashboardData     = "/dashboard"
)

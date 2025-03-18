package database

const GetUserByEmail = `SELECT id, first_name, last_name, username, email,coalesce(profile_url, '') as profile_url FROM system_users WHERE email = $1 LIMIT 1`
const CreateUser = `INSERT INTO system_users (first_name, last_name, email, username, password) VALUES (:first_name, :last_name, :email,:username,:password)`
const LoginUsernameQuery = `SELECT id, username, password, status, first_name, last_name,email FROM system_users WHERE username = $1 LIMIT 1`
const UpdateUserProfile = `UPDATE system_users SET first_name = :first_name, last_name = :last_name, email = :email, profile_url = :profile_url, password=:password WHERE username = :username`

// CreateCategory setup
const CreateCategory = `INSERT INTO category (name, description, type, icon) VALUES (:name, :description, :type, :icon)`
const GetCategories = `SELECT name, description, type, icon FROM category limit $1 offset $2`
const GetCategory = `SELECT name, description, type, icon FROM category WHERE name = $1`
const UpdateCategory = `UPDATE category SET name = :name, description = :description, type = :type, icon = :icon WHERE name = :name`
const DeleteCategory = `DELETE FROM category WHERE name = :name`

// CreateAccount transactions
const CreateAccount = `INSERT INTO accounts (user_id, name, balance) VALUES (:user_id, :name, :balance)`
const GetAccounts = `SELECT id, user_id, name, balance FROM accounts WHERE user_id = $1 limit $2 offset $3`
const GetAccount = `SELECT id, user_id, name, balance FROM accounts WHERE id = $1`
const UpdateAccount = `UPDATE accounts SET name = :name, balance = :balance WHERE id = :id`
const InsertTransaction = `INSERT INTO transactions (amount, description, date, type, category_id, user_id) VALUES (:amount, :description, :date, :type, :category_id, :account_id, :user_id)`

package database

const GetUserByEmail = `SELECT first_name, last_name, username, email,coalesce(profile_url, '') as profile_url FROM system_users WHERE email = $1 LIMIT 1`

const CreateUser = `INSERT INTO system_users (first_name, last_name, email, username, password) VALUES (:first_name, :last_name, :email,:username,:password)`

const LoginUsernameQuery = `SELECT username, password, status, first_name, last_name,email FROM system_users WHERE username = $1 LIMIT 1`
const UpdateUserProfile = `UPDATE system_users SET first_name = :first_name, last_name = :last_name, email = :email, profile_url = :profile_url WHERE username = :username`

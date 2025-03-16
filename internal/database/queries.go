package database

const GetUserByEmail = `SELECT * FROM system_users WHERE email = $1 LIMIT 1`
const CreateUser = `INSERT INTO system_users (first_name, last_name, email, username, password) VALUES (:first_name, :last_name, :email,:username,:password)`

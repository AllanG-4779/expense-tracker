ALTER TABLE transactions
DROP CONSTRAINT transactions_account_id_fkey, -- Drop the new FK constraint
    DROP COLUMN account_id, -- Remove account_id column
    ADD COLUMN user_id INT NOT NULL REFERENCES system_users(id) ON DELETE CASCADE; -- Restore user_id with FK

ALTER TABLE transactions DROP CONSTRAINT transactions_user_id_fkey,
    DROP COLUMN user_id,
    ADD COLUMN account_id INT NOT NULL REFERENCES accounts(id) ON DELETE CASCADE;

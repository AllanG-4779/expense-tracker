CREATE TABLE system_users (
                              id SERIAL PRIMARY KEY ,
                              first_name VARCHAR(20) NOT NULL,
                              last_name VARCHAR(20) NOT NULL,
                              email VARCHAR(100) NOT NULL UNIQUE ,
                              username VARCHAR(20) NOT NULL UNIQUE ,
                              profile_url VARCHAR(100) NULL,
                              password VARCHAR(256) NOT NULL,
                              status BOOLEAN DEFAULT FALSE,
                              created_at TIMESTAMPTZ DEFAULT now(),
                              updated_at TIMESTAMPTZ DEFAULT now(),
                              deleted BOOLEAN DEFAULT FALSE
);

CREATE TABLE accounts (
                          id SERIAL PRIMARY KEY,
                          user_id INT NOT NULL REFERENCES system_users(id) ON DELETE CASCADE,
                          name VARCHAR(50) NOT NULL,
                          balance DECIMAL(10,2) NOT NULL,
                          created_at TIMESTAMPTZ DEFAULT now(),
                          updated_at TIMESTAMPTZ DEFAULT now(),
                          deleted BOOLEAN DEFAULT FALSE
);

CREATE TABLE category (
                          id SERIAL PRIMARY KEY,
                          name VARCHAR(50) NOT NULL UNIQUE,
                          icon VARCHAR(100) NOT NULL,
                          type VARCHAR(10) NOT NULL check ( TYPE in ('expense', 'income') ),
                          description VARCHAR(100) NULL,
                          created_at TIMESTAMPTZ DEFAULT now(),
                          updated_at TIMESTAMPTZ DEFAULT now(),
                          deleted BOOLEAN DEFAULT FALSE
);

CREATE TABLE transactions (
                              id SERIAL PRIMARY KEY,
                              amount DECIMAL(10,2) NOT NULL,
                              description VARCHAR(100) NOT NULL,
                              date DATE NOT NULL DEFAULT  CURRENT_DATE,
                              type VARCHAR(10) NOT NULL check ( type in ('expense', 'income') ),
                              category_id INT NOT NULL REFERENCES category(id) ON DELETE CASCADE,
                              user_id INT NOT NULL REFERENCES system_users(id) ON DELETE CASCADE,
                              created_at TIMESTAMPTZ DEFAULT now(),
                              updated_at TIMESTAMPTZ DEFAULT now(),
                              deleted BOOLEAN DEFAULT FALSE
);

CREATE TABLE budget (
                        id SERIAL PRIMARY KEY,
                        amount DECIMAL(10,2) NOT NULL,
                        balance DECIMAL(10,2) NOT NULL,
                        category_id INT NOT NULL REFERENCES category(id) ON DELETE CASCADE,
                        user_id INT NOT NULL REFERENCES system_users(id) ON DELETE CASCADE,
                        start_date DATE NOT NULL,
                        end_date DATE NOT NULL,
                        created_at TIMESTAMPTZ DEFAULT now(),
                        updated_at TIMESTAMPTZ DEFAULT now(),
                        deleted BOOLEAN DEFAULT FALSE
);



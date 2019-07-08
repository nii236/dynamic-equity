CREATE TABLE pies (
    id SERIAL PRIMARY KEY,
    currency INTEGER NOT NULL,
    non_cash_multiplier INTEGER NOT NULL,
    cash_multiplier INTEGER NOT NULL,
    commission_rate_percent INTEGER NOT NULL,
    royalty_rate_percent INTEGER NOT NULL,
    finders_fee_percent INTEGER NOT NULL
);

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email TEXT NOT NULL,
    username TEXT NOT NULL
);

CREATE TABLE pies_users (
    user_id INTEGER NOT NULL REFERENCES users(id),
    pie_id INTEGER NOT NULL REFERENCES pies(id),
    PRIMARY KEY (user_id, pie_id)
);

CREATE TABLE chunks (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id),
    pie_id INTEGER REFERENCES pies(id),
    amount INTEGER NOT NULL,
    approved BOOLEAN NOT NULL,
    archived BOOLEAN NOT NULL,
    comment TEXT
);

CREATE TABLE projects (
    id SERIAL PRIMARY KEY,
    pie_id INTEGER REFERENCES pies(id),
    name TEXT NOT NULL
);

CREATE TABLE roles (
    id SERIAL PRIMARY KEY,
    role TEXT NOT NULL,
    action TEXT NOT NULL,
    target TEXT NOT NULL
);

CREATE TABLE users_roles (
    user_id INTEGER NOT NULL REFERENCES users(id),
    role_id INTEGER NOT NULL REFERENCES roles(id),
    PRIMARY KEY (user_id, role_id)
);

-- Drop Foreign Key constraints --
ALTER TABLE IF EXISTS accounts DROP CONSTRAINT IF EXISTS owner_currency_key;

ALTER TABLE IF EXISTS accounts DROP CONSTRAINT IF EXISTS accounts_owner_fkey;

-- Drop Tables --

DROP TABLE IF EXISTS users;
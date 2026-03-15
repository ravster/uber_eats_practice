-- Dumping this all in a simple SQL file because we work with SQL. There's no reason to hide from it.
-- I'm trying to make this as idempotent as possible. That way it can be run on a separate machine easily.
-- When things go into production, my preference is for these commands to be run by teammates, and
-- then removed. Any new devs who want to work on the program then have to get an anonymized and
-- truncated backup of the database and start working. Alternatively, they just get a schema-only backup
-- of the prod-DB and fill it up with data themselves (or with some sort of non-PII seed file). Easy.

CREATE TABLE IF NOT EXISTS cuisines (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS restaurants (
	id SERIAL PRIMARY KEY,
	name VARCHAR(50) NOT NULL UNIQUE,
	description VARCHAR(140) NOT NULL
);

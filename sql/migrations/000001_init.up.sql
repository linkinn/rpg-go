-- ---------------------------
-- Table structure for account
-- ---------------------------
CREATE TABLE accounts (
    id SERIAL PRIMARY KEY,
    nickname VARCHAR(45) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    last_active TIMESTAMP,
    last_ip VARCHAR(20),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- ---------------------------
-- Table structure for clan
-- ---------------------------
CREATE TABLE clans (
    id SERIAL PRIMARY KEY,
    name VARCHAR(45) NOT NULL UNIQUE,
    motto VARCHAR(500),
    level INTEGER DEFAULT 1,
    reputation_score INTEGER DEFAULT 0,
    leader_id SERIAL NOT NULL,
    max_online_member INTEGER DEFAULT 0,
    prev_max_online_member INTEGER DEFAULT 0,
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- ---------------------------
-- Table structure for clan_notice
-- ---------------------------
CREATE TABLE clans_notice (
    id SERIAL PRIMARY KEY,
    clan_id INTEGER REFERENCES clans(id),
    notice VARCHAR(500) NOT NULL,
    is_active BOOLEAN DEFAULT false,
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- ---------------------------
-- Table structure for race
-- ---------------------------
CREATE TABLE races (
    id SERIAL PRIMARY KEY,
    name VARCHAR(45) NOT NULL UNIQUE,
    description VARCHAR(500),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- ---------------------------
-- Table structure for class
-- ---------------------------
CREATE TABLE classes (
    id SERIAL PRIMARY KEY,
    race_id INTEGER REFERENCES races(id),
    name VARCHAR(45) NOT NULL UNIQUE,
    acc INTEGER NOT NULL,
    str INTEGER NOT NULL,
    con INTEGER NOT NULL,
    dex INTEGER NOT NULL,
    men INTEGER NOT NULL,
    wit INTEGER NOT NULL,
    critical INTEGER NOT NULL,
    evasion INTEGER NOT NULL,
    magic_attack INTEGER NOT NULL,
    magic_defense INTEGER NOT NULL,
    magic_speed INTEGER NOT NULL,
    power_attack INTEGER NOT NULL,
    power_defense INTEGER NOT NULL,
    power_speed INTEGER NOT NULL,
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- ---------------------------
-- Table structure for character
-- ---------------------------
CREATE TYPE gender AS ENUM ('m', 'w');

CREATE TABLE characters (
    id SERIAL PRIMARY KEY,
    account_id INTEGER REFERENCES accounts(id),
    clan_id INTEGER REFERENCES clans(id),
    class_id INTEGER REFERENCES classes(id),
    name VARCHAR(45) NOT NULL UNIQUE,
    level INTEGER DEFAULT 0,
    maxHp INTEGER DEFAULT 0,
    curHp INTEGER DEFAULT 0,
    maxMp INTEGER DEFAULT 0,
    curMp INTEGER DEFAULT 0,
    acc INTEGER DEFAULT 0,
    str INTEGER DEFAULT 0,
    con INTEGER DEFAULT 0,
    dex INTEGER DEFAULT 0,
    men INTEGER DEFAULT 0,
    wit INTEGER DEFAULT 0,
    critical INTEGER DEFAULT 0,
    evasion INTEGER DEFAULT 0,
    magic_attack INTEGER DEFAULT 0,
    magic_defense INTEGER DEFAULT 0,
    magic_speed INTEGER DEFAULT 0,
    power_attack INTEGER DEFAULT 0,
    power_defense INTEGER DEFAULT 0,
    power_speed INTEGER DEFAULT 0,
    run_speed INTEGER DEFAULT 0,
    walker_speed INTEGER DEFAULT 0,
    gender gender,
    exp VARCHAR(20),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- ---------------------------
-- Table structure for weapon type
-- ---------------------------
CREATE TABLE weapons_type (
    id SERIAL PRIMARY KEY,
    name VARCHAR(45) NOT NULL UNIQUE,
    description VARCHAR(500),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- ---------------------------
-- Table structure for weapon
-- ---------------------------
CREATE TABLE weapons (
    id SERIAL PRIMARY KEY,
    weapon_type_id INTEGER REFERENCES weapons_type(id),
    name VARCHAR(45) NOT NULL UNIQUE,
    price INTEGER NOT NULL,
    weight INTEGER NOT NULL,
    acc INTEGER NOT NULL,
    str INTEGER NOT NULL,
    con INTEGER NOT NULL,
    dex INTEGER NOT NULL,
    men INTEGER NOT NULL,
    wit INTEGER NOT NULL,
    critical INTEGER NOT NULL,
    magic_attack INTEGER NOT NULL,
    magic_speed INTEGER NOT NULL,
    power_attack INTEGER NOT NULL,
    power_speed INTEGER NOT NULL,
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- ---------------------------
-- Table structure for armor type
-- ---------------------------
CREATE TABLE armors_type (
    id SERIAL PRIMARY KEY,
    name VARCHAR(45) NOT NULL UNIQUE,
    description VARCHAR(500) NOT NULL UNIQUE,
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- ---------------------------
-- Table structure for armor
-- ---------------------------
CREATE TABLE armors (
    id SERIAL PRIMARY KEY,
    armor_type_id INTEGER REFERENCES armors_type(id),
    name VARCHAR(45) NOT NULL UNIQUE,
    body_part VARCHAR(30) NOT NULL,
    price INTEGER NOT NULL,
    weight INTEGER NOT NULL,
    maxHp INTEGER NOT NULL,
    maxMp INTEGER NOT NULL,
    evasion INTEGER NOT NULL,
    magic_defense INTEGER NOT NULL,
    power_defense INTEGER NOT NULL,
    run_speed INTEGER NOT NULL,
    walker_speed INTEGER NOT NULL,
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

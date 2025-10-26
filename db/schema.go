package db

const Schema = `
CREATE TABLE IF NOT EXISTS packs (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	user_id INTEGER NOT NULL,
	pack_name TEXT NOT NULL,
	pack_title TEXT NOT NULL,
	pack_type TEXT NOT NULL,
	pack_link TEXT NOT NULL,
	sticker_count INTEGER NOT NULL,
	created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
	UNIQUE(user_id, pack_name)
);

CREATE INDEX IF NOT EXISTS idx_user_id ON packs(user_id);
CREATE INDEX IF NOT EXISTS idx_created_at ON packs(created_at);

CREATE TABLE IF NOT EXISTS users (
	user_id INTEGER PRIMARY KEY,
	username TEXT,
	first_name TEXT,
	last_name TEXT,
	language_code TEXT,
	is_active INTEGER DEFAULT 1,
	created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
	last_seen_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_users_last_seen ON users(last_seen_at);
CREATE INDEX IF NOT EXISTS idx_users_is_active ON users(is_active);
`


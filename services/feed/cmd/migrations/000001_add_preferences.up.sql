CREATE TABLE IF NOT EXISTS author_preferences (
	user_id BIGINT NOT NULL,
	name_author TEXT NOT NULL,
	priority INT CHECK (priority >= 1 AND priority <= 3),
	created_at TIMESTAMP DEFAULT NOW(),
	PRIMARY KEY (user_id, name_author)
);

CREATE TABLE IF NOT EXISTS category_preferences (
	user_id BIGINT NOT NULL,
	name_category TEXT NOT NULL,
	priority INT CHECK (priority >= 1 AND priority <= 3),
	created_at TIMESTAMP DEFAULT NOW(),
	PRIMARY KEY (user_id, name_category)
);

CREATE TABLE IF NOT EXISTS publisher_preferences (
	user_id BIGINT NOT NULL,
	name_publisher TEXT NOT NULL,
	priority INT CHECK (priority >= 1 AND priority <= 3),
	created_at TIMESTAMP DEFAULT NOW(),
	PRIMARY KEY (user_id, publisher_name)
);

CREATE INDEX idx_author_prefs_user ON author_preferences (user_id);
CREATE INDEX idx_category_prefs_user ON category_preferences (user_id);
CREATE INDEX idx_publisher_prefs_user ON publisher_preferences (user_id);
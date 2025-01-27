DROP INDEX IF EXISTS idx_author_prefs_user;
DROP INDEX IF EXISTS idx_category_prefs_user;
DROP INDEX IF EXISTS idx_publisher_prefs_user;

DROP TABLE IF EXISTS author_preferences;
DROP TABLE IF EXISTS category_preferences;
DROP TABLE IF EXISTS publisher_preferences;

DROP TYPE IF EXISTS preference_type;
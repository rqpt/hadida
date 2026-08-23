DROP INDEX IF EXISTS idx_users_verification_token;

ALTER TABLE users
    DROP COLUMN IF EXISTS is_verified,
    DROP COLUMN IF EXISTS verification_token,
    DROP COLUMN IF EXISTS verification_expires_at;

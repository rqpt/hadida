ALTER TABLE users
    ADD COLUMN is_verified BOOLEAN NOT NULL DEFAULT FALSE,
    ADD COLUMN verification_token VARCHAR(64),
    ADD COLUMN verification_expires_at TIMESTAMPTZ;

CREATE INDEX idx_users_verification_token ON users(verification_token)
    WHERE verification_token IS NOT NULL;

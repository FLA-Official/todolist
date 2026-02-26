CREATE TABLE projects (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    key VARCHAR(10) NOT NULL UNIQUE,
    description TEXT,
    owner_id INT NOT NULL,
    partner INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    end_at TIMESTAMP NULL,

    CONSTRAINT fk_owner
        FOREIGN KEY (owner_id)
        REFERENCES users(id)
        ON DELETE CASCADE,

    CONSTRAINT fk_partner
        FOREIGN KEY (partner)
        REFERENCES users(id)
        ON DELETE SET NULL
);
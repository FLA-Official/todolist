CREATE TABLE project_members (
    id SERIAL PRIMARY KEY,
    project_id INT NOT NULL,
    user_id INT NOT NULL,
    role TEXT DEFAULT 'member' CHECK (role IN ('admin', 'member', 'viewer')),  -- Adjust roles as needed for your app
    joined_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (project_id, user_id),  -- Prevents duplicate memberships
    FOREIGN KEY (project_id) REFERENCES projects(id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE project_members (
    id SERIAL PRIMARY KEY,

    project_id INT NOT NULL,
    user_id INT NOT NULL,

    role VARCHAR(20) NOT NULL DEFAULT 'member',

    joined_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_project_member_project
        FOREIGN KEY (project_id)
        REFERENCES projects(id)
        ON DELETE CASCADE,

    CONSTRAINT fk_project_member_user
        FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON DELETE CASCADE,

    CONSTRAINT unique_project_user
        UNIQUE (project_id, user_id),

    CONSTRAINT role_check
        CHECK (role IN ('owner', 'admin', 'member'))
);
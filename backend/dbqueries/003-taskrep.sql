CREATE TABLE tasks (
    id SERIAL PRIMARY KEY,
    project_id INT NOT NULL,
    title TEXT NOT NULL,
    description TEXT,
    status VARCHAR(20) NOT NULL,
    priority VARCHAR(20) NOT NULL,
    assignee_id INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    end_at TIMESTAMP,

    CONSTRAINT fk_project
        FOREIGN KEY (project_id)
        REFERENCES projects(id)
        ON DELETE CASCADE,

    CONSTRAINT fk_assignee
        FOREIGN KEY (assignee_id)
        REFERENCES users(id)
        ON DELETE SET NULL
);
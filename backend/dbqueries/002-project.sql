CREATE TABLE projects (
	id SERIAL PRIMARY KEY,
	name TEXT NOT NULL,
	description TEXT,
	owner_id INT,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

	CONSTRAINT fk_project_owner
	FOREIGN KEY (owner_id)
	REFERENCES users(id)
	ON DELETE CASCADE
);
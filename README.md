# 🚀 TodoList (Jira-style Task Manager)

A full-stack **Todo / Task Management Application** built to organize tasks efficiently with a clean UI and scalable backend.

---

## 📸 Preview

Unfortunately, no SS yet as I am trying to update it a bit more and looks more stable so stay tuned and please accept my apologies.
<!-- Add screenshots here -->
<!-- ![App Screenshot](./screenshots/app.png) -->

---

## ✨ Features

- ✅ Create, update, and delete tasks
- 📂 Organize tasks into projects
- 🏷️ Task properties (status, due date, etc.)
- ⚡ Fast UI with React + Vite
- 🔧 REST API backend (Go)
- 📦 Scalable project structure

---

## 🛠️ Tech Stack

### Frontend
- React (Vite)
- React Router
- CSS / Tailwind (if used)

### Backend
- Go (Golang)
- REST API

### Database
- PostgreSQL (or planned)

---

## 📁 Project Structure

todolist/
│
├── frontend/ # React app
├── backend/ # Go API server
├── db/ # Database scripts/schema
└── README.md


---

## ⚙️ Getting Started

### 1️⃣ Clone the repo

```bash
git clone https://github.com/FLA-Official/todolist.git
cd todolist

cd frontend
npm install
npm run dev

# if you get
## vite is not recognized
#Then run 
npm install vite

### Run Backend 
cd backend
go mod tidy
go run main.go

4️⃣ Environment Variables

Create .env file in backend:

PORT=8080
DB_URL=your_database_url

| Method | Endpoint   | Description   |
| ------ | ---------- | ------------- |
| GET    | /tasks     | Get all tasks |
| POST   | /tasks     | Create a task |
| PUT    | /tasks/:id | Update a task |
| DELETE | /tasks/:id | Delete a task |


🧠 Purpose

This project is built to:

Practice full-stack development

Learn React + Go integration

Understand REST API design

Simulate a Jira-like system

🔮 Future Improvements

🔐 Authentication (JWT)

👥 Multi-user system

📊 Dashboard & analytics

🧩 Drag-and-drop tasks

☁️ Deployment (Docker + CI/CD)

⭐ Support

If you like this project, give it a ⭐ on GitHub!

📜 License

MIT License

👨‍💻 Author

Farhan Labeeb Apon
GitHub: https://github.com/FLA-Official
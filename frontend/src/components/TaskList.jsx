// src/components/TaskList.jsx
import React from "react";

const TaskList = ({ tasks, onDelete }) => {
  if (!tasks || tasks.length === 0) return <p style={{ textAlign: "center" }}>No tasks yet!</p>;

  return (
    <ul style={{ listStyle: "none", padding: 0 }}>
      {tasks.map((task) => (
        <li
          key={task.id}
          style={{
            border: "1px solid #ccc",
            borderRadius: "5px",
            padding: "1rem",
            marginBottom: "1rem",
            backgroundColor: "#fff",
          }}
        >
          <h3>{task.title}</h3>
          <p>{task.description}</p>
          <small>
            Created: {new Date(task.createdtime).toLocaleString()} | End:{" "}
            {new Date(task.endtime).toLocaleDateString()}
          </small>
          <br />
          <button
            onClick={() => onDelete(task.id)}
            style={{
              marginTop: "0.5rem",
              padding: "0.3rem 0.6rem",
              backgroundColor: "#ff4d4f",
              color: "#fff",
              border: "none",
              borderRadius: "3px",
              cursor: "pointer",
            }}
          >
            Delete
          </button>
        </li>
      ))}
    </ul>
  );
};

export default TaskList;

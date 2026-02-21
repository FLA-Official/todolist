import React, { useEffect, useState } from "react";
import TaskForm from "../components/TaskForm";
import TaskList from "../components/TaskList";

const Home = () => {
  const [tasks, setTasks] = useState([]);
  const [editTask, setEditTask] = useState(null);

  useEffect(() => {
    let isMounted = true;

    const loadTasks = async () => {
      try {
        const res = await fetch("http://localhost:8080/tasks");
        const data = await res.json();
        console.log("Fetched tasks:", data);
        if (isMounted) {
          setTasks(data);
        }
      } catch (err) {
        console.error("Error fetching tasks:", err);
      }
    };

    loadTasks();

    return () => {
      isMounted = false;
    };
  }, []);

  const refreshTasks = async () => {
    try {
      const res = await fetch("http://localhost:8080/tasks");
      const data = await res.json();
      console.log("Fetched tasks:", data);
      setTasks(data);
    } catch (err) {
      console.error("Error fetching tasks:", err);
    }
  };

  // Handle Create OR Update
  const handleSubmit = async (task) => {
    try {
      if (editTask) {
        // UPDATE existing task
        await fetch(`http://localhost:8080/tasks/${editTask.id}`, {
          method: "PUT",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify(task),
        });

        setEditTask(null); // exit edit mode
      } else {
        // CREATE new task
        await fetch("http://localhost:8080/tasks", {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify(task),
        });
      }

      refreshTasks(); // refresh list
    } catch (err) {
      console.error("Error saving task:", err);
    }
  };

  // Delete task
  const handleDelete = async (id) => {
    try {
      await fetch(`http://localhost:8080/tasks/${id}`, {
        method: "DELETE",
      });
      refreshTasks();
    } catch (err) {
      console.error("Error deleting task:", err);
    }
  };

  return (
    <div style={{ padding: "2rem", maxWidth: "800px", margin: "0 auto" }}>
      <h1 style={{ textAlign: "center" }}>Task Manager</h1>

      <TaskForm
        onCreate={handleSubmit}
        editTask={editTask}
      />

      <TaskList
        tasks={tasks}
        onDelete={handleDelete}
        onEdit={setEditTask}
      />
    </div>
  );
};

export default Home;

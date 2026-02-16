import { useState } from "react";

const TaskForm = ({ onCreate }) => {
  const [form, setForm] = useState({
    title: "",
    description: "",
    endtime: "",
  });

  const handleChange = (e) => {
    setForm({ ...form, [e.target.name]: e.target.value });
  };

  const handleSubmit = (e) => {
    e.preventDefault();

    const payload = {
      title: form.title,
      description: form.description,
      endtime: new Date(form.endtime).toISOString(),
    };

    onCreate(payload);
    setForm({ title: "", description: "", endtime: "" });
  };

  return (
    <form onSubmit={handleSubmit} className="form">
      <input
        name="title"
        placeholder="Title"
        value={form.title}
        onChange={handleChange}
        required
      />

      <input
        name="description"
        placeholder="Description"
        value={form.description}
        onChange={handleChange}
        required
      />

      <input
        type="datetime-local"
        name="endtime"
        value={form.endtime}
        onChange={handleChange}
        required
      />

      <button type="submit">Add Task</button>
    </form>
  );
};

export default TaskForm;

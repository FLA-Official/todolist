import { useState } from "react";
import { register as registerAPI } from "../api/auth";

const Register = () => {
  const [data, setData] = useState({ username: "", fullname: "", gmail: "", password: "" });

  const handleSubmit = async () => {
    await registerAPI(data);
    alert("Registered successfully!");
  };

  return (
    <div className="auth">
      <h2>Register</h2>
      <input placeholder="Username" onChange={(e) => setData({ ...data, username: e.target.value })} />
      <input placeholder="Full Name" onChange={(e) => setData({ ...data, fullname: e.target.value })} />
      <input placeholder="Email" onChange={(e) => setData({ ...data, gmail: e.target.value })} />
      <input type="password" placeholder="Password" onChange={(e) => setData({ ...data, password: e.target.value })} />
      <button onClick={handleSubmit}>Register</button>
    </div>
  );
};

export default Register;
import { useState, useContext } from "react";
import { login as loginAPI } from "../api/auth";
import { AuthContext } from "../context/AuthContext";
import Register from "./Register";

const Login = () => {
  const { login } = useContext(AuthContext);
  const [data, setData] = useState({ gmail: "", password: "" });

  const handleSubmit = async () => {
    const res = await loginAPI(data);
    const result = await res.json();
    login(result.token);
  };

  return (
    <div className="auth">
      <h2>Login</h2>
      <input placeholder="Email" onChange={(e) => setData({ ...data, gmail: e.target.value })} />
      <input type="password" placeholder="Password" onChange={(e) => setData({ ...data, password: e.target.value })} />
      <button onClick={handleSubmit}>Login</button>
      <Register />
    </div>
  );
};

export default Login;
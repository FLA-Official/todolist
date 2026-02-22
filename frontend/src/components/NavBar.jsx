import { useContext } from "react";
import { AuthContext } from "../context/AuthContext";

const Navbar = () => {
  const { logout } = useContext(AuthContext);

  return (
    <div className="navbar">
      <h2>ToDo</h2>
      <button onClick={logout}>Logout</button>
    </div>
  );
};

export default Navbar;
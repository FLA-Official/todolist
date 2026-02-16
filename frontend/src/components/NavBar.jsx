// src/components/Navbar.jsx
import React from "react";

const Navbar = () => {
  return (
    <nav style={navStyle}>
      <h1>To-Do</h1>
    </nav>
  );
};

const navStyle = {
  backgroundColor: "#282c34",
  color: "white",
  padding: "1rem",
  textAlign: "center",
};

export default Navbar;

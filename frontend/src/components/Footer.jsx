// src/components/Footer.jsx
import React from "react";

const Footer = () => {
  return (
    <footer style={footerStyle}>
      <p>
        <a
          href="https://github.com/FLA-Official"
          target="_blank"
          rel="noopener noreferrer"
          style={{ color: "#61dafb", textDecoration: "none" }}
        >
          My GitHub
        </a>
      </p>
    </footer>
  );
};

const footerStyle = {
  marginTop: "2rem",
  padding: "1rem",
  textAlign: "center",
  borderTop: "1px solid #ccc",
};

export default Footer;

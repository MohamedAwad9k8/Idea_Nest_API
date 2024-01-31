import React from "react";
import { Link } from "react-router-dom";

const MainPage = ({ loggedUser }) => {
  return (
    <div className="main-page-container">
      <div className="main-page-content">
        <h1 className="main-page-title">WELCOME TO IDEANEST's API APP</h1>
        {loggedUser && (
          <h2 className="main-page-welcome">Welcome, {loggedUser.name}!</h2>
        )}
        <div className="main-page-buttons">
          <Link to="/signup" className="main-page-button">
            <button>Sign Up</button>
          </Link>
          <Link to="/login" className="main-page-button">
            <button>Login</button>
          </Link>
        </div>
      </div>
    </div>
  );
};

export default MainPage;

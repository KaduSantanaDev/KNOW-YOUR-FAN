import React from 'react';
import background from '../assets/images/background.png';
import './Nav.css';

const Nav = () => {
  return (
    <nav className="nav-container">
      {/* <div className="nav-background">
        <img src={background} alt="background" />
      </div> */}
      <div className="nav-content">
        <div className="nav-logo">
          <img src={background} alt="Logo Know Your Fan" className="nav-logo-img" />
          <h1>Know Your Fan</h1>
        </div>
        <div className="nav-links">
          <a href="/" className="nav-link">Home</a>
          <a href="/cadastro" className="nav-link">Cadastro</a>
          <a href="/sobre" className="nav-link">Sobre</a>
        </div>
      </div>
    </nav>
  );
};

export default Nav; 
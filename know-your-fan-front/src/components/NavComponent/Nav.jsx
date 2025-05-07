import React from 'react';
import background from '../../assets/images/background.png';
import './Nav.css';
import { useNavigate } from 'react-router-dom';

const Nav = () => {
  const navigate = useNavigate();
  return (
    <nav className="nav-container">
      <div className="nav-content">
        <div className="nav-logo">
          <img src={background} alt="Logo Know Your Fan" className="nav-logo-img" />
          <h1>Know Your Fan</h1>
        </div>
        <div className="nav-links">
          <a onClick={() => navigate('/')} className="nav-link">Home</a>
          <a onClick={() => navigate('/cadastro')} className="nav-link">Cadastro</a>
          <a onClick={() => navigate('/sobre')} className="nav-link">Sobre</a>
        </div>
      </div>
    </nav>
  );
};

export default Nav; 
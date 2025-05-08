import React from 'react';
import './Home.css';
import { useNavigate } from 'react-router-dom';
export default function Home() {
  const navigate = useNavigate();
  return (
    <>
      <div className="hero-container" id="hero">
        <div className="hero-card">
          <h2>Bem-vindo ao Nosso Serviço</h2>
          <p className="hero-subtitle">
            Conecte-se, valide e gerencie seus documentos com facilidade.
          </p>
          <button onClick={() => navigate('/cadastro')} className="cta-button">
            Começar Agora
          </button>
        </div>
      </div>


      <section id="features" className="features-section">
        <h3>Recursos</h3>
        <div className="features-grid">
          <div className="feature-card">
            <h4>Validação Rápida</h4>
            <p>Valide documentos em segundos com nosso sistema automatizado.</p>
          </div>
          <div className="feature-card">
            <h4>Notificações via Email</h4>
            <p>Receba alertas e atualizações diretamente no seu e-mail.</p>
          </div>
          <div className="feature-card">
            <h4>Dashboard Intuitivo</h4>
            <p>Acompanhe todos os seus documentos em um painel centralizado.</p>
          </div>
        </div>
      </section>

      <section id="contact" className="contact-section">
        <h3>Contato</h3>
        <form className="login-form" action="#" method="post">
          <div className="input-group">
            <label htmlFor="name">Nome</label>
            <input type="text" id="name" name="name" placeholder="Seu nome" required />
          </div>
          <div className="input-group">
            <label htmlFor="email">Email</label>
            <input type="email" id="email" name="email" placeholder="Seu email" required />
          </div>
          <button type="submit" className="login-button">Enviar Mensagem</button>
        </form>
      </section>
      <footer className="footer">
        <p>&copy; 2025 Sua Empresa. Todos os direitos reservados.</p>
      </footer>
    </>
  );
}
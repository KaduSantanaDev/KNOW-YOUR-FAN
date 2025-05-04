import React, { useState } from 'react';
import './App.css';

export default function App() {
  const [username, setUsername] = useState('');
  const [email, setEmail] = useState('');
  const [image, setImage] = useState(null);
  const [cpf, setcpf] = useState(null);

  const [street, setStreet] = useState('');
  const [number, setNumber] = useState('');
  const [complement, setComplement] = useState('');
  const [neighborhood, setNeighborhood] = useState('');
  const [city, setCity] = useState('');
  const [state, setState] = useState('');
  const [zip, setZip] = useState('');



  const handleSubmit = (e) => {
    e.preventDefault();
    console.log('Usuário:', username);
    console.log('Senha:', password);
  };

  return (
    <div className="login-container">
      <form onSubmit={handleSubmit} className="login-form">
        <h2>Cadastro</h2>
        <div className="input-group">
          <label htmlFor="name">Nome completo:</label>
          <input
            type="text"
            id="name"
            value={name}
            onChange={(e) => setUsername(e.target.value)}
            required
          />
        </div>
        <div className="input-group">
          <label htmlFor="email">email:</label>
          <input
            type="email"
            id="email"
            value={email}
            onChange={(e) => setPassword(e.target.value)}
            required
          />
        </div>
        <div className="input-group">
          <label htmlFor="cpf">cpf:</label>
          <input
            type="text"
            id="cpf"
            value={cpf}
            onChange={(e) => setPassword(e.target.value)}
            required
          />
        </div>

        <div className="input-group">
          <label htmlFor="image">Documento:</label>
          <input
            type="file"
            id="image"
            accept="image/*"
            onChange={(e) => setImage(e.target.files[0])}
          />
        </div>

        <h3 style={{ marginTop: '1em' }}>Endereço</h3>

        <div className="input-group">
          <label htmlFor="street">Rua:</label>
          <input
            type="text"
            id="street"
            value={street}
            onChange={(e) => setStreet(e.target.value)}
            required
          />
        </div>

        <div className="input-group">
          <label htmlFor="number">Número:</label>
          <input
            type="text"
            id="number"
            value={number}
            onChange={(e) => setNumber(e.target.value)}
            required
          />
        </div>

        <div className="input-group">
          <label htmlFor="complement">Complemento:</label>
          <input
            type="text"
            id="complement"
            value={complement}
            onChange={(e) => setComplement(e.target.value)}
          />
        </div>

        <div className="input-group">
          <label htmlFor="neighborhood">Bairro:</label>
          <input
            type="text"
            id="neighborhood"
            value={neighborhood}
            onChange={(e) => setNeighborhood(e.target.value)}
            required
          />
        </div>

        <div className="input-group">
          <label htmlFor="city">Cidade:</label>
          <input
            type="text"
            id="city"
            value={city}
            onChange={(e) => setCity(e.target.value)}
            required
          />
        </div>

        <div className="input-group">
          <label htmlFor="state">Estado:</label>
          <input
            type="text"
            id="state"
            value={state}
            onChange={(e) => setState(e.target.value)}
            required
          />
        </div>

        <div className="input-group">
          <label htmlFor="zip">CEP:</label>
          <input
            type="text"
            id="zip"
            value={zip}
            onChange={(e) => setZip(e.target.value)}
            required
          />
        </div>
        <button type="submit" className="login-button">Entrar</button>
      </form>
    </div>
  );
}

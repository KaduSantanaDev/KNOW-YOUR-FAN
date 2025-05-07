import Nav from '../NavComponent/Nav';
import { useState } from 'react';
import { useLocation } from 'react-router-dom';


export default function Document() {
    const location = useLocation();
    const { name, status, id, email } = location.state;
    const [afterStatus, setAfterStatus] = useState(status);
    console.log(status);

    const [showToast, setShowToast] = useState(false);
    const [toastMessage, setToastMessage] = useState('');
    
    console.log({"ID": id});
    const handleValidateDocument = async (e) => {
        e.preventDefault();
        try {
          const res = await fetch(`http://localhost:3031/api/v1/clients?id=${id}`, {
            method: 'GET',
          });
      
          if (!res.ok) {
            setToastMessage('Erro ao validar documento, enviamos um email com mais informações');
            setShowToast(true);
          }
      
          const data = await res.json();
          if (data && data.client) {
            const status = data.client.status ? 'Aprovado' : 'Pendente';
            setAfterStatus(status);
          } else {
            console.warn('Resposta inesperada:', data);
            setAfterStatus(null);
          }
        } catch (error) {
          console.error('Erro ao validar documento:', error);
          setAfterStatus(null);
        }
      };

    return (
        <>
        {showToast && (
            <Toast message={toastMessage} onClose={() => setShowToast(false)} />
        )}
        <Nav className="navbar-container" />
      <div className="d-flex justify-content-center align-items-center" style={{ minHeight: '100vh' }}>
        <div className="card" style={{ width: '18rem' }}>
          <div className="card-body">
            <h5 className="card-title">Nome: {name}</h5>
            <p className="card-text">
              Status: {afterStatus}
            </p>
            <p className="card-text">
              Email: {email}
            </p>
            {afterStatus === 'Aprovado' && (
              <a href="#" className="btn btn-primary disabled cursor-not-allowed">Documento Validado</a>
            )}
            {afterStatus === 'Pendente' && (
              <a className="btn btn-primary" onClick={handleValidateDocument}>Validar Documento</a>
            )}
          </div>
        </div>
      </div>
      </>
    );
  }
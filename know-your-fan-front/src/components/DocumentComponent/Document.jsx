import { useState, useEffect } from 'react';
import { useLocation, useNavigate } from 'react-router-dom';
import Toast from '../ToastComponent/Toast';
import './Document.css';

export default function Document() {
  const location = useLocation();
  const navigate = useNavigate();
  

  useEffect(() => {
    if (!location.state) {
      navigate('/');
    }
  }, [location, navigate]);

  if (!location.state) return null;

  const { name, status, id, email } = location.state;
  const [afterStatus, setAfterStatus] = useState(status);
  const [showToast, setShowToast] = useState(false);
  const [toastMessage, setToastMessage] = useState('');

  const handleValidateDocument = async (e) => {
    e.preventDefault();
    try {
      const res = await fetch(`http://localhost:3031/api/v1/clients?id=${id}`, {
        method: 'GET',
      });

      if (!res.ok) {
        throw new Error();
      }

      const data = await res.json();
      if (data && data.client) {
        const status = data.client.status ? 'Aprovado' : 'Pendente';
        setAfterStatus(status);
        if (status === 'Aprovado') {
          setToastMessage('Documento validado com sucesso');
        } else {
          setToastMessage('Erro ao validar documento, enviamos um email com mais informações. Ou tente novamente.');
        }
      } else {
        setToastMessage('Erro ao validar documento, enviamos um email com mais informações');
      }
    } catch (error) {
      setToastMessage('Erro ao validar documento, enviamos um email com mais informações');
    } finally {
      setShowToast(true);
    }
  };

  return (
    <>
      <Toast message={toastMessage} show={showToast} onClose={() => setShowToast(false)} />
      <div className="document-container">
        <div className="document-card">
          <h2>Validação de Documento</h2>
          <p><strong>Nome:</strong> {name || 'Não possui nome'}</p>
          <p><strong>Email:</strong> {email || 'Não possui email'}</p>
          <p><strong>Status:</strong> {afterStatus || 'Pendente'}</p>
          {afterStatus === 'Aprovado' ? (
            <button className="disabled-button" disabled>Documento Validado</button>
          ) : (
            <button className="login-button" onClick={handleValidateDocument}>Validar Documento</button>
          )}
        </div>
      </div>
    </>
  );
}

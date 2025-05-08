import { useState, useEffect } from 'react';
import { useLocation, useNavigate } from 'react-router-dom';
import Toast from '../ToastComponent/Toast';
import './Document.css';

export default function Document() {
  const location = useLocation();
  const navigate = useNavigate();

  useEffect(() => {
    if (!location.state) navigate('/');
  }, [location, navigate]);

  if (!location.state) return null;

  const { name, status: initialStatus, id, email } = location.state;
  const [isValidated, setIsValidated] = useState(initialStatus);
  const [showToast, setShowToast] = useState(false);
  const [toastMessage, setToastMessage] = useState('');
  const [clickCount, setClickCount] = useState(0);

  const fetchValidation = async () => {
    try {
      const res = await fetch(`http://localhost:3031/api/v1/clients?id=${id}`, {
        method: 'GET',
      });
      if (!res.ok) throw new Error();

      const data = await res.json();
      if (data?.client) {
        const valid = Boolean(data.client.status);
        setIsValidated(valid);
      } else {
        setToastMessage('Erro ao validar documento, enviamos um email com mais informações.');
        setShowToast(true); 
      }
    } catch {
      setToastMessage('Erro ao validar documento, enviamos um email com mais informações.');
      setShowToast(true);
    }
  };

  const handleRevalidate = async (e) => {
    e.preventDefault();
    const newCount = clickCount + 1;
    setClickCount(newCount);

    if (newCount >= 3) {
      setToastMessage('Você tentou validar 3 vezes. Por favor, verifique seu email para mais informações.');
      setShowToast(true);
    } else {
      await fetchValidation();
    }
  };

  useEffect(() => {
    fetchValidation();
  }, [id]);

  return (
    <>
      <Toast
        message={toastMessage}
        show={showToast}
        onClose={() => setShowToast(false)}
      />

      <div className="document-container">
        <div className="document-card">
          <h2>Validação de Documento</h2>
          <p><strong>Nome:</strong> {name || 'Não possui nome'}</p>
          <p><strong>Email:</strong> {email || 'Não possui email'}</p>
          <p><strong>Status:</strong> {isValidated ? 'Aprovado' : 'Pendente'}</p>

          <button
            className={isValidated ? 'disabled-button' : 'login-button'}
            onClick={handleRevalidate}
            disabled={isValidated}
          >
            {isValidated ? 'Documento Validado' : 'Revalidar Documento'}
          </button>
        </div>
      </div>
    </>
  );
}

import React, { useEffect } from 'react';
import './Toast.css';

const Toast = ({ message, show, onClose }) => {
  useEffect(() => {
    if (show) {
      const timer = setTimeout(() => {
        onClose();
      }, 3000); // Auto close after 3 seconds

      return () => clearTimeout(timer);
    }
  }, [show, onClose]);

  if (!show) return null;

  return (
    <div className={`toast ${show ? 'show' : ''}`}>
      <div className="toast-content">
        <p>{message}</p>
      </div>
    </div>
  );
};

export default Toast; 
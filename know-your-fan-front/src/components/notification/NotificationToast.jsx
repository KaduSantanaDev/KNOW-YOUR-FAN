import { useEffect } from "react";

export default function NotificationToast({ message, show, onClose, duration = 3000 }) {
  useEffect(() => {
    if (show) {
      const timer = setTimeout(() => {
        onClose();
      }, duration);
      return () => clearTimeout(timer);
    }
  }, [show, duration, onClose]);

  return (
    show && (
      <div className="fixed bottom-6 right-6 bg-green-500 text-white px-4 py-2 rounded-lg shadow-lg z-50 transition-opacity duration-300">
        {message}
      </div>
    )
  );
}

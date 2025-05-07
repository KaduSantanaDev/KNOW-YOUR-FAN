import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import App from './App';
import Document from './components/DocumentComponent/Document';

export default function AppRouter() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<App />} />
        <Route path="/document" element={<Document />} />
      </Routes>
    </Router>
  );
}
import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import Register from './components/RegisterComponent/Register';
import Document from './components/DocumentComponent/Document';
import Layout from './Layout';
import Home from './components/HomeComponent/Home';

export default function AppRouter() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<Layout />}>
          <Route index element={<Home />} />
          <Route path="/cadastro" element={<Register />} />
          <Route path="/document" element={<Document />} />
        </Route>
      </Routes>
    </Router>
  );
}

// src/App.js
import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import DataTableComponent from './components/DataTableComponent';
import KomisiPage from './components/KomisiPage';
import NavBar from './components/NavBar';
import './App.css';

function App() {
  return (
    <Router>
      <NavBar />
      <Routes>
        <Route path="/" element={<DataTableComponent />} />
        <Route path="/komisi" element={<KomisiPage />} />
      </Routes>
    </Router>
  );
}

export default App;

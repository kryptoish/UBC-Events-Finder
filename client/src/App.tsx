import React from 'react';
import {Routes, Route } from 'react-router-dom';
import Home from '../Home';
import AuthCallback from './AuthCallback';

const App: React.FC = () => {
  return (
    <Routes>
      <Route path="/" element={<Home />} />
      <Route path="/auth/callback" element={<AuthCallback />} />
    </Routes>
  );
};  
export default App;

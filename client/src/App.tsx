import React, { useEffect, useState } from 'react';
import { getGreeting } from './api/api';

const App: React.FC = () => {
  const [message, setMessage] = useState<string>('');
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    getGreeting()
      .then(data => setMessage(data.message))
      .catch(err => {
        console.error('Error fetching greeting:', err);
        setError('Failed to load greeting');
      });
  }, []);

  console.log('API Base URL:', import.meta.env.VITE_API_URL);

  if (error) return <div>{error}</div>;
  return <div>{message || "Loading..."}</div>;
};

export default App;

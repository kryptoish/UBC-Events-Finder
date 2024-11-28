import React, { useEffect, useState } from 'react';
import { getKey} from './api/api';

const AuthCallback: React.FC = () => {
  const [message, setMessage] = useState<string>('');
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const searchParams = new URLSearchParams(location.search);
    const code = searchParams.get('code');

    if (code) {
    getKey(code)
      .then(data => setMessage(data.token))
      .catch(err => {
        console.error('Error fetching greeting:', err);
        setError('Failed to authenticate');
      });
    } else {
      setError('Code not found');
    }
  }, [location.search]);

  if (error) return <div>{error}</div>;
  return <div>{message || "Loading callback..."}</div>;
};

export default AuthCallback;
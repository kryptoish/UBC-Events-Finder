import React, { useEffect, useState } from 'react';
import { getGreeting } from './src/api/api';

const Home: React.FC = () => {
  const [message, setMessage] = useState<string>('');
  const [imageUrl, setImageUrl] = useState<string>('');
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    getGreeting()
      .then(data => { 
        setMessage(data.message)
        setImageUrl(data.imageUrl)
      })
      .catch(err => {
        console.error('Error fetching greeting:', err);
        setError('Failed to load greeting');
      });
  }, []);

  if (error) return <div>{error}</div>;

  return (
    <div>
      <p>{message || "Loading..."}</p>
      {imageUrl && <img src={imageUrl} alt="first Image" style={{ maxWidth: '50%', height: 'auto' }} />}
    </div>
  );
};

export default Home;
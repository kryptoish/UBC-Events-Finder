import React, { useEffect, useState } from 'react';
import { getGreeting } from './api/api.ts';

const App: React.FC = () => {
  const [message, setMessage] = useState<string>('');

  useEffect(() => {
    getGreeting().then(data => setMessage(data.message)).catch(console.error);
  }, []);

  return <div>{message ? message : "Loading..."}</div>;
};

export default App;
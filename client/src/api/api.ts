import axios from 'axios';

const api = axios.create({
  baseURL: import.meta.env.VITE_API_URL || 'https://ubc-events-finder.onrender.com/', // Vite uses `import.meta.env`
});

export const getGreeting = async () => {
  const response = await api.get('/api/main');
  return response.data;
};

export const getKey = async () => {
  const response = await api.get('/api/main');
  return response.data;
};
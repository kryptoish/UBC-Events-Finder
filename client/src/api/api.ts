import axios from 'axios';

const api = axios.create({
  baseURL: import.meta.env.VITE_API_URL, || 'https://ubc-events-finder.onrender.com/' // Vite uses `import.meta.env`  'http://localhost:8080'
});

export const getGreeting = async () => {
  const response = await api.get('/api/main');
  return response.data;
};

export const getKey = async (code: string) => {
  const response = await api.get(`auth/callback?code=${code}`);
  return response.data;
};
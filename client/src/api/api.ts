import axios from 'axios';

const api = axios.create({
  baseURL: 'https://ubc-events-finder.onrender.com/' /* 'http://localhost:8080' */ //import.meta.env.VITE_API_URL, || // Vite uses `import.meta.env`  
});

export const getGreeting = async () => {
  const response = await api.get('/api/main');
  return response.data;
};

export const getKey = async (code: string) => {
  const response = await api.get(`auth/callback?code=${code}`);
  return response.data;
};
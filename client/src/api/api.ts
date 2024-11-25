import axios from 'axios';

const api = axios.create({
  baseURL: import.meta.env.VITE_API_URL || 'http://localhost:8080', // Vite uses `import.meta.env`
});

export const getGreeting = async () => {
  const response = await api.get('/api/greeting');
  return response.data;
};

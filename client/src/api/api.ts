import axios from 'axios';

const api = axios.create({
  baseURL: process.env.REACT_APP_API_URL || 'http://localhost:8080',
});

export const getGreeting = async () => {
  const response = await api.get('/api/greeting');
  return response.data;
};
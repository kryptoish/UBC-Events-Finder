import axios from 'axios';

export interface EventData {
    id: string;
    caption: string;
    media_url: string;
    permalink: string;
    username: string;
    food: string;
    date: string;
    time: string;
    location: string;
}

const api = axios.create({
  baseURL: /* 'https://ubc-events-finder.onrender.com/' */ 'http://localhost:8080'  //import.meta.env.VITE_API_URL, || // Vite uses `import.meta.env`  
});

export const getGreeting = async (): Promise<EventData[]> => {
  try {
    const response = await api.get<EventData[]>('/api/main');
    return response.data;
  } catch (error) {
    console.error('Error fetching events:', error);
    throw new Error('Unable to retrieve events.');
  }
};

export const getKey = async (code: string) => {
  const response = await api.get(`auth/callback?code=${code}`);
  return response.data;
};

import './index.css'
import React, { useEffect, useState } from 'react';
import { getGreeting, EventData } from './api/api';

const Home: React.FC = () => {
  const [events, setEvents] = useState<EventData[]>([]);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
      const getEvents = async () => {
          try {
              const data = await getGreeting();
              const validEvents = data.filter(event => {
                  const eventDate = new Date(`${event.date}T${event.time}`);
                  return eventDate > new Date(); // Discard expired events
              });
              setEvents(validEvents);
          } catch (error) {
              setError('Failed to load events. Please try again later.');
          }
      };

      getEvents();
  }, []);
  /*
  const [events, setEvents] = useState<EventData[]>([]);
  const [caption, setMessage] = useState<string>('');
  const [media_url, setImageUrl] = useState<string>('');
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    getGreeting()
      .then(data => { 
        setMessage(data.caption)
        setImageUrl(data.media_url)
      })
      .catch(err => {
        console.error('Error fetching greeting:', err);
        setError('Failed to load greeting');
      });
  }, []);*/

  if (error) return <div>{error}</div>;

  return (
    <div>
      <p>{events[0].caption || "Loading..."}</p>
      {events[0].media_url && <img src={events[0].media_url} alt="first Image" style={{ maxWidth: '50%', height: 'auto' }} />}
    </div>
  );
};

export default Home;
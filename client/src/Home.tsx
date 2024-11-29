import React, { useEffect, useState } from 'react';
import axios from 'axios';
import './index.css';

interface EventData {
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
    baseURL: 'https://ubc-events-finder.onrender.com/',
});

const CACHE_EXPIRATION_MS = 3600000; // 1 hour

const fetchEvents = async (): Promise<EventData[]> => {
  const cachedData = localStorage.getItem('eventsCache');
  const cacheTimestamp = localStorage.getItem('cacheTimestamp');

  if (cachedData && cacheTimestamp && Date.now() - parseInt(cacheTimestamp) < CACHE_EXPIRATION_MS) {
      return JSON.parse(cachedData);
  }

  try {
      const response = await api.get<{ data: EventData[] }>('/');
      const events = response.data.data;
      localStorage.setItem('eventsCache', JSON.stringify(events));
      localStorage.setItem('cacheTimestamp', Date.now().toString());
      return events;
  } catch (error) {
      console.error('Error fetching events:', error);
      throw new Error('Unable to retrieve events.');
  }
};

const Home: React.FC = () => {
  const [events, setEvents] = useState<EventData[]>([]);
  const [selectedEvent, setSelectedEvent] = useState<EventData | null>(null);
  const [filters, setFilters] = useState({ username: 'all', food: 'all' });
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const getEvents = async () => {
      try {
        const data = await fetchEvents();
        console.log('Fetched events:', data); // Debug log
        setEvents(data); // Ensure only valid data is set
      } catch (err) {
        console.error(err);
        setError('Failed to load events. Please try again later.');
      }
    };
  
    getEvents();
  }, []);

  const filteredEvents = events
  ? events.filter((event) => {
      const usernameMatch = filters.username === 'all' || event.username.toLowerCase() === filters.username;
      const foodMatch = filters.food === 'all' || event.food.toLowerCase() === filters.food;
      return usernameMatch && foodMatch;
    })
  : [];

  const timeFormat = (time: string): string => {
    if (!time) return 'N/A';
    const [hour, minute] = time.split(":").map(Number);
    const meridiem = hour >= 12 ? 'PM' : 'AM';
    const formattedTime = hour % 12 || 12;
    return `${formattedTime}:${minute.toString().padStart(2, '0')} ${meridiem}`;
  };

  console.log('Rendering events:', events)
  return (
    <div className="app">
      <header>
        <h1>UBC Events Finder</h1>
      </header>
      {error ? (
        <div className="error-message">
          <p>{error}</p>
        </div>
      ) : (
        <>
          <div className="filters">
            <select
              onChange={(e) => setFilters({ ...filters, username: e.target.value })}
              value={filters.username}
            >
              <option value="all">All Usernames</option>
              {Array.from(new Set((events || []).map((event) => event.username.toLowerCase()))).map((username) => (
                <option key={username} value={username}>
                  {username}
                </option>
              ))}
            </select>
            <select
              onChange={(e) => setFilters({ ...filters, food: e.target.value })}
              value={filters.food}
            >
              <option value="all">All Foods</option>
              {Array.from(new Set(events.map((event) => event.food.toLowerCase()))).map((food) => (
                <option key={food} value={food}>
                  {food || 'No Specific Food'}
                </option>
              ))}
            </select>
          </div>

          <div className="events-container">
            {filteredEvents.length === 0 ? (
              <p>No events available. Check back later!</p>
            ) : (
              filteredEvents.map((event) => (
                <div
                  key={event.id}
                  className="event-card"
                  onClick={() => setSelectedEvent(event)}
                >
                  <img
                    src={event.media_url || './assets/default.png'}
                    alt={event.caption || 'N/A'}
                    onError={(e) => (e.currentTarget.src = './assets/default.png')}
                  />
                  <div className="event-details">
                    <p><strong>{event.caption || 'N/A'}</strong></p>
                    <p>
                      {event.date || 'No Date Provided'} at {timeFormat(event.time || 'N/A')}
                    </p>
                  </div>
                </div>
              ))
            )}
          </div>
        </>
      )}

      {selectedEvent && (
        <div className="modal" onClick={() => setSelectedEvent(null)}>
          <div className="modal-content" onClick={(e) => e.stopPropagation()}>
            <img
              src={selectedEvent.media_url || './assets/default.png'}
              alt={selectedEvent.caption || 'N/A'}
              onError={(e) => (e.currentTarget.src = './assets/default.png')}
            />
            <h2>{selectedEvent.caption || 'N/A'}</h2>
            <p><strong>Date:</strong> {selectedEvent.date || 'N/A'}</p>
            <p><strong>Time:</strong> {timeFormat(selectedEvent.time || 'N/A')}</p>
            <p><strong>Location:</strong> {selectedEvent.location || 'N/A'}</p>
            <p><strong>Username:</strong> {selectedEvent.username || 'N/A'}</p>
            <p><strong>Food:</strong> {selectedEvent.food || 'N/A'}</p>
            <button onClick={() => setSelectedEvent(null)}>Close</button>
          </div>
        </div>
      )}
    </div>
  );
};

export default Home;

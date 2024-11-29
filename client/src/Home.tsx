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
      const response = await api.get<EventData[]>('/');
      localStorage.setItem('eventsCache', JSON.stringify(response.data));
      localStorage.setItem('cacheTimestamp', Date.now().toString());
      return response.data;
  } catch (error) {
      console.error('Error fetching events:', error);
      throw new Error('Unable to retrieve events.');
  }
};

const Home: React.FC = () => {
  const [events, setEvents] = useState<EventData[]>([]);
  const [selectedEvent, setSelectedEvent] = useState<EventData | null>(null);
  const [filters, setFilters] = useState({ club: 'all', food: 'all' });
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
      const getEvents = async () => {
          try {
              const data = await fetchEvents();
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

  const filteredEvents = events.filter((event) => {
      const clubMatches = filters.club === 'all' || event.username.toLowerCase() === filters.club;
      const foodMatches = filters.food === 'all' || event.food.toLowerCase() === filters.food;
      return clubMatches && foodMatches;
  });

  const getUniqueValues = (key: keyof EventData) => {
      const uniqueValues = Array.from(new Set(events.map((event) => event[key]?.toLowerCase() || 'N/A')));
      uniqueValues.sort();
      return uniqueValues;
  };

  const handleFilterChange = (filterKey: keyof typeof filters, value: string) => {
      setFilters({ ...filters, [filterKey]: value });
  };

  const timeFormat = (time: string): string => {
      if (!time) return 'N/A';
      const [hour, minute] = time.split(":").map(Number);
      const meridiem = hour >= 12 ? 'PM' : 'AM';
      const formattedTime = hour % 12 || 12;
      return `${formattedTime}:${minute.toString().padStart(2, '0')} ${meridiem}`;
  };

  const calculateEndTime = (startTime: string): string => {
      if (!startTime) return 'N/A';
      const [hour, minute] = startTime.split(":").map(Number);
      const endHour = (hour + 1) % 24;
      return `${endHour.toString().padStart(2, '0')}:${minute}`;
  };

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
                          onChange={(e) => handleFilterChange('club', e.target.value)}
                          value={filters.club}
                      >
                          <option value="all">All Clubs</option>
                          {getUniqueValues('username').map((club) => (
                              <option key={club} value={club}>
                                  {club}
                              </option>
                          ))}
                      </select>
                      <select
                          onChange={(e) => handleFilterChange('food', e.target.value)}
                          value={filters.food}
                      >
                          <option value="all">All Foods</option>
                          {getUniqueValues('food').map((food) => (
                              <option key={food} value={food}>
                                  {food}
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
                                      src={`/assets/${event.food || 'default'}.png`}
                                      alt={event.caption || 'N/A'}
                                      onError={(e) => (e.currentTarget.src = './assets/default.png')}
                                  />
                                  <div className="event-details">
                                      <p><strong>{event.caption || 'N/A'}</strong></p>
                                      <p>
                                          {event.date || 'N/A'} at {timeFormat(event.time || 'N/A')}
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
                      <div className="image-wrapper">
                          <img
                              src={selectedEvent.media_url || './assets/default.png'}
                              alt={selectedEvent.caption || 'N/A'}
                              onError={(e) => (e.currentTarget.src = './assets/default.png')}
                          />
                      </div>
                      <h2>{selectedEvent.caption || 'N/A'}</h2>
                      <p>
                          <strong>Date:</strong> {selectedEvent.date || 'N/A'}
                      </p>
                      <p>
                          <strong>Time:</strong> {timeFormat(selectedEvent.time || 'N/A')}
                      </p>
                      <p>
                          <strong>Location:</strong> {selectedEvent.location || 'N/A'}
                      </p>
                      <p>
                          <strong>Club:</strong> {selectedEvent.username || 'N/A'}
                      </p>
                      <p>
                          <strong>Food:</strong> {selectedEvent.food || 'N/A'}
                      </p>
                      <p>
                          <strong>Original Post:</strong>{' '}
                          {selectedEvent.permalink ? (
                              <a href={selectedEvent.permalink} target="_blank" rel="noopener noreferrer">
                                  View Post
                              </a>
                          ) : 'N/A'}
                      </p>
                      <button onClick={() => setSelectedEvent(null)}>Close</button>
                  </div>
              </div>
          )}
      </div>
  );
};

export default Home;

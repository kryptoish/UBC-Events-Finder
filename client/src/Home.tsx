import React, { useEffect, useState } from "react";
import { Cacheables } from "cacheables";
import axios from "axios";
import "./index.css";

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

// Cache instance
const cache = new Cacheables({
  logTiming: true,
  log: true,
});

const Home: React.FC = () => {
  const [events, setEvents] = useState<EventData[]>([]);
  const [filteredEvents, setFilteredEvents] = useState<EventData[]>([]);
  const [selectedEvent, setSelectedEvent] = useState<EventData | null>(null);
  const [filters, setFilters] = useState({ username: "all", food: "all" });
  const [error, setError] = useState<string | null>(null);

  const apiURL = "https://ubc-events-finder.onrender.com/";

  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await cache.cacheable(
          () => axios.get<{ data: EventData[] }>(apiURL),
          apiURL,
          { cachePolicy: "max-age", maxAge: 3600000 } // 1 hour cache
        );

        const data = response?.data?.data || [];
        setEvents(data);
        setFilteredEvents(data);
      } catch (err) {
        console.error("Error fetching events:", err);
        setError("Failed to load events. Please try again later.");
      }
    };

    fetchData();
  }, []);

  // Update filtered events whenever filters or events change
  useEffect(() => {
    const applyFilters = () => {
      const { username, food } = filters;

      const filtered = events.filter((event) => {
        const usernameMatch =
          username === "all" || event.username.toLowerCase() === username;
        const foodMatch =
          food === "all" || event.food.toLowerCase() === food;
        return usernameMatch && foodMatch;
      });

      setFilteredEvents(filtered);
    };

    applyFilters();
  }, [filters, events]);

  // Utility functions
  const timeFormat = (time: string): string => {
    if (!time) return "N/A";

    const [hour, minute] = time.split(":").map(Number);
    if (isNaN(hour) || isNaN(minute)) return "N/A";

    const meridiem = hour >= 12 ? "PM" : "AM";
    const formattedHour = hour % 12 || 12;
    return `${formattedHour}:${minute.toString().padStart(2, "0")} ${meridiem}`;
  };

  const yummyDate = (date: string): string => {
    if (!date) return "N/A";

    const [year, month, day] = date.split("-").map(Number);
    const months = [
      "January", "February", "March", "April", "May", "June",
      "July", "August", "September", "October", "November", "December",
    ];

    return `${months[month - 1]} ${day}, ${year}`;
  };

  return (
    <div className="app">
      <header>
        <h1>UBC Events Finder</h1>
      </header>
      {error ? (
        <div className="error-message">{error}</div>
      ) : (
        <>
          <div className="filters">
            <select
              onChange={(e) =>
                setFilters((prev) => ({ ...prev, username: e.target.value }))
              }
              value={filters.username}
            >
              <option value="all">All Usernames</option>
              {[...new Set(events.map((e) => e.username.toLowerCase()))].map(
                (username) => (
                  <option key={username} value={username}>
                    {username}
                  </option>
                )
              )}
            </select>
            <select
              onChange={(e) =>
                setFilters((prev) => ({ ...prev, food: e.target.value }))
              }
              value={filters.food}
            >
              <option value="all">All Foods</option>
              {[...new Set(events.map((e) => e.food.toLowerCase()))].map(
                (food) => (
                  <option key={food} value={food}>
                    {food}
                  </option>
                )
              )}
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
                    src={event.media_url || "./assets/default.png"}
                    alt={event.caption}
                    onError={(e) =>
                      (e.currentTarget.src = "./assets/default.png")
                    }
                  />
                  <div className="event-details">
                    <p><strong>{event.caption}</strong></p>
                    <p>
                      {yummyDate(event.date)} at {timeFormat(event.time)}
                    </p>
                  </div>
                </div>
              ))
            )}
          </div>
        </>
      )}

      {selectedEvent && (
        <div
          className="modal"
          onClick={() => setSelectedEvent(null)}
        >
          <div
            className="modal-content"
            onClick={(e) => e.stopPropagation()}
          >
            <img
              src={selectedEvent.media_url || "./assets/default.png"}
              alt={selectedEvent.caption}
              onError={(e) =>
                (e.currentTarget.src = "./assets/default.png")
              }
            />
            <h2>{selectedEvent.caption}</h2>
            <p><strong>Date:</strong> {yummyDate(selectedEvent.date)}</p>
            <p><strong>Time:</strong> {timeFormat(selectedEvent.time)}</p>
            <p><strong>Location:</strong> {selectedEvent.location}</p>
            <p><strong>Username:</strong> {selectedEvent.username}</p>
            <p><strong>Food:</strong> {selectedEvent.food}</p>
            <button onClick={() => setSelectedEvent(null)}>Close</button>
          </div>
        </div>
      )}
    </div>
  );
};

export default Home;


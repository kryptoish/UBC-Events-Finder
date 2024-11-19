# Dior Architecture

## Models

### EventModel
- **Responsibility**: 
  - The `EventModel` stores and organizes data on food events, including:
    - Event name
    - Location
    - Time
    - Other relevant details
  - It retrieves data by scraping Instagram posts and newsletter announcements from UBC faculty and sub-organizations.

- **Location**: Server

- **Communication**: 
  - Only the `EventController` communicates with this model.
  - The `EventController` can request:
    - General event details (name, date, time, faculty).
    - Specific content from Instagram posts, including images.

---

### FilterModel
- **Responsibility**: 
  - The `FilterModel` applies filters to the list of events displayed in the `FeedView`, sorting based on user-defined criteria such as:
    - Date
    - Organization
    - Food type

- **Location**: Server

- **Communication**: 
  - Communicates with:
    - **EventModel**: Retrieves all scraped data and filters for relevant events.
    - **FeedView**: Supplies filtered data to display upcoming events.

---

## Controllers

### EventController
- **Responsibility**: 
  - The `EventController` manages data retention and processing for all components.

- **Location**: Both Server and Client side.

- **Communication**: 
  - Interacts with all other components, including:
    - **EventModel**: Retrieves raw event data and updates when necessary.
    - **FeedView**: Sends processed lists of events for display and updates on new events or filters.
    - **ErrorView**: Communicates error messages to handle issues during data retrieval.
    - **EventDetailView**: Sends specific event details upon selection by the user.
    - **FilterModel**: Applies filters and updates the `FeedView` with results.

---

## Views

### FeedView
- **Responsibility**: 
  - Displays a list of all upcoming free food events, serving as the main interface for students to:
    - Browse events
    - Apply filters
    - Select events for details

- **Location**: Client side

- **Communication**: 
  - **EventController**: Requests and receives event lists and updates.
  - **EventDetailView**: Navigates to show detailed event information upon selection.
  - **ErrorView**: Triggers to display error messages if loading issues occur.

---

### ErrorView
- **Responsibility**: 
  - Handles and displays error messages, ensuring users are informed of issues (e.g., network failures).

- **Location**: Client side

- **Communication**: 
  - **EventController**: Receives error messages regarding event data fetching.
  - **FeedView**: Redirects to the `FeedView` once issues are resolved.

---

### EventDetailView
- **Responsibility**: 
  - Displays comprehensive details about a specific event, aiding students in their decision-making process.

- **Location**: Client side

- **Communication**: 
  - **EventModel**: Fetches full details of the selected event.
  - **FeedView**: Returns users to the main feed after viewing event details.

---

## Supporting Services

### Instagram API
- **Description**: 
  - An external service enabling the application to fetch social media posts related to free food events at UBC. 
  - It ensures access to the latest information about food opportunities on campus.

---

### Vercel
- **Description**: 
  - A cloud platform utilized to host the backend database and manage server-side functions of the web application.

---

## Key Pages and Their Functions

### Main Event Feed
- **Appearance**: 
  - A scrollable list of event cards, each displaying:
    - Event name
    - Date
    - Time
    - Location
    - Small image of the food offered

- **Functionality**: 
  - Displays upcoming free food events in chronological order.
  - Allows students to filter events based on preferences (e.g., type of food, location).
  - Enables bookmarking of events for notifications.

---

### Event Details Page
- **Appearance**: 
  - A full-page view with:
    - A large image (event photo or generic food graphic).
    - Detailed information about the event (host organization, times, food description).

- **Functionality**: 
  - Provides context about the event and links to the original social media post when applicable.

---

### Event Submission Form
- **Appearance**: 
  - A simple form containing fields for:
    - Event name
    - Date
    - Location
    - Food description

- **Functionality**: 
  - Allows event organizers to submit new events for visibility.
  - Reviews submissions to ensure only valid events are posted.

---


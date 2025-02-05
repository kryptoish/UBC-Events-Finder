# RELEASE 1.0

### Access the Web App
- **Web App:** [UBC Events Finder](https://ubc-events-finder.vercel.app/)
- **Backend API:** [UBC Events Finder Backend](https://ubc-events-finder.onrender.com/)

> Note: The backend server may take up to **~4 minutes** to start if inactive for over 15 minutes. This is handled gracefully on the frontend.

Changes are automatically deployed when commits are pushed to the repository.

### Video Tutorial
A walkthrough video is available to guide users on navigating the website: *https://youtu.be/Dssr83o3zvk*

---

## Frontend

### New Features
1. **Dynamic Updates:** 
   - The site dynamically updates based on a JSON file fetched from the backend.
   - Events are displayed in order of their proximity to the current time, making it easier to find upcoming events.

2. **Filtering Options:** 
   - Users can filter free food events by **club** and **food type**.
   - Expired posts and events yet to occur are also available for reference.

3. **Event Details:**
   - Each event card includes:
     - **Google Calendar Integration**
     - **Event Time & Location**
     - **Link to the Original Instagram Post**

4. **Responsive Design:**
   - The website adapts seamlessly across devices, providing a smooth user experience on both laptops and phones.

5. **Improved User Experience:**
   - Simplifies the process of finding free food events without searching multiple Instagram accounts.

### Error Handling
- **Data Retrieval Failures:**
  - Cached events are stored for 1 hour and displayed if the backend cannot fetch a new JSON file.
- **Incomplete Captions:**
  - Prompts the user if critical information like the date or location is missing in a caption.
- **Testing Coverage:**
  - Extensive error-handling tests were conducted using **Jest** to ensure robustness.
  - Manual tests were performed by altering Instagram post data to verify site updates.

### Edge Cases
- If no events match the selected filters:
  - A message informs the user to check back later.
- Outdated events:
  - Automatically removed to ensure relevance.

### Exception Handling
- If both the backend and cache fail:
  - Displays an error message informing the user to check back later.

---

## Backend

### New Features
1. **JSON File Updates:** 
   - Regularly updated with data from specific UBC Instagram accounts.
   - Uses **regex** and keyword filtering to identify food-related posts.

2. **Data Parsing:**
   - Each post provides:
     - `ID`, `Caption`, `Media Image`, `Post Link`, `Username`, `Food Type`, `Date`, `Time`, and `Location`.
   - Non-food-related posts are excluded.

3. **Caching:**
   - Data from the previous run is cached and displayed if the server is down.

4. **Server Hosting:**
   - Free hosting leads to a 15-minute spin-down due to inactivity.
   - When active, new data is retrieved every time the page is refreshed.

### Testing
- **High Coverage:** 
  - Go's standard library testing achieved **95% statement coverage** for data processing functions.
  - Run tests with `client/src/api/test.sh`.
- **Instagram API:**
  - Functions interacting with single-use tokens were tested manually with test users on a local setup.

### Error Handling
- **HTTP Status Handling:** 
  - Errors are passed to callback functions, ensuring user feedback.
- **Unparsable Data:**
  - Returns empty strings and directs users to examine the original post for details.

### Edge Cases
- **NLP Limitations:**
  - Regex and keyword-based parsing have natural limitations.
- **Empty Instagram Accounts:**
  - Handled gracefully without causing errors.

---

## Future Improvements
- **Performance Optimization:**
  - Only retrieve posts from the last 6 months to reduce processing load.
- **Deep Learning Integration:**
  - Train a model to improve field parsing.
- **Database Support:**
  - Use PostgreSQL for API key storage and dynamic account addition.
- **Account Sign-Ups:**
  - Provide a link for Instagram accounts to join the service.
- **API Permissions:**
  - Partner with more Instagram accounts for direct API access, reducing reliance on web scraping.

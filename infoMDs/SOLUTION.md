## Design Specification for Team Dior Sausage’s Solution to Food Insecurity at UBC

---

### Problem

Many students at the University of British Columbia (UBC) experience food insecurity due to the rising costs of food on campus. Despite the presence of campus events that frequently offer free food, many students miss out on these opportunities simply because they are unaware of them. This lack of information not only leads to students going hungry but also results in significant food waste at these events.

While UBC provides a food bank, the associated stigma deters some students from using it. Other commercial solutions, such as *Too Good To Go*, are not consistently reliable, as they offer limited availability and may not always meet students' nutritional needs. Therefore, UBC students need an accessible, real-time solution to connect them with free food opportunities on campus, helping to support student welfare and reduce campus food waste.

---

### Solution

Our proposed solution is an interactive web-based application called **FeedUBC**. This web-page will provide students with updates and resources about free food events on campus. It will be designed to be  **simple** and **focused on immediate use cases**, allowing students to access valuable information with minimal effort. This web-page will address the problem by connecting students with food resources through these primary features:

#### Core Features:

1. **Event Feed**:
   - The page’s main screen will present a **scrollable feed** of upcoming events that offer free food. Events are organized chronologically, with current or near-term events displayed at the top.
   - **Filtering options** let students filter events by **location, type of food, and event duration**.


---

### Detailed Design Specification

#### 1. Main Event Feed Screen

- **Layout**: The main screen shows a simple list of upcoming free food events, each displayed in a card format with essential details (e.g., event name, location, time, food type).
- **Interaction**:
  - **Click on an event** to view its full details.
  - **Filter options** (location, food type, event duration) accessible via a button at the top.
  - **Bookmark** an event with a star icon, saving it to the “Bookmarks” section.

#### 2. Event Details Screen

- **Layout**: Each event detail screen displays:
  - Event name, location, start and end times, type of food offered, and any description available.


---

### Mockup Screens

- **Main Event Feed Screen**: Lists all upcoming events.

![Event details mock sketch](images/eventMainFeed.png)

- **Event Details Screen**: Shows complete event details and options for bookmarking.

![Event details mock sketch](images/EventDetails.png)

---

### Error Handling

1. **Event Unavailability**:
   - If no events are currently available, the main screen will display a friendly message, such as “No free food events are available right now. Check back soon!”

2. **Network Issues**:
   - In cases of no connectivity, the app will display a message: “Network unavailable. Please check your connection.”

---

### Technology Stack 

1. **TypeScript**:
   - **Purpose**: The primary web interface for users to access and interact with FeedUBC.
   - **Why**: TypeScript improves reliability and code quality with static typing.

2. **Go**:
   - **Purpose**: The primary backend for scraping data from emails, instagram posts, and webpages.
   - **Why**: Go allows for highly efficient web scraping due to its simplicity, performance, and strong support for concurrency. Libraries such as Colly or GoQuery simplify the development process.

3. **Instagram API**:
   - **Purpose**: Supports parsing from social media, enabling events posted on instagram to be viewed on the web page.
   - **Why**: Instagram API offers easy access to publicly available data, enhancing user experience.

#### Additional Tools and Services

1. **GitHub**:
   - **Purpose**: Source code management, task management, collaboration for development.
   - **Why**: GitHub facilitates version control, team collaboration, issue tracking, and documentation management in Markdown.


2. **Miro** (for Design Mockups):
   - **Purpose**: Provides visual mockups of FeedUBC screens and user flows.
   - **Why**: Miro allows for quick, collaborative design, supporting iterative feedback and alignment on design requirements.

---

### Summary of Technologies and Tools

- **Frontend**: TypeScript
- **Backend**: Go
- **APIs**: Instagram API
- **Design Tools**: Miro, Github

This technology stack is chosen for its ease of integration, cost-effectiveness, scalability, and quick deployment, allowing us to efficiently build and deploy a functional prototype of FeedUBC within the 4-week timeline.

---

### Conclusion

**FeedUBC** aims to provide a real-time, reliable platform that connects UBC students to free food events, helping to address food insecurity while minimizing food waste on campus. The simple, anonymous, and focused interaction design of FeedUBC will make it accessible and appealing to a wide range of students. By implementing this MVP, we can offer immediate value and address a pressing need among UBC’s student body.

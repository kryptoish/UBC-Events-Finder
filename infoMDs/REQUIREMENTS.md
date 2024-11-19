# FeedUBC Requirements Specification

---

## Frontend Requirements

### Core Features

- **Scrollable Event Feed**
  - The main screen must display a **scrollable feed** of upcoming free food events.
  
- **Event Detail Screen**
  - Every event detail screen must display the following information:
    - **Event Name**
    - **Location**
    - **Time**
    - **Food Offered**
  
- **Filtering Functionality**
  - Implement features that allow users to filter events based on:
    - **Location**
    - **Type of Food**
    - **Event Duration**
  
- **Responsive Design**
  - The website must be designed to be used seamlessly across a multitude of devices, including:
    - **Computers**
    - **Mobile Devices**

### User Interface

- **Layout Components**
  - **Scrollable List of Events**
    - A dynamically updating list showcasing all upcoming free food events.
  
  - **Filter Button**
    - Positioned prominently at the **top of the screen** for easy access to filtering options.
  
  - **Header**
    - Includes the application **logo** and **navigation links** to different sections of the website, such as:
      - **Home**
      - **Bookmarks**
      - **Submit Event**

### TypeScript Usage Requirements

- **TypeScript for Frontend Code**
  - **All frontend code** must be written in **TypeScript**. This ensures:
    - **Type Safety**: Reduces runtime errors by catching type-related issues during development.
    - **Enhanced Code Quality**: Improves maintainability and scalability of the codebase.
    - **Improved Developer Experience**: Provides better tooling support, such as autocompletion and refactoring tools.

---

## Backend Requirements

### Data Scraping

- **Instagram API Integration**
  - Integrate the **Instagram API** to parse through UBC faculty accounts for posts highlighting events with free food.
  
- **Multiple Data Sources**
  - The backend must fetch event data from various sources, including:
    - **Designated Email Accounts**
    - **Instagram Posts**
    - **Web Pages**

### Structuring Data

- **Data Organization**
  - The backend must structure and categorize data based on the following parameters:
    - **Date**: Ensures events are listed chronologically.
    - **Location**: Categorizes events by their physical location on campus.
    - **Type of Event**: Differentiates events based on the nature of the free food offered.

### GitHub Collaboration

- **Version Control and Issue Management**
  - Utilize **GitHub** for:
    - **Collaboration**: Facilitates team collaboration through branching and pull requests.
    - **Issue Tracking**: Manages bugs, feature requests, and other project-related tasks efficiently.

### Concurrency and Performance

- **Go’s Concurrency Features**
  - Leverage Go’s concurrency mechanisms to enhance performance, specifically by using libraries such as:
    - **Colly**
    - **GoQuery**
  
- **Parallel Scraping**
  - Implement **parallel scraping** techniques to ensure:
    - **Fast Data Retrieval**: Minimizes the time taken to collect and process event data.
    - **Efficient Resource Utilization**: Optimizes the use of system resources during data scraping.

---

## Error Handling

- **No Events Available**
  - **User Message**: When no events are available, the main screen must display a friendly message:
    > “No free food events are available right now. Check back soon!”
  
- **Network Unavailable**
  - **User Message**: If the user experiences network issues, the main screen must inform them:
    > “Network unavailable. Please check your connection.”

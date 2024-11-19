# Project Coordination and Testing Plan

## Coordination of Work
- **Project Manager**: Ahmad Khan
  - **Responsibilities**:
    - Ensure all team members are on top of their tasks.
    - Provide assistance to members facing difficulties in completing tasks.
  - **Meetings**:
    - Weekly meetings to assess progress.
    - Ahmad will schedule a time and place that accommodates all team members.

---

## Communication Tools
- **Primary Tools**:
  - **GitHub Issues**: For task management and tracking bugs.
  - **Discord**: Main communication platform.
    - **Reason for Choice**: 
      - Organizes conversations into channels for better topic tracking.
- **Alternative Platforms**:
  - Instagram
  - Snapchat
  - iMessage

---

## Component Ownership
| Component            | Owner              |
|----------------------|-------------------|
| FilterModel          | Ahmad Khan        |
| EventModel           | Krish Thakur      |
| ErrorView            | Sayyam Singla     |
| EventDetailView      | Sayyam Singla     |
| FeedView             | Hamzah Chaudhry   |
| EventController      | Quinn             |

---

## Timeline
| Milestone                           | Date     | Description                                           |
|-------------------------------------|----------|-------------------------------------------------------|
| Milestone 1: Frontend Development   | Nov 1    | Start development of frontend components.             |
| Milestone 2: Backend Development     | Nov 1    | Begin development of backend components.              |
| Milestone 3: Half-way Point Review   | Nov 15   | Assess progress and address any issues; adjust plans. |
| Milestone 4: Complete Core Features  | Nov 29   | Finalize core features for frontend and backend.     |
| Milestone 5: Testing Phase           | Dec 3    | Begin testing and bug-finding phase.                 |
| Milestone 6: Finalization             | Dec 6    | Launch web page and finalize documentation.          |

---

## Requirements Verification

### Main Screen Scrollable Feed of Events
- **Verification Plan**: 
  - Each requirement will be checked and validated through a thorough testing plan.

### Acceptance Testing
- **Frontend Requirements**:
  - **Scrollable Event Feed**: 
    - **Test**: Ensure the main screen displays events in a scrollable list in chronological order.
    - **Pass Criteria**: Events load properly, are scrollable, and in the correct order.
  
  - **Event Detail Screen**: 
    - **Test**: Every event detail screen should display the event name, location, time, and food offered.
    - **Pass Criteria**: All details are visible correctly and error-free.
  
  - **Filtering Functionality**: 
    - **Test**: Verify location, type of food, and event duration filters function correctly.
    - **Pass Criteria**: Events feed updates promptly based on filter changes.

  - **Responsive Design**: 
    - **Test**: Check the website's appearance on desktops and mobile devices.
    - **Pass Criteria**: Layout adjusts smoothly depending on screen size.

- **Backend Requirements**:
  - **Scraping and Integrating Data**: 
    - **Test**: Verify integration with the Instagram API fetches event posts accurately.
    - **Pass Criteria**: Events are parsed and stored correctly without duplication or data loss.

  - **API Development**: 
    - **Test**: Ensure all API endpoints respond correctly to valid and invalid requests.
    - **Pass Criteria**: Proper HTTP status codes are returned; data integrity is maintained.

---

## Testing Strategies

### Automated Testing
- **Unit Tests**: 
  - Write unit tests for components using testing frameworks like Jest (frontend) and Go's testing package (backend).
  
- **Integration Tests**: 
  - Ensure various application components interact as expected.

- **End-to-End (E2E) Tests**: 
  - Simulate user interactions from start to finish to verify expected behavior.

### Manual Testing
- **Usability Testing**: 
  - Conduct user sessions to identify usability issues.

- **Responsive Testing**: 
  - Manually verify the application's appearance on different devices and browsers.

- **Security Testing**: 
  - Manually check for vulnerabilities not covered by automated tests.

### Code Reviews
- **Peer Review**: 
  - Implement mandatory code reviews through GitHub Pull Requests.

### Continuous Verification
- **Continuous Integration (CI)**: 
  - Set up GitHub Actions for automated testing on pushes and pull requests.

- **Automated Tests**: 
  - Run unit and integration tests automatically to ensure functionality remains intact.

### Code Quality Tools
- **Linting**: 
  - Configure ESLint for front-end TypeScript code and GolangCI-Lint for Go backend code to maintain consistency.

- **Static Analysis**: 
  - Use tools like SonarQube for early detection of potential issues.

---

This structured Markdown format provides clear guidance for project coordination, communication, component ownership, timelines, and testing strategies, ensuring that all team members can easily follow and contribute to the project’s success.

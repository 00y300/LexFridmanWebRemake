# Project TODO List

This document outlines our tasks for building a NextJS web application with a Golang backend.

---

## 1. Overall Project Setup

- [x] **Establish Project Structure**
  - [x] Set up the NextJS framework for the front end.
  - [x] Set up the Golang server environment for the backend.

- [ ] **Styling & Layout**
  - [ ] Define a consistent layout for the webpage using NextJS.
  - [ ] Establish global style guidelines (colors, fonts, etc.).

---

## 2. Home (Index) Page

- [x] **Navigation & Interactivity**
  - [x] Create buttons in the hero section that link to internal pages.
  - [x] Ensure responsive design and smooth transitions between pages.

---

## 3. Lectures Page

- [ ] **API Research & Integration**
  - [ ] Evaluate APIs to pull videos from a playlist.
  - [x] Research the YouTube API (Golang implementation) for video retrieval.

- [ ] **Backend Development**
  - [x] **Initial Server Setup**
    - [x] Create a basic Golang backend that renders a simple text response to verify connectivity with NextJS.
  - [ ] **Video Data Retrieval**
    - [~] Implement logic to pull videos using the YouTube API.
    - [x] Implement the logic to pull videos from a YouTube playlist by either:
      - [ ] Given a **YouTube Channel ID** and **Playlist ID** retrieve the videos of that specific playlist on the channel.
    - [ ] Explore database options to store video transcripts.
  
- [ ] **Frontend & API Communication**
  - [x] Develop a GET endpoint on the backend for video data.
  - [x] Connect the NextJS frontend to this endpoint to fetch and display video information.
  - [ ] Get the appropriate YouTube Playlist to display on the correct pages.
    - [ ] Get the **Lecture playlist** and its videos on the Lectures page.
    - [ ] Get the **Podcast playlist** and its videos on the Podcast page.
  - [ ] Define and implement custom queries for filtering videos and transcripts.

---

## 4. Future Enhancements & Considerations

- [ ] **Deployment & Testing**
  - [ ] Plan the deployment strategy (consider Docker, CI/CD pipelines).
  - [ ] Set up comprehensive testing for both frontend and backend components.

- [ ] **Error Handling & Logging**
  - [ ] Implement robust error handling on the server.
  - [ ] Integrate logging to monitor application performance and issues.

- [ ] **UI/UX Improvements**
  - [ ] Gather user feedback for iterative UI/UX enhancements.
  - [ ] Refine the interface based on testing results.

---

*This TODO list is a living document. Regularly update it as tasks are completed or new tasks are identified.*

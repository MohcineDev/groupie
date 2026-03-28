# 🎵 Groupie Trackers

Groupie Trackers is a web application that consumes and processes data from a provided API to display information about music artists and bands in a user-friendly interface. The project focuses on backend development in Go, data manipulation, and building interactive client-server features.

---

## 📌 Project Overview

The application retrieves and organizes data from an API composed of four main endpoints:

- **Artists**  
  Contains details about bands and artists:
  - Name
  - Image
  - Year of activity start
  - First album release date
  - Members

- **Locations**  
  Provides information about past and upcoming concert locations.

- **Dates**  
  Lists concert dates (past and upcoming).

- **Relation**  
  Links artists with their respective concert dates and locations.

---

## 🚀 Features

- Display artist and band information in multiple formats:
  - Cards / Blocks
  - Tables / Lists
  - Dedicated detail pages
  - (Optional) Graphs or visualizations

- Interactive **client-server communication**:
  - Trigger events (e.g., filtering, searching, or dynamic loading)
  - Fetch data from the server using request-response mechanisms

- Event-driven functionality:
  - Respond to user actions (clicks, filters, etc.)
  - Dynamically update the UI

---

## 🛠️ Tech Stack

- **Backend:** Go (Golang)
- **Frontend:** HTML, CSS (optional JavaScript)
- **Data Format:** JSON
- **Architecture:** Client-Server (RESTful API consumption)

---

## ⚙️ Installation & Usage

1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/groupie-trackers.git
   cd groupie-trackers

2. Run the server:
go run main.go

3. Open your browser:
http://localhost:8080

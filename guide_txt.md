# Groupie Trackers Project File Structure
## (MAKE SURE TO PUSH ANY WORK ON YOUR OWN BRANCH)
## 1. Go Files (Backend)
- **`main.go`**  
  This file starts the server and handles which page to show when a user visits different parts of the website.

- **`handlers.go`**  
  Contains the logic to fetch artist, concert locations, and dates from the API and then show this data on your website.

- **`api.go`**  
  Handles communication with the external API, bringing in the data you need (e.g., artist info, concert dates).

- **`models.go`**  
  Defines the structure of the data. For example, it defines an "Artist" (name, image, first album, etc.) or a "Concert" (location, date, etc.).

- **`utils.go`** (optional)  
  Extra helper functions for tasks like converting data formats or handling errors can go here.

## 2. HTML Files (Templates for Web Pages)
- **`index.html`**  
  The homepage that lists bands/artists and basic details about them.

- **`artist.html`**  
  Displays detailed information about a single artist, including their image, history, and concert details.

- **`locations.html`**  
  Displays concert locations for past and upcoming events.

- **`dates.html`**  
  Lists concert dates, including which artists are performing and where.

- **`layout.html`** (optional)  
  A template for shared parts of the website, like the header and footer, to avoid repeating the same HTML on every page.

## 3. CSS Files (Styling)
- **`style.css`**  
  Controls how the website looks: colors, fonts, layout, etc.

- **`responsive.css`** (optional)  
  Helps make the site look good on phones and tablets, ensuring it's mobile-friendly.

## 4. Test Files (Testing Your Code)
- **`api_test.go`**  
  Tests if API calls are working correctly and if data is being fetched properly.

- **`handlers_test.go`**  
  Tests the code that serves the correct pages to the user, ensuring everything works as expected.

- **`models_test.go`** (optional)  
  Tests whether the data for artists and concerts is being stored and managed correctly.

## 5. Other
- **`go.mod`**  
  A configuration file for managing your project's Go dependencies.

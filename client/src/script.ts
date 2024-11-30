import { Cacheables } from "cacheables";

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

document.addEventListener("DOMContentLoaded", () => {
    const jsonURL = "https://ubc-events-finder.onrender.com/";

    const loadingSpinner = document.getElementById("loading-spinner")!;
    const errorMessage = "See post for more information.";
    const eventsContainer = document.getElementById("events-container")!;
    const usernameFilter = document.querySelector('#username-filter') as HTMLSelectElement;
    const foodFilter = document.querySelector('#food-filter') as HTMLSelectElement;
    const modal = document.getElementById("event-modal")!;
    const modalDetails = document.getElementById("modal-details")!;
    const modalImage = document.getElementById("modal-image");
    const modalCloseButton = document.getElementById("modal-close");

    let fetchedEvents: EventData[] = [];

    // Fetch JSON data from backend with caching
    async function fetchJSON() {
        if (!eventsContainer) {
            console.error("Events container element not found.");
            return;
        }

        // Show the loading spinner
        loadingSpinner.classList.remove("hidden");

        try {
            const apiResponse = await cache.cacheable(
                () => fetch(jsonURL).then((res) => res.json()),
                jsonURL,
                {
                    cachePolicy: "max-age",
                    maxAge: 3600000, // Cache expires after 1 hour
                }
            );

            // Extract the `data` field from the API response
            if (!apiResponse || !apiResponse.data || !Array.isArray(apiResponse.data)) {
                throw new Error("Invalid API response structure. Expected { data: [...] }");
            }

            const data: EventData[] = apiResponse.data; // Get the array of events
            console.log("Fetched data:", data); // Debugging log

            fetchedEvents = data;

            populateFilters(data); // Pass the array to populateFilters
            renderEvents(data);    // Pass the array to renderEvents
        } catch (error) {
            console.error("Error fetching JSON:", error);
            eventsContainer.innerHTML = "Failed to load events. Please try again later.";
        } finally {
            // Hide the loading spinner
            loadingSpinner.classList.add("hidden");
        }
    }



    // Render events dynamically
    function renderEvents(events: EventData[]) {
        if (!eventsContainer) {
            console.error("Events container element not found.");
            return;
        }

        eventsContainer.innerHTML = ""; // Clear previous event cards
        if (events.length === 0) {
            eventsContainer.innerHTML = "Currently no Free Food Events are available. Check back later!";
            eventsContainer.style.fontSize = "30px";
            return;
        }

        events.forEach((event) => {
            const eventCard = document.createElement("div");
            eventCard.className = "event-card";

            const img = document.createElement("img");
            img.src = event.media_url;
            img.alt = event.caption;

            // Set default media_url if specified media_url is not found
            img.onerror = () => {
                img.src = './assets/default.png';
            };

            const details = document.createElement("div");
            details.className = "event-details";

            const formattedTime = timeFormat(event.time);
            details.innerHTML = `
                <p><strong>${event.caption}</strong></p>
                <p>${yummyDate(event.date)} at ${formattedTime}</p>
            `;

            // Add click event to open modal
            eventCard.addEventListener("click", () => {
                openModal(event);
            });

            eventCard.appendChild(img);
            eventCard.appendChild(details);
            eventsContainer.appendChild(eventCard);
        });
    }

    // Open modal and populate it with event details
    function openModal(event: EventData) {
        const modal = document.getElementById("event-modal")!;
        const modalImage = document.getElementById("event-image") as HTMLImageElement;
        const modalTitle = document.getElementById("event-title")!;
        const modalDate = document.getElementById("event-date")!;
        const modalTime = document.getElementById("event-time")!;
        const modalLocation = document.getElementById("event-location")!;
        const modalUsername = document.getElementById("event-username")!;
        const modalFood = document.getElementById("event-food")!;
        const modalLink = document.getElementById("event-link") as HTMLAnchorElement;

        // Populate Image Section
        modalImage.src = event.media_url || "./assets/default.png";
        modalImage.alt = event.caption || "Event Image";

        // Populate Details Section
        modalTitle.textContent = event.caption || "Event Title";
        modalDate.textContent = event.date || "See post for more information.";
        modalTime.textContent = event.time || "No time provided";
        modalLocation.textContent = event.location || "No location provided";
        modalUsername.textContent = event.username || "N/A";
        modalFood.textContent = event.food || "N/A";
        modalLink.href = event.permalink || "#";
        modalLink.target = "_blank"; // Open in new tab
        modalLink.rel = "noopener noreferrer";

        // Add-to-Calendar Button
        const calendarButton = document.createElement("add-to-calendar-button");
        calendarButton.setAttribute("name", "FREE FOOD brought to you by FeedUBC!");
        calendarButton.setAttribute("description", event.caption || "Event Description");
        calendarButton.setAttribute("startDate", event.date || "2024-01-01");
        calendarButton.setAttribute("startTime", event.time || "12:00 PM");
        calendarButton.setAttribute("endTime", calculateEndTime(event.time));
        calendarButton.setAttribute("timeZone", "America/Vancouver");
        calendarButton.setAttribute("location", event.location || "Event Location");
        calendarButton.setAttribute("options", "'Google'");
        calendarButton.setAttribute("inline", "true");

        // Append the calendar button to the modal details section
        const calendarContainer = document.querySelector(".calendar-button");
        calendarContainer?.replaceChildren(calendarButton); // Replace existing content with the new button

        // Show Modal
        modal.classList.add("visible");
    }



    // Close modal
    function closeModal() {
        if (!modal) {
            console.error("Modal element not found.");
            return;
        }

        modal.classList.remove("visible");
    }

    // Filter events by username and food
    function applyFilters() {
        if (!usernameFilter || !foodFilter) {
            console.error("Filter elements not found.");
            return;
        }
    
        // Show loading spinner
        loadingSpinner.classList.remove("hidden");
    
        const selectedUsername = usernameFilter.value.toLowerCase();
        const selectedFood = foodFilter.value.toLowerCase();
    
        const filteredEvents = fetchedEvents.filter((event) => {
            const usernameMatch = selectedUsername === "all" || event.username.toLowerCase() === selectedUsername;
            const foodMatch = selectedFood === "all" || event.food.toLowerCase() === selectedFood;
            return usernameMatch && foodMatch;
        });
    
        // Render filtered events
        renderEvents(filteredEvents);
    
        // Hide loading spinner
        setTimeout(() => loadingSpinner.classList.add("hidden"), 300); // Add a small delay for smooth UX
    }
    
    // Populate filters with unique values
    function populateFilters(data: EventData[]) {
        if (!usernameFilter || !foodFilter) {
            console.error("Filter elements not found.");
            return;
        }

        // Get unique usernames and foods while filtering out empty strings
        const usernames = getUniqueValues(data, "username").filter((username) => username.trim() !== "");
        const foods = getUniqueValues(data, "food").filter((food) => food.trim() !== "");

        // Set default "All" option
        usernameFilter.innerHTML = `<option value="all">All</option>`;
        foodFilter.innerHTML = `<option value="all">All</option>`;

        // Add unique usernames to the username filter
        usernames.forEach((username) => {
            const option = document.createElement("option");
            option.value = username;
            option.textContent = username;
            usernameFilter.appendChild(option);
        });

        // Add unique foods to the food filter
        foods.forEach((food) => {
            const option = document.createElement("option");
            option.value = food;
            option.textContent = food;
            foodFilter.appendChild(option);
        });
    }


    function getUniqueValues(data: EventData[], key: keyof EventData): string[] {
        if (!Array.isArray(data)) {
            console.error("Expected an array but received:", data);
            return [];
        }
        return Array.from(new Set(data.map((item) => item[key]?.toLowerCase() || ""))).sort();
    }


    // Convert 24-hour time to 12-hour format
    function timeFormat(time: string): string {
        if (!time) return errorMessage; // Handle missing or empty time

        const [hour, minute] = time.split(":").map(Number);
        if (isNaN(hour) || isNaN(minute)) return errorMessage; // Handle invalid time format

        const meridiem = hour >= 12 ? "PM" : "AM";
        const formattedHour = hour % 12 || 12;
        return `${formattedHour}:${minute.toString().padStart(2, "0")} ${meridiem}`;
    }


    // Calculate end time by adding 1 hour to start time
    function calculateEndTime(startTime: string): string {
        const [hour, minute] = startTime.split(":").map(Number);
        const endHour = (hour + 1) % 24;
        return `${endHour.toString().padStart(2, "0")}:${minute.toString().padStart(2, "0")}`;
    }

    // return date in yummy form
    function yummyDate(date: string): string {
        if (!date) return "No Date Provided"; // Handle missing or empty date

        const [year, month, day] = date.split("-").map(Number);
        if (isNaN(year) || isNaN(month) || isNaN(day)) return errorMessage; // Handle invalid date format

        const months = [
            "January", "February", "March", "April", "May", "June",
            "July", "August", "September", "October", "November", "December"
        ];

        const monthString = months[month - 1];
        if (!monthString) return "Invalid Date"; // Handle invalid month

        return `${monthString} ${day}, ${year}`;
    }


    // Attach event listeners
    usernameFilter.addEventListener("change", applyFilters);
    foodFilter.addEventListener("change", applyFilters);
    modalCloseButton?.addEventListener("click", closeModal);
    modal.addEventListener("click", (e) => {
        if (e.target === modal) closeModal();
    });
    fetchJSON(); // Fetch and render data on page load


});

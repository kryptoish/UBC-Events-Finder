
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


document.addEventListener("DOMContentLoaded", () => {
    const jsonURL = "https://ubc-events-finder.onrender.com/";
    
    const CACHE_KEY = "cachedEvents";
    const CACHE_TIMESTAMP_KEY = "cachedTimestamp";
    const CACHE_EXPIRY_TIME = 3600000; // 1 hour in milliseconds

    const loadingSpinner = document.getElementById("loading-spinner")!;
    const errorMessage = "See post for more information.";
    const eventsContainer = document.getElementById("events-container")!;
    const usernameFilter = document.querySelector('#username-filter') as HTMLSelectElement;
    const foodFilter = document.querySelector('#food-filter') as HTMLSelectElement;
    const modal = document.getElementById("event-modal")!;
    const modalDetails = document.getElementById("modal-details")!;
    const modalImage = document.getElementById("modal-image");
    const modalCloseButton = document.getElementById("modal-close");
    const showExpiredToggle = document.getElementById('expired-toggle') as HTMLInputElement;

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
            const cachedData = localStorage.getItem(CACHE_KEY);
            const cachedTimestamp = localStorage.getItem(CACHE_TIMESTAMP_KEY);

            const now = Date.now();

            // Check if cache exists and is not expired
            if (cachedData && cachedTimestamp && now - parseInt(cachedTimestamp) < CACHE_EXPIRY_TIME) {
                console.log("Using cached data");
                fetchedEvents = JSON.parse(cachedData);
                populateFilters(fetchedEvents);
                renderEvents(fetchedEvents);
            } else {
                console.log("Fetching data from API");
                const response = await fetch(jsonURL);
                const apiResponse = await response.json();

                if (!apiResponse || !apiResponse.data || !Array.isArray(apiResponse.data)) {
                    throw new Error("Invalid API response structure. Expected { data: [...] }");
                }

                fetchedEvents = apiResponse.data;

                // Save fetched data and timestamp to localStorage
                localStorage.setItem(CACHE_KEY, JSON.stringify(fetchedEvents));
                localStorage.setItem(CACHE_TIMESTAMP_KEY, now.toString());

                populateFilters(fetchedEvents);
                renderEvents(fetchedEvents);
            }}
            catch (error) {
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

        const showExpired = showExpiredToggle.checked;
        const currentDateTime = new Date();

        eventsContainer.innerHTML = "";

        if (events.length === 0) {
            eventsContainer.innerHTML = "Currently no Free Food Events are available. Check back later!";
            eventsContainer.style.fontSize = "30px";
            return;
        }

        events.forEach((event) => {
            const eventDateTime = event.date && event.time ? new Date(`${event.date}T${event.time}`) : null;
            const isExpired = eventDateTime ? eventDateTime < currentDateTime : false;

            if (isExpired && !showExpired) return;

            const eventCard = document.createElement("div");
            eventCard.className = "event-card";
            eventCard.style.position = "relative";

            const img = document.createElement("img");
            img.src = event.media_url;
            img.alt = event.caption;

            img.onerror = () => {
                img.src = "/src/assets/default.png";
            };

            const details = document.createElement("div");
            details.className = "event-details";

            const formattedTime = event.time ? timeFormat(event.time) : null;
            const formattedDate = event.date ? yummyDate(event.date) : null;

            // Update to handle missing date or time
            if (!formattedDate || !formattedTime) {
                details.innerHTML = `<p><strong>${event.caption}</strong></p><p>${errorMessage}</p>`;
            } else {
                details.innerHTML = `
                    <p><strong>${event.caption}</strong></p>
                    <p>${formattedDate} at ${formattedTime}</p>
                `;
            }

            if (isExpired) {
                const expiredOverlay = document.createElement("div");
                expiredOverlay.className = "expired-overlay";

                const expiredImage = document.createElement("img");
                expiredImage.src = "/src/assets/expired.png";
                expiredImage.alt = "Expired";

                expiredOverlay.appendChild(expiredImage);
                eventCard.appendChild(expiredOverlay);
            }

            eventCard.addEventListener("click", () => {
                openModal(event);
            });

            eventCard.appendChild(img);
            eventCard.appendChild(details);
            eventsContainer.appendChild(eventCard);
        });
    }

    // Add event listener for the toggle button
    showExpiredToggle.addEventListener("change", () => {
        applyFilters(); // Reapply filters when toggle state changes
    });

      
      
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
        modalImage.src = event.media_url || "src/assets/default.png";
        modalImage.alt = event.caption || "Event Image";

        // Populate Details Section
        modalTitle.textContent = event.caption || errorMessage;
        modalDate.textContent = event.date || errorMessage;
        modalTime.textContent = event.time || errorMessage;
        modalLocation.textContent = event.location || errorMessage;
        modalUsername.textContent = event.username || errorMessage;
        modalFood.textContent = event.food || errorMessage;
        modalLink.href = event.permalink || errorMessage;
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
        if (!date) return errorMessage; // Handle missing or empty date

        const [year, month, day] = date.split("-").map(Number);
        if (isNaN(year) || isNaN(month) || isNaN(day)) return errorMessage; // Handle invalid date format

        const months = [
            "January", "February", "March", "April", "May", "June",
            "July", "August", "September", "October", "November", "December"
        ];

        const monthString = months[month - 1];
        if (!monthString) return errorMessage; // Handle invalid month

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

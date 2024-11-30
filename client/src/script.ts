// Interface defining the structure of an event
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
    // Constants for configuration and caching
    const jsonURL = "https://ubc-events-finder.onrender.com/";
    const CACHE_KEY = "cachedEvents";
    const CACHE_TIMESTAMP_KEY = "cachedTimestamp";
    const CACHE_EXPIRY_TIME = 3600000; // 1 hour in milliseconds

    // DOM element references
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

    /**
     * Fetches JSON data from the backend with caching.
     */
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

            // Use cached data if it exists and is not expired
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
            }
        } catch (error) {
            console.error("Error fetching JSON:", error);
            eventsContainer.innerHTML = "Failed to load events. Please try again later.";
        } finally {
            // Hide the loading spinner
            loadingSpinner.classList.add("hidden");
        }
    }

    /**
     * Renders event cards dynamically.
     */
    function renderEvents(events: EventData[]) {
        if (!eventsContainer) {
            console.error("Events container element not found.");
            return;
        }

        const showExpired = showExpiredToggle.checked;
        const currentDateTime = new Date();

        // Clear existing events
        eventsContainer.innerHTML = "";

        // Handle no events case
        if (events.length === 0) {
            eventsContainer.innerHTML = "Currently no Free Food Events are available. Check back later!";
            eventsContainer.style.fontSize = "30px";
            return;
        }

        events.forEach((event) => {
            const eventDateTime = event.date && event.time ? new Date(`${event.date}T${event.time}`) : null;
            const isExpired = eventDateTime ? eventDateTime < currentDateTime : false;

            // Skip expired events if toggle is off
            if (isExpired && !showExpired) return;

            const eventCard = document.createElement("div");
            eventCard.className = "event-card";
            eventCard.style.position = "relative";

            const img = document.createElement("img");
            img.src = event.media_url;
            img.alt = event.caption;

            // Fallback for missing images
            img.onerror = () => {
                img.src = "/assets/default.png";
            };

            const details = document.createElement("div");
            details.className = "event-details";

            const formattedTime = event.time ? timeFormat(event.time) : null;
            const formattedDate = event.date ? yummyDate(event.date) : null;

            // Show default message if date or time is missing
            if (!formattedDate || !formattedTime) {
                details.innerHTML = `<p><strong>${event.caption}</strong></p><p>${errorMessage}</p>`;
            } else {
                details.innerHTML = `
                    <p><strong>${event.caption}</strong></p>
                    <p>${formattedDate} at ${formattedTime}</p>
                `;
            }

            // Add expired overlay for expired events
            if (isExpired) {
                const expiredOverlay = document.createElement("div");
                expiredOverlay.className = "expired-overlay";

                const expiredImage = document.createElement("img");
                expiredImage.src = "/assets/expired.png";
                expiredImage.alt = "Expired";

                expiredOverlay.appendChild(expiredImage);
                eventCard.appendChild(expiredOverlay);
            }

            // Attach click event to open modal
            eventCard.addEventListener("click", () => {
                openModal(event);
            });

            eventCard.appendChild(img);
            eventCard.appendChild(details);
            eventsContainer.appendChild(eventCard);
        });
    }

    /**
     * Opens a modal to display event details.
     */
    function openModal(event: EventData) {
        if (!modal) return;

        const modalImage = document.getElementById("event-image") as HTMLImageElement;
        const modalTitle = document.getElementById("event-title")!;
        const modalDate = document.getElementById("event-date")!;
        const modalTime = document.getElementById("event-time")!;
        const modalLocation = document.getElementById("event-location")!;
        const modalUsername = document.getElementById("event-username")!;
        const modalFood = document.getElementById("event-food")!;
        const modalLink = document.getElementById("event-link") as HTMLAnchorElement;

        modalImage.src = event.media_url || "/assets/default.png";
        modalImage.alt = event.caption || "Event Image";

        modalTitle.textContent = event.caption || errorMessage;
        modalDate.textContent = event.date || errorMessage;
        modalTime.textContent = event.time || errorMessage;
        modalLocation.textContent = event.location || errorMessage;
        modalUsername.textContent = event.username || errorMessage;
        modalFood.textContent = event.food || errorMessage;
        modalLink.href = event.permalink || "#";
        modalLink.target = "_blank";
        modalLink.rel = "noopener noreferrer";

        modal.classList.add("visible");
    }

    /**
     * Closes the modal.
     */
    function closeModal() {
        if (!modal) return;
        modal.classList.remove("visible");
    }

    /**
     * Converts 24-hour time to 12-hour format.
     */
    function timeFormat(time: string): string {
        const [hour, minute] = time.split(":").map(Number);
        const meridiem = hour >= 12 ? "PM" : "AM";
        const formattedHour = hour % 12 || 12;
        return `${formattedHour}:${minute.toString().padStart(2, "0")} ${meridiem}`;
    }

    /**
     * Converts a date string to a more readable format.
     */
    function yummyDate(date: string): string {
        const [year, month, day] = date.split("-").map(Number);
        const months = ["January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"];
        return `${months[month - 1]} ${day}, ${year}`;
    }

    // Attach event listeners
    usernameFilter.addEventListener("change", applyFilters);
    foodFilter.addEventListener("change", applyFilters);
    modalCloseButton?.addEventListener("click", closeModal);
    modal.addEventListener("click", (e) => {
        if (e.target === modal) closeModal();
    });

    // Fetch data and render events on page load
    fetchJSON();
});

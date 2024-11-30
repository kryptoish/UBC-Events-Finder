// import { fetchJSON, timeFormat, yummyDate, applyFilters } from "../src/script";

// jest.mock("./yourModule", () => ({
//     renderEvents: jest.fn(),
// }));

// global.fetch = jest.fn();

// describe("fetchJSON", () => {
//     beforeEach(() => {
//         fetch.mockClear();
//     });

//     it("should fetch data and cache it", async () => {
//         const mockData = {
//             data: [
//                 {
//                     id: "1",
//                     caption: "Test Event",
//                     media_url: "https://example.com/image.jpg",
//                     permalink: "https://example.com/event",
//                     username: "testuser",
//                     food: "Pizza",
//                     date: "2024-11-29",
//                     time: "18:00",
//                     location: "UBC",
//                 },
//             ],
//         };
//         fetch.mockResolvedValueOnce({
//             json: async () => mockData,
//         });

//         const cacheMock = {
//             cacheable: jest.fn(() => Promise.resolve(mockData)),
//         };

//         const result = await fetchJSON("https://test-url.com", cacheMock);

//         expect(fetch).toHaveBeenCalledWith("https://test-url.com");
//         expect(result).toEqual(mockData);
//     });

//     it("should handle fetch errors gracefully", async () => {
//         fetch.mockRejectedValueOnce(new Error("Fetch error"));

//         const cacheMock = {
//             cacheable: jest.fn(() => Promise.reject(new Error("Fetch error"))),
//         };

//         await expect(fetchJSON("https://test-url.com", cacheMock)).rejects.toThrow(
//             "Fetch error"
//         );
//     });
// });

// describe("timeFormat", () => {
//     it("should format time correctly", () => {
//         expect(timeFormat("13:45")).toBe("1:45 PM");
//         expect(timeFormat("00:00")).toBe("12:00 AM");
//         expect(timeFormat("12:00")).toBe("12:00 PM");
//     });

//     it("should return error message for invalid time", () => {
//         expect(timeFormat("")).toBe("See post for more information.");
//         expect(timeFormat("invalid")).toBe("See post for more information.");
//     });
// });

// describe("yummyDate", () => {
//     it("should format date correctly", () => {
//         expect(yummyDate("2024-11-29")).toBe("November 29, 2024");
//         expect(yummyDate("2024-01-01")).toBe("January 1, 2024");
//     });

//     it("should return error message for invalid date", () => {
//         expect(yummyDate("")).toBe("No Date Provided");
//         expect(yummyDate("invalid")).toBe("See post for more information.");
//     });
// });

// describe("applyFilters", () => {
//     it("should filter events based on username and food", () => {
//         const mockEvents = [
//             { username: "user1", food: "Pizza" },
//             { username: "user2", food: "Burger" },
//         ];

//         const usernameFilter = { value: "user1" };
//         const foodFilter = { value: "Pizza" };

//         const mockLoadingSpinner = { classList: { remove: jest.fn(), add: jest.fn() } };

//         applyFilters(mockEvents, usernameFilter, foodFilter, mockLoadingSpinner);

//         expect(renderEvents).toHaveBeenCalledWith([{ username: "user1", food: "Pizza" }]);
//     });
// }); 

import { fetchJSON, timeFormat, yummyDate, applyFilters } from '../src/script.ts';

global.fetch = jest.fn() as jest.Mock;

describe("fetchJSON", () => {
    beforeEach(() => {
        (fetch as jest.Mock).mockClear(); // Properly clear the mock
    });

    it("should fetch data and cache it", async () => {
        const mockData = {
            data: [
                {
                    id: "1",
                    caption: "Test Event",
                    media_url: "https://example.com/image.jpg",
                    permalink: "https://example.com/event",
                    username: "testuser",
                    food: "Pizza",
                    date: "2024-11-29",
                    time: "18:00",
                    location: "UBC",
                },
            ],
        };
        (fetch as jest.Mock).mockResolvedValueOnce({
            json: async () => mockData,
        });

        const cacheMock = {
            cacheable: jest.fn(() => Promise.resolve(mockData)),
        };

        const result = await fetchJSON("https://test-url.com", cacheMock);

        expect(fetch).toHaveBeenCalledWith("https://test-url.com");
        expect(result).toEqual(mockData);
    });

    it("should handle fetch errors gracefully", async () => {
        (fetch as jest.Mock).mockRejectedValueOnce(new Error("Fetch error"));

        const cacheMock = {
            cacheable: jest.fn(() => Promise.reject(new Error("Fetch error"))),
        };

        await expect(fetchJSON("https://test-url.com", cacheMock)).rejects.toThrow(
            "Fetch error"
        );
    });
});
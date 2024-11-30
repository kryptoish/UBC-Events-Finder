"use strict";
// import { fetchJSON } from "/src/script.ts"; // Adjust the import to the correct path
var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
var __generator = (this && this.__generator) || function (thisArg, body) {
    var _ = { label: 0, sent: function() { if (t[0] & 1) throw t[1]; return t[1]; }, trys: [], ops: [] }, f, y, t, g = Object.create((typeof Iterator === "function" ? Iterator : Object).prototype);
    return g.next = verb(0), g["throw"] = verb(1), g["return"] = verb(2), typeof Symbol === "function" && (g[Symbol.iterator] = function() { return this; }), g;
    function verb(n) { return function (v) { return step([n, v]); }; }
    function step(op) {
        if (f) throw new TypeError("Generator is already executing.");
        while (g && (g = 0, op[0] && (_ = 0)), _) try {
            if (f = 1, y && (t = op[0] & 2 ? y["return"] : op[0] ? y["throw"] || ((t = y["return"]) && t.call(y), 0) : y.next) && !(t = t.call(y, op[1])).done) return t;
            if (y = 0, t) op = [op[0] & 2, t.value];
            switch (op[0]) {
                case 0: case 1: t = op; break;
                case 4: _.label++; return { value: op[1], done: false };
                case 5: _.label++; y = op[1]; op = [0]; continue;
                case 7: op = _.ops.pop(); _.trys.pop(); continue;
                default:
                    if (!(t = _.trys, t = t.length > 0 && t[t.length - 1]) && (op[0] === 6 || op[0] === 2)) { _ = 0; continue; }
                    if (op[0] === 3 && (!t || (op[1] > t[0] && op[1] < t[3]))) { _.label = op[1]; break; }
                    if (op[0] === 6 && _.label < t[1]) { _.label = t[1]; t = op; break; }
                    if (t && _.label < t[2]) { _.label = t[2]; _.ops.push(op); break; }
                    if (t[2]) _.ops.pop();
                    _.trys.pop(); continue;
            }
            op = body.call(thisArg, _);
        } catch (e) { op = [6, e]; y = 0; } finally { f = t = 0; }
        if (op[0] & 5) throw op[1]; return { value: op[0] ? op[1] : void 0, done: true };
    }
};
Object.defineProperty(exports, "__esModule", { value: true });
// global.fetch = jest.fn();
// // testing fetching from json file: 
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
var script_ts_1 = require("./script.ts");
jest.mock("./yourModule", function () { return ({
    renderEvents: jest.fn(),
}); });
global.fetch = jest.fn();
describe("fetchJSON", function () {
    beforeEach(function () {
        fetch.mockClear();
    });
    it("should fetch data and cache it", function () { return __awaiter(void 0, void 0, void 0, function () {
        var mockData, cacheMock, result;
        return __generator(this, function (_a) {
            switch (_a.label) {
                case 0:
                    mockData = {
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
                    fetch.mockResolvedValueOnce({
                        json: function () { return __awaiter(void 0, void 0, void 0, function () { return __generator(this, function (_a) {
                            return [2 /*return*/, mockData];
                        }); }); },
                    });
                    cacheMock = {
                        cacheable: jest.fn(function () { return Promise.resolve(mockData); }),
                    };
                    return [4 /*yield*/, (0, script_ts_1.fetchJSON)("https://test-url.com", cacheMock)];
                case 1:
                    result = _a.sent();
                    expect(fetch).toHaveBeenCalledWith("https://test-url.com");
                    expect(result).toEqual(mockData);
                    return [2 /*return*/];
            }
        });
    }); });
    it("should handle fetch errors gracefully", function () { return __awaiter(void 0, void 0, void 0, function () {
        var cacheMock;
        return __generator(this, function (_a) {
            switch (_a.label) {
                case 0:
                    fetch.mockRejectedValueOnce(new Error("Fetch error"));
                    cacheMock = {
                        cacheable: jest.fn(function () { return Promise.reject(new Error("Fetch error")); }),
                    };
                    return [4 /*yield*/, expect((0, script_ts_1.fetchJSON)("https://test-url.com", cacheMock)).rejects.toThrow("Fetch error")];
                case 1:
                    _a.sent();
                    return [2 /*return*/];
            }
        });
    }); });
});
describe("timeFormat", function () {
    it("should format time correctly", function () {
        expect((0, script_ts_1.timeFormat)("13:45")).toBe("1:45 PM");
        expect((0, script_ts_1.timeFormat)("00:00")).toBe("12:00 AM");
        expect((0, script_ts_1.timeFormat)("12:00")).toBe("12:00 PM");
    });
    it("should return error message for invalid time", function () {
        expect((0, script_ts_1.timeFormat)("")).toBe("See post for more information.");
        expect((0, script_ts_1.timeFormat)("invalid")).toBe("See post for more information.");
    });
});
describe("yummyDate", function () {
    it("should format date correctly", function () {
        expect((0, script_ts_1.yummyDate)("2024-11-29")).toBe("November 29, 2024");
        expect((0, script_ts_1.yummyDate)("2024-01-01")).toBe("January 1, 2024");
    });
    it("should return error message for invalid date", function () {
        expect((0, script_ts_1.yummyDate)("")).toBe("No Date Provided");
        expect((0, script_ts_1.yummyDate)("invalid")).toBe("See post for more information.");
    });
});
describe("applyFilters", function () {
    it("should filter events based on username and food", function () {
        var mockEvents = [
            { username: "user1", food: "Pizza" },
            { username: "user2", food: "Burger" },
        ];
        var usernameFilter = { value: "user1" };
        var foodFilter = { value: "Pizza" };
        var mockLoadingSpinner = { classList: { remove: jest.fn(), add: jest.fn() } };
        (0, script_ts_1.applyFilters)(mockEvents, usernameFilter, foodFilter, mockLoadingSpinner);
        expect(renderEvents).toHaveBeenCalledWith([{ username: "user1", food: "Pizza" }]);
    });
});

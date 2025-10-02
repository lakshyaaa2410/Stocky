# Stocky

Stocky is a hypothetical platform that rewards users with shares of Indian stocks (like Reliance, TCS, Infosys) for actions such as onboarding, referrals, or trading milestones. Users receive full stock units, while Stocky handles purchases from NSE/BSE, including brokerage, taxes, and regulatory fees, which are tracked internally.

The system also updates the INR valuation of user holdings hourly via a hypothetical price service.

## Tech Stack

**Backend:** GoLang (Gin, and GORM)

**Logging:** Logrus

**Database:** PostgreSQL

**API Client:** Postman

## Endpoints

-   **POST `/reward`**  
    Adds a reward for a user. Prevents duplicates for actions like "Onboarding," calculates deductions (Brokerage 5%, GST 18%), and logs all transactions in the ledger.

-   **GET `/today-stocks/:userId`**  
    Fetches all rewards for a user for the current day.

-   **GET `/stats/:userId`**  
    Fetches user statistics, including total rewards and stock performance.

-   **GET `/historical-inr/:userId`**  
    Retrieves the historical stock value in INR for a user.

-   **PUT `/update-stock-prices`**  
    Manually triggers an update for stock prices.

-   **GET `/portfolio/:userId`**  
    Retrieves the current portfolio of a user, including stocks and cash balance.

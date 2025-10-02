# Stocky

Stocky is a hypothetical platform that rewards users with shares of Indian stocks (like Reliance, TCS, Infosys) for actions such as onboarding, referrals, or trading milestones. Users receive full stock units, while Stocky handles purchases from NSE/BSE, including brokerage, taxes, and regulatory fees, which are tracked internally.

The system also updates the INR valuation of user holdings hourly via a hypothetical price service.

---

## Tech Stack

-   **Backend:** GoLang (Gin, and GORM)
-   **Logging:** Logrus
-   **Database:** PostgreSQL
-   **API Client:** Postman

---

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

---

## Setup & Installation

Follow these steps to clone and run the Stocky backend on your local machine:

1. **Clone the repository**

    ```bash
    git clone https://github.com/lakshyaaa2410/Stocky.git
    cd Stocky

    ```

2. **Install Go dependencies**

    ```bash
    go mod tidy

    ```

3. **Set up PostgreSQL**

-   Install PostgreSQL locally (or use a managed instance like Neon).

-   Create a database (e.g. assignment as per your config).

-   Create a .env file at the project root with database credentials and other configuration values.

-   Example .env
    ```bash
    DB_URL=postgres://username:password@localhost:5432/stocky?sslmode=disable
    ```

4. **Run database migrations**

-   On Windows:

    ```bash
    go run .\migration\dbMigration.go

    ```

-   On Linux / Mac:
    ```bash
    go run ./migration/dbMigration.go
    ```

5. **Start the server**

    ```bash
    go run main.go

    ```

6. **Test the APIs**
   Use the Postman collection or your preferred API client to hit endpoints.

---

## Deployment

The Stocky backend is deployed on **Render**, and the PostgreSQL database is hosted on **Neon**. This allows you to access the API endpoints without running the backend locally.

-   **Backend:** Hosted on Render – automatically builds and deploys from the GitHub repository.  
    [![Visit Render](https://img.shields.io/badge/Render-Stocky-blue?logo=render)](https://stocky-ffqq.onrender.com)

    > **Note:** Render may take some time to respond on the first request due to cold start.

-   **Database:** Hosted on Neon – a managed PostgreSQL service providing secure access to the Stocky database.
-   **Environment Variables:** Ensure the `.env` file is configured with the database URL and any other secrets required by the backend.

You can directly test the API endpoints using the Postman collection:

[![Run in Postman](https://run.pstmn.io/button.svg)](https://www.postman.com/jashn2/stocky-stock-rewards/overview)

---

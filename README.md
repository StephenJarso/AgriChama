🌾 AgriChama

AgriChama is a digital micro-economy engine designed to empower Village Savings and Loan Associations (VSLAs), smallholder farmers, and communities in disaster-prone regions. It bridges the gap between traditional communal ledger-based savings and modern financial inclusion using a high-performance Go backend and a responsive React frontend.
🚀 Technical Stack
Backend (The Engine)

    Language: Go (Golang) 1.22+

    Framework: Gin Gonic (High-performance HTTP routing)

    ORM: GORM (v2) with PostgreSQL

    Database: PostgreSQL (ACID compliant for financial integrity)

    Auth: JWT (JSON Web Tokens) with Role-Based Access Control (RBAC)

Frontend (The Interface)

    Framework: React 18 (Vite)

    Styling: Tailwind CSS v4 (Bento Grid Layout)

    Animation: Framer Motion

    Icons: Lucide React

📂 Project Structure (Domain-Based)

The project follows Domain-Driven Design (DDD) principles to ensure that business logic is isolated and scalable.
Plaintext

agrichama/
├── backend/
│   ├── cmd/api/main.go          # Entry point & Dependency Injection
│   ├── internal/
│   │   ├── domain/              # BOUNDED CONTEXTS
│   │   │   ├── treasury/        # Savings, Deposits, Group Balances
│   │   │   ├── lending/         # Loan Lifecycle & TrustScore Logic
│   │   │   ├── soko/            # Marketplace & Inventory
│   │   │   └── risk/            # Disaster Alerts & NGO Tasks
│   │   ├── platform/            # Infrastructure (DB, Redis, Logger)
│   │   └── middleware/          # Auth, CORS, RBAC
│   └── migrations/              # Versioned SQL migrations
│
├── frontend/
│   ├── src/
│   │   ├── app/                 # Providers & Global Router
│   │   ├── domains/             # MIRRORED BUSINESS MODULES
│   │   │   ├── treasury/        # <BalanceCard />, <DepositForm />
│   │   │   ├── lending/         # <LoanRequest />, <TrustMeter />
│   │   │   └── soko/            # <MarketplaceGrid />, <ProductCard />
│   │   ├── shared/              # Design System (Bento Components)
│   │   └── layouts/             # DashboardLayout, AuthLayout
│   └── tailwind.config.js
│
└── docker-compose.yml           # Full stack: Go + Postgres + Redis

🏛 Architecture & Design Philosophy
1. Financial Integrity

    Immutable Ledger: Every transaction (saving or repayment) is recorded as a unique event. We use Soft Deletes (DeletedAt)—financial data is never permanently removed from the database.

    ACID Transactions: All money movements (e.g., disbursing a loan from the treasury) are wrapped in Postgres Transactions via GORM to prevent data corruption.

2. Aesthetic: "Modern Earthy Resilience"

    Palette: Olive Primary (#5A5A40), Warm Background (#f5f5f0), Sepia Accents (#fbf5e9).

    Layout: Bento-grid inspired dashboards with 24px rounded cards to provide a trustworthy, organized user experience for non-technical users.

🛠 Role-Based Access Control (RBAC)
Role	Responsibility
Leader	VSLA Administrator; approves loans and manages group settings.
Member	Community user; saves, borrows, and reports risk.
NGO	Partner; deploys impact tasks and emergency relief funding.
Buyer	Public; purchases agricultural produce from the Soko marketplace.
🚦 GitHub Workflow & Team Rules

    Branching:

        main: Production-only (Locked).

        develop: Integration branch for features.

        feat/domain-task: Individual feature branches.

    Commits: Use Conventional Commits (e.g., feat(lending): add interest calculation logic).

    Code Review: Every Pull Request requires at least one peer approval and must pass go test ./....

    Security: Never commit .env files. Ensure all PII (Personally Identifiable Information) is handled according to local data protection laws.

📝 Legal & NDAs

    IP Ownership: All code and design assets in this repository are the property of the AgriChama project.

    Confidentiality: Team members are prohibited from sharing the proprietary TrustScore algorithm or member financial data with third parties.

⚡ Quick Start (Development)

    Clone the Repo:
    Bash

    git clone https://github.com/your-org/agrichama.git
    cd agrichama

    Environment Setup:

        Copy .env.example to .env in both backend/ and frontend/ folders.

    Spin up the Infrastructure:

Bash

    docker-compose up -d
    ```

4.  **Run Backend:**
    ```bash
    cd backend && go run cmd/api/main.go
    ```

5.  **Run Frontend:**
    ```bash
    cd frontend && npm run dev
    ```

---
**Maintained by:** AgriChama Dev Team  
**Last Updated:** May 2026
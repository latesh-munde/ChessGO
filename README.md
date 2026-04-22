# ♟️ Chess Server (Go)

A production-structured backend chess engine built in Go, designed with clean architecture, separation of concerns, and extensibility in mind.
Includes REST APIs, full chess rule validation, and an embedded frontend UI.

---

# 🚀 Features

- Full chess rule enforcement (legal moves, check, checkmate, stalemate)
- Support for special moves:
  - Castling
  - En-passant
  - Pawn promotion (extendable)

- RESTful API for game management
- Embedded frontend (HTML/CSS/JS)
- In-memory game storage (thread-safe)
- Clean layered architecture (API → Engine → Game → Core)

---

# 📁 Project Structure

```
chess-server/
│
├── cmd/server/main.go
│
├── internal/
│   ├── api/
│   ├── game/
│   ├── engine/
│   ├── pieces/
│   ├── core/
│   └── store/
│
├── go.mod
└── README.md
```

---

# 🧠 Architecture Overview

```
api → engine → game → core
        ↓
      pieces
```

### Responsibilities

- **API Layer** → Handles HTTP requests and responses
- **Engine Layer** → Contains all chess rules and validations
- **Game Layer** → Maintains game state and applies moves
- **Core Layer** → Basic primitives (board, square, color)
- **Pieces Layer** → Movement rules per piece
- **Store Layer** → Game persistence (in-memory)

✔ No cyclic dependencies
✔ Clear separation of concerns
✔ Easily extendable (DB, auth, multiplayer)

---

# 📦 Module Breakdown

## 🚀 cmd/server

- **main.go** → Entry point; starts server, initializes routes and store

---

## 🌐 internal/api

- **handlers.go** → Maps HTTP requests to game/engine logic
- **router.go** → Defines REST endpoints
- **middleware.go** → Logging and request middleware
- **ui.go** → Serves embedded frontend using Go `embed`

### 🎨 Frontend (`internal/api/web`)

- **index.html** → UI layout (board + controls)
- **style.css** → Chessboard and UI styling
- **app.js** → API calls, rendering, user interaction

---

## ♟️ internal/game (State Management)

- **game.go** → Game lifecycle and move execution
- **state.go** → Game state (board, turn, rights, etc.)
- **move.go** → Move parsing and representation
- **rules.go** → Applies moves (no validation)

---

## 🧠 internal/engine (Chess Logic)

- **attack_map.go** → Calculates attacked squares
- **check.go** → Detects check
- **simulate.go** → Simulates moves safely
- **validator.go** → Validates legal moves
- **enpassant.go** → En-passant logic
- **castling.go** → Castling validation
- **mate.go** → Checkmate and stalemate detection

---

## ♜ internal/pieces (Movement Rules)

- **piece.go** → Piece interface and types
- Individual files define movement rules:
  - king.go
  - queen.go
  - rook.go
  - bishop.go
  - knight.go
  - pawn.go

---

## 🧩 internal/core (Primitives)

- **color.go** → Player colors and helpers
- **square.go** → Board coordinates
- **board.go** → Board representation and utilities

---

## 💾 internal/store

- **store.go** → Storage interface
- **memory.go** → Thread-safe in-memory implementation

---

# 🔌 API Endpoints (Example)

```
POST   /games           → Create new game
GET    /games/{id}      → Get game state
POST   /games/{id}/move → Make a move
POST   /games/{id}/resign → Resign game
```

---

# ▶️ Running the Project

### 1. Clone Repository

```
git clone https://github.com/latesh-munde/ChessGO.git
cd chess-server
```

### 2. Run Server

```
go run cmd/server/main.go
```

### 3. Open in Browser

```
http://localhost:8080
```

---

# 🧪 Design Highlights

- Thread-safe game handling using mutex
- Pure rule engine (stateless validation logic)
- Clean separation of mutation vs validation
- Frontend decoupled from backend logic
- Easily replaceable storage layer

---

# 🚧 Future Improvements

- Database integration (PostgreSQL / Redis)
- WebSocket support for real-time multiplayer
- Authentication & user sessions
- Game history and replay system
- AI opponent integration
- Time controls (chess clock)

---

# 📌 Key Takeaway

This project is designed to demonstrate:

- Strong backend architecture
- Clean code organization
- Real-world system design principles
- Deep understanding of domain logic (chess rules)

---

# 👨‍💻 Author

Built as a portfolio-grade backend system to showcase production-level Go design and system thinking.

---

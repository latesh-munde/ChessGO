const PIECES = {
  King:   { 0: "♔", 1: "♚" },
  Queen: { 0: "♕", 1: "♛" },
  Rook:  { 0: "♖", 1: "♜" },
  Bishop:{ 0: "♗", 1: "♝" },
  Knight:{ 0: "♘", 1: "♞" },
  Pawn:  { 0: "♙", 1: "♟" }
};

let gameId = null;
let selected = null;
let legalSquares = [];

function sq(rank, file) {
  return String.fromCharCode(97 + file) + (rank + 1);
}

async function api(path, options = {}) {
  return fetch(path, options).then(r => r.json());
}

async function renderBoard() {
  const status = document.getElementById("status");

  // create game once
  if (!gameId) {
    const g = await api("/games", { method: "POST" });
    gameId = g.id;
  }

  const state = await api(`/games/${gameId}/state`);

  // ----- GAME END STATUS -----
  if (state.checkmate) {
    status.innerText = "CHECKMATE";
  } else if (state.stalemate) {
    status.innerText = "STALEMATE";
  } else if (state.inCheck) {
    status.innerText = "CHECK";
  } else {
    status.innerText = "Your move";
  }

  document.getElementById("whiteClock").innerText =
  formatTime(state.whiteTime);
document.getElementById("blackClock").innerText =
  formatTime(state.blackTime);

document.getElementById("whiteClock").classList.toggle(
  "active",
  state.turn === 0
);
document.getElementById("blackClock").classList.toggle(
  "active",
  state.turn === 1
);


  const board = document.getElementById("board");
  board.innerHTML = "";

  const last = state.lastMove || {};
  const fromLast = last.from;
  const toLast = last.to;

  for (let r = 7; r >= 0; r--) {
    for (let f = 0; f < 8; f++) {
      const div = document.createElement("div");
      div.className = "square " + ((r + f) % 2 === 0 ? "white" : "black");

      const square = sq(r, f);

      // highlight last move
      if (square === fromLast || square === toLast) {
        div.classList.add("last");
      }

      const piece = state.board[r][f];

      // highlight king in check
      if (
        state.inCheck &&
        piece &&
        piece.type === "King" &&
        piece.color === state.turn
      ) {
        div.classList.add("check");
      }

      // click handling
      div.onclick = async () => {
        if (selected && legalSquares.includes(square)) {
          await fetch(`/games/${gameId}/move`, {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({ from: selected, to: square })
          });

          selected = null;
          legalSquares = [];
          renderBoard();
          return;
        }

        if (!piece) return;

        selected = square;
        legalSquares = await api(`/games/${gameId}/legal?from=${square}`);
        renderBoard();
      };

      if (piece) {
        div.textContent = PIECES[piece.type][piece.color];
      }

      if (legalSquares.includes(square)) {
        div.classList.add("legal");
      }
      if (selected === square) {
        div.classList.add("selected");
      }

      board.appendChild(div);
    }
  }
}

function formatTime(ms) {
  const total = Math.max(0, Math.floor(ms / 1000));
  const min = Math.floor(total / 60);
  const sec = total % 60;
  return `${min}:${sec.toString().padStart(2, "0")}`;
}

renderBoard();
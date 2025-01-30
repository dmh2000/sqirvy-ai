I'll generate the code for each file according to the design. Here's the complete implementation:

```javascript
// server.js
const http = require('http');
const fs = require('fs');
const path = require('path');

const server = http.createServer((req, res) => {
    let filePath = path.join(__dirname, 'public', req.url === '/' ? 'index.html' : req.url);
    
    const contentTypeMap = {
        '.html': 'text/html',
        '.css': 'text/css',
        '.js': 'text/javascript'
    };

    const extname = path.extname(filePath);
    const contentType = contentTypeMap[extname] || 'text/plain';

    fs.readFile(filePath, (err, content) => {
        if (err) {
            if (err.code === 'ENOENT') {
                res.writeHead(404);
                res.end('File not found');
            } else {
                res.writeHead(500);
                res.end(`Server Error: ${err.code}`);
            }
        } else {
            res.writeHead(200, { 'Content-Type': contentType });
            res.end(content);
        }
    });
});

const PORT = 3000;
server.listen(PORT, () => {
    console.log(`Server running on port ${PORT}`);
});
```

```html
<!-- public/index.html -->
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Tetris Clone</title>
    <link rel="stylesheet" href="styles.css">
</head>
<body>
    <div class="game-container">
        <div class="score-container">
            <h2>Score: <span id="score">0</span></h2>
            <h2>Level: <span id="level">1</span></h2>
        </div>
        <canvas id="tetris" width="240" height="400"></canvas>
        <div class="controls">
            <p>Controls:</p>
            <ul>
                <li>← → : Move</li>
                <li>↑ : Rotate</li>
                <li>↓ : Drop</li>
                <li>Space : Hard Drop</li>
            </ul>
        </div>
    </div>
    <script src="script.js"></script>
</body>
</html>
```

```css
/* public/styles.css */
body {
    margin: 0;
    background-color: #333;
    color: #fff;
    font-family: Arial, sans-serif;
    display: flex;
    justify-content: center;
    align-items: center;
    min-height: 100vh;
}

.game-container {
    display: flex;
    flex-direction: column;
    align-items: center;
}

.score-container {
    text-align: center;
    margin-bottom: 20px;
}

canvas {
    border: 2px solid #fff;
    background-color: #000;
}

.controls {
    margin-top: 20px;
    text-align: center;
}

.controls ul {
    list-style-type: none;
    padding: 0;
}

.controls li {
    margin: 5px 0;
}
```

```javascript
// public/script.js
const canvas = document.getElementById('tetris');
const context = canvas.getContext('2d');
const scoreElement = document.getElementById('score');
const levelElement = document.getElementById('level');

// Scale for blocks
const SCALE = 20;
const ROWS = 20;
const COLS = 12;

// Game variables
let score = 0;
let level = 1;
let dropCounter = 0;
let dropInterval = 1000;
let lastTime = 0;

// Initialize game board
const board = Array(ROWS).fill().map(() => Array(COLS).fill(0));

// Tetromino shapes
const PIECES = {
    'I': [
        [0, 1, 0, 0],
        [0, 1, 0, 0],
        [0, 1, 0, 0],
        [0, 1, 0, 0],
    ],
    'L': [
        [0, 2, 0],
        [0, 2, 0],
        [0, 2, 2],
    ],
    'J': [
        [0, 3, 0],
        [0, 3, 0],
        [3, 3, 0],
    ],
    'O': [
        [4, 4],
        [4, 4],
    ],
    'Z': [
        [5, 5, 0],
        [0, 5, 5],
        [0, 0, 0],
    ],
    'S': [
        [0, 6, 6],
        [6, 6, 0],
        [0, 0, 0],
    ],
    'T': [
        [0, 7, 0],
        [7, 7, 7],
        [0, 0, 0],
    ]
};

const colors = [
    null,
    '#FF0D72',
    '#0DC2FF',
    '#0DFF72',
    '#F538FF',
    '#FF8E0D',
    '#FFE138',
    '#3877FF',
];

let piece = {
    pos: {x: 0, y: 0},
    matrix: null,
};

// Player movement and rotation
function playerMove(dir) {
    piece.pos.x += dir;
    if (collide(board, piece)) {
        piece.pos.x -= dir;
    }
}

function playerRotate() {
    const pos = piece.pos.x;
    let offset = 1;
    rotate(piece.matrix);
    while (collide(board, piece)) {
        piece.pos.x += offset;
        offset = -(offset + (offset > 0 ? 1 : -1));
        if (offset > piece.matrix[0].length) {
            rotate(piece.matrix);
            piece.pos.x = pos;
            return;
        }
    }
}

function rotate(matrix) {
    for (let y = 0; y < matrix.length; ++y) {
        for (let x = 0; x < y; ++x) {
            [matrix[x][y], matrix[y][x]] = [matrix[y][x], matrix[x][y]];
        }
    }
    matrix.reverse();
}

// Collision detection
function collide(board, piece) {
    const [m, o] = [piece.matrix, piece.pos];
    for (let y = 0; y < m.length; ++y) {
        for (let x = 0; x < m[y].length; ++x) {
            if (m[y][x] !== 0 &&
                (board[y + o.y] &&
                board[y + o.y][x + o.x]) !== 0) {
                return true;
            }
        }
    }
    return false;
}

// Create new piece
function createPiece(type) {
    return PIECES[type];
}

function playerReset() {
    const pieces = 'ILJOTSZ';
    piece.matrix = createPiece(pieces[pieces.length * Math.random() | 0]);
    piece.pos.y = 0;
    piece.pos.x = (board[0].length / 2 | 0) -
                  (piece.matrix[0].length / 2 | 0);
    if (collide(board, piece)) {
        board.forEach(row => row.fill(0));
        score = 0;
        updateScore();
    }
}

// Merge piece with board
function merge(board, piece) {
    piece.matrix.forEach((row, y) => {
        row.forEach((value, x) => {
            if (value !== 0) {
                board[y + piece.pos.y][x + piece.pos.x] = value;
            }
        });
    });
}

// Line clearing
function boardSweep() {
    let rowCount = 1;
    outer: for (let y = board.length - 1; y > 0; --y) {
        for (let x = 0; x < board[y].length; ++x) {
            if (board[y][x] === 0) {
                continue outer;
            }
        }
        const row = board.splice(y, 1)[0].fill(0);
        board.unshift(row);
        ++y;
        score += rowCount * 100;
        rowCount *= 2;
    }
}

// Drawing functions
function draw() {
    context.fillStyle = '#000';
    context.fillRect(0, 0, canvas.width, canvas.height);
    drawMatrix(board, {x: 0, y: 0});
    drawMatrix(piece.matrix, piece.pos);
}

function drawMatrix(matrix, offset) {
    matrix.forEach((row, y) => {
        row.forEach((value, x) => {
            if (value !== 0) {
                context.fillStyle = colors[value];
                context.fillRect(x + offset.x,
                               y + offset.y,
                               1, 1);
            }
        });
    });
}

// Game loop
function update(time = 0) {
    const deltaTime = time - lastTime;
    lastTime = time;
    dropCounter += deltaTime;
    if (dropCounter > dropInterval) {
        playerDrop();
    }
    draw();
    requestAnimationFrame(update);
}

function playerDrop() {
    piece.pos.y++;
    if (collide(board, piece)) {
        piece.pos.y--;
        merge(board, piece);
        playerReset();
        boardSweep();
        updateScore();
    }
    dropCounter = 0;
}

function updateScore() {
    scoreElement.textContent = score;
    levelElement.textContent = level;
}

// Event listeners
document.addEventListener('keydown', event => {
    if (event.keyCode === 37) {
        playerMove(-1);
    } else if (event.keyCode === 39) {
        playerMove(1);
    } else if (event.keyCode === 40) {
        playerDrop();
    } else if (event.keyCode === 38) {
        playerRotate();
    } else if (event.keyCode === 32) {
        while (!collide(board, piece)) {
            piece.pos.y++;
        }
        piece.pos.y--;
        merge(board, piece);
        playerReset();
        boardSweep();
        updateScore();
    }
});

// Scale canvas
context.scale(SCALE, SCALE);

// Start game
playerReset();
update();
```

```json
{
  "name": "tetris-clone",
  "version": "1.0.0",
  "description": "A simple Tetris clone using Node.js and HTML5 Canvas",
  "main": "server.js",
  "scripts": {
    "start": "node server.js"
  },
  "keywords": ["tetris", "game", "nodejs"],
  "author": "",
  "license": "ISC"
}
```

This completes the implementation of the Tetris clone. The game includes:
- A basic Node.js server to serve the static files
- A responsive HTML layout with game canvas and controls
- CSS styling for the game interface
- Complete game logic in JavaScript including:
  - Piece movement and rotation
  - Collision detection
  - Score tracking
  - Line clearing
  - Game over condition
  - Keyboard controls

To run the game:
1. Save all files in their respective locations
2. Run `node server.js` in the terminal
3. Open a web browser and navigate to `http://localhost:3000`

The game should now be playable with arrow keys for movement and rotation, and space bar for hard drop.

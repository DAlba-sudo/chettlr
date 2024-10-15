function createSquare(piece_switch, board, num) {
    let square = document.createElement('div');

    square.style.background = (num % 2 == 0) ? "#00cc66" : "#f2f2f2";
    square.style.padding = "10px";
    square.style.textAlign = "center";
    square.style.margin = 0;

    square.style.fontSize = "200%";
    switch (piece_switch) {
        case 'r':
            square.innerHTML = "&#9820";
            break;
        case 'R':
            square.innerHTML = "&#9814";
            break;
        case 'p':
            square.innerHTML = "&#9823";
            break;
        case 'P':
            square.innerHTML = "&#9817";
            break;
        case 'n':
            square.innerHTML = "&#9822";
            break;
        case 'N':
            square.innerHTML = "&#9816";
            break;
        case 'b':
            square.innerHTML = "&#9821";
            break;
        case 'B':
            square.innerHTML = "&#9815";
            break;
        case 'k':
            square.innerHTML = "&#9818";
            break;
        case 'K':
            square.innerHTML = "&#9812";
            break;
        case 'q':
            square.innerHTML = "&#9819";
            break;
        case 'Q':
            square.innerHTML = "&#9813";
            break;
        default:
            square.innerHTML = " ";
            break;
    }

    board.append(square);
}

function renderBoards() {
    let boards = document.querySelectorAll("chess-board");
    for (const board of boards) {
        board.style.display = "grid";
        board.style.gridTemplateColumns = "1fr 1fr 1fr 1fr 1fr 1fr 1fr 1fr";
        board.style.gridTemplateRows = "1fr 1fr 1fr 1fr 1fr 1fr 1fr 1fr";
        board.style.gap = "0px 0px";
        board.style.maxWidth = "min-content";

        let position = board.getAttribute("fen");
        if (position == null) {
            continue;
        }
        
        // parse the position and place items their respective square
        // according to the FEN notation
        let square_number = 1;
        for (const char of position) {
            if (char == '/') {
                square_number ++;
                continue;
            } else if (char == '-' || char == ' ') {
                break;
            }

            let number = parseInt(char, 10)
            console.log("Number: " + number);
            if (isNaN(number)) {
                createSquare(char, board, square_number);
                square_number++;
            } else {
                for (let i = 0; i < number; i++) {
                    createSquare(char, board, square_number);
                    square_number++;
                }
            }
        }
    }
}

function setupCarousels() {
    let carousels = document.querySelectorAll("chess-carousel");
    for (const carousel of carousels) {
        let boards = carousel.querySelectorAll("chess-board");
        
        // configures left button to do the proper carousel action
        let left = carousel.querySelector(".carousel-left");
        left.addEventListener("click", (ev) => {
            let num = parseInt(left.parentElement.parentElement.getAttribute("step")) - 1;
            if (num == 0) {
                return;
            }
            boards[num].style.display = "none";
            boards[(num - 1) % boards.length].style.display = "grid";
            left.parentElement.parentElement.setAttribute("step", num);
        });

        // configures right button to do the proper carousel action
        let right = carousel.querySelector(".carousel-right");
        right.addEventListener("click", (ev) => {
            let num = parseInt(right.parentElement.parentElement.getAttribute("step")) - 1;
            if (num == boards.length - 1) {
                return;
            }

            boards[num].style.display = "none";
            boards[(num + 1) % boards.length].style.display = "grid";
            right.parentElement.parentElement.setAttribute("step", `${num + 2}`);
        });

        for (let i = 0; i < boards.length; i++) {
            if (i == 0) {
                continue;
            }
            let board = boards[i];
            board.style.display = "none";
        }
    }
}

function main() {
    renderBoards();
    setupCarousels();
}

main();
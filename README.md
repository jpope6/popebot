# popebot

![Logo](popebot.png)

## Overview
Popebot is a UCI compatible chess engine written entirely in Golang. This project is a work in progress and is not yet completed, but if you would like to do play it you can do so [here](https://lichess.org/@/popebot).

Note: If the bot is not online and you want to play it, send me a message!

## Board Representation

The board is represented using a Bitboard approach. Each square on the board is represented by a bit in a 64-bit unsigned integer. Two bitboards are used to represent the position for pieces of each color. Twelve bitboards represent different piece types and colors (pawn, knight, bishop, rook, queen, and king). These bitboards are encapsulated in a Position struct, allowing for quick and efficient access to manipulate game positions. This approach enables efficient move generation and minimizes memory usage within the application.

## Move Generation

Move generation is performed during initialization and stored for future lookup. The program iterates through each possible piece and square on the board, storing legal moves for each piece. For bishops, rooks, and queens, all possible moves for every possible occupancy are generated and stored in lookup tables. During search, the program checks if the moves for each piece and its current square are legal and adds them to a move list array for searching.

## Evaluation

Each piece is assigned a predefined score reflecting its value and importance in the game. To assess the positional strength of each piece, the program utilizes piece square tables. These tables assign a score to each square on the board for each piece type, categorized into middle game and end game tables. This allows the program to prioritize different strategies based on the game state.

## Search

The search algorithm is the core of the chessbot's decision-making process. The negamax algorithm with alpha-beta pruning is employed, along with the MVV/LVA (Most Valuable Victim/Least Valuable Attacker) heuristic.

Negamax is a variant of the minimax algorithm, using recursive tree traversal to explore all possible moves and evaluate every position. Unlike minimax, negamax does not switch between minimum and maximum scores; it seeks the best score on each turn. Alpha-beta pruning is used to reduce the number of nodes traversed by tracking alpha and beta values and pruning branches that will not affect the final result.

MVV/LVA prioritizes moves that capture the most valuable piece with the least valuable piece. This heuristic enhances alpha-beta pruning by sorting moves using the quicksort algorithm, significantly reducing the number of nodes searched in many positions.

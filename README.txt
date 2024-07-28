TIC-TAC-TOE game, using Golang (specifically gin package(library) for BE features) and Postman.

The game runs on this module/EXE (which operates as the BE server) and the UI is done with Postman GET requests.

GAME FLOW
---------------------------------------------------------------------------------------------------------------
To run the server, you can add an argument of specific port number to run with, otherwise default port is 3333.

Once the server is up, start game by doing the following HTTP GET request in Postman: URL_PATH/start. For example:
http://localhost:3333/start

This will initiate the game grid, as well as letting either X or O player to begin.
The grid spots is set by numbers (from 1 to 9). Each number represents a different cell in the 3x3 grid.
The starting grid will appear like this:
1     2     3
4     5     6
7     8     9

After the game has been initialized, begin game with the following HTTP GET request in Postman: 
URL_PATH/player-type?mark=number(1,9).
Examples:
For player X wanting to spot cell 3, the HTTP GET request will be: http://localhost:3333/player-x?mark=3
For player O wanting to spot cell 8, the HTTP GET request will be: http://localhost:3333/player-o?mark=8
This will change the grid, marking the spot desired with "X" or "O".

X player example:
1     2     X
4     5     6
7     8     9

O player example:
1     2     3
4     5     6
7     O     9


The game ends if:
1. One of the players (X or O) manages to fill one of the 8 win rows.
2. Grid is full of X and O marks (with no numbers), indicating a draw.


Unauthorized requests:
1. A value in the "mark" query that is not a number or is a number but is higher than 9 or lower than 1.
2. One of the players tries to take a turn twice (each player only gets 1 turn).
3. A player tries to mark an already marked cell.

If one of the above unauthorized requests occurs, an appropriate error will be given as a response.
In addition, the turn won't pass to the next player.


At any point of wishing to restart the game, simply make the "start" HTTP GET request mentioned above.


CODE-RELATED
--------------------------------------------------------------------------------------------------------------
1. The grid is represented as a slice (array) of string, and each cell holds a value between 1-9.
2. There are 4 global variables which indicate game conditions (GameOver status, X/O player turn and the grid)
3. Besides the gin package (which handles the BE capabilities), I used 2 other packages:
    a. "os" package, to check runline arguments.
    b. "strconv" package, which is used to convert strings to int and vice-versa for various conditions.
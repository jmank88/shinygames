# Lightsout

A simple Lightsout game using the shiny gui package (golang.org/x/exp/shiny).

Loosely ported from TI-68K C source found here: http://www.ticalc.org/archives/files/fileinfo/318/31860.html

This is a WIP.

Includes 50 levels. The command line flag 'level' may be used to begin at a level other than 1.


# How to play

Each level begins with a different configuration of lights 'on' (white). The goal is to turn all lights 'off' (black).

Clicking any light on the board will toggle its state, and its vertical and horizontal neighbor's states.

![Toggle gif](/gifs/toggle.gif)

A level is completed when all lights are turned off, and the game will advance to the next level.

![Next gif](/gifs/next.gif)


# TODO

- [ ] reduce flicker/delay on resize
- [ ] random levels
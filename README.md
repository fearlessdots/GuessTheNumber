# GuessTheNumber

A simple command-line game written in Go where the player tries to guess a randomly generated number within a certain range.

## Getting Started

To play the game, run this:

```shell
git clone https://github.com/your-username/GuessTheNumber.git
cd GuessTheNumber
make build
./GuessTheNumber
```

# Gameplay

## Number of chances

The number of chances in the game refers to the maximum number of attempts the player has to guess the correct number before the game is over. By default, the game allows the player five chances to guess the number correctly. However, the game allows the player to choose a custom number of chances at the start of the game. The game will provide feedback after each guess, letting the player know if the guess was too high or too low, and how many chances are left.

It's worth noting that the number of chances can affect the difficulty of the game. For instance, in the case of the Medium and Hard levels, it may be advisable to increase the maximum number of chances to give the player a better chance of guessing the correct number. Conversely, reducing the number of chances can make the game more challenging and increase the level of difficulty. Ultimately, the number of chances is up to the player to choose based on their preferred level of challenge.

## Game levels

The game has three default levels: Easy, Medium, and Hard. The maximum number for each level is set to 15, 25, and 35, respectively.

- Easy: The maximum number is 15, which means you have a smaller range of numbers to choose from. This level is great for beginners or for those who want a quick and easy game.
- Medium: The maximum number is 25, which gives you a wider range of numbers to choose from than the easy level. This level is perfect for those who have played the game before and want a bit of a challenge.
- Hard: The maximum number is 35, which means you have an even wider range of numbers to choose from than the medium level. This level is the most challenging of the default levels and is perfect for those who are experienced players and want a real challenge.

You can always choose the custom option to set a maximum number of your choice.

> In the case of the Medium and Hard levels, I recommend increasing the maximum number of chances to a higher value.

## Steps

1. Upon starting the game, the user is presented with two menus:
  - The first one is to choose the maximum number of chances. If the user wants so, a custom number can be inserted. The default value is `5`.
  - The second one is to select a game level or choose `Custom` to set a custom maximum number.
2. Guess a number within the range of 1 and the maximum number.
3. After each guess, the game will tell you if your guess was too high or too low and show how many chances you have left.
4. If you guess the number correctly, you win! The game will ask if you want to play again.
5. If you don't guess the number correctly within the maximum number of chances, you lose. The game will ask if you want to play again.

## License
GuessTheNumber is licensed under the GPL-3.0 license.

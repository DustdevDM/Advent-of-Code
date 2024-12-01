using AdventOfCode_Day2.src;

Console.ForegroundColor = ConsoleColor.Green;
Console.WriteLine("Advent of Code 2023 Day 2\n");
Console.ResetColor();

string input = File.ReadAllText("./Input.txt");

List<string> inputLines  = input.Split(new char[] { '\n' }, StringSplitOptions.RemoveEmptyEntries).ToList();
List<Game> games = new List<Game>();

inputLines.ForEach(line => games.Add(new Game(line)));
int resultValue = 0;
int powerValue = 0;

foreach (Game game in games)
{
    Console.ForegroundColor = game.IsPossible ? ConsoleColor.Green : ConsoleColor.Red;
    Console.Write($"Game {game.GameId}: {game.IsPossible}; ");
    Console.ForegroundColor = ConsoleColor.Blue;
    Console.Write($"Power: {game.Power}\n");
    Console.ResetColor();
    if (game.IsPossible)
        resultValue += game.GameId;

    powerValue += game.Power;
}


Console.ForegroundColor = ConsoleColor.Yellow;
Console.WriteLine("Result First Star: " + resultValue);
Console.ForegroundColor = ConsoleColor.Blue;
Console.WriteLine("Result Second Star: " + powerValue);
Console.ResetColor();

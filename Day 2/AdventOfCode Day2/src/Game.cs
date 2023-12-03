using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace AdventOfCode_Day2.src
{
    public class Game
    {
        public int GameId { get; set; }

        public bool IsPossible { get; set; }

        public List<CubeReveal> CubeReveals { get; set;}

        public int Power { get; set; }

        public Game(string inputLine)
        {
            this.gameLineParser(inputLine);
        }

        private void gameLineParser(string inputLine)
        {
            this.GameId = extractGameId(inputLine);
            this.CubeReveals = extractCubeReveals(inputLine);
            SetGamePossibleValue();
        } 

        private int extractGameId(string inputLine)
        {
            return int.Parse(inputLine.Split(": ").First().Replace("Game ", string.Empty));
        }

        private List<CubeReveal> extractCubeReveals(string inputLine)
        {
            List<CubeReveal> cubeReveals = new List<CubeReveal>();
            foreach (var extractetLine in inputLine.Split(": ").Last().Split("; "))
            {
                cubeReveals.Add(new CubeReveal(extractetLine));
            };
            return cubeReveals;

        }

        private void SetGamePossibleValue()
        {
            int highestRedCubes = 0;
            int highestGreenCubes = 0;
            int highestBlueCubes = 0;

            foreach (CubeReveal cubeReveal in CubeReveals)
            {
                highestRedCubes = cubeReveal.RedCubes > highestRedCubes ? cubeReveal.RedCubes : highestRedCubes;
                highestGreenCubes = cubeReveal.GreenCubes > highestGreenCubes ? cubeReveal.GreenCubes : highestGreenCubes;
                highestBlueCubes = cubeReveal.BlueCubes > highestBlueCubes ? cubeReveal.BlueCubes : highestBlueCubes;
            }

            this.IsPossible = true;

            if (highestRedCubes > 12)
                this.IsPossible = false;

            if (highestGreenCubes > 13)
                this.IsPossible = false;

            if (highestBlueCubes > 14)
                this.IsPossible = false;

            this.Power = highestRedCubes * highestGreenCubes * highestBlueCubes;
        }
    }
}

namespace AdventOfCode_Day2.src
{
    public class CubeReveal
    {
        public int RedCubes { get; set; }
        public int GreenCubes { get; set; }
        public int BlueCubes { get; set; }

        public CubeReveal(string revealInput)
        {
            this.RedCubes = 0;
            this.GreenCubes = 0;
            this.BlueCubes = 0;

            foreach (var revealSet in revealInput.Split(", "))
            {
                extractAndAddRevealSet(revealSet.Replace("\r", ""));
            };
        }

        private void extractAndAddRevealSet(string revealSetInput)
        {
            int cubeCount = int.Parse(revealSetInput.Split(" ").First());
            string cubeColor = revealSetInput.Split(" ").Last();

            switch (cubeColor)
            {
                case "red":
                    this.RedCubes += cubeCount;
                    break;
                case "green":
                    this.GreenCubes += cubeCount;
                    break;
                case "blue":
                    this.BlueCubes += cubeCount;
                    break;
                default:
                    throw new Exception("Unable to determine Cubecolor");
            }
        }
    }
}

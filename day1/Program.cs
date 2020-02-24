using System;
using System.IO;

namespace day1
{
    class Program
    {
        static void Main(string[] args)
        {
            using (StreamReader reader = new StreamReader("part-1-input.txt"))
            {
                string input;
                int weightSum = 0;

                while ((input = reader.ReadLine()) != null)
                {
                    int weight = int.Parse(input);

                    weightSum += (weight / 3) - 2;
                }

                Console.WriteLine("Part 1 Fuel requirements: {0}", weightSum);
            }

            using (StreamReader reader = new StreamReader("part-2-input.txt"))
            {
                string input;
                int totalWeight = 0;

                while ((input = reader.ReadLine()) != null)
                {
                    int weight = (int.Parse(input) / 3) - 2;

                    while (weight > 0)
                    {
                        totalWeight += weight;
                        weight = (weight / 3) - 2;
                    }
                }

                Console.WriteLine("Part 2 Fuel requirements: {0}", totalWeight);
                Console.ReadLine();
            }
        }
    }
}

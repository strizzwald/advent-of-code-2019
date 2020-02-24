using System;
using System.IO;

namespace day1
{
    class Program
    {
        static void Main(string[] args)
        {
            using (StreamReader reader = new StreamReader("input.txt"))
            {
                string input;
                int weightSum = 0;

                while ((input = reader.ReadLine()) != null)
                {
                    int weight = int.Parse(input);

                    weightSum += (weight / 3) - 2;
                }

                Console.Write("Fuel requirements: {0}", weightSum);
                Console.ReadLine();
            }
        }
    }
}

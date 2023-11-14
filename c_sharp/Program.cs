namespace MyApp
{
    internal class Program
    {
        static void Main(string[] args)
        {
            var readPath = args[0];
            Console.WriteLine(day_three.day_three.rucksack_compartments(readPath));
        }
    }
}

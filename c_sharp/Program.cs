namespace MyApp
{
    internal class Program
    {
        static void Main(string[] args)
        {
            var readPath = args[0];
            Console.WriteLine(day_two.day_two.RockPaperScissors(readPath));
        }
    }
}

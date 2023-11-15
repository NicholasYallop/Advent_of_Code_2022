namespace MyApp
{
    internal class Program
    {
        static void Main(string[] args)
        {
            var readPath = args[0];
            Console.WriteLine(day_four.day_four.overlapping_assignments(readPath));
        }
    }
}

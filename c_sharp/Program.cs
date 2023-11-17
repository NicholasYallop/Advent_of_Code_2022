namespace MyApp
{
    internal class Program
    {
        static void Main(string[] args)
        {
            var readPath = args[0];
            Console.WriteLine(day_five.day_five.crate_arrangement(readPath));
        }
    }
}

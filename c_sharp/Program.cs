namespace MyApp
{
    internal class Program
    {
        static void Main(string[] args)
        {
            var readPath = args[0];

            var lines = File.ReadLines(readPath);
            
            var largest_elf_bag = 0;
            var elf_bag_buffer = 0;
            foreach (var line in lines){
                if (line == ""){
                    largest_elf_bag = Math.Max(largest_elf_bag, elf_bag_buffer);
                    elf_bag_buffer = 0;
                    continue;
                }
                if (int.TryParse(line, out var fooditem)){
                        elf_bag_buffer += fooditem;
                }
            }
            largest_elf_bag = Math.Max(largest_elf_bag, elf_bag_buffer);

            Console.WriteLine(largest_elf_bag);
        }
    }
}

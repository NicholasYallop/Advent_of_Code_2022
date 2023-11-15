namespace day_four{
    public static class day_four{
        public static int overlapping_assignments(string readPath){
            var result = 0;
            var lines = File.ReadLines(readPath);

            foreach (var line in lines){
                var eleves = line.Split(',');

                var elf_one = eleves[0].Split('-');
                var start_one = int.Parse(elf_one[0]);
                var end_one = int.Parse(elf_one[1]);

                var elf_two = eleves[1].Split('-');
                var start_two = int.Parse(elf_two[0]);
                var end_two = int.Parse(elf_two[1]);

                if (start_one == start_two) result++;
                else if (start_one > start_two) { if (end_two >= start_one) result++; }
                else if (end_one >= start_two) result++;
            }

            return result;
        }
    }
}

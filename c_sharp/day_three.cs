namespace day_three{
    public static class day_three{
        public static int priority(char c){
            int i = c - 'a' + 1;
            return i > 0 ? i : c - 'A' + 27;
        }

        public static int rucksack_compartments(string readPath){
            var result = 0;
            var lines = File.ReadLines(readPath);

            List<char> chars = new();
            var wrapped_elf_count = -1;
            foreach(var line in lines){
                wrapped_elf_count = (wrapped_elf_count+1)%3;
                switch (wrapped_elf_count){
                    case 0:
                        chars = line.ToList();
                        break;
                    case 1:
                        chars = chars.Where(c => line.Contains(c)).ToList();
                        break;
                    case 2:
                        chars = chars.Where(c => line.Contains(c)).ToList();
                        result += priority(chars.First());
                        break;
                }
            }

            return result;
        }
    }
}

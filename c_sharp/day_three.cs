namespace day_three{
    public static class day_three{
        public static int priority(char c){
            int i = c - 'a' + 1;
            return i > 0 ? i : c - 'A' + 27;
        }

        public static int rucksack_compartments(string readPath){
            var result = 0;
            var lines = File.ReadLines(readPath);

            foreach(var line in lines){
                var length = line.Length;

                var compartment_one = line.Take(length/2);
                var compartment_two = line.Skip(length/2).Take(length/2);

                List<char> tested = new();
                foreach(var char_one in compartment_one){
                    if (!tested.Contains(char_one)){
                        tested.Add(char_one);

                        foreach(var char_two in compartment_two){
                            if (char_one==char_two){
                                result += priority(char_one);
                                break;
                            }
                        }
                    }
                }
            }

            return result;
        }
    }
}

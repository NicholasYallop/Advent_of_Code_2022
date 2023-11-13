namespace day_two{
    public static class day_two{
        private enum their_choices{A, B, C}
        private enum my_choices{X, Y, Z}

        public static int RockPaperScissors(string readPath){
            var result = 0;
            var lines = File.ReadLines(readPath);

            foreach(var line in lines){
                var choices = line.Split(" ");
                var their_choice = (int)Enum.Parse(typeof(their_choices), choices[0]);
                var my_choice = (int)Enum.Parse(typeof(my_choices), choices[1]);
                
                int result_score;
                if (my_choice == their_choice){
                    // draw
                    result_score = 3;
                }else if ((their_choice+1)%3 == my_choice){
                    // I win
                    result_score = 6;
                }else{
                    // I lose
                    result_score = 0;
                }

                result += result_score + my_choice + 1;
            }

            return result;
        }
    }
}

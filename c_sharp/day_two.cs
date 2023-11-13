namespace day_two{
    public static class day_two{
        private enum TheirChoices{A, B, C}
        private enum Outcomes{X=0, Y=3, Z=6}

        public static int RockPaperScissors(string readPath){
            var result = 0;
            var lines = File.ReadLines(readPath);

            foreach(var line in lines){
                var choices = line.Split(" ");
                var their_choice = (int)Enum.Parse(typeof(TheirChoices), choices[0]);
                var desired_outcome = Enum.Parse(typeof(Outcomes), choices[1]);
                
                int my_choice;
                switch (desired_outcome){
                    case Outcomes.X:
                        my_choice = (their_choice+2)%3;
                        break;
                    case Outcomes.Y:
                        my_choice = their_choice;
                        break;
                    case Outcomes.Z:
                        my_choice = (their_choice+1)%3;
                        break;
                    default:
                        my_choice = 0;
                        break;
                }

                result += (int)desired_outcome + my_choice + 1;
            }

            return result;
        }
    }
}

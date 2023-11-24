namespace day_six
{
    public static class day_six
    {
        public static int Repeating_Characters(string input)
        {
            var char_buffer = "";
            var index = 0;

            foreach (var input_char in input){
                index ++;
                if (char_buffer.Contains(input_char)){
                        char_buffer = char_buffer.Substring(
                                char_buffer.IndexOf(input_char)+1
                                );

                        char_buffer = char_buffer + input_char;
                }else if (char_buffer.Length == 13){
                    return index; 
                }else{
                    char_buffer = char_buffer + input_char;
                }
            }
            
            return 0;
        }
    }
}

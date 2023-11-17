namespace day_five{
    public static class day_five{
        public static string crate_arrangement(string readPath){
            var lines = File.ReadLines(readPath);
            var finding_start = true;

            List<string> start_line_buffer = new();


             var crates = new List<List<char>>();
            foreach (var line in lines){
                if (line == ""){
                    var columns = start_line_buffer.Last().Split(" ").Where(l => l.Length>0).Count();
                    while (crates.Count < columns){
                        crates.Add(new List<char>());
                    }

                    for(var i=start_line_buffer.Count-2 ; i>=0 ; i--){
                        for (var j=1 ; j<start_line_buffer[i].Length ; j++){
                            if (j%4==1 && start_line_buffer[i][j] != ' '){
                                crates[j/4].Add(start_line_buffer[i][j]);
                            }
                        }
                    }

                    finding_start = false;
                    continue;
                }

                if (finding_start){
                    start_line_buffer.Add(line);
                    continue;
                }

                var command = line.Split(' ');
                if (!(
                        int.TryParse(command[1], out var count)
                        && int.TryParse(command[3], out var from_column)
                        && int.TryParse(command[5], out var to_column)
                    )){
                    throw new InvalidOperationException();
                }

                var from = from_column-1;
                var to = to_column-1;
                
                for (var i=0 ; i<count ; i++){
                    crates[to] = crates[to].Append(crates[from].Last()).ToList();
                    crates[from].RemoveAt(crates[from].Count-1);
                }
            }

            return crates.Aggregate("", (agg, list) =>
                        agg + list.Last()
                    );
        }

        public static void Print(this List<List<char>> list){
            var count = 0;
            list.ForEach(list =>{
                    count++;
                    Console.WriteLine( count + 
                            list.Aggregate("", (agg, c) => agg + " " + c)
                            );
            });
            Console.WriteLine("");
        }
    }
}



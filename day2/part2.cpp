#include <iostream>
#include <fstream>
#include <regex>
#include <vector>
#include <string>

using namespace std;

// g++ -std=c++11  -o main part2.cpp && ./main

vector<string> split(const string str, const string regex_str)
{
    regex regexz(regex_str);
    return {sregex_token_iterator(str.begin(), str.end(), regexz, -1), sregex_token_iterator()};
}

int main()
{
    ifstream file("input.txt");
    int sum = 0;
    if (file.is_open())
    {
        string line;
        while (getline(file, line))
        {
            // split line by : + space
            vector<string> parts = split(line, ":\\s");

            // split first part by space
            vector<string> game = split(parts[0], "\\s");
            int game_num = stoi(game[1]);
            int red = 0, green = 0, blue = 0;

            // split second part by ; + space
            vector<string> sets = split(parts[1], ";\\s");
            for (int i = 0; i < sets.size(); i++)
            {
                // split each set by , + space
                vector<string> set = split(sets[i], ",\\s");
                for (int j = 0; j < set.size(); j++)
                {
                    // split count and cube by space
                    vector<string> cubes = split(set[j], "\\s");
                    int count = stoi(cubes[0]);
                    if (cubes[1] == "red" && count > red)
                        red = count;
                    else if (cubes[1] == "green" && count > green)
                        green = count;
                    else if (cubes[1] == "blue" && count > blue)
                        blue = count;
                }
            }

            int power = red * green * blue;
            sum += power;
        }
    }

    cout << "sum: " << sum << endl;
    return 0;
}
#include <iostream>
#include <fstream>
#include <regex>
#include <vector>
#include <string>

using namespace std;

// g++ -std=c++11  -o main part1.cpp && ./main

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

            // split second part by ; + space
            vector<string> sets = split(parts[1], ";\\s");
            int error = 0;
            for (int i = 0; i < sets.size(); i++)
            {
                // split each set by , + space
                vector<string> set = split(sets[i], ",\\s");
                for (int j = 0; j < set.size(); j++)
                {
                    // split count and cube by space
                    vector<string> cubes = split(set[j], "\\s");
                    int count = stoi(cubes[0]);
                    if (cubes[1] == "red" && count > 12)
                    {
                        error = 1;
                        break;
                    }
                    else if (cubes[1] == "green" && count > 13)
                    {
                        error = 1;
                        break;
                    }
                    else if (cubes[1] == "blue" && count > 14)
                    {
                        error = 1;
                        break;
                    }
                }

                if (error)
                    break;
            }

            if (!error)
                sum += game_num;
        }
    }

    cout << "sum: " << sum << endl;
    return 0;
}
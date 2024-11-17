#include <iostream>
#include <fstream>
#include <string>
#include <vector>
#include <regex>

using namespace std;

// g++ -std=c++11 -o main part2.cpp && ./main

vector<string> split(const string str, const string regex_str)
{
    regex regexz(regex_str);
    return {sregex_token_iterator(str.begin(), str.end(), regexz, -1), sregex_token_iterator()};
}

int main()
{
    ifstream file("input.txt");
    int sum = 0;
    vector<int> bucket;
    if (file.is_open())
    {

        string line;
        while (getline(file, line))
        {
            vector<string> data = split(line, ":\\s+");
            int card_num = stoi(split(data[0], "\\s+")[1]) - 1;
            vector<string> parts = split(data[1], "\\s+\\|\\s+");
            vector<string> part1 = split(parts[0], "\\s+");
            vector<string> part2 = split(parts[1], "\\s+");

            // convert part1 & part2 to int vectors and sort them
            vector<int> p1, p2;
            for (int i = 0; i < part1.size(); i++)
            {
                p1.push_back(stoi(part1[i]));
            }

            for (int i = 0; i < part2.size(); i++)
            {
                p2.push_back(stoi(part2[i]));
            }

            sort(p1.begin(), p1.end());
            sort(p2.begin(), p2.end());

            // check how many elements in p2 are in p1
            // assume no duplicates
            int count = 0, i = 0, j = 0;
            while (i <= p1.size() - 1 && j <= p2.size() - 1)
            {
                if (p1[i] == p2[j])
                {
                    count++;
                    i++;
                    j++;
                }
                else if (p1[i] < p2[j])
                {
                    i++;
                }
                else
                {
                    j++;
                }
            }

            // allocate enough space for bucket
            if (card_num + count + 1 > bucket.size())
            {
                vector<int> space(card_num + count + 1 - bucket.size(), 1);
                bucket.insert(end(bucket), begin(space), end(space));
            }

            // push count of copies into bucket
            for (int i = card_num + 1; i < card_num + count + 1; i++)
            {
                bucket[i] += bucket[card_num];
            }
        }

        for (int i = 0; i < bucket.size(); i++)
            sum += bucket[i];
        cout << "sum: " << sum << endl;
    }

    return 0;
}
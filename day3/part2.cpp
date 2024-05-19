#include <iostream>
#include <fstream>
#include <string>
#include <cctype>
#include <vector>
#include <map>

using namespace std;

// g++ -std=c++11 -o main part2.cpp && ./main

vector<int> rowEnum = {0, 0, -1, 1, -1, -1, 1, 1};
vector<int> colEnum = {-1, 1, 0, 0, -1, 1, -1, 1};

int main()
{
    ifstream file("input.txt");
    int sum = 0;

    if (file.is_open())
    {
        string line;
        vector<string> docs;
        map<string, pair<int, int>> geared;

        // read all lines into a vector
        while (getline(file, line))
        {
            docs.push_back(line);
        }

        // iterate over each line to check for valid numbers
        for (int i = 0; i < docs.size(); i++)
        {
            string num_str = "";
            bool ok = false;
            string star_pos;
            for (int j = 0; j < docs[i].size(); j++)
            {
                if (isdigit(docs[i][j]) && ok)
                {
                    num_str += docs[i][j];
                }
                else if (isdigit(docs[i][j]) && !ok)
                {
                    num_str += docs[i][j];
                    for (int k = 0; k < 8; k++)
                    {
                        int newX = i + rowEnum[k];
                        int newY = j + colEnum[k];
                        if (newX >= 0 && newX < docs.size() && newY >= 0 && newY < docs[i].size() && docs[newX][newY] == '*')
                        {
                            star_pos = to_string(newX) + "," + to_string(newY); // store the position of the *
                            ok = true;
                            break;
                        }
                    }
                }

                // check if we have a valid number
                if (!isdigit(docs[i][j]) || j == docs[i].size() - 1) // make sure to handle end of line
                {
                    if (ok)
                    {
                        if (geared.find(star_pos) == geared.end())
                            geared[star_pos] = make_pair(stoi(num_str), -1);
                        else if (geared[star_pos].second != -1) // * has more than 2 adjacent part numbers
                            geared[star_pos].first = -2;
                        else // found the second part number for *
                            geared[star_pos].second = stoi(num_str);
                    }

                    num_str = "";
                    ok = false;
                }
            }
        }

        for (auto it = geared.begin(); it != geared.end(); it++)
        {
            if (it->second.first != -2 && it->second.second != -1)
                sum += it->second.first * it->second.second;
        }

        cout << "sum: " << sum << endl;
    }

    return 0;
}
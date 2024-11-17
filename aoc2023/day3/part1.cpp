#include <iostream>
#include <fstream>
#include <string>
#include <cctype>
#include <vector>

using namespace std;

// g++ -std=c++11 -o main part1.cpp && ./main

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
                        if (newX >= 0 && newX < docs.size() && newY >= 0 && newY < docs[i].size() && docs[newX][newY] != '.' && !isdigit(docs[newX][newY]))
                        {
                            ok = true;
                            break;
                        }
                    }
                }

                // check if we have a valid number
                if (!isdigit(docs[i][j]) || j == docs[i].size() - 1) // make sure to handle end of line
                {
                    if (ok)
                        sum += stoi(num_str);
                    num_str = "";
                    ok = false;
                }
            }
        }

        cout << "sum: " << sum << endl;
    }

    return 0;
}
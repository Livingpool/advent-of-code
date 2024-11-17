#include <iostream>
#include <string>
#include <fstream>
#include <cctype>
#include <vector>

using namespace std;

int find_first_digit(string line, int n)
{
    for (int i = 0; i < n; i++)
    {
        if (isdigit(line[i]))
        {
            return line[i] - '0';
        }
        if (line[i] == 'o' && i + 2 <= n - 1 && strcmp("one", line.substr(i, 3).c_str()) == 0)
        {
            return 1;
        }
        if (line[i] == 't' && i + 2 <= n - 1 && strcmp("two", line.substr(i, 3).c_str()) == 0)
        {
            return 2;
        }
        if (line[i] == 't' && i + 4 <= n - 1 && strcmp("three", line.substr(i, 5).c_str()) == 0)
        {
            return 3;
        }
        if (line[i] == 'f' && i + 3 <= n - 1 && strcmp("four", line.substr(i, 4).c_str()) == 0)
        {
            return 4;
        }
        if (line[i] == 'f' && i + 3 <= n - 1 && strcmp("five", line.substr(i, 4).c_str()) == 0)
        {
            return 5;
        }
        if (line[i] == 's' && i + 2 <= n - 1 && strcmp("six", line.substr(i, 3).c_str()) == 0)
        {
            return 6;
        }
        if (line[i] == 's' && i + 4 <= n - 1 && strcmp("seven", line.substr(i, 5).c_str()) == 0)
        {
            return 7;
        }
        if (line[i] == 'e' && i + 4 <= n - 1 && strcmp("eight", line.substr(i, 5).c_str()) == 0)
        {
            return 8;
        }
        if (line[i] == 'n' && i + 3 <= n - 1 && strcmp("nine", line.substr(i, 4).c_str()) == 0)
        {
            return 9;
        }
    }

    return -1;
}

int find_last_digit(string line, int n)
{
    for (int i = n - 1; i >= 0; i--)
    {
        if (isdigit(line[i]))
        {
            return line[i] - '0';
        }
        if (i - 2 >= 0 && line[i - 2] == 'o' && strcmp("one", line.substr(i - 2, 3).c_str()) == 0)
        {
            return 1;
        }
        if (i - 2 >= 0 && line[i - 2] == 't' && strcmp("two", line.substr(i - 2, 3).c_str()) == 0)
        {
            return 2;
        }
        if (i - 4 >= 0 && line[i - 4] == 't' && strcmp("three", line.substr(i - 4, 5).c_str()) == 0)
        {
            return 3;
        }
        if (i - 3 >= 0 && line[i - 3] == 'f' && strcmp("four", line.substr(i - 3, 4).c_str()) == 0)
        {
            return 4;
        }
        if (i - 3 >= 0 && line[i - 3] == 'f' && strcmp("five", line.substr(i - 3, 4).c_str()) == 0)
        {
            return 5;
        }
        if (i - 2 >= 0 && line[i - 2] == 's' && strcmp("six", line.substr(i - 2, 3).c_str()) == 0)
        {
            return 6;
        }
        if (i - 4 >= 0 && line[i - 4] == 's' && strcmp("seven", line.substr(i - 4, 5).c_str()) == 0)
        {
            return 7;
        }
        if (i - 4 >= 0 && line[i - 4] == 'e' && strcmp("eight", line.substr(i - 4, 5).c_str()) == 0)
        {
            return 8;
        }
        if (i - 3 >= 0 && line[i - 3] == 'n' && strcmp("nine", line.substr(i - 3, 4).c_str()) == 0)
        {
            return 9;
        }
    }

    return -1;
}

int main(int argc, char *argv[])
{
    ifstream file("input.txt");
    vector<int> pos(12, -1);

    int sum = 0;
    if (file.is_open())
    {
        string line;
        while (getline(file, line))
        {
            int n = line.size();
            int f = find_first_digit(line, n);
            int l = find_last_digit(line, n);

            if (f == -1 || l == -1)
            {
                cout << "Error: could not find first or last digit" << endl;
                return 1;
            }
            else
            {
                sum += f * 10 + l;
            }
        }
    }

    cout << sum << endl;
    return 0;
}
#include <iostream>
#include <fstream>
#include <regex>
#include <vector>
#include <string>

using namespace std;

vector<string> split(const string str, const string regex_str)
{
    regex regexz(regex_str);
    return {sregex_token_iterator(str.begin(), str.end(), regexz, -1), sregex_token_iterator()};
}

int main() {
    ifstream file("input.txt");
    int lowest_location_num = INT_MAX;


    if (file.is_open()) {
        string first_line;
        getline(file, first_line);
        vector<string> seeds = split(split(first_line, ":\\s+")[1], "\\s+");

        for (int i = 0; i < 7; i++) {
            string name; 
            getline(file, name);

            string line;
            vector<vector<string>> 
        }
    }

    return 0;
}
#include <iostream>
#include <fstream>
#include <cctype>

using namespace std;

int main(int argc, char *argv[]) {
    ifstream file("input.txt");
    int sum = 0;
    if (file.is_open()) {
        string line;
        while (getline(file, line)) {
            int start = 0;
            while (1) {
                if (isdigit(line[start])) {
                    break;
                }
                start++;
            }
            int end = line.size() - 1;
            while (1) {
                if (isdigit(line[end])) {
                    break;
                }
                end--;
            }

            sum += (line[start] - '0') * 10 + (line[end] - '0');
        }
        file.close();
    }
    cout << sum << endl;
    return 0;
}
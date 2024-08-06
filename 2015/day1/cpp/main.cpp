// Santa was hoping for a white Christmas, but his weather machine's "snow"
// function is powered by stars, and he's fresh out! To save Christmas, he needs
// you to collect fifty stars by December 25th.
//
// Collect stars by helping Santa solve puzzles. Two puzzles will be made
// available on each day in the Advent calendar; the second puzzle is unlocked
// when you complete the first. Each puzzle grants one star. Good luck!
//
// Here's an easy puzzle to warm you up.
//
// Santa is trying to deliver presents in a large apartment building, but he
// can't find the right floor - the directions he got are a little confusing. He
// starts on the ground floor (floor 0) and then follows the instructions one
// character at a time.
//
// An opening parenthesis, (, means he should go up one floor, and a closing
// parenthesis, ), means he should go down one floor.
//
// The apartment building is very tall, and the basement is very deep; he will
// never find the top or bottom floors.
//
// For example:
//
//     (()) and ()() both result in floor 0.
//     ((( and (()(()( both result in floor 3.
//     ))((((( also results in floor 3.
//     ()) and ))( both result in floor -1 (the first basement level).
//     ))) and )())()) both result in floor -3.
//
// To what floor do the instructions take Santa?

#include <fstream>
#include <iostream>
#include <limits>
#include <ostream>
#include <string>
using namespace std;

class part1 {
public:
  void processFile(const string &filename) {

    // Open the input file named "input.txt"
    ifstream inputFile("santa.txt");
    // Check if the file is successfully opened
    if (!inputFile.is_open()) {
      cerr << "There was an error opening the file" << endl;
    }

    // Declare a string variable to store each
    // line of the file
    string line;
    long long floor = 0;
    long long openCount = 0, closeCount = 0;

    // Read each line of the file and print it to the
    // standard output stream
    cout << "File content:" << endl;
    while (getline(inputFile, line)) {
      cout << line << endl;
      for (char c : line) {
        if (c == '(') {
          floor++;
          openCount++;
          if (floor > numeric_limits<int>::max()) {
            cout << "Warning: floor exceeded max int value" << endl;
          }
        } else if (c == ')') {
          floor--;
          closeCount++;
          if (floor < numeric_limits<int>::min()) {
            cout << "Warning: floor exceeded min int value" << endl;
          }
        }
      }
    }

    cout << "Total '(' count: " << openCount << endl;
    cout << "Total ') count: " << closeCount << endl;
    cout << "Final floor value: " << floor << endl;
    inputFile.close();
  };
};
int main(int argc, char *argv[]) {
  part1 solver;
  solver.processFile("santa.txt");
  return 1;
}

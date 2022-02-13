#include <iostream>
#include <vector>
#include <string>
#include <sstream>

int main() {
  std::string line;
  getline(std::cin, line);
  std::istringstream iss(line);
  int total;
  iss >> total;
  for (int n; iss >> n;) {
    total *= n;
  }
  std::cout << total << std::endl;
}

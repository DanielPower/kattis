#include <iostream>
#include <sstream>

int main() {
  std::string line;
  getline(std::cin, line);
  std::cout << ((line == "OCT 31" || line == "DEC 25") ? "yup" : "nope") << std::endl;
}

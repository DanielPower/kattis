#include <iostream>

int main() {
  int value;
  int total = 0;
  std::cin >> value;
  for (int value; std::cin >> value;) {
    total += value;
  }
  std::cout << total << std::endl;
}

#include <iostream>
#include <string>
#include <math.h>

int main() {
  std::string p;
  std::cin >> p;
  int total = 0;
  for (;std::cin >> p;) {
    int power = p.back() - '0';
    p.pop_back();
    int number = std::stoi(p);
    total += pow(number, power);
  }
  std::cout << total << std::endl;
}

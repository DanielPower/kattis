#include <algorithm>
#include <iostream>

int gcd(int a, int b) {
  if (b == 0) {
    return a;
  }
  return gcd(b, a % b);
}

int main() {
  int year, a, b;
  int lowest = 500000000;
  std::cin >> year;
  for (; std::cin >> year >> a >> b;) {
    int common = (a * b) / gcd(a, b);
    lowest = std::min(lowest, year + common);
  }
  std::cout << lowest << std::endl;
}

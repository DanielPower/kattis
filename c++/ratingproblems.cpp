#include <iostream>

int main() {
  int n, k;
  std::cin >> n >> k;
  float total = 0;
  for (int v; std::cin >> v;) {
    total += v;
  }
  std::cout << (total - 3*(n-k))/n  << " " << (total + 3*(n-k))/n << std::endl;
}

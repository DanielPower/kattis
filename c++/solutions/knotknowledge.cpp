#include <iostream>
#include <iterator>
#include <sstream>
#include <bits/stdc++.h>
#include <string>

int main() {
  int n;
  std::cin >> n;
  std::unordered_set<std::string> knots;
  for (int i=0; i<n; i++) {
    std::string knot;
    std::cin >> knot;
    knots.insert(knot);
  }
  for (std::string knot; std::cin >> knot;) {
    knots.erase(knot);
  }
  std::cout << *std::begin(knots) << std::endl;
}

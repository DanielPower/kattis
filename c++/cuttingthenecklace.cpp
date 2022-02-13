#include <iostream>
#include <deque>

int main() {
  int n, k;
  std::deque<int> weights;
  std::cin >> n >> k;
  for (int i=0; i<k; i++) {
    int w;
    std::cin >> w;
    weights.push_front(w);
  }

  int total_weight = 0;
  for (int weight : weights) {
    total_weight += weight;
  }

  if (total_weight % n != 0) {
    std::cout << "NO" << std::endl;
    return 0;
  }

  int weight_per_person = total_weight / n;
  for (int i=0; i<k; i++) {
    int current_weight = 0;
    bool possible = true;
    for (int weight : weights) {
      current_weight += weight;
      if (current_weight > weight_per_person) {
        possible = false;
        break;
      }
      if (current_weight == weight_per_person) {
        current_weight = 0;
      }
    }
    if (possible && current_weight == 0) {
      std::cout << "YES" << std::endl;
      return 0;
    }
    weights.push_front(weights.back());
    weights.pop_back();
  }
  std::cout << "NO" << std::endl;
  return 0;
}

#include <iostream>
#include <queue>

int main() {
  int count;
  for (; std::cin >> count;) {
    std::queue<int> queue = {};
    std::priority_queue<int> priority_queue = {};
    std::vector<int> stack = {};
    bool is_queue = true;
    bool is_priority_queue = true;
    bool is_stack = true;
    for (int i = 0; i < count; i++) {
      int operation, element;
      std::cin >> operation >> element;
      if (operation == 1) {
        if (is_queue) {
          queue.push(element);
        }
        if (is_priority_queue) {
          priority_queue.push(element);
        }
        if (is_stack) {
          stack.push_back(element);
        }
      } else {
        if (is_queue) {
          if (queue.empty() || queue.front() != element) {
            is_queue = false;
          } else {
            queue.pop();
          }
        }
        if (is_priority_queue) {
          if (priority_queue.empty() || priority_queue.top() != element) {
            is_priority_queue = false;
          } else {
            priority_queue.pop();
          }
        }
        if (is_stack) {
          if (stack.empty() || stack.back() != element) {
            is_stack = false;
          } else {
            stack.pop_back();
          }
        }
      }
    }
    if (!is_queue && !is_priority_queue && !is_stack) {
      std::cout << "impossible\n";
    } else if (is_queue && !is_priority_queue && !is_stack) {
      std::cout << "queue\n";
    } else if (!is_queue && is_priority_queue && !is_stack) {
      std::cout << "priority queue\n";
    } else if (!is_queue && !is_priority_queue && is_stack) {
      std::cout << "stack\n";
    } else {
      std::cout << "not sure\n";
    }
  }
}

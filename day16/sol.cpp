#include <cstdint>
#include <fstream>
#include <iostream>
#include <regex>
#include <string>
#include <unordered_map>
#include <unordered_set>

struct Valve {
  int flow;
  std::string name;

  Valve(const std::string name, const int flow) : name(name), flow(flow) {}
};

std::unordered_map<std::string, Valve *> valves;

std::unordered_map<std::string, std::vector<std::string>> graph;

std::unordered_map<std::uint64_t, int> DP;

int f(std::string cur_pos, std::unordered_set<std::string> opened_valves,
      int time) {
  if (time == 0) {
    return 0;
  }

  std::uint64_t bitmask = 0;
  for (const auto &[key, value] : graph) {
    if (opened_valves.count(key) == 0) {
      bitmask = bitmask << 1;
    } else {
      bitmask = (bitmask << 1) + 1;
    }
  }

  bitmask += (cur_pos[0] - 'A') * 26 + cur_pos[1] - 'A';

  bitmask <<= time;

  if (DP.count(bitmask)) {
    return DP[bitmask];
  }

  int ans = 0;
  Valve cur_valve = *valves[cur_pos];

  if (opened_valves.count(cur_pos) == 0 && cur_valve.flow > 0) {
    std::unordered_set<std::string> n_valves(opened_valves);
    n_valves.insert(cur_pos);
    ans = std::max(ans, (time - 1) * cur_valve.flow +
                            f(cur_valve.name, n_valves, time - 1));
  }

  for (auto child : graph[cur_pos]) {
    Valve child_valve = *valves[child];
    ans = std::max(ans, f(child_valve.name, opened_valves, time - 1));
  }

  DP[bitmask] = ans;

  return ans;
}

int main() {
  std::ifstream nameFileout;
  nameFileout.open("sample.txt");
  std::string line;

  const std::string delimeter = ", ";
  const std::regex pattern(
      R"(Valve ([A-Z]+) has flow rate=([0-9]+); tunnels? leads? to valves? ([\s\S]+))");

  while (std::getline(nameFileout, line)) {
    std::smatch match;
    std::regex_match(line, match, pattern);

    std::string token;
    size_t pos = 0;

    std::string cur_valve = match[1];
    int flow_rate = std::stoi(match[2]);
    std::string children = match[3];

    // std::cout << cur_valve << flow_rate << children << std::endl;

    valves[cur_valve] = new Valve(cur_valve, flow_rate);

    while ((pos = children.find(delimeter)) != std::string::npos) {
      token = children.substr(0, pos);
      graph[cur_valve].push_back(token);
      children = children.substr(pos + delimeter.length());
    }

    if (!children.empty()) {
      graph[cur_valve].push_back(children);
    }
  }

  int time = 30;
  std::cout << f("AA", std::unordered_set<std::string>(), time) << std::endl;

  // auto print_key_value = [](const auto &key, const auto &value) {
  //   std::cout << "Key:[" << key << "] Value:[" << &value << "]\n";
  //   for (auto &child : value) {
  //     std::cout << child << " " << std::endl;
  //   }
  // };
  //
  // for (const auto &[key, value] : graph)
  //   print_key_value(key, value);
}

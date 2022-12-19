#include <cstdint>
#include <fstream>
#include <iostream>
#include <ostream>
#include <regex>
#include <set>
#include <string>

struct Point {
  int x;
  int y;
  int d;
  Point(const int x, const int y) : x(x), y(y) {}

  bool operator<(const Point &pos) const {
    if (x == pos.x) {
      return y < pos.y;
    }
    return x < pos.x;
  }
  bool operator==(const Point &pos) const { return x == pos.x && y == pos.y; }
};

int main() {
  std::string item_name;
  std::ifstream nameFileout;
  nameFileout.open("sample.txt");
  std::string line;

  const std::regex pattern(
      R"(Sensor at x=(-?[0-9]+), y=(-?[0-9]+): closest beacon is at x=(-?[0-9]+), y=(-?[0-9]+))");

  std::set<Point> sensors;
  std::set<Point> beacons;
  std::set<Point> res;

  const std::uint64_t grid_max = 4000000;
  const int grid_min = 0;

  while (std::getline(nameFileout, line)) {
    std::smatch match;
    std::regex_match(line, match, pattern);

    Point sensor = Point(std::stoi(match[1]), std::stoi(match[2]));
    res.insert(sensor);

    const auto beacon = Point(std::stoi(match[3]), std::stoi(match[4]));
    beacons.insert(beacon);
    res.insert(beacon);

    int dist = std::abs(sensor.x - beacon.x) + std::abs(sensor.y - beacon.y);
    sensor.d = dist;
    sensors.insert(sensor);

    for (int d = 0; d <= dist + 1; d++) {
      Point left_up(sensor.x - dist - 1 + d, sensor.y + d);
      Point right_up(sensor.x + dist + 1 - d, sensor.y + d);
      Point left_down(sensor.x - dist - 1 + d, sensor.y - d);
      Point right_down(sensor.x + dist + 1 - d, sensor.y - d);

      res.insert(left_up);
      res.insert(right_up);
      res.insert(left_down);
      res.insert(right_down);
    }
  }

  for (Point res : res) {
    int found = 1;
    if (res.x < 0 || res.x > grid_max || res.y < 0 || res.y > grid_max) {
      continue;
    }
    for (Point sensor : sensors) {
      if (std::abs(res.x - sensor.x) + std::abs(res.y - sensor.y) <= sensor.d) {
        found = 0;
        break;
      }
    }
    if (found == 1) {
      const std::uint64_t tuning_freq = (grid_max * res.x) + res.y;
      std::cout << res.x << ", " << res.y << std::endl;
      std::cout << tuning_freq << std::endl;
      break;
    }
  }

  // std::cout << "Results:" << std::endl;
  // for (Point p : res) {
  //   std::cout << p.x << ", " << p.y << std::endl;
  // }
  return 0;
}

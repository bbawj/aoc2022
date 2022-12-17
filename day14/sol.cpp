#include "iostream"
#include <algorithm>
#include <fstream>
#include <set>
#include <string>
#include <vector>

class Pos {
public:
  int x;
  int y;

  Pos(int x, int y) {
    this->x = x;
    this->y = y;
  }

  bool operator<(const Pos &pos) const {
    return std::minmax(x, y) < std::minmax(pos.x, pos.y);
  }
  bool operator==(const Pos &pos) const {
    return std::minmax(x, y) == std::minmax(pos.x, pos.y);
  }
};

int main() {
  std::string item_name;
  std::ifstream nameFileout;
  nameFileout.open("input.txt");
  std::string line;

  std::vector<Pos> coords;
  std::set<Pos> filled;

  while (std::getline(nameFileout, line)) {
    std::string delimiter = " -> ";
    std::string pos_delimeter = ",";

    size_t pos = 0;
    std::string token;
    while ((pos = line.find(delimiter)) != std::string::npos) {
      token = line.substr(0, pos);
      std::string x = token.substr(0, token.find(pos_delimeter));
      std::string y =
          token.substr(token.find(pos_delimeter) + pos_delimeter.length());

      Pos rock(std::stoi(x), std::stoi(y));
      coords.push_back(rock);

      line = line.substr(pos + delimiter.length());
    }
    std::string x = line.substr(0, line.find(pos_delimeter));
    std::string y =
        line.substr(line.find(pos_delimeter) + pos_delimeter.length());
    Pos rock(std::stoi(x), std::stoi(y));
    coords.push_back(rock);
    coords.push_back(Pos(-1, -1));
  }

  for (int i = 1; i < coords.size(); i++) {
    Pos cur = coords[i];
    Pos prev = coords[i - 1];
    if (cur.x == -1 || prev.x == -1) {
      continue;
    }

    if (cur.y == prev.y) {
      for (int j = std::min(cur.x, prev.x); j <= std::max(cur.x, prev.x); j++) {
        Pos rock(j, cur.y);
        // std::cout << rock.x << ", " << rock.y << std::endl;
        filled.insert(rock);
      }
    }

    if (cur.x == prev.x) {
      for (int j = std::min(cur.y, prev.y); j <= std::max(cur.y, prev.y); j++) {
        Pos rock(cur.x, j);
        // std::cout << rock.x << ", " << rock.y << std::endl;
        filled.insert(rock);
      }
    }
  }

  // std::cout << filled.size() << std::endl;
  // for (Pos pos : coords) {
  //   std::cout << pos.x << ", " << pos.y << std::endl;
  // }

  // sand
  int maxY = filled.rbegin()->y;
  Pos start(500, 0);
  Pos sand = start;
  int begin = filled.size();
  while (1) {
    // std::cout << sand.x << ", " << sand.y << std::endl;
    if (sand.y == maxY + 1) {
      filled.insert(sand);
      sand = start;
      continue;
    }
    Pos down(sand.x, sand.y + 1);
    if (filled.count(down) == 0) {
      sand = down;
      continue;
    }

    Pos downLeft(sand.x - 1, sand.y + 1);
    if (filled.count(downLeft) == 0) {
      sand = downLeft;
      continue;
    }

    Pos downRight(sand.x + 1, sand.y + 1);
    if (filled.count(downRight) == 0) {
      sand = downRight;
      continue;
    }

    filled.insert(sand);
    if (sand.x == 500 && sand.y == 0) {
      break;
    }
    sand = start;
  }
  std::cout << maxY << std::endl;
  std::cout << filled.size() - begin << std::endl;
}

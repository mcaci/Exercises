#include <boost/filesystem.hpp>
#include <fstream>
#include <sstream>

#include "Utils.hpp"

boost::filesystem::path getDataPath() {
  return boost::filesystem::current_path() / "./";
}

std::vector<std::string> loadLinesFromFile(const std::string& iInputFile) {

  std::vector<std::string> aData;
  auto aPath = getDataPath();

  if (boost::filesystem::is_directory(aPath)) {
    auto aFilePath = aPath / iInputFile;
    std::ifstream aIfs(aFilePath.string().c_str());
    if (aIfs.good()) {
      std::string aLine;
      while (std::getline(aIfs, aLine)) {
        aData.push_back(aLine);
      }
    }
  }
  else {
    ERROR("Path to 'data' folder do not exists. Looking for '" << aPath << "'");
  }
  return aData;
}

std::string loadStringFromFile(const std::string& iInputFile) {
  std::vector<std::string> aLines = loadLinesFromFile(iInputFile);
  std::stringstream aData;
  for (auto& aLine : aLines)
  {
    aData << aLine << "\n";
  }
  return aData.str();
}

#include "Kata.hpp"

const std::string swallowed(const int index) {
  char aDelimiter = index == 1 || index == 2 ? ';' : ',';
  return "She swallowed the " + aAnimals[index] + " to catch the " + aAnimals[index-1] + aDelimiter + "\n";
}

const std::string lastVerseLine(const int index) {
  const std::string p = index != 0 ? "P" : "p";
  return "I don't know why she swallowed a fly - " + p + "erhaps she'll die!\n\n";
}

const std::string beginVerse(const int index) {
    std::string aConjunction = index == 4 ? "that" : "who";
    return "There was an old lady " + aConjunction + " swallowed a " + aAnimals[index] + ";\n";
}

const std::string createVerse(const int verseIndex) {
    std::string aVerse = beginVerse(verseIndex);
    if (verseIndex > 1) {
      aVerse += aAnimalVerses[verseIndex];
    }
    if (verseIndex != aAnimals.size() - 1) {
      for (int index = verseIndex; index > 0; index--) {
        if (index == 1) {
            aVerse += aAnimalVerses[index];
        }
        aVerse += swallowed(index);
      }
      aVerse += lastVerseLine(verseIndex);
    }
    return aVerse;
}

std::string Kata::produceSong() {
  std::string song;
  for (int i = 0; i <= 7; i++) {
    song += createVerse(i);   
  } 
  return song;
}

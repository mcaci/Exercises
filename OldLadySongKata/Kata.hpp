#ifndef KATA_HPP_
#define KATA_HPP_

#include <string>
#include <vector>

const std::string aFlyVerse = "";
const std::string aSpiderVerse = "That wriggled and jiggled and tickled inside her!\n";
const std::string aBirdVerse = "How absurd to swallow a bird!\n";
const std::string aCatVerse = "Fancy that she swallowed a cat!\n";
const std::string aDogVerse = "What a hog, to swallow a dog!\n";
const std::string aGoatVerse = "She just opened her throat and swallowed a goat!\n";
const std::string aCowVerse = "I don't know how she swallowed a cow!\n";
const std::string aHorseVerse = "...She's dead, of course!\n";

const std::vector<std::string> aAnimals = {"fly", "spider", "bird", "cat", "dog", "goat", "cow", "horse"};
const std::vector<std::string> aAnimalVerses = {aFlyVerse, aSpiderVerse, aBirdVerse, aCatVerse, aDogVerse, aGoatVerse, aCowVerse, aHorseVerse};

class Kata {

public:
  std::string produceSong();
};

#endif

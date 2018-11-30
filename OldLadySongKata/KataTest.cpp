#include "gtest/gtest.h"

#include "Kata.hpp"
#include "Utils.hpp"

namespace test
{

TEST(KataTest, test_whenNothing_shouldBeTrue) {
  auto aExpectedString = loadStringFromFile("song.txt");
  Kata aKata;
  ASSERT_EQ(aExpectedString, aKata.produceSong());
}

} // namespace test


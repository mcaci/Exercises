package kata.potter;

public class Basket {

  int[] books = new int[5];

  public void add(int num) {
    books[num - 1]++;
  }

  public int getSize() {
    int count = 0;
    for (int i = 0; i < 5; i++) {
      count += books[i];
    }
    return count;
  }

  public int getNumberOfDifferentBooks() {
    int count = 0;
    for (int i = 0; i < 5; i++) {
      if (books[i] > 0) {
        count++;
      }
    }
    return count;
  }

  public Basket copy(){
    int[] copy = new int[5];
    copy[0] = books[0];
    copy[1] = books[1];
    copy[2] = books[2];
    copy[3] = books[3];
    copy[4] = books[4];
    Basket result = new Basket();
    result.books = copy;
    return result;
  }

}

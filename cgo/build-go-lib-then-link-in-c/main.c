#include <stdio.h>
#include "libsum.h"

int main() {
  printf("%d\n", sum(1, 2));

  int n = 10;
  printf("------ Now showing first %d primes ------\n", 1000);
  showPrimes(n);
}

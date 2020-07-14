#include <stdio.h>
#include "libsum.h"

int main() {
  printf("%d\n", sum(1, 2));

  int n = 100000;
  printf("------ Now showing first %d primes ------", 1000);
  showPrimes(n);
}

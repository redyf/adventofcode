#include <stdio.h>
#include <stdlib.h>

int main(int argc, char *argv[]) {
  // Create a pointer to the file
  FILE *fptr;
  int ch;
  int total = 0;

  // Read its content
  fptr = fopen("floor.txt", "r");

  if (fptr == NULL) {
    printf("Error reading file!");
    exit(1);
  }

  while ((ch = fgetc(fptr)) != EOF) {
    if (ch == '(') {
      total++;
    } else if (ch == ')') {
      total--;
    }
  }
  printf("%d\n", total);

  // Close file
  fclose(fptr);

  return EXIT_SUCCESS;
}

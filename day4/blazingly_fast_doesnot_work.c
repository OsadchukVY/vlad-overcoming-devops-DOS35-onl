/* WORKS ONLY WITH HOST FILESYSTEM*/

#define _FILE_OFFSET_BITS 64
#define _GNU_SOURCE
#include <fcntl.h>
#include <stdio.h>
#include <stdlib.h>
#include <sys/stat.h>
#include <sys/types.h>
#include <unistd.h>

void close_f(int *fd) {
  if (*fd >= 0)
    close(*fd);
}

int main(int argc, char* argv[]) {

  if (argc != 3) {
    fprintf(stderr, "Usage: %s <source> <destination>\n", argv[0]);
    exit(EXIT_FAILURE);
  }

  size_t in_f_size, out_f_size;
  struct stat in_stat, out_stat;

  int fd_in_file  __attribute__((cleanup(close_f)))  = open(argv[1], O_RDONLY);
  if ((fd_in_file == -1) && (fstat(fd_in_file, &in_stat) == -1)) {
    perror("FAILED TO OPEN SHARED FILE");
    exit(1);
  }

  in_f_size = in_stat.st_size;

  int fd_out_file __attribute__((cleanup(close_f))) = open(argv[2],O_CREAT|O_WRONLY|O_TRUNC,0644);
  if (fd_out_file == -1) {
    perror("FAILED TO CREATE FILE");
    exit(1);
  }

  do {
    out_f_size = copy_file_range(fd_in_file, NULL, fd_out_file, NULL, in_f_size, 0); 
    if (out_f_size == -1) {
      perror("FAILED TO COPY FILE");
      return EXIT_FAILURE;
    }

    in_f_size -= out_f_size;

  } while((in_f_size > 0) && (out_f_size > 0));

  return EXIT_SUCCESS;
}

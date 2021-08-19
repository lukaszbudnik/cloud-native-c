#include <stdio.h>
#include <string.h>
#include "response.h"
#include "storage.h"

int main(int argc, char **argv)
{
  if (argc != 3)
  {
    fprintf(stderr, "This app requires exactly 2 parameters first is the path to local file and second is the S3 bucket.\n");
    fprintf(stderr, "Example:\n%s app.c my-bucket-name\n", argv[0]);
    return WRONG_PARAMETERS_ERROR;
  }

  int res = PutObject(argv[1], argv[2]);

  if (res > 0)
  {
    fprintf(stderr, "There was an error uploading the file.\n");
    return res;
  }

  char *content = GetObject(argv[1], argv[2]);
  if (content == NULL)
  {
    return FILE_DOWNLOAD_ERROR;
  }

  printf("Content of the uploaded file is:\n%s\n", content);
  Free(content);

  return SUCCESS;
}

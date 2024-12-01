#include <stdio.h>
#include <stdlib.h>

char *readFile(const char *filePath) {
	FILE *file = fopen(filePath, "r");
	if (!file) {
		return NULL;
	}

	// Get file size
	if (fseek(file, 0, SEEK_END) != 0) {
		fclose(file);
		return NULL;
	}

	long fileSize = ftell(file);
	if (fileSize == -1) {
		fclose(file);
		return NULL;
	}

	rewind(file);

	// Allocate memory for the contents
	char *contents = malloc(fileSize + 1); // +1 for the null terminator
	if (!contents) {
		fclose(file);
		return NULL;
	}

	// Read file contents
	size_t bytesRead = fread(contents, 1, fileSize, file);
	fclose(file);

	// Ensure null termination
	contents[bytesRead] = '\0';

	return contents;
}

int main(void) {
	return 0;
}

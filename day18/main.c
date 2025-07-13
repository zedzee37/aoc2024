#include <ctype.h>
#include <stddef.h>
#include <stdint.h>
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int to_1d(int x, int y) {
	return x * 70 + y;
}

typedef struct {
	bool is_wall;
	int g_cost;
	int h_cost;
} AStarCell;

typedef AStarCell * Grid;

Grid create_grid() {
	AStarCell *grid = malloc(sizeof(AStarCell) * 70 * 70);
	if (!grid) {
		perror("Coult not initialize the grid!");
		return NULL;
	}
	return grid;
}

void populate_grid(Grid grid, char *file_contents, int bytes) {
	int file_pos = 0;
	size_t file_size = strlen(file_contents);

	int x = -1;
	int y = -1;
	int byte_count = 0;
	while (byte_count < bytes) {
		char ch = file_contents[file_pos];
		file_pos++;

		if (!isdigit(ch)) {
			continue;
		}
		
		char buf[3] = "00";
		buf[2] = '\0';
		
		buf[0] = ch;
		
		ch = file_contents[file_pos];
		if (isdigit(ch)) {
			buf[1] = ch;
		}

		int value = atoi(buf);
		
		if (x == -1) {
			x = value;
		} else if (y == -1) {
			y = value;
			grid[to_1d(x, y)].is_wall = true;
			byte_count++;
			x = -1;
			y = -1;
		}
	}
}

void print_grid(Grid grid) {
	for (int y = 0; y < 70; y++) {
		for (int x = 0; x < 70; y++) {
			int idx = to_1d(x, y);
			AStarCell *cell = &grid[idx];

			if (cell->is_wall) {
				printf("#");
			} else {
				printf(".");
			}
		}
		printf("\n");
	}
}

char *read_file(const char *file_path) {
	FILE *fp = fopen(file_path, "r");
	if (!fp) {
		return NULL;
	}

	fseek(fp, SEEK_SET, SEEK_END);
	size_t size = ftell(fp);
	fseek(fp, SEEK_SET, SEEK_SET);

	char *file_contents = malloc(sizeof(char) * size + 1);
	
	fread(file_contents, sizeof(char), size, fp);
	file_contents[size] = '\0';

	fclose(fp);

	return file_contents;
}

int main() {
	char *file_contents = read_file("input.txt");
	Grid grid = create_grid();

	populate_grid(grid, file_contents, 1024);
	print_grid(grid);

	free(file_contents);
	free(grid);
	return 0;
}

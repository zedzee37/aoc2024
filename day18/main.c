#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

typedef struct {
	uint64_t first_bits;
	uint8_t last_bits;
} GridRow;

bool grid_row_is_corrupted(GridRow *row, int x) {
	x++;
	if (x > 70) {
		perror("x > 70 found");
		exit(-1);
	}
	if (x > 64) {
		int x_diff = x - 64;
		uint8_t target_bit = 1 << x_diff;
		return row->last_bits & target_bit;
	}
	
	uint64_t target_bit = (uint64_t)1 << x;
	return row->first_bits & target_bit;
}

void grid_row_set(GridRow *row, int x) {
	x++;
    if (x > 70 || x < 0) {
        perror("grid_row_set: x out of bounds");
        exit(-1);
    }

    if (x > 64) {
        int x_diff = x - 64;
        row->last_bits |= 1 << x_diff;
    } else {
        row->first_bits |= (uint64_t)1 << x;
    }
}

GridRow *populate_grid(int steps) {
	GridRow *rows = malloc(sizeof(GridRow) * 70);
}

int main() {
	GridRow *row = malloc(sizeof(GridRow));
	return 0;
}

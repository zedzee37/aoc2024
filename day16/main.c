#include <stdint.h>
#include <stdlib.h>

typedef struct {
    int32_t x;
    int32_t y;
} Vec2;

Vec2 add(Vec2 a, Vec2 b) {
    Vec2 res = { .x = a.x + b.x, .y = a.y + b.y };
    return res;
}

uint32_t manhattan_distance(Vec2 a, Vec2 b) {
    return abs(a.x - b.x) + abs(a.y - b.y);
}

Vec2 right(Vec2 dir) {
    Vec2 res = { .x = -dir.y, .y = dir.x };
    return res;
}

Vec2 left(Vec2 dir) {
    Vec2 res = { .x = dir.y, .y = -dir.x };
    return res;
}

typedef struct {
    Vec2 pos;
    Vec2 dir;
    uint32_t cost; 
} Node;

int main() {
    return 0;
}

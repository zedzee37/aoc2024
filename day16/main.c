#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>
#include <string.h>

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

Node test_node(uint32_t cost) {
    Vec2 v = { 0, 0 };
    Node n = { v, v, cost };
    return n;
}

typedef struct {
    Node *nodes;
    size_t capacity;
    size_t len;
} PriorityQueue;

PriorityQueue *new_queue() {
    PriorityQueue *queue = malloc(sizeof(PriorityQueue));
    
    queue->nodes = malloc(sizeof(Node) * 10);
    queue->capacity = 10;
    queue->len = 0;

    return queue;
}

void queue_free(PriorityQueue *queue) {
    free(queue->nodes);
    free(queue);
}

uint32_t queue_parent(uint32_t idx) {
    return (idx - 1) / 2;
}

uint32_t queue_right(uint32_t idx) {
    return 2 * idx + 2;
}

uint32_t queue_left(uint32_t idx) {
    return 2 * idx + 1;
}

void check_heap(PriorityQueue *queue) {
    uint32_t idx = queue->len - 1;
    while (idx != 0) {
        Node current = queue->nodes[idx];
        uint32_t above = queue_parent(idx);

        Node above_node = queue->nodes[above];
        if (above_node.cost > current.cost) {
            queue->nodes[idx] = above_node;
            queue->nodes[above] = current;
        } else {
            break;
        }
        idx = above;
    }
}

void queue_insert(PriorityQueue *queue, Node *node) {
    if (queue->capacity >= queue->len - 1) {
        queue->capacity *= 2;
        queue->nodes = realloc(queue->nodes, sizeof(Node) * queue->capacity);
    }

    queue->nodes[queue->len++] = *node;
    check_heap(queue);
}

void check_heap_down(PriorityQueue *queue, uint32_t i) {
    uint32_t idx = i;
    uint32_t left = queue_left(idx);
    uint32_t right = queue_left(idx);
    uint32_t cur_cost = queue->nodes[idx].cost;

    if (left < queue->len && queue->nodes[left].cost < cur_cost) {
        idx = left; 
    }
    if (right < queue->len && queue->nodes[right].cost < cur_cost) {
        idx = right; 
    }
    if (idx != i) {
        Node temp = queue->nodes[i];
        queue->nodes[i] = queue->nodes[idx];
        queue->nodes[idx] = temp;
        
        check_heap_down(queue, idx);
    }
}

Node queue_pop(PriorityQueue *queue) {
    Node ret = queue->nodes[0];
    queue->nodes[0] = queue->nodes[queue->len - 1];
    queue->len--;

    check_heap_down(queue, 0);
    return ret;
}

char **read_input(const char *file_path) {
    FILE *file_handle = fopen(file_path, "r");
    if (!file_handle) {
        return NULL;
    }

    fseek(file_handle, 0, SEEK_END);
    size_t size = ftell(file_handle);
    fseek(file_handle, 0, SEEK_SET);

    fclose(file_handle); 
}

int main() {
    PriorityQueue *queue = new_queue();

    return 0;
}

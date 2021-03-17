#include "binarySearchTree.h"
#include <stdio.h>

int main() {
    BinarySearchTree *tree = NULL;

    int data[5][2] = {{3, 3},
                      {1, 1},
                      {4, 4},
                      {2, 2},
                      {5, 5}};

    int i;
    for (i = 0; i < sizeof(int[5][2]) / sizeof(int[2]); i++) {
        insert(&tree, data[i][0], data[i][1]);
    }
    midOrder(tree);

    // 删除结点
    int count = 0, key = 0;

    printf("Please input count:");
    scanf("%d", &count);

    while(count--) {
        printf("Please input key:");
        scanf("%d", &key);
        delete(&tree, key);
        midOrder(tree);
    }
    return 0;
}


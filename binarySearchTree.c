// 二叉搜索树的实现，包括插入，删除，查找，遍历
#include <stdio.h>
#include <stdlib.h>
#include "binarySearchTree.h"

/**
 * 获取树中最大值结点，一定是叶子结点
 * @param tree
 * @return
 */
BinarySearchNode *getMaxNode(BinarySearchTree *tree) {
    if (NULL == tree) {
        return NULL;
    }
    BinarySearchNode *p = tree;
    while (NULL != p->right) {
        p = p->right;
    }
    return p;
}

/**
 * 获取树中最小值结点，一定是叶子结点
 * @param tree
 * @return
 */
BinarySearchNode *getMinNode(BinarySearchTree *tree) {
    if (NULL == tree) {
        return NULL;
    }
    BinarySearchNode *p = tree;
    while (NULL != p->left) {
        p = p->left;
    }
    return p;
}

/**
 * 二叉搜索树插入
 * @param tree
 * @param key
 * @param value
 */
void insert(BinarySearchTree **tree, int key, int value) {
    if (NULL == *tree) { // 如果是空树，则创建根结点，并指向它
        *tree = (BinarySearchNode *) malloc(sizeof(BinarySearchNode));
        (*tree)->key = key;
        (*tree)->value = value;
        (*tree)->left = NULL;
        (*tree)->right = NULL;
        return;
    }
    // p指向当前结点
    BinarySearchNode *p = *tree, *parent = NULL;
    while (NULL != p) {
        if (key == p->key) { // 如果key和当前结点的key相同，则覆盖value并返回
            p->value = value;
            return;
        } else if (key < p->key) { // 如果key小于当前结点，则移动到左子树
            parent = p;
            p = p->left;
        } else { // 如果key大于当前结点，则移动到右子树
            parent = p;
            p = p->right;
        }
    }
    // key和p指针所指的结点进行比较
    if (key < parent->key) {
        parent->left = (BinarySearchNode *) malloc(sizeof(BinarySearchNode));
        parent->left->key = key;
        parent->left->value = value;
        parent->left->left = NULL;
        parent->left->right = NULL;
    } else {
        parent->right = (BinarySearchNode *) malloc(sizeof(BinarySearchNode));
        parent->right->key = key;
        parent->right->value = value;
        parent->right->left = NULL;
        parent->right->right = NULL;
    }
}

/**
 * 二叉搜索树删除操作
 * @param tree
 * @param key
 */
void delete(BinarySearchTree **tree, int key) {
    if (NULL == *tree) {
        return; // 如果是一颗空树则直接返回
    }

    // parent指向当前结点的父结点，p指向当前结点
    BinarySearchNode *parent = NULL, *p = *tree;
    while (NULL != p) {
        if (key < p->key) {
            parent = p;
            p = p->left;
        } else if (key > p->key) {
            parent = p;
            p = p->right;
        } else {
            if (NULL != p->left && NULL != p->right) { // 若当前结点，左子树、右子树都存在
                BinarySearchNode *pMax = getMaxNode(p->left);
                p->key = pMax->key;
                p->value = pMax->value;
                delete(&p->left, pMax->key);
            } else if (NULL != p->left) { // 若左子树存在
                BinarySearchNode *pMax = getMaxNode(p->left);
                p->key = pMax->key;
                p->value = pMax->value;
                delete(&p->left, pMax->key);
            } else if (NULL != p->right) { // 若右子树存在
                BinarySearchNode *pMin = getMinNode(p->right);
                p->key = pMin->key;
                p->value = pMin->value;
                delete(&p->right, pMin->key);
            } else { // 若当前结点是叶子结点
                if (NULL == parent) {
                    *tree = NULL;
                    free(p);
                    p = NULL;
                } else {
                    if (key < parent->key) {
                        parent->left = NULL;
                        free(p);
                        p = NULL;
                    } else {
                        parent->right = NULL;
                        free(p);
                        p = NULL;
                    }
                }
            }
            break;
        }
    }
}

/**
 * 二叉搜索树的查找操作
 * @param tree
 * @param key
 * @return
 */
unsigned int find(BinarySearchTree *tree, int key) {
    if (NULL == tree) {
        return -1;// 如果是空树直接返回-1，表示未找到
    }
    BinarySearchNode *p = tree;
    while (NULL != p) {
        if (key == p->key) { // 如果跟当前结点关键字相等则直接返回
            return p->value;
        } else if (key < p->key) { // 如果小于当前结点关键字，则移动指针到左孩子
            p = p->left;
        } else { // 如果大于当前结点，则移动指针到右孩子
            p = p->right;
        }
    }
    return -1;
}

/**
 * 先序遍历--递归实现
 * @param tree
 */
void preOrder(BinarySearchTree *tree) {
    if (NULL == tree) {
        return;
    }
    printf("%d\n", tree->key);
    preOrder(tree->left);
    preOrder(tree->right);
}

/**
 * 中序遍历--递归实现
 * @param tree
 */
void midOrder(BinarySearchTree *tree) {
    if (NULL == tree) {
        return;
    }
    midOrder(tree->left);
    printf("%d\n", tree->key);
    midOrder(tree->right);
}

/**
 * 后序遍历
 * @param tree
 */
void afterOrder(BinarySearchTree *tree) {
    if (NULL == tree) {
        return;
    }
    afterOrder(tree->left);
    afterOrder(tree->right);
    printf("%d\n", tree->key);
}

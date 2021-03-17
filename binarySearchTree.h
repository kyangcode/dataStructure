// 树结点的定义
struct Node {
    int key; // 关键字
    unsigned int value; // 数据部分
    struct Node *left; // 左结点的指针
    struct Node *right; // 右结点的指针
};
// 二叉搜索树结点类型
typedef struct Node BinarySearchNode;
// 二叉搜索树类型
typedef struct Node BinarySearchTree;

/**
 * 获取树中最大值结点，一定是叶子结点
 * @param tree
 * @return
 */
BinarySearchNode *getMaxNode(BinarySearchTree *tree);

/**
 * 获取树中最小值结点，一定是叶子结点
 * @param tree
 * @return
 */
BinarySearchNode *getMinNode(BinarySearchTree *tree);

/**
 * 二叉搜索树插入
 * @param tree
 * @param key
 * @param value
 */
void insert(BinarySearchTree **tree, int key, int value);

/**
 * 二叉搜索树删除操作
 * @param tree
 * @param key
 */
void delete(BinarySearchTree **tree, int key);

/**
 * 二叉搜索树的查找操作
 * @param tree
 * @param key
 * @return
 */
unsigned find(BinarySearchTree *tree, int key);

/**
 * 先序遍历--递归实现
 * @param tree
 */
void preOrder(BinarySearchTree *tree);

/**
 * 中序遍历--递归实现
 * @param tree
 */
void midOrder(BinarySearchTree *tree);

/**
 * 后序遍历
 * @param tree
 */
void afterOrder(BinarySearchTree *tree);
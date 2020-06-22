# 带mem的subtree示例
解决多个子树重用，导致running状态混乱问题。
修改方法是修改subtree实现，将原来的多个tick改为只用一个主tick。

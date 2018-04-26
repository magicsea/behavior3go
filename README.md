# behavior3go
golang behavior tree,from http://behavior3.com
## 简介
带在线编辑器的行为树，可使用官方的在线编辑器编辑逻辑节点。  
使用js版本翻译，保持和原版的编辑器数据格式一致。   
此行为树和一般的行为树略有不同，行为树结构只保持一份无状态，状态记录在黑板里（一般行为树每个对象一份树结构，树结构保存状态）。  
[>>原版编辑器<<](http://editor.behavior3.com/#/dash/home/)  
原版好像挂了，可以用[编译好的桌面版](https://pan.baidu.com/s/1L68N-AODskOKrw-0vCYKJQ)
## 示例
在loader文件夹里，参考loader_test.go。  
通过tree.json文件与编辑器关联，可使用编辑器导入导出此文件内容。
## 完整示例
[io类游戏示例](https://github.com/magicsea/h5game/tree/master/server)  
bin/b3.json为行为树的数据，在编辑器中导入树就可以还原工程，如图。  

![image](https://github.com/magicsea/behavior3go/blob/master/b3_simple1.png)

## 网页版编辑器本地搭建方法
- 下载源码https://github.com/behavior3/behavior3editor到本地(工程目录)
- 安装nodejs,npm
- 安装bower: npm install -g bower
- 安装依赖包:cd到工程目录下 npm install 然后 bower install
- 安装gulp:npm install --global gulp
- 运行:在工程目录下 gulp serve
- 客户端用浏览器打开 http://127.0.0.1:8000
## 联系
QQ群:285728047

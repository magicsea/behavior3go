# behavior3go
golang behavior tree,from https://github.com/behavior3
## 简介
带在线编辑器的行为树，可使用官方的在线编辑器编辑逻辑节点。  
使用js版本翻译，保持和原版的编辑器数据格式一致。   
此行为树和一般的行为树略有不同，行为树结构只保持一份无状态，状态记录在黑板里（一般行为树每个对象一份树结构，树结构保存状态）。  
[专用的编辑器分支版本](https://github.com/magicsea/behavior3editor)   
[重新部署的WEB版编辑器](http://47.101.48.70/b3/#/dash/home)  
[编译好的桌面版](https://github.com/magicsea/behavior3editor/releases)
## 示例 Examples
- [load_from_tree](https://github.com/magicsea/behavior3go/tree/master/examples/load_from_tree)  ：从导出的树文件加载示例
- [load_from_project](https://github.com/magicsea/behavior3go/tree/master/examples/load_from_project) ：从导出的工程文件加载示例
- [load_from_rawproject](https://github.com/magicsea/behavior3go/tree/master/examples/load_from_rawproject) ：从原生工程文件加载示例
- [subtree](https://github.com/magicsea/behavior3go/tree/master/examples/subtree) ：子树的使用示例（需要[专用的编辑器分支版本](https://github.com/magicsea/behavior3editor)）

## 完整示例
[io类游戏示例](https://github.com/magicsea/h5game/tree/master/server)  
bin/b3.json为行为树的数据，在编辑器中导入树就可以还原工程，如图。  

![image](https://github.com/magicsea/behavior3go/blob/master/b3_simple1.png)

## 网页版编辑器本地搭建方法
- 下载源码到本地(工程目录) https://github.com/behavior3/behavior3editor
- 安装nodejs,npm
- 安装bower: npm install -g bower
- 安装依赖包:cd到工程目录下 npm install 然后 bower install
- 安装gulp:npm install --global gulp
- 运行:在工程目录下 gulp serve
- 客户端用浏览器打开 http://127.0.0.1:8000
- 自己部署web版本，只需要把生成的build目录放到自己的tomcat,IIS目录就可以浏览器访问

## 更新
* 添加子树支持 SubTree 节点，需要编辑器修改node导出category字段

## 联系
QQ群:285728047

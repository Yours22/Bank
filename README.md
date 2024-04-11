# Bank BackEnd

学习写一个银行的后端。目的是练习和熟悉工具链。

更详细的笔记记录在`Note`中，请自行查阅，自认为会有所帮助。

## 开发环境
windows11 wsl2 

## 数据库 
使用dbdiagram.io 和 docker postgresql:12-alpine  编写和设计数据库。  

### 备注  
127.0.0.1  
端口 5432  
用户名 postgres  
密码 password  
暂时不放在服务器上测试。   
进入postgres的CLI命令行界面
```bash
docker exec -it [container名字，目前为postgres12] psql -U [数据库用户] -l
```
其他操作可以参考这篇博客。  
[Docker筆記 - 進入Container，建立並操作 PostgreSQL Container](https://medium.com/alberthg-docker-notes/docker%E7%AD%86%E8%A8%98-%E9%80%B2%E5%85%A5container-%E5%BB%BA%E7%AB%8B%E4%B8%A6%E6%93%8D%E4%BD%9C-postgresql-container-d221ba39aaec)   
数据库的操作参考[菜鸟教程](www.runoob.com)即可。

原文链接

    https://blog.csdn.net/why444216978/article/details/119827014?utm_medium=distribute.pc_aggpage_search_result.none-task-blog-2~aggregatepage~first_rank_ecpm_v1~rank_v31_ecpm-4-119827014.pc_agg_new_rank&utm_term=docker+%E6%90%AD%E5%BB%BAetcd%E9%9B%86%E7%BE%A4&spm=1000.2123.3001.4430

make_image.sh是我对一些操作的集成。

先创建docker-compose.yml

再运行build.sh

环境变量配置

进入某一台容器

/ # vi ~/.bash_profile

export ENDPOINTS=etcd0:2380,etcd1:2380,etcd1:2380

alias list='etcdctl --write-out=table --endpoints=$ENDPOINTS member list'
alias status='etcdctl --write-out=table --endpoints=$ENDPOINTS endpoint status'
alias health='etcdctl --write-out=table --endpoints=$ENDPOINTS endpoint health'

source ~/.bash_profile

查看集群成员列表和各项指标：

/ # list
+------------------+---------+-------+-------------------+----------------------------+------------+
|        ID        | STATUS  | NAME  |    PEER ADDRS     |        CLIENT ADDRS        | IS LEARNER |
+------------------+---------+-------+-------------------+----------------------------+------------+
| ade526d28b1f92f7 | started | etcd1 | http://etcd1:2380 | http://192.168.1.105:23791 |      false |
| cf1d15c5d194b5c9 | started | etcd0 | http://etcd0:2380 | http://192.168.1.105:23790 |      false |
| d282ac2ce600c1ce | started | etcd2 | http://etcd2:2380 | http://192.168.1.105:23792 |      false |
+------------------+---------+-------+-------------------+----------------------------+------------+



/ # status
+------------+------------------+---------+---------+-----------+------------+-----------+------------+--------------------+--------+
|  ENDPOINT  |        ID        | VERSION | DB SIZE | IS LEADER | IS LEARNER | RAFT TERM | RAFT INDEX | RAFT APPLIED INDEX | ERRORS |
+------------+------------------+---------+---------+-----------+------------+-----------+------------+--------------------+--------+
| etcd0:2380 | cf1d15c5d194b5c9 |   3.5.0 |   25 kB |     false |      false |         2 |         33 |                 33 |        |
| etcd1:2380 | ade526d28b1f92f7 |   3.5.0 |   33 kB |     false |      false |         2 |         33 |                 33 |        |
| etcd1:2380 | ade526d28b1f92f7 |   3.5.0 |   33 kB |     false |      false |         2 |         33 |                 33 |        |
+------------+------------------+---------+---------+-----------+------------+-----------+------------+--------------------+--------+



/ # health
+------------+--------+-------------+-------+
|  ENDPOINT  | HEALTH |    TOOK     | ERROR |
+------------+--------+-------------+-------+
| etcd1:2380 |   true |  9.294369ms |       |
| etcd0:2380 |   true | 17.445272ms |       |
| etcd1:2380 |   true | 24.301921ms |       |
+------------+--------+-------------+-------+


# 添加root用户才可以开启权限校验
etcdctl user add root

# 开启权限校验
etcdctl auth enable

# 关闭权限校验
etcdctl auth disable

# 添加test前缀的读角色
etcdctl --user=root role add test_r

# 添加test前缀的读写角色
etcdctl --user=root role add test_rw

# 查看角色列表
etcdctl --user=root role list

# 添加test前缀的读用户
etcdctl --user=root user add test_r

# 给test_r角色授予/test/前缀key的读权限
etcdctl --user=root role grant-permission test_r read /test/ --prefix=true

# 给test_rw角色授予/test/前缀key的读写权限
etcdctl --user=root role grant-permission test_rw readwrite /test/ --prefix=true

# 添加test前缀的读写用户
etcdctl --user=root user add test_rw

# 查看用户列表
etcdctl --user=root user list

# 为test_r用户添加test_r角色
etcdctl --user=root user grant-role test_r test_r

# 为test_rw用户添加test_rw角色
etcdctl --user=root user grant-role test_rw test_rw
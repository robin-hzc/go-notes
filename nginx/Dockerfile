# 第一行先定义基础镜像，表示当前镜像文件是基于哪个进行编辑的.
FROM debian:latest
# 指定镜像维护者的信息.
MAINTAINER @Owen 280500569@qq.com
# 先备份原源配置文件
#RUN cp /etc/apt/sources.list /etc/apt/sources.list.bak
# 然后将debian的apt源改成中科大源
RUN sed -i 's/deb.debian.org/mirrors.ustc.edu.cn/g' /etc/apt/sources.list
RUN sed -i 's|security.debian.org/debian-security|mirrors.ustc.edu.cn/debian-security|g' /etc/apt/sources.list
RUN apt clean && apt update -y
#RUN  apt update
RUN  echo "dash dash/sh boolean false" | debconf-set-selections
RUN  DEBIAN_FRONTEND=noninteractive dpkg-reconfigure dash
# 初始版的nginx需要的依赖：gcc zlib1g zlib1g-dev openssl libssl-dev libpcre3 libpcre3-dev
# libgd-dev和libexif-dev为ngx_http_image_filter_module需要的依赖
#RUN apt -y install vim
RUN apt -y install gcc
RUN apt -y install zlib1g zlib1g-dev
RUN apt -y install openssl libssl-dev
RUN apt -y install libpcre3 libpcre3-dev
RUN apt -y install libgd-dev
RUN apt -y install libexif-dev
RUN apt -y install make
# 将dockerfile上下文目录中的ngx_http_image_filter_module.tar.gz文件拷贝进容器根目录并解压,我没有找到这个模块的下载地址，所以没开启
#COPY ngx_http_image_filter_module.tar.gz /ngx_http_image_filter_module.tar.gz
#RUN tar zxf ngx_http_image_filter_module.tar.gz
# 将dockerfile上下文目录中的nginx-1.16.1.tar.gz文件拷贝进容器根目录并解压
COPY docker-test/nginx/nginx-1.21.4.tar.gz /nginx-1.21.4.tar.gz
RUN tar zxf nginx-1.21.4.tar.gz
# 创建nginx用户，yum方式安装无需做此步骤，因为默认yum安装会自动创建nginx用户
# 咱们下面命令以nginx用户运行，因此我们需要在容器中创建nginx用户
# -M(不创建主目录) -s(不允许登录)
RUN useradd -M -s /sbin/nologin nginx
# 切换至nginx目录
WORKDIR /nginx-1.21.4/
COPY docker-test/nginx/build.sh /nginx-1.21.4/
RUN chmod +rwx build.sh&&./build.sh

# 将access.log日志重定向到docker的标准输出
# 将error.log日志重定向到docker的标准错误，这样docker logs 的时候就可以看到日志了
# ln -sf a b 表示 b链接向a
RUN ln -sf /dev/stdout /usr/local/nginx/logs/access.log && ln -sf /dev/stderr /usr/local/nginx/logs/error.log
RUN mkdir -p /usr/local/nginx/conf.d
WORKDIR /
# 删除安装使用的文件和文件夹
#RUN rm -rf ngx_http_image_filter_module.tar.gz
RUN rm -rf nginx-1.21.4.tar.gz
RUN rm -rf ngx_http_image_filter_module
RUN rm -rf nginx-1.21.4

# 定义向外暴露的端口号，多个端口用空格做间隔，启动容器时便于使用-p将此端口向宿主机端口映射
EXPOSE 80/tcp 443/tcp

# 解释 -> '-g' 'daemon off'
# nginx默认是以后台模式启动的，Docker未执行自定义的CMD之前，nginx的pid是1，
# 执行到CMD之后，nginx就在后台运行，bash或sh脚本的pid变成了1，
# 所以一旦执行完自定义CMD，nginx容器也就退出了，为了保持nginx的容器不退出，应该关闭nginx后台运行

CMD ["/usr/local/nginx/sbin/nginx", "-g", "daemon off;"]
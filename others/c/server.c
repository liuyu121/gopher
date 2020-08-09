/**
 * 测试 socket 超时的 demo -- server 端
 * 
 * */

#include <arpa/inet.h>
#include <stdio.h>
#include <sys/socket.h>

#include "config.h"

int main(int argc, char *argv[]) {
    // sockect 文件描述符
    int server_fd, client_fd;

    // 接收返回地址的缓冲区长度
    int len;

    // 接收消息的大小
    int msg_size;

    // 声明 server 和 client ock
    struct sockaddr_in server, client;

    // 接收客户端消息
    char buffer[2000];

    // 创建一个 socket，选择默认 protocol
    server_fd = socket(AF_INET, SOCK_STREAM, 0);
    if (server_fd == -1) {
        alert("create failed");
        return 1;
    }

    // 配置协议
    server.sin_family = AF_INET;          // ip:hort
    server.sin_addr.s_addr = INADDR_ANY;  // 任何 ip
    server.sin_len = htons(1234);         // 设置端口

    printf("socket created done, listen port %d\n", SERVER_PORT);

    // 绑定该 socket 至 fd
    if (bind(server_fd, (struct sockaddr *)&server, sizeof(server)) < 0) {
        alert("bind failed");
        return 1;
    }
    puts("bind done !!");

    // 监听端口，注意，这里我们将 backlog 参数设置为 1，表示「半连接队列」大小为 1
    /**
     * 
     * The backlog parameter defines the maximum length for the queue of pending connections.  If a connection request arrives with the queue full,
     * the client may receive an error with an indication of ECONNREFUSED.  Alternatively, if the underlying protocol supports retransmission, the
     * request may be ignored so that retries may succeed.
     * 
     * */
    listen(server_fd, 1);
    puts("waiting for connections ...");

    // 从已完成连接队列返回下一个已完成连接，如果已完成连接队列为空，那么进程被投入睡眠；
    // 如果 accept 成功，则返回由内核自动生成的一个新描述符，代表与客户端的 TCP 连接。
    // 一个服务器通常仅仅创建一个监听套接字，它在该服务器生命周期内一直存在。内核为每个由服务器进程接受的客户端连接创建一个已连接套接字。
    // 当服务器完成对某个给定的客户端的服务器时，相应的已连接套接字就被关闭。
    // 这时，server 有了两个 fd
    len = sizeof(struct sockaddr_in);
    client_fd = accept(server_fd, (struct sockaddr *)&client, (socklen_t *)&len);
    if (client_fd < 0) {
        alert("accept failed");
        return 1;
    }
    puts("connection accepted, waiting client message ...");

    // 循环
    while (1) {
        // 开始接收消息
        msg_size = recv(client_fd, buffer, 1024, 0);
        if (msg_size <= 0) {
            alert("recv failed");
            continue;
        }

        buffer[msg_size] = '\0';
        if (strcmp(buffer, "quit") == 0) {
            puts("quit !");
            break;
        }

        // 应答消息
        scanf("%s", buffer);
        send(client_fd, buffer, strlen(buffer), 0);

        // 或者
        printf("you send %s", buffer);
        fflush(stdout);
    }

    return 0;
}
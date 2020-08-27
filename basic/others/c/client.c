/**
 * 测试 socket 超时的 demo -- client 端
 * 
 * */

#include <arpa/inet.h>
#include <errno.h>
#include <stdio.h>
#include <string.h>
#include <sys/socket.h>
#include <sys/time.h>

#include "config.h"

void alert(const char *pStr) {
    printf("errno [%d] errmsg [%s] , detail %s\n", errno, strerror(errno), *pStr);
}

int main(int argc, char *argv[]) {
    int sock;
    struct sockaddr_in server;
    char message[1000], server_reply[2000];

    sock = socket(AF_INET, SOCK_STREAM, 0);
    if (sock == -1) {
        alert("create failed");
        return 1;
    }
    puts("socket created done");

    server.sin_addr.s_addr = inet_addr("127.0.0.1");
    server.sin_family = AF_INET;
    server.sin_port = htons(8888);

    struct timeval timeout;
    timeout.tv_sec = 10;
    timeout.tv_usec = 0;

    if (setsockopt(sock, SOL_SOCKET, SO_RCVTIMEO, (char *)&timeout,
                   sizeof(timeout)) < 0) {
        printf("setsockopt failed\n");
        return 1;
    }

    if (setsockopt(sock, SOL_SOCKET, SO_SNDTIMEO, (char *)&timeout,
                   sizeof(timeout)) < 0) {
        printf("setsockopt failed\n");
        return 1;
    }

    //Connect to remote server
    if (connect(sock, (struct sockaddr *)&server, sizeof(server)) < 0) {
        printf("errON[%d]errMsg[%s]\n", errno, strerror(errno));
        perror("connect failed. Error");
        return 1;
    }

    puts("Connected\n");

    //keep communicating with server
    while (1) {
        printf("Enter message : ");
        scanf("%s", message);

        //Send some data
        if (send(sock, message, strlen(message), 0) < 0) {
            puts("Send failed");
            return 1;
        }

        //Receive a reply from the server
        if (recv(sock, server_reply, 2000, 0) < 0) {
            puts("recv failed");
            break;
        }

        puts("Server reply :");
        puts(server_reply);
    }

    close(sock);
    return 0;
}

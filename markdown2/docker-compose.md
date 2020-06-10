# docker compose

## 多个容器组合

A docker-compose.yml file is a YAML file that defines how Docker containers should behave in production.

> docker-compose.yml

```yml
version: "3"
services:
  web:
    # replace username/repo:tag with your name and image details
    image: username/repo:tag
    deploy:
      replicas: 5
      resources:
        limits:
          cpus: "0.1"
          memory: 50M
      restart_policy:
        condition: on-failure
    ports:
      - "4000:80"
    networks:
      - webnet
networks: webnet:yml
```

```java
public class Main {
    public static void main(String[] args) {
        Date date = new Date(1590508800000L);

        System.out.println(date);
    }
}

```
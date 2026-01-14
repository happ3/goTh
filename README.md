
启动项目时需要先启动redis  redis文件夹已经直接放到项目中
```
.\redis-server.exe .\redis.windows.conf
```

然后再postMan中导入 [博客postman请求](博客postman请求.json)

或者通过curl请求

```
查询用户信息
curl --location --request GET 'http://localhost:8080/user/userInfo?id=1' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyLCJ1c2VyX25hbWUiOiJoYXBwMyIsImV4cCI6MTc2ODM4ODI1OCwiaWF0IjoxNzY4Mzg3MzU4fQ.01ohQ1jKE2K2AUZ_tk6NkHccxl2T0oRSEFc75JuQBzk' \
--data-raw ''

注册用户
curl --location --request POST 'http://localhost:8080/user/registerUser' \
--header 'Content-Type: application/json' \
--data-raw '{"username":"happ6","password":"123456789","email":"809301236@qq.com"}
'

登录

curl --location --request POST 'http://localhost:8080/user/login' \
--form 'username="happ3"' \
--form 'password="123456789"'



退出登录
curl --location --request POST 'http://localhost:8080/user/logout' \
--header 'refresh_token: 9787a24e-fe2a-41a6-a5f4-117595b478cc' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyLCJ1c2VyX25hbWUiOiJoYXBwMyIsImV4cCI6MTc2ODM3NzU3NywiaWF0IjoxNzY4Mzc2Njc3fQ.EfyuxK5IWP9q8igV8MOYJP9Q3tC8xGEXkzASDzDQNwQ'



刷新token
curl --location --request POST 'http://localhost:8080/user/refreshToken' \
--header 'refresh_token: ee845e3f-7eb4-49bf-a55a-8de6e5ea0099' \
--form 'username="happ3"' \
--form 'password="123456789"'


用户认证
curl --location --request POST 'http://localhost:8080/user/authentication' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyLCJ1c2VyX25hbWUiOiJoYXBwMyIsImV4cCI6MTc2ODM4ODI1OCwiaWF0IjoxNzY4Mzg3MzU4fQ.01ohQ1jKE2K2AUZ_tk6NkHccxl2T0oRSEFc75JuQBzk' \
--header 'Content-Type: application/json' \
--data-raw '{"id":3}
'

新增博客
curl --location --request POST 'http://localhost:8080/posts/addPosts' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyLCJ1c2VyX25hbWUiOiJoYXBwMyIsImV4cCI6MTc2ODM4NTExNSwiaWF0IjoxNzY4Mzg0MjE1fQ.DNkWr3Q-fXQQgrk5fDRAqlDUYn-PanQS4HGAnhdpjMA' \
--header 'Content-Type: application/json' \
--data-raw '{"content":"这是一篇博客9，很好看","title":"博客8"}
'

查看博客详情

curl --location --request POST 'http://localhost:8080/posts/getPosts' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyLCJ1c2VyX25hbWUiOiJoYXBwMyIsImV4cCI6MTc2ODM4NDEzMiwiaWF0IjoxNzY4MzgzMjMyfQ.7j-EUHdnEbZ-Qa_SYy22pMt1W1mizi6op1RuqxnbUaE' \
--header 'Content-Type: application/json' \
--data-raw '{"id":1}
'

删除博客
curl --location --request POST 'http://localhost:8080/posts/delPost' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyLCJ1c2VyX25hbWUiOiJoYXBwMyIsImV4cCI6MTc2ODM4NTExNSwiaWF0IjoxNzY4Mzg0MjE1fQ.DNkWr3Q-fXQQgrk5fDRAqlDUYn-PanQS4HGAnhdpjMA' \
--header 'Content-Type: application/json' \
--data-raw '{"id":6}
'

更新博客
curl --location --request POST 'http://localhost:8080/posts/updatePost' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyLCJ1c2VyX25hbWUiOiJoYXBwMyIsImV4cCI6MTc2ODM4NTExNSwiaWF0IjoxNzY4Mzg0MjE1fQ.DNkWr3Q-fXQQgrk5fDRAqlDUYn-PanQS4HGAnhdpjMA' \
--header 'Content-Type: application/json' \
--data-raw '{"id":3,"content":"这是一篇博客2221，很好看","title":"博客8"}
'

新增评论
curl --location --request POST 'http://localhost:8080/comments/addComments' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyLCJ1c2VyX25hbWUiOiJoYXBwMyIsImV4cCI6MTc2ODM4NjA3OCwiaWF0IjoxNzY4Mzg1MTc4fQ.GJdKmuQP_Tslkap-PlDSnj6wSfhmyxV9OIh717jncho' \
--header 'Content-Type: application/json' \
--data-raw '{"content":"评论这是一篇博客9，很好看","postId":2}
'

查看文章下所有评论
curl --location --request POST 'http://localhost:8080/comments/findCommentByPostId' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyLCJ1c2VyX25hbWUiOiJoYXBwMyIsImV4cCI6MTc2ODM4NjA3OCwiaWF0IjoxNzY4Mzg1MTc4fQ.GJdKmuQP_Tslkap-PlDSnj6wSfhmyxV9OIh717jncho' \
--header 'Content-Type: application/json' \
--data-raw '{"postId":2}
'




```

建表语句
```
/*
SQLyog Ultimate v12.2.6 (64 bit)
MySQL - 8.0.30 : Database - vlog
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`vlog` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;

USE `vlog`;

/*Table structure for table `comments` */

DROP TABLE IF EXISTS `comments`;

CREATE TABLE `comments` (
  `id` int NOT NULL AUTO_INCREMENT,
  `content` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `userId` int DEFAULT NULL,
  `postId` int DEFAULT NULL,
  `createAt` timestamp NULL DEFAULT NULL,
  `updateAt` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='博客评论';

/*Data for the table `comments` */

insert  into `comments`(`id`,`content`,`userId`,`postId`,`createAt`,`updateAt`) values 
(1,'博客1写的真好',1,1,'2026-01-13 15:46:55',NULL),
(2,'博客2写的也不错',2,2,'2026-01-13 15:47:09',NULL),
(3,'博客1我觉得一般',2,1,'2026-01-12 15:47:27',NULL),
(4,'博客2我也觉得写的很好',1,2,'2026-01-15 15:47:59',NULL),
(5,'评论这是一篇博客9，很好看',2,2,NULL,NULL);

/*Table structure for table `posts` */

DROP TABLE IF EXISTS `posts`;

CREATE TABLE `posts` (
  `id` int NOT NULL AUTO_INCREMENT,
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `content` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `userId` int DEFAULT NULL,
  `createAt` timestamp NULL DEFAULT NULL,
  `updateAt` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='博客';

/*Data for the table `posts` */

insert  into `posts`(`id`,`title`,`content`,`userId`,`createAt`,`updateAt`) values 
(1,'博客1','博客1内容',1,'2026-01-13 15:48:36',NULL),
(2,'博客2','博客2内容',1,'2026-01-13 15:48:39',NULL),
(3,'博客8','这是一篇博客2221，很好看',2,'2026-01-14 17:15:03','2026-01-14 17:15:03');

/*Table structure for table `user` */

DROP TABLE IF EXISTS `user`;

CREATE TABLE `user` (
  `id` int NOT NULL AUTO_INCREMENT,
  `username` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `pwz` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `email` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `auth` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '0' COMMENT '0未认证 1已认证',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;

/*Data for the table `user` */

insert  into `user`(`id`,`username`,`pwz`,`email`,`auth`) values 
(1,'xiaofang','$2a$10$NZkUO.PdgC7HJ4Kwh6rrhec54rJA8caUpaDcQgNXbPtrgF2qqqzmW','809301236@qq.com','1'),
(2,'happ3','$2a$10$Hq0O2uuAbpSGJwhm4MFvxug4CYSNEW7dnjbFFgSLMr4qpSTdBrxsK','809301236@qq.com','1'),
(6,'happ6','$2a$10$.oNmK8EEqNikTJUozLCFjeFC8UZPBJOoGGSsaLKqXwBgHY1yM6tom','809301236@qq.com','0');

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

```
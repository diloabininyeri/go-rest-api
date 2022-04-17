-- MySQL dump 10.13  Distrib 8.0.28, for Linux (x86_64)
--
-- Host: localhost    Database: payments
-- ------------------------------------------------------
-- Server version	8.0.28

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `balances`
--
create  database payments;
use payments;
DROP TABLE IF EXISTS `balances`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `balances` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `amount` int NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=55 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping sql for table `balances`
--

LOCK TABLES `balances` WRITE;
/*!40000 ALTER TABLE `balances` DISABLE KEYS */;
INSERT INTO `balances` VALUES (1,2,100,'2022-04-14 14:23:22',NULL),(2,5,200,'2022-04-14 14:23:22',NULL),(3,3,455,'2022-04-14 14:23:22',NULL),(4,7,250,'2022-04-15 08:36:06',NULL),(5,7,250,'2022-04-15 08:36:11',NULL),(6,7,250,'2022-04-15 08:36:13',NULL),(7,7,250,'2022-04-15 08:36:14',NULL),(8,7,250,'2022-04-15 08:36:50',NULL),(9,7,250,'2022-04-15 09:03:03',NULL),(10,2,-250,'2022-04-15 09:26:45',NULL),(11,8,158,'2022-04-15 09:50:40',NULL),(12,8,158,'2022-04-15 09:50:41',NULL),(13,8,158,'2022-04-15 09:50:42',NULL),(14,8,-158,'2022-04-15 09:50:55',NULL),(15,9,158,'2022-04-15 11:46:10',NULL),(16,9,-5,'2022-04-15 11:46:26',NULL),(17,9,-5,'2022-04-15 11:46:30',NULL),(18,9,-5,'2022-04-15 12:07:12',NULL),(19,9,-5,'2022-04-15 12:07:25',NULL),(20,8,-158,'2022-04-15 12:16:31',NULL),(21,8,-158,'2022-04-15 12:23:40',NULL),(22,9,-158,'2022-04-15 12:23:48',NULL),(23,9,-158,'2022-04-15 12:24:27',NULL),(24,9,-158,'2022-04-15 12:29:11',NULL),(25,9,-158,'2022-04-15 12:54:19',NULL),(26,9,-158,'2022-04-15 12:55:41',NULL),(27,9,-158,'2022-04-15 12:55:46',NULL),(28,9,-158,'2022-04-15 12:57:07',NULL),(29,9,-158,'2022-04-15 12:57:16',NULL),(30,9,-158,'2022-04-15 12:58:02',NULL),(31,9,-158,'2022-04-15 12:58:16',NULL),(32,9,-158,'2022-04-15 12:58:19',NULL),(33,19,200,'2022-04-15 12:58:33',NULL),(34,19,200,'2022-04-15 12:59:33',NULL),(35,19,200,'2022-04-15 12:59:36',NULL),(36,12,200,'2022-04-15 12:59:44',NULL),(37,12,200,'2022-04-15 13:01:35',NULL),(38,12,200,'2022-04-15 13:01:56',NULL),(39,12,200,'2022-04-15 13:02:18',NULL),(40,12,200,'2022-04-15 13:03:38',NULL),(41,12,200,'2022-04-15 13:03:49',NULL),(42,12,200,'2022-04-15 13:03:54',NULL),(43,12,200,'2022-04-15 13:12:30',NULL),(44,12,200,'2022-04-15 13:12:55',NULL),(45,12,200,'2022-04-15 13:12:58',NULL),(46,12,200,'2022-04-15 13:14:17',NULL),(47,12,200,'2022-04-15 13:14:33',NULL),(48,12,200,'2022-04-15 13:14:50',NULL),(49,13,200,'2022-04-15 13:15:37',NULL),(50,13,200,'2022-04-15 13:15:51',NULL),(51,13,200,'2022-04-15 13:15:56',NULL),(52,13,200,'2022-04-15 13:15:59',NULL),(53,13,-250,'2022-04-15 13:19:54',NULL),(54,13,-125,'2022-04-15 13:20:02',NULL);
/*!40000 ALTER TABLE `balances` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-04-17  8:57:16
-- MySQL dump 10.13  Distrib 8.0.23, for osx10.16 (x86_64)
--
-- Host: localhost    Database: telkomselbe
-- ------------------------------------------------------
-- Server version	8.0.23

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


CREATE DATABASE telkomselbe;

USE telkomselbe;
--
-- Table structure for table `customer`
--

DROP TABLE IF EXISTS `customer`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `customer` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `fullname` varchar(100) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `msisdn` bigint unsigned NOT NULL,
  `referral_code` varchar(10) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `created` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `customer_UN` (`msisdn`),
  UNIQUE KEY `customer_ref_UN` (`referral_code`)
) ENGINE=InnoDB AUTO_INCREMENT=93 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `customer`
--

LOCK TABLES `customer` WRITE;
/*!40000 ALTER TABLE `customer` DISABLE KEYS */;
INSERT INTO `customer` VALUES (1,'asdas',123123,NULL,'2021-08-10 18:01:35',NULL),(12,'asdas',1231234,'TEST125','2021-08-10 18:01:35',NULL),(13,'asdas',12312345,'TEST124','2021-08-10 18:01:35',NULL),(14,'asdas',6282124911591,'TEST123','2021-08-10 18:01:35',NULL),(15,'Test',6282123333333,NULL,'2021-08-10 18:01:35',NULL),(22,'asdas',6282124911592,NULL,'2021-08-10 18:01:35',NULL),(23,'asdas',6282124911593,NULL,'2021-08-10 18:01:35',NULL),(41,'asdas',6282124911594,NULL,'2021-08-10 18:01:35',NULL),(42,'asdas',6282124911595,NULL,'2021-08-10 18:01:35',NULL),(47,'asdas',628212491159,NULL,'2021-08-10 18:01:35',NULL),(50,'asdas',628212491152,'5FBBD','2021-08-10 18:01:35',NULL),(68,'asdas',628212491153,NULL,'2021-08-11 03:18:04',NULL),(69,'asdas',628212491154,'6P84O','2021-08-11 03:18:18',NULL),(71,'asdas',628212491155,'1HUWE','2021-08-11 03:42:00',NULL),(72,'asdas',628212491156,'A6ATM','2021-08-11 03:42:06',NULL),(73,'asdas',628212491157,'ZXKQA','2021-08-11 03:42:11',NULL),(75,'asdas',628212491158,'YB4J3','2021-08-11 03:54:11',NULL),(78,'asdas',628212491160,'ZMUM3','2021-08-11 03:55:15',NULL),(80,'asdas',628212491161,'I00K5','2021-08-11 03:59:58',NULL),(82,'asdas',628212491162,'SRUIG','2021-08-11 04:42:12',NULL),(83,'asdas',628212491163,'LFO2K','2021-08-11 04:44:03',NULL),(85,'asdas',628212491164,'VFHV6','2021-08-11 04:44:34',NULL),(86,'asdas',628212491165,'YC8AN','2021-08-11 04:47:53',NULL),(88,'asdas',628212491166,'PFWL7','2021-08-11 04:49:52',NULL),(89,'asdas',628212491167,'4YWTI','2021-08-11 04:50:42',NULL),(90,'asdas',628212491168,'Z39NH','2021-08-11 05:22:17',NULL),(92,'asdas',628212491169,'OWB42','2021-08-11 05:22:32',NULL);
/*!40000 ALTER TABLE `customer` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `provisioning_reward`
--

DROP TABLE IF EXISTS `provisioning_reward`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `provisioning_reward` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `customer_id` bigint unsigned DEFAULT NULL,
  `offer_id` varchar(20) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `created` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `status` varchar(20) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `yearmonth` mediumint unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `provisioning_reward_UN` (`customer_id`,`yearmonth`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `provisioning_reward`
--

LOCK TABLES `provisioning_reward` WRITE;
/*!40000 ALTER TABLE `provisioning_reward` DISABLE KEYS */;
INSERT INTO `provisioning_reward` VALUES (1,14,'12 GB','2021-08-11 10:52:08','ONPROGRESS',202108),(2,88,'12 GB','2021-08-11 10:52:08','ONPROGRESS',202108);
/*!40000 ALTER TABLE `provisioning_reward` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `referral_mapping`
--

DROP TABLE IF EXISTS `referral_mapping`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `referral_mapping` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `from_id` bigint unsigned DEFAULT NULL,
  `to_id` bigint unsigned DEFAULT NULL,
  `created` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `referral_mapping_UN` (`from_id`,`to_id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `referral_mapping`
--

LOCK TABLES `referral_mapping` WRITE;
/*!40000 ALTER TABLE `referral_mapping` DISABLE KEYS */;
INSERT INTO `referral_mapping` VALUES (1,14,80,'2021-08-11 03:59:58'),(2,14,82,'2021-08-11 04:42:12'),(3,14,83,'2021-08-11 04:44:03'),(4,0,85,'2021-08-11 04:44:34'),(5,0,86,'2021-08-11 04:47:53'),(6,14,88,'2021-08-11 04:49:52'),(7,88,89,'2021-08-11 04:50:42'),(8,88,90,'2021-08-11 05:22:17'),(9,88,92,'2021-08-11 05:22:32');
/*!40000 ALTER TABLE `referral_mapping` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `rule_reward`
--

DROP TABLE IF EXISTS `rule_reward`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `rule_reward` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `referral_min_total` bigint unsigned DEFAULT NULL,
  `referral_max_total` bigint unsigned DEFAULT NULL,
  `offer_id` varchar(50) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `created` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `rule_reward`
--

LOCK TABLES `rule_reward` WRITE;
/*!40000 ALTER TABLE `rule_reward` DISABLE KEYS */;
INSERT INTO `rule_reward` VALUES (1,1,1,'2 GB','2021-08-11 07:52:22',NULL),(2,2,5,'12 GB','2021-08-11 07:52:22',NULL),(3,6,9999999,'20 GB','2021-08-11 07:52:22',NULL);
/*!40000 ALTER TABLE `rule_reward` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `summary_referral_monthly`
--

DROP TABLE IF EXISTS `summary_referral_monthly`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `summary_referral_monthly` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `customer_id` bigint unsigned DEFAULT NULL,
  `yearmonth` mediumint unsigned DEFAULT NULL,
  `referral_total` bigint unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `summary_referral_monthly_UN` (`yearmonth`,`customer_id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `summary_referral_monthly`
--

LOCK TABLES `summary_referral_monthly` WRITE;
/*!40000 ALTER TABLE `summary_referral_monthly` DISABLE KEYS */;
INSERT INTO `summary_referral_monthly` VALUES (1,1,202102,0),(2,14,202108,2),(3,0,202108,0),(4,86,202108,0),(5,88,202108,2),(6,89,202108,0),(7,90,202108,0),(8,92,202108,0);
/*!40000 ALTER TABLE `summary_referral_monthly` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping routines for database 'telkomselbe'
--
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2021-08-11 17:58:30

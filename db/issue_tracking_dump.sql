mysqldump: [Warning] Using a password on the command line interface can be insecure.
-- MySQL dump 10.13  Distrib 8.0.42, for Linux (x86_64)
--
-- Host: localhost    Database: issue_tracking_system
-- ------------------------------------------------------
-- Server version	8.0.42

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
-- Table structure for table `issues`
--

DROP TABLE IF EXISTS `issues`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `issues` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `key` varchar(255) NOT NULL,
  `summary` varchar(255) NOT NULL,
  `description` text NOT NULL,
  `project_key` varchar(255) NOT NULL,
  `reporter` varchar(255) NOT NULL,
  `assignee` varchar(255) NOT NULL,
  `status` enum('open','in_progress','resolved') NOT NULL DEFAULT 'open',
  `issueType` enum('bug','task','story') NOT NULL DEFAULT 'task',
  `createdAt` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updatedAt` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `sprint_id` int unsigned DEFAULT NULL,
  `started_at` timestamp NULL DEFAULT NULL,
  `finished_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `key` (`key`),
  KEY `project_key` (`project_key`),
  KEY `fk_issues_sprint` (`sprint_id`),
  CONSTRAINT `fk_issues_sprint` FOREIGN KEY (`sprint_id`) REFERENCES `sprints` (`id`) ON DELETE SET NULL,
  CONSTRAINT `issues_ibfk_1` FOREIGN KEY (`project_key`) REFERENCES `projects` (`project_key`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=34 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `issues`
--

LOCK TABLES `issues` WRITE;
/*!40000 ALTER TABLE `issues` DISABLE KEYS */;
INSERT INTO `issues` VALUES (1,'TP8-001','test summary','test descriptions','TP8','john.doe@example.com','jane.doe@example.com','open','bug','2025-04-15 16:38:17','2025-04-15 16:38:17',NULL,NULL,NULL),(3,'TP7-001','test summary','test descriptions','TP7','john.doe@example.com','jane.doe@example.com','open','bug','2025-04-15 16:41:26','2025-04-15 16:41:26',NULL,NULL,NULL),(4,'TP7-002','test summary','test descriptions','TP7','john.doe@example.com','jane.doe@example.com','open','bug','2025-04-27 15:28:36','2025-04-28 16:32:02',NULL,NULL,NULL),(5,'TP7-003','test summary','test descriptions','TP7','john.doe@example.com','jane.doe@example.com','in_progress','bug','2025-04-27 16:28:24','2025-04-27 16:28:24',NULL,NULL,NULL),(6,'TP7-004','test summary','test descriptions','TP7','john.doe@example.com','jane.doe@example.com','in_progress','bug','2025-04-27 16:28:48','2025-04-27 16:28:48',NULL,NULL,NULL),(7,'TP7-005','test summary','test descriptions','TP7','john.doe@example.com','jane.doe@example.com','in_progress','bug','2025-04-27 18:24:03','2025-04-27 18:24:03',NULL,NULL,NULL),(8,'TP7-006','test summary','test descriptions','TP7','john.doe@example.com','jane.doe@example.com','in_progress','bug','2025-04-28 10:59:44','2025-04-28 10:59:44',NULL,NULL,NULL),(9,'TP7-007','test summary','test descriptions','TP7','john.doe@example.com','jane.doe@example.com','in_progress','bug','2025-04-28 14:02:45','2025-04-28 14:02:45',NULL,NULL,NULL),(10,'TP7-008','test summary','test descriptions','TP7','john.doe@example.com','jane.doe@example.com','in_progress','bug','2025-04-28 14:03:00','2025-04-28 14:03:00',NULL,NULL,NULL),(11,'TP7-009','test summary','test descriptions','TP7','john.doe@example.com','jane.doe@example.com','in_progress','bug','2025-04-30 18:53:13','2025-04-30 18:53:13',NULL,NULL,NULL),(12,'TP7-010','test summary','test descriptions','TP7','john.doe@example.com','jane.doe@example.com','in_progress','bug','2025-04-30 18:53:35','2025-04-30 18:53:35',NULL,NULL,NULL),(13,'TP7-011','test summary','test descriptions','TP7','john.doe@example.com','jane.doe@example.com','in_progress','bug','2025-04-30 18:58:07','2025-04-30 18:58:07',NULL,NULL,NULL),(28,'TP7-012','test summary','test descriptions','TP7','john.doe@example.com','jane.doe@example.com','in_progress','bug','2025-04-30 19:58:05','2025-04-30 19:58:05',NULL,NULL,NULL),(31,'TP(-001','dfdsf','asdfsfd','TP(','sdfsd','dsfdsf','resolved','bug','2025-04-30 20:05:48','2025-05-01 10:08:35',NULL,'2025-04-30 20:06:19','2025-04-30 20:16:16'),(32,'TP(-002','test summary','test descriptions','TP(','john.doe@example.com','jane.doe@example.com','in_progress','bug','2025-05-01 10:34:02','2025-05-01 10:34:02',NULL,NULL,NULL),(33,'TP(-003','test summary','test descriptions','TP(','john.doe@example.com','jane.doe@example.com','resolved','bug','2025-05-01 10:34:34','2025-05-01 10:35:00',NULL,'2025-05-01 10:34:46','2025-05-01 10:35:00');
/*!40000 ALTER TABLE `issues` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `project_assignments`
--

DROP TABLE IF EXISTS `project_assignments`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `project_assignments` (
  `project_id` int NOT NULL,
  `user_id` int unsigned NOT NULL,
  `role` varchar(255) DEFAULT 'member',
  `assigned_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`project_id`,`user_id`),
  KEY `user_id` (`user_id`),
  CONSTRAINT `project_assignments_ibfk_1` FOREIGN KEY (`project_id`) REFERENCES `projects` (`id`) ON DELETE CASCADE,
  CONSTRAINT `project_assignments_ibfk_2` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `project_assignments`
--

LOCK TABLES `project_assignments` WRITE;
/*!40000 ALTER TABLE `project_assignments` DISABLE KEYS */;
INSERT INTO `project_assignments` VALUES (1,1,'member','2025-04-28 20:20:31'),(1,2,'member','2025-04-29 11:37:58'),(1,3,'developer','2025-04-29 12:04:06'),(2,1,'member','2025-04-29 11:58:14'),(2,2,'member','2025-04-29 11:40:45'),(2,3,'member','2025-04-29 12:05:58'),(4,2,'scrum_master','2025-04-30 20:58:20');
/*!40000 ALTER TABLE `project_assignments` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `project_scope`
--

DROP TABLE IF EXISTS `project_scope`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `project_scope` (
  `id` int NOT NULL AUTO_INCREMENT,
  `project_key` varchar(255) NOT NULL,
  `scope_id` int NOT NULL,
  PRIMARY KEY (`id`),
  KEY `project_key` (`project_key`),
  KEY `scope_id` (`scope_id`),
  CONSTRAINT `project_scope_ibfk_1` FOREIGN KEY (`project_key`) REFERENCES `projects` (`project_key`) ON DELETE CASCADE,
  CONSTRAINT `project_scope_ibfk_2` FOREIGN KEY (`scope_id`) REFERENCES `scopes` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `project_scope`
--

LOCK TABLES `project_scope` WRITE;
/*!40000 ALTER TABLE `project_scope` DISABLE KEYS */;
INSERT INTO `project_scope` VALUES (5,'TP8',6),(6,'TP7',7),(7,'TP7',8),(12,'TP(',6);
/*!40000 ALTER TABLE `project_scope` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `projects`
--

DROP TABLE IF EXISTS `projects`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `projects` (
  `id` int NOT NULL AUTO_INCREMENT,
  `project_key` varchar(255) NOT NULL,
  `name` varchar(255) NOT NULL,
  `description` text NOT NULL,
  `project_lead` int unsigned DEFAULT NULL,
  `issue_count` int NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `wip_limit` int NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `project_key` (`project_key`),
  KEY `project_lead` (`project_lead`),
  CONSTRAINT `projects_ibfk_1` FOREIGN KEY (`project_lead`) REFERENCES `users` (`id`) ON DELETE SET NULL
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `projects`
--

LOCK TABLES `projects` WRITE;
/*!40000 ALTER TABLE `projects` DISABLE KEYS */;
INSERT INTO `projects` VALUES (1,'TP7','Test Project 7','Test project description',1,12,'2025-03-28 13:16:21',0),(2,'TP8','Test Project 7','Test project description',1,1,'2025-04-15 16:23:11',2),(4,'TP(','ese','df',1,3,'2025-04-30 20:04:17',10);
/*!40000 ALTER TABLE `projects` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `schema_migrations`
--

DROP TABLE IF EXISTS `schema_migrations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `schema_migrations` (
  `version` bigint NOT NULL,
  `dirty` tinyint(1) NOT NULL,
  PRIMARY KEY (`version`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `schema_migrations`
--

LOCK TABLES `schema_migrations` WRITE;
/*!40000 ALTER TABLE `schema_migrations` DISABLE KEYS */;
INSERT INTO `schema_migrations` VALUES (20250428195934,0);
/*!40000 ALTER TABLE `schema_migrations` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `scopes`
--

DROP TABLE IF EXISTS `scopes`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `scopes` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `description` text NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `scopes`
--

LOCK TABLES `scopes` WRITE;
/*!40000 ALTER TABLE `scopes` DISABLE KEYS */;
INSERT INTO `scopes` VALUES (1,'Cross-functional Initiative','This scope involves multiple teams across different business areas.','2025-04-15 16:05:47'),(4,'Cross-functional Initiative 2','This scope involves multiple teams across different business areas.','2025-04-15 16:07:03'),(6,'Cross-functional Initiative 3','This scope involves multiple teams across different business areas.','2025-04-15 16:18:12'),(7,'Test Scope from UI','LOLOL','2025-04-30 08:05:08'),(8,'Sprint Alpha','Spring Alpha Description','2025-04-30 08:17:40'),(9,'TEs','sdf','2025-04-30 08:27:44'),(10,'Test','add','2025-04-30 08:44:10');
/*!40000 ALTER TABLE `scopes` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sprint_issues`
--

DROP TABLE IF EXISTS `sprint_issues`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sprint_issues` (
  `issue_id` int unsigned NOT NULL,
  `sprint_id` int unsigned NOT NULL,
  PRIMARY KEY (`issue_id`,`sprint_id`),
  KEY `sprint_id` (`sprint_id`),
  CONSTRAINT `sprint_issues_ibfk_1` FOREIGN KEY (`issue_id`) REFERENCES `issues` (`id`) ON DELETE CASCADE,
  CONSTRAINT `sprint_issues_ibfk_2` FOREIGN KEY (`sprint_id`) REFERENCES `sprints` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sprint_issues`
--

LOCK TABLES `sprint_issues` WRITE;
/*!40000 ALTER TABLE `sprint_issues` DISABLE KEYS */;
INSERT INTO `sprint_issues` VALUES (6,2);
/*!40000 ALTER TABLE `sprint_issues` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sprints`
--

DROP TABLE IF EXISTS `sprints`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `sprints` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `description` text,
  `start_date` date NOT NULL,
  `end_date` date NOT NULL,
  `project_key` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `project_key` (`project_key`),
  CONSTRAINT `sprints_ibfk_1` FOREIGN KEY (`project_key`) REFERENCES `projects` (`project_key`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sprints`
--

LOCK TABLES `sprints` WRITE;
/*!40000 ALTER TABLE `sprints` DISABLE KEYS */;
INSERT INTO `sprints` VALUES (2,'Sprint 1','This is the first sprint focusing on core features.','2025-05-01','2025-05-14','TP7','2025-04-30 18:28:11','2025-04-30 18:28:11');
/*!40000 ALTER TABLE `sprints` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `standups`
--

DROP TABLE IF EXISTS `standups`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `standups` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `project_key` varchar(255) NOT NULL,
  `start_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `end_time` timestamp NULL DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `project_key` (`project_key`),
  CONSTRAINT `standups_ibfk_1` FOREIGN KEY (`project_key`) REFERENCES `projects` (`project_key`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=23 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `standups`
--

LOCK TABLES `standups` WRITE;
/*!40000 ALTER TABLE `standups` DISABLE KEYS */;
INSERT INTO `standups` VALUES (1,'TP8','2025-04-15 16:23:22',NULL,'2025-04-15 16:23:22'),(2,'TP7','2025-04-27 15:47:16','2025-04-27 16:09:27','2025-04-27 15:47:16'),(3,'TP7','2025-04-27 18:03:39',NULL,'2025-04-27 18:03:39'),(4,'TP7','2025-04-27 18:03:44','2025-04-27 18:04:08','2025-04-27 18:03:44'),(5,'TP7','2025-04-27 18:04:12',NULL,'2025-04-27 18:04:12'),(6,'TP7','2025-04-27 18:15:45',NULL,'2025-04-27 18:15:45'),(7,'TP7','2025-04-27 18:20:47',NULL,'2025-04-27 18:20:47'),(8,'TP7','2025-04-27 18:24:06',NULL,'2025-04-27 18:24:06'),(9,'TP7','2025-04-28 10:36:41',NULL,'2025-04-28 10:36:41'),(10,'TP7','2025-04-28 10:40:21','2025-04-28 10:44:24','2025-04-28 10:40:21'),(11,'TP7','2025-04-28 10:55:20','2025-04-28 10:55:47','2025-04-28 10:55:20'),(14,'TP7','2025-04-28 10:59:27','2025-04-28 10:59:30','2025-04-28 10:59:27'),(15,'TP7','2025-04-28 10:59:46','2025-04-28 10:59:49','2025-04-28 10:59:46'),(16,'TP7','2025-04-28 13:42:49','2025-04-28 13:43:29','2025-04-28 13:42:49'),(17,'TP7','2025-04-28 13:57:46','2025-04-28 13:58:27','2025-04-28 13:57:46'),(18,'TP7','2025-04-28 14:02:31','2025-04-28 14:02:55','2025-04-28 14:02:31'),(19,'TP7','2025-04-28 14:03:02',NULL,'2025-04-28 14:03:02'),(20,'TP7','2025-04-28 14:03:35',NULL,'2025-04-28 14:03:35'),(22,'TP7','2025-04-30 19:23:20','2025-04-30 19:23:23','2025-04-30 19:23:20');
/*!40000 ALTER TABLE `standups` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `firstName` varchar(255) NOT NULL,
  `lastName` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `createdAt` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `email` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (1,'Max','Marston','max.marston@test.com','$2a$10$vzdUsM063NeXALZkDt7uM.rOIqyUaTJU2kSwk.suiLWqhOElc3sDO','2025-03-28 13:16:09'),(2,'Holly','Briggs','holly.briggs@test.com','$2a$10$UKAcq.6bYJu98W4W28sxAuq.kJZB8m.ZXxu6OINDtNvMiYipAdTjS','2025-04-29 11:34:48'),(3,'Greg','Hall','greg.hall@test.com','$2a$10$r1D9hBjj8HxnJquO2.HZFufB/3alQHeXwbbYKwTImbIkpW6c7FBuG','2025-04-29 12:03:56');
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2025-05-07 21:40:24

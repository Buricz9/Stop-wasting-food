-- --------------------------------------------------------
-- Host:                         127.0.0.1
-- Wersja serwera:               11.3.2-MariaDB - mariadb.org binary distribution
-- Serwer OS:                    Win64
-- HeidiSQL Wersja:              12.6.0.6765
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


-- Zrzut struktury bazy danych recipedatabase
CREATE DATABASE IF NOT EXISTS `recipedatabase` /*!40100 DEFAULT CHARACTER SET armscii8 COLLATE armscii8_general_ci */;
USE `recipedatabase`;

-- Zrzut struktury tabela recipedatabase.availableingredients
CREATE TABLE IF NOT EXISTS `availableingredients` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `recipe_id` int(11) NOT NULL,
  `ingredient_name` varchar(255) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `availableingredients_ibfk_1` (`recipe_id`),
  CONSTRAINT `availableingredients_ibfk_1` FOREIGN KEY (`recipe_id`) REFERENCES `recipes` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=armscii8 COLLATE=armscii8_general_ci;

-- Zrzucanie danych dla tabeli recipedatabase.availableingredients: ~0 rows (około)

-- Zrzut struktury tabela recipedatabase.history_of_inputs
CREATE TABLE IF NOT EXISTS `history_of_inputs` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `historyOfIngredients` text NOT NULL,
  `historyOfNumber` int(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=50 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Zrzucanie danych dla tabeli recipedatabase.history_of_inputs: ~0 rows (około)

-- Zrzut struktury tabela recipedatabase.missingingredients
CREATE TABLE IF NOT EXISTS `missingingredients` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `recipe_id` int(11) NOT NULL,
  `ingredient_name` text NOT NULL,
  PRIMARY KEY (`id`),
  KEY `FK_missingingredients_recipes` (`recipe_id`),
  CONSTRAINT `FK_missingingredients_recipes` FOREIGN KEY (`recipe_id`) REFERENCES `recipes` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=armscii8 COLLATE=armscii8_general_ci;

-- Zrzucanie danych dla tabeli recipedatabase.missingingredients: ~0 rows (około)

-- Zrzut struktury tabela recipedatabase.recipes
CREATE TABLE IF NOT EXISTS `recipes` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `history_id` int(11) NOT NULL,
  `title` varchar(255) NOT NULL,
  `carbs` float NOT NULL,
  `proteins` float NOT NULL,
  `calories` float NOT NULL,
  PRIMARY KEY (`id`),
  KEY `FK_recipes_history_of_inputs` (`history_id`),
  CONSTRAINT `FK_recipes_history_of_inputs` FOREIGN KEY (`history_id`) REFERENCES `history_of_inputs` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB AUTO_INCREMENT=29 DEFAULT CHARSET=armscii8 COLLATE=armscii8_general_ci;

-- Zrzucanie danych dla tabeli recipedatabase.recipes: ~0 rows (około)

/*!40103 SET TIME_ZONE=IFNULL(@OLD_TIME_ZONE, 'system') */;
/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=IFNULL(@OLD_SQL_NOTES, 1) */;

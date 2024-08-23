-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Aug 23, 2024 at 05:47 AM
-- Server version: 10.4.32-MariaDB
-- PHP Version: 8.2.12

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `bukuin`
--

-- --------------------------------------------------------

--
-- Table structure for table `books`
--

CREATE TABLE `books` (
  `id_book` int(11) NOT NULL,
  `id_user` int(10) NOT NULL,
  `title` varchar(255) NOT NULL,
  `author` varchar(255) NOT NULL,
  `description` varchar(255) NOT NULL,
  `launch_year` year(4) NOT NULL,
  `isbn` varchar(20) NOT NULL,
  `cover_image_url` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `books`
--

INSERT INTO `books` (`id_book`, `id_user`, `title`, `author`, `description`, `launch_year`, `isbn`, `cover_image_url`) VALUES
(5, 11, 'To Kill a Mockingbird', 'Harper Lee', 'A novel about racial injustice in the Deep South.', '1960', '9780060935467', 'https://m.media-amazon.com/images/I/71FxgtFKcQL.jpg'),
(8, 11, 'Test', 'Mighty Doyok', 'Doyok Otoy Ali', '2024', '9789909911231', '-'),
(9, 11, '1984', 'George Orwell', ' A dystopian novel set in a totalitarian society ruled by the Party and its leader, Big Brother.', '1949', '9780451524935', 'https://m.media-amazon.com/images/I/61NAx5pd6XL._AC_UF1000,1000_QL80_.jpg'),
(11, 11, 'The Great Gatsby', 'F. Scott Fitzgerald', 'A novel set in the Roaring Twenties that critiques the American Dream through the life of Jay Gatsby.', '1925', '9780743273565', 'https://upload.wikimedia.org/wikipedia/id/2/26/TheGreatGatsby2012Poster.jpg'),
(22, 11, 'Dilan 1991', 'Pidi Baiq', 'The sequel to \"Dilan 1990,\" this book continues the love story between Dilan and Milea, exploring their relationship and the challenges they face.', '2022', '9786026625664', 'https://ebooks.gramedia.com/ebook-covers/31755/big_covers/ID_MIZ2016MTH03DDADT1_B.jpg'),
(23, 11, 'The Lying Game', 'Sabrina Benaim', 'A powerful collection of poetry exploring themes of identity, mental health, and personal growth.', '2022', '9786237049512', '-'),
(24, 11, 'Tes Buku', 'Kibo', 'Tes', '2000', '9787878787878', '-');

-- --------------------------------------------------------

--
-- Table structure for table `roles`
--

CREATE TABLE `roles` (
  `id_role` int(11) NOT NULL,
  `role_name` varchar(20) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `roles`
--

INSERT INTO `roles` (`id_role`, `role_name`) VALUES
(1, 'Admin'),
(2, 'User');

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id_user` int(11) NOT NULL,
  `id_role` int(11) DEFAULT NULL,
  `name` varchar(50) NOT NULL,
  `username` varchar(10) NOT NULL,
  `password` varchar(255) NOT NULL,
  `email` varchar(30) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id_user`, `id_role`, `name`, `username`, `password`, `email`) VALUES
(11, 1, 'Mimin', 'Admin2', '$2a$10$zi0NOhV8WI1Y0gAdVsaRP.zOW6SsjblAFshnH6TtATj4ryJQVMzMe', 'masadmin@tes.com'),
(12, 2, 'Manto', 'User404', '$2a$10$uB1zkczWpuSlpYb2Mudil.449YH6jGe9l6PrZS4kJ53ZDmZf4odsW', 'masmanto@tes.com'),
(13, 2, 'Udin', 'Udin', '$2a$10$jNLLkNmwsdo1soG0TzUsPOCsj/ALhUxHwIOvCfqDU3FwzLeclGOcq', 'udinhengker@gmail.com'),
(21, 2, 'udin1', 'Udin15', '$2a$10$yklBHrUO7n32xGw3DsTf2O1iez2nHt6ZrlVutEAPWlTo0DHwypw7e', 'udin@nganga.com');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `books`
--
ALTER TABLE `books`
  ADD PRIMARY KEY (`id_book`),
  ADD UNIQUE KEY `isbn` (`isbn`),
  ADD KEY `id_user` (`id_user`);

--
-- Indexes for table `roles`
--
ALTER TABLE `roles`
  ADD PRIMARY KEY (`id_role`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id_user`),
  ADD KEY `fk_users_roles` (`id_role`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `books`
--
ALTER TABLE `books`
  MODIFY `id_book` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=25;

--
-- AUTO_INCREMENT for table `roles`
--
ALTER TABLE `roles`
  MODIFY `id_role` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id_user` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=22;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `books`
--
ALTER TABLE `books`
  ADD CONSTRAINT `books_ibfk_1` FOREIGN KEY (`id_user`) REFERENCES `users` (`id_user`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Constraints for table `users`
--
ALTER TABLE `users`
  ADD CONSTRAINT `fk_users_roles` FOREIGN KEY (`id_role`) REFERENCES `roles` (`id_role`) ON DELETE CASCADE ON UPDATE CASCADE;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;

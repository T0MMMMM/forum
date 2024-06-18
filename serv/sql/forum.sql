CREATE TABLE IF NOT EXISTS `users` (
  `id` integer PRIMARY KEY,
  `username` varchar(255),
  `email` varchar(255),
  `password` varchar(255),
  `created_at` varchar(255)
);
CREATE TABLE IF NOT EXISTS `categories` (
  `id` integer PRIMARY KEY,
  `name` varchar(255),
  `description` text
);
CREATE TABLE IF NOT EXISTS `topics` (
  `id` integer PRIMARY KEY,
  `categoryID` integer,
  `userID` integer,
  `title` varchar(255),
  `content` text,
  `created_at` varchar(255),
  `status` varchar(255),
  `visible` bool,
  `like` integer,
  `dislike` integer,
  FOREIGN KEY (`categoryID`) REFERENCES `categories` (`id`),
  FOREIGN KEY (`userID`) REFERENCES `users` (`id`)
);
CREATE TABLE IF NOT EXISTS `answers` (
  `id` integer PRIMARY KEY,
  `topicID` integer,
  `userID` integer,
  `content` text,
  `created_at` varchar(255),
  `status` varchar(255),
  `visible` bool,
  `like` integer,
  `dislike` integer,
  FOREIGN KEY (`topicID`) REFERENCES `topics` (`id`),
  FOREIGN KEY (`userID`) REFERENCES `users` (`id`)
);
CREATE TABLE IF NOT EXISTS `messages` (
  `id` integer PRIMARY KEY,
  `senderID` integer,
  `recipientID` integer,
  `content` text,
  `visible` bool,
  `created_at` varchar(255),
  FOREIGN KEY (`senderID`) REFERENCES `users` (`id`),
  FOREIGN KEY (`recipientID`) REFERENCES `users` (`id`)
);
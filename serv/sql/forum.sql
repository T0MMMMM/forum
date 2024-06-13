CREATE TABLE IF NOT EXISTS `users` (
  `id` integer PRIMARY KEY,
  `username` varchar(255),
  `email` varchar(255),
  `password` varchar(255),
  `created_at` timestamp,
  `updated_at` timestamp
);
CREATE TABLE IF NOT EXISTS `categories` (
  `id` integer PRIMARY KEY,
  `name` varchar(255),
  `description` text,
  `created_at` timestamp,
  `updated_at` timestamp
);
CREATE TABLE IF NOT EXISTS `topics` (
  `id` integer PRIMARY KEY,
  `categoryID` integer,
  `userID` integer,
  `title` varchar(255),
  `content` text,
  `created_at` timestamp,
  `updated_at` timestamp,
  FOREIGN KEY (`categoryID`) REFERENCES `categories`(`id`),
  FOREIGN KEY (`userID`) REFERENCES `users`(`id`)
);
CREATE TABLE IF NOT EXISTS `Posts` (
  `id` integer PRIMARY KEY,
  `topicID` integer,
  `userID` integer ,
  `content` text,
  `created_at` timestamp,
  `updated_at` timestamp,
  FOREIGN KEY (`topicID`) REFERENCES `topics`(`id`),
  FOREIGN KEY (`userID`) REFERENCES `users`(`id`)
);

CREATE TABLE `users` (
  `id` integer PRIMARY KEY,
  `username` varchar(255),
  `email` varchar(255),
  `password` varchar(255),
  `created_at` timestamp,
  `updated_at` timestamp
);

CREATE TABLE `categories` (
  `id` integer PRIMARY KEY,
  `name` varchar(255),
  `description` text,
  `created_at` timestamp,
  `updated_at` timestamp
);

CREATE TABLE `topics` (
  `id` integer PRIMARY KEY,
  `categoryID` integer,
  `userID` integer,
  `title` varchar(255),
  `content` text,
  `created_at` timestamp,
  `updated_at` timestamp
);

CREATE TABLE `Posts` (
  `id` integer PRIMARY KEY,
  `topicID` integer,
  `userID` integer,
  `content` text,
  `created_at` timestamp,
  `updated_at` timestamp
);

ALTER TABLE `topics` ADD FOREIGN KEY (`categoryID`) REFERENCES `categories` (`id`);

ALTER TABLE `topics` ADD FOREIGN KEY (`userID`) REFERENCES `users` (`id`);

ALTER TABLE `Posts` ADD FOREIGN KEY (`topicID`) REFERENCES `topics` (`id`);

ALTER TABLE `Posts` ADD FOREIGN KEY (`userID`) REFERENCES `users` (`id`);

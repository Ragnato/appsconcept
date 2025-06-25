CREATE DATABASE IF NOT EXISTS fizzbuzz;

CREATE TABLE IF NOT EXISTS fizzbuzz.fizzbuzz_requests (
    `id` INT AUTO_INCREMENT PRIMARY KEY,
    `int1` INT NOT NULL,
    `int2` INT NOT NULL,
    `limit_val` INT NOT NULL,
    `str1` VARCHAR(50) NOT NULL,
    `str2` VARCHAR(50) NOT NULL
);

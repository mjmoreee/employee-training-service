# Employee Training Service [Prototype]
The Employee Training Service Platform is a comprehensive software solution designed to streamline and enhance the employee training process within organizations. This platform empowers businesses to efficiently manage, deliver, and track employee training programs, ultimately fostering continuous learning and development among their workforce.

## The Resons
Given the rapidly evolving digital landscape, digital security concerns often take precedence, it is undeniable that human resources frequently represent the most significant vulnerability in security, this project aims to assist companies in enhancing their employees' awareness of digital security, furthermore, with the enactment of Law Number 27 of 2022 concerning Personal Data Protection by the government, it has become increasingly important for businesses to pay greater attention to digital security issues.

## The Process

1. **Planning and Design:** 
   Initial phase involving project planning and conceptual design.
   
2. **UI/UX Design using Figma:**
   Creation of user interface and user experience designs using Figma, a collaborative design tool.
   
3. **Entity-Relationship Diagram (ERD) Design:**
   Development of the entity-relationship diagram to represent the database structure.
      ## MySQL Docker Setup

      ### Introduction

      This document provides instructions for setting up a MySQL database using Docker, creating a database, and defining a table schema for an employee management system.

      ### Prerequisites

      - Docker installed on your system.

      ### Step 1: Run MySQL in Docker

      To run MySQL in a Docker container, use the following command:
      
      ```bash
      docker run --name employee-training-service -e MYSQL_ROOT_PASSWORD=12345678 -d mysql
      ```

      This command creates a Docker container named "employee-training-service" with a root password of "12345678."

      ### Step 2: Create a Database

      After running the MySQL container, access it using the following command:

      ```bash
      docker exec -it employee-training-service bash
      ```

      Now, you can interact with the MySQL server. Inside the MySQL command-line interface, use the following SQL syntax to create a new database called "employee_training_service" and switch to it:

      ```sql
      CREATE DATABASE employee_training_service;
      USE employee_training_service;
      ```

      ### Step 3: Create Employee Table

      You can create a table named "employee" with the following SQL syntax:

      ```sql
      CREATE TABLE employee (
      employee_id INT NOT NULL AUTO_INCREMENT,
      first_name VARCHAR(100) NOT NULL,
      last_name VARCHAR(100),
      description VARCHAR(255),
      gender ENUM('Male', 'Female', 'Other') NOT NULL,
      role VARCHAR(100) NOT NULL,
      department VARCHAR(100) NOT NULL,
      hired_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
      email VARCHAR(100) NOT NULL,
      telephone VARCHAR(15) NOT NULL,
      ranking INT NOT NULL DEFAULT 0,
      points INT NOT NULL DEFAULT 0,
      image_path VARCHAR(255) DEFAULT '/assets/images/employee-default.jpg',
      PRIMARY KEY (employee_id),
      UNIQUE KEY email_unique (email)
      ) ENGINE = InnoDB;
      ```

      This SQL statement creates an "employee" table with various fields such as employee ID, name, description, gender, role, department, hire date, email, telephone, ranking, and points. It also sets default values for certain fields and ensures that the "email" field is unique.

      You now have a MySQL database running in a Docker container with an "employee" table ready to store employee information.

      ### Step 4: Sample Data Entry for Employee Table

      Below is a set of sample data entries for the "employee" table, these entries represent fictional employee records with varying rankings, starting with the highest ranking (1) and points (1000).

      ```sql
      INSERT INTO employee (first_name, last_name, description, gender, role, department, email, telephone, ranking, points)
      VALUES
      ('Jack', 'Smith', 'Senior Engineer', 'Male', 'Engineer', 'IT', 'jack.smith@talentforge.id', '+1 (123) 456-7890', 1, 1000),
      ('John', 'Doe', 'Product Manager', 'Male', 'Manager', 'Management', 'john.doe@talentforge.id', '+1 (987) 654-3210', 2, 950),
      ('Jane', 'Johnson', 'Marketing Specialist', 'Female', 'Specialist', 'Marketing', 'jane.johnson@talentforge.id', '+1 (555) 123-4567', 3, 900),
      ('Michael', 'Brown', 'Software Developer', 'Male', 'Developer', 'IT', 'michael.brown@talentforge.id', '+1 (111) 222-3333', 4, 850),
      ('Emily', 'Wilson', 'HR Coordinator', 'Female', 'Coordinator', 'HR', 'emily.wilson@talentforge.id', '+1 (999) 888-7777', 5, 800),
      ('James', 'Anderson', 'Financial Analyst', 'Male', 'Analyst', 'Finance', 'james.anderson@talentforge.id', '+1 (777) 888-9999', 6, 750),
      ('Olivia', 'Davis', 'Marketing Manager', 'Female', 'Manager', 'Marketing', 'olivia.davis@talentforge.id', '+1 (123) 456-7890', 7, 700),
      ('William', 'Taylor', 'Sales Representative', 'Male', 'Sales', 'Sales', 'william.taylor@talentforge.id', '+1 (111) 222-3333', 8, 650),
      ('Ava', 'Clark', 'Software Engineer', 'Female', 'Engineer', 'IT', 'ava.clark@talentforge.id', '+1 (555) 123-4567', 9, 600),
      ('Noah', 'Wilson', 'Project Manager', 'Male', 'Manager', 'Management', 'noah.wilson@talentforge.id', '+1 (987) 654-3210', 10, 550),
      ('Sophia', 'Johnson', 'Graphic Designer', 'Female', 'Designer', 'Design', 'sophia.johnson@talentforge.id', '+1 (777) 888-9999', 11, 500),
      ('Liam', 'Miller', 'Database Administrator', 'Male', 'Administrator', 'IT', 'liam.miller@talentforge.id', '+1 (999) 888-7777', 12, 450),
      ('Isabella', 'Moore', 'HR Manager', 'Female', 'Manager', 'HR', 'isabella.moore@talentforge.id', '+1 (123) 456-7890', 13, 400),
      ('Benjamin', 'Johnson', 'Marketing Analyst', 'Male', 'Analyst', 'Marketing', 'benjamin.johnson@talentforge.id', '+1 (111) 222-3333', 14, 350),
      ('Mia', 'Garcia', 'Sales Manager', 'Female', 'Manager', 'Sales', 'mia.garcia@talentforge.id', '+1 (555) 123-4567', 15, 300),
      ('Elijah', 'Brown', 'Software Tester', 'Male', 'Tester', 'IT', 'elijah.brown@talentforge.id', '+1 (987) 654-3210', 16, 250),
      ('Charlotte', 'Davis', 'Content Writer', 'Female', 'Writer', 'Content', 'charlotte.davis@talentforge.id', '+1 (777) 888-9999', 17, 200),
      ('Lucas', 'Jones', 'Network Engineer', 'Male', 'Engineer', 'IT', 'lucas.jones@talentforge.id', '+1 (999) 888-7777', 18, 150),
      ('Amelia', 'Martinez', 'Marketing Coordinator', 'Female', 'Coordinator', 'Marketing', 'amelia.martinez@talentforge.id', '+1 (123) 456-7890', 19, 100),
      ('Henry', 'Smith', 'Financial Manager', 'Male', 'Manager', 'Finance', 'henry.smith@talentforge.id', '+1 (111) 222-3333', 20, 50),
      ('Luna', 'Hernandez', 'Sales Analyst', 'Female', 'Analyst', 'Sales', 'luna.hernandez@talentforge.id', '+1 (555) 123-4567', 21, 45),
      ('Alexander', 'Wilson', 'Front-end Developer', 'Male', 'Developer', 'IT', 'alexander.wilson@talentforge.id', '+1 (123) 456-7890', 22, 40),
      ('Grace', 'Anderson', 'Digital Marketing Specialist', 'Female', 'Specialist', 'Marketing', 'grace.anderson@talentforge.id', '+1 (111) 222-3333', 23, 35),
      ('Logan', 'Garcia', 'Network Administrator', 'Male', 'Administrator', 'IT', 'logan.garcia@talentforge.id', '+1 (777) 888-9999', 24, 30),
      ('Chloe', 'Brown', 'HR Analyst', 'Female', 'Analyst', 'HR', 'chloe.brown@talentforge.id', '+1 (555) 123-4567', 25, 25),
      ('Owen', 'Williams', 'UI/UX Designer', 'Male', 'Designer', 'Design', 'owen.williams@talentforge.id', '+1 (999) 888-7777', 26, 20),
      ('Zoe', 'Johnson', 'Systems Analyst', 'Female', 'Analyst', 'IT', 'zoe.johnson@talentforge.id', '+1 (123) 456-7890', 27, 15),
      ('Mason', 'Gonzalez', 'Sales Coordinator', 'Male', 'Coordinator', 'Sales', 'mason.gonzalez@talentforge.id', '+1 (111) 222-3333', 28, 10),
      ('Sofia', 'Lee', 'Back-end Developer', 'Female', 'Developer', 'IT', 'sofia.lee@talentforge.id', '+1 (555) 123-4567', 29, 5),
      ('Ethan', 'Hall', 'Marketing Assistant', 'Male', 'Assistant', 'Marketing', 'ethan.hall@talentforge.id', '+1 (777) 888-9999', 30, 1);
      ```

4. **Coding with Go (Golang):**
   Implementation of the project by writing code using the Go programming language.

5. **Testing:**
   Thorough testing of the project to ensure functionality and reliability.

## The Results
On Progress

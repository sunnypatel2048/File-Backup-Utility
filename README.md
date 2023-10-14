# File Backup Utility

A simple Go utility for backing up files and directories from a source folder to a destination folder.

## Table of Contents

- [Introduction](#introduction)
- [Features](#features)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
- [Usage](#usage)
- [How It Works](#how-it-works)
- [Tests](#tests)
- [Contributing](#contributing)
- [License](#license)

## Introduction

This Go File Backup Utility is designed to copy files and directories from a source directory to a destination directory, including subdirectories. It provides a simple way to create backups of your data.

## Features

- Backup files and directories.
- Recursively copy subdirectories.
- Error handling for various scenarios.

## Getting Started

### Prerequisites

- Go (Golang) should be installed on your system. If it's not, you can download it from [here](https://golang.org/dl/).

### Installation

1. Clone this repository to your local machine:

   ```bash
   git clone https://github.com/yourusername/go-file-backup-utility.git
   ```
   
2. Switch to the project directory:

    ```bash
    cd File-Backup-Utility
    ```

## Usage

1. Open a terminal and navigate to the project directory.

2. Run the utility:

    ```bash
    go run main.go
    ```

3. Follow the on-screen prompts to provide source and destination paths.

4. The utility will create a backup in the destination directory.

## How It Works

The utility copies files and directories from the source to the destination using Go's os and io packages. It recursively traverses subdirectories, creating backups of the entire structure.

## Tests

To run tests, use the following command in the project directory:

  ```bash
  go test
  ```

## Contributing

Contributions are welcome! If you'd like to contribute to this project, please follow these steps:

1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Make your changes and commit them with descriptive messages.
4. Push your branch to your forked repository.
5. Create a pull request to the main repository.

## License

TThis project is licensed under the [MIT License](LICENSE). You can view the full license text [here](https://opensource.org/licenses/MIT).

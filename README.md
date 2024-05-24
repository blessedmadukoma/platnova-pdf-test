
# Platnova PDF Generator

My submission for the Platnova assessment Go application that generates a Transaction Statement in PDF. The generated PDF is saved as `account_statement.pdf`.

## Table of Contents

- [Platnova PDF Generator](#platnova-pdf-generator)
  - [Table of Contents](#table-of-contents)
  - [Installation](#installation)
  - [Usage](#usage)
  - [Project Structure](#project-structure)
  - [Dependencies](#dependencies)
  - [Contributing](#contributing)
  - [License](#license)

## Installation

1. **Clone the repository:**

    ```sh
    git clone https://github.com/blessedmadukoma/platnova-pdf-test.git

    cd platnova-pdf-test
    ```

2. **Install dependencies:**

    Make sure you have Go installed. Then, install the required Go packages by running:

    ```sh
    go mod tidy
    ```

## Usage

1. **Place your image:**

    Ensure you have an image named `Revolut logo.png` in the `images` directory. If you use a different name or location, update the code accordingly.

2. **Run the application:**

    ```sh
    go run *.go
    ```

3. **Check the output:**

    After running the application, a file named `account_statement.pdf` will be generated in your project directory.

## Project Structure

```
platnova-pdf-test/
├── images/
│   └── Revolut logo.png
├── main.go
├── helpers.go
└── README.md
```

- `images/`: Directory to store the image used in the PDF header.
- `main.go`: The main Go file containing the logic for generating the PDF.
- `helpers.go`: The helpers file is to generate utils and extras needed such as color
- `README.md`: Documentation file.

## Dependencies

- [Maroto](https://github.com/johnfercher/maroto): A PDF creation library for Go.

## Contributing

Contributions are welcome! Please follow these steps to contribute:

1. Fork the repository.
2. Create a new branch (`git checkout -b feature-branch`).
3. Make your changes.
4. Commit your changes (`git commit -m 'Add some feature'`).
5. Push to the branch (`git push origin feature-branch`).
6. Open a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
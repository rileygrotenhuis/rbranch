# rbranch

rbranch is a CLI tool built using Go and Bubbletea for those who are too lazy to type out those super long and annoying git branch names. Making your life maybe 2% easier, you can now simply use the `rbranch` command with a few additional flags to easily perform simple git branch operations.

## Prerequisites

Before you can run this application, ensure you have the following software installed on your system:

1. **Go**: This application requires the latest version of Go to run. You can download and install Go from the [official website](https://go.dev/dl/). Ensure you have the latest version of Go by running the following command::

    ```bash
    go version
    ```

## Installation

To get started with this application, follow these steps:

1. **Clone the Repository**: Start by cloning this repository to your local machine. You can do this by running the following command in your terminal:

    ```bash
    git clone https://github.com/rileygrotenhuis/rbrnach.git
    ```

2. **Navigate to the Project Directory**: Change your working directory to the newly cloned repository:

    ```bash
    cd rbrnach
    ```

3. **Install Dependencies**: Now you need to install the project dependencies. Run the following command:

   ```bash
   go mod tidy
   ```

4. **Build Application**: At this point you will want to build the executable for rbranch by running the following command:

    ```bash
    go build
    ```

5. **Make Globally Accessible**: The last step in the installation process is to move the newly created executable to your `$PATH`. Run the following command:

    ```bash
    sudo mv rbranch /usr/local/bin
    ```

## Usage

### Checkout a Branch

To view a list of all available branches and select one of them to checkout, you can use the default command:

```bash
rbranch
```

### Copy a Branch

To view a list of all available branches and select one of them to copy to your clipboard, you will need to append the copy flag

```bash
rbranch -c
```

### Delete a Branch

To view a list of all available branches and select one of them to delete, you will need to append the deletion flag

```bash
rbranch -d
```
A simple CLI application to view and interact with JSON files, demonstrating my ability to manage and update dependencies using [Renovate](https://github.com/renovatebot/renovate).

## What is this repository about?

This project serves as a showcase of my proficiency with dependency management tools, specifically **Renovate**. The repository intentionally uses an outdated version of the popular Go CLI library [cobra](https://github.com/spf13/cobra) in its dependencies. The goal is to leverage the Renovate app to automatically detect the outdated dependency, open a pull request, and help me upgrade cobra to its latest version.

## Key Objectives

- **Demonstrate Renovate setup:** Use a `renovate.json` configuration file to instruct Renovate on how to monitor and update dependencies.
- **Showcase dependency updates:** Start with an intentionally outdated version of cobra, let Renovate open PRs, and merge those updates.
- **Provide a simple CLI tool:** The CLI allows users to view JSON data from the command line, making the example practical and easy to test.

## How it works

1. **Initial Setup:** The repository contains a Go CLI app using an outdated version of cobra.
2. **Renovate Configuration:** The `renovate.json` file tells Renovate how to track dependencies.
3. **Automated PRs:** Renovate will scan the project, detect the outdated cobra version, and open a PR to update it.
4. **Review and Merge:** I will review and merge the PR, demonstrating a standard dependency update workflow.

## Getting Started

### Prerequisites

- Go (>=1.18)
- [cobra](https://github.com/spf13/cobra) (the version will be managed by Renovate)

### Installation

Clone the repository:

```bash
git clone https://github.com/yashpal2104/renovate-sample-task.git
cd renovate-sample-task
```

Install dependencies:

```bash
go mod tidy
```

### Usage

To build the CLI:

```bash
go build -o fruitcli
```

To list all fruits from the sample JSON:

```bash
./fruitcli list
```

To filter fruits by color:

```bash
./fruitcli filter Red
```

To filter fruits by size:

```bash
./fruitcli size Large
```

> **Note:** The CLI reads from the included `json_file.json` by default.

### Example

```bash
./fruitcli list
./fruitcli filter Yellow
./fruitcli size Small
```
## External Dependency

This project uses the [cobra](https://github.com/spf13/cobra) library for building the CLI. Cobra is a widely used Go library for creating powerful modern CLI applications.

### Security Example

Cobra has had security advisories in the past. For example:

- [CVE-2023-3978](https://nvd.nist.gov/vuln/detail/CVE-2023-3978):  
  A vulnerability in cobra before v1.7.0 allowed attackers to execute arbitrary code via maliciously crafted input to the `Args` function.

This repository demonstrates how [Renovate](https://github.com/renovatebot/renovate) can help keep dependencies like cobra up to date, automatically opening pull requests when security vulnerabilities are discovered and patched.

## Dependency Management with Renovate

- The project includes a `renovate.json` file to configure Renovate.
- Renovate will automatically propose updates for dependencies—starting with cobra.
- This process demonstrates automated dependency management and continuous maintenance.

## Contributing

Pull requests are welcome! For major changes, please open an issue first to discuss what you would like to change.

## License

[Apache 2.0](LICENSE)

---

**Note:** This repository is primarily for demonstration purposes to highlight automated dependency management using Renovate.
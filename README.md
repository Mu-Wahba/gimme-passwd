# Password Generator

This project provides a simple password generator utility written in Go. It allows users to generate random passwords of specified lengths using a defined character set and save all generated passwords to a specific file with the time stamp and a description 

## Installation

To use this project, ensure you have Go installed on your machine.
Clone the repository:
```bash
git clone https://github.com/mu-wahba/gimme-passwd
cd gimme-passwd

```

Configure the environment for your needs, copy the example environment file:
```bash
cp .env.example .env
```

## Usage

```bash
go run -o gimme-passwd
```

## Examples
To generate a password with the default 8 characters:
```bash
./gimme-passwd 
```

To generate a password of length between 8 and 20 characters:
```bash
./gimme-passwd -l 12
```

To generate a password with specific length and character sets:
```bash
./gimme-passwd -l 16 -d "this pasword for my blog"
```


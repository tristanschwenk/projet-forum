# EZYO/FORUM

Forum projet

## Getting Started

### Requirements

- Go 1.13+
- Go modules

### Running the project

 - ``npm run start`` - Start the API
 - ``npm run watch`` - Start the API and watch changes


### Troubleshooting

> #### godotenv: command not found
> This error occurs because the GO path is not in your env PATH, to fix use this command
> ```bash
> echo -e "export PATH=\"\$PATH:\$HOME/go/bin\"\n" >> ~/.bashrc && source ~/.bashrc
> ```
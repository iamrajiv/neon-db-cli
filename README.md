<div align="center">
<img src="https://neon.tech/favicon/favicon.png" height="auto" width="200" />
<br />
<h1>Neon Database CLI</h1>
<p>
CLI tool to interact with the Neon Database APIs
</p>
<a href="https://github.com/iamrajiv/neon-db-cli/network/members"><img src="https://img.shields.io/github/forks/iamrajiv/neon-db-cli?color=0969da&style=for-the-badge" height="auto" width="auto" /></a>
<a href="https://github.com/iamrajiv/neon-db-cli/stargazers"><img src="https://img.shields.io/github/stars/iamrajiv/neon-db-cli?color=0969da&style=for-the-badge" height="auto" width="auto" /></a>
<a href="https://github.com/iamrajiv/neon-db-cli/blob/main/LICENSE"><img src="https://img.shields.io/github/license/iamrajiv/neon-db-cli?color=0969da&style=for-the-badge" height="auto" width="auto" /></a>
</div>

## About

The NeonDB CLI Tool is a command-line interface that enables interaction with the Neon Database APIs.

[Neon](https://neon.tech/) is a fully managed serverless PostgreSQL solution. Neon offers modern developer features such as serverless computing, branching, bottomless storage, and more. Learn more about Neon [here](https://neon.tech/docs/introduction).

The Neon API allows you to programmatically manage your Neon account and its objects, including:

- API keys
- Projects
- Branches
- Compute endpoints
- Databases
- Roles

To authenticate with the Neon API, you need to include an API key token in the Authorization header of your API request.

The Neon API uses `Bearer Token Authentication`. Therefore, the Authorization header should be formatted as follows:

- `H 'Authorization: Bearer $NEON_API_KEY'`

For example:

```shell
curl 'https://console.neon.tech/api/v2/projects' \
  -H 'Accept: application/json' \
  -H 'Authorization: Bearer $NEON_API_KEY' \
  -H 'Content-Type: application/json'

```

To try an API request in the browser from this page, click the `Try It` button on any Neon API endpoint, and enter the token in the `Bearer` prompt. You can select the endpoint from the sidebar.

The [Get a list of projects](https://api-docs.neon.tech/reference/listprojects) endpoint is a simple one to try as it only requires an API key token.

A Neon API key is a randomly-generated 64-bit token that remains valid until it is revoked.

To create an API key, follow these steps:

1. Go to the **[Developer Settings](https://console.neon.tech/app/settings/api-keys)** page in the Neon Console.
2. Provide a name for the API key.
3. Click **Create** and copy the generated token.

Store the token in a secure location. Once you leave the `Developer Settings` page in the Neon Console, the token will not be visible or copyable.

For more information on managing Neon API keys, refer to the [Manage API keys](https://neon.tech/docs/manage/api-keys) section in the Neon documentation.

The project's folder structure is as follows:

```shell
.
├── .env
├── LICENSE
├── Makefile
├── README.md
├── apis
│   ├── api_keys
│   │   └── api_keys.go
│   └── projects
│       └── projects.go
├── assets
│   └── neon-db-cli.svg
├── cmd
│   └── cmd.go
├── go.mod
├── go.sum
├── main.go
└── utils
    └── utils.go
```

## Usage

1. Follow the instructions **[here](https://neon.tech/docs/get-started-with-neon/setting-up-a-project)** to set up a project in Neon.
2. Configure the database connection details in a **`.env`** file.
3. Run the command **`make all`** to install all the dependencies, remove the previous binary, and build a new binary named **`neondb`**.
4. Use the command **`./neondb`** to run the CLI tool.

When we run `./neondb`, we get the following output:

```shell
➜  neon-db-cli git:(main) ✗ ./neondb
Neon Database CLI

Usage:
  neondb [command]

Available Commands:
  completion   Generate the autocompletion script for the specified shell
  createapikey Create a new API key for your Neon account
  help         Help about any command
  listapikeys  List API keys for your Neon account
  revokeapikey Revoke an API key for your Neon account

Flags:
  -h, --help   help for neondb

Use "neondb [command] --help" for more information about a command.
```

### Instructions for setting environment variables

To run the project locally, create a `.env` file and add the following environment variables:

- **`API_KEY`**: API key for your Neon account.

You can find your API key on the [Developer Settings](https://console.neon.tech/app/settings/api-keys) page.

## Demonstration

Before the demonstration, I set up the new project and configured the database connection details in a `.env` file.

When I run the available commands, I get the following output:

#### completion

```shell
➜  neon-db-cli git:(main) ✗ ./neondb completion
Generate the autocompletion script for neondb for the specified shell.
See each sub-command's help for details on how to use the generated script.

Usage:
  neondb completion [command]

Available Commands:
  bash        Generate the autocompletion script for bash
  fish        Generate the autocompletion script for fish
  powershell  Generate the autocompletion script for powershell
  zsh         Generate the autocompletion script for zsh

Flags:
  -h, --help   help for completion

Use "neondb completion [command] --help" for more information about a command.
```

#### createapikey

```shell
➜  neon-db-cli git:(main) ✗ ./neondb createapikey --keyname=key2
{"id":000002,"key":"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"}
```

#### help

```shell
➜  neon-db-cli git:(main) ✗ ./neondb help
Neon Database CLI

Usage:
  neondb [command]

Available Commands:
  completion   Generate the autocompletion script for the specified shell
  createapikey Create a new API key for your Neon account
  help         Help about any command
  listapikeys  List API keys for your Neon account
  revokeapikey Revoke an API key for your Neon account

Flags:
  -h, --help   help for neondb

Use "neondb [command] --help" for more information about a command.
```

#### listapikeys

```shell
➜  neon-db-cli git:(main) ✗ ./neondb listapikeys
[{"id":000002,"name":"key2","created_at":"2023-05-17T07:42:45Z","last_used_at":null,"last_used_from_addr":""},{"id":000001,"name":"key1","created_at":"2023-05-16T19:12:11Z","last_used_at":"2023-05-17T07:43:51Z","last_used_from_addr":"49.37.25.219"}]
```

#### revokeapikey

```shell
➜  neon-db-cli git:(main) ✗ ./neondb revokeapikey --keyid=000002
{"id":000002,"name":"key2","revoked":true,"last_used_at":null,"last_used_from_addr":""}
```

## License

[MIT](https://github.com/iamrajiv/neon-db-cli/blob/main/LICENSE)

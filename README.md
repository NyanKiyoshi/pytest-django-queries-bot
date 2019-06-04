# pytest-django-queries-bot
An integration of [pytest-django-queries](https://github.com/NyanKiyoshi/pytest-django-queries) into GitHub pull requests.

## Workflow
### Opening a Pull Request
The GitHub webhook endpoint of the bot will receive an event from GitHub notifying about the pull request creation. Those information will be stored for further verification, e.g., checking whether the bot should be expecting a given SHA1 commit hash results or not.

### Sending a Diff File
On a pull request, once the CI, e.g. Travis was done analyzing the performances and comparing the performance, it will send the results to the `diff` endpoint of the bot.

[![https://i.imgur.com/VtrLu8M.png](https://i.imgur.com/VtrLu8M.png)](https://i.imgur.com/VtrLu8M.png)

### Sending a Results File
On a push event (mostly master), once the CI, e.g. Travis or Circle CI was done analyzing the performances of the branch it will send the raw JSON results to the `upload` endpoint of the bot to allow other pull requests to compare their results with other base branches (master or else).

## Installing the bot
### Requirements
- `bash`
- `npm`
- `go 1.1x` (`brew install go`)
- `go-dep` (`brew install dep` or `apt-get install go-dep`)

### Installation
1. `npm i`
1. `export PATH="$PATH:$(npm bin)"`

### Integrating with GitHub
TBA.

[![https://i.imgur.com/Cs9Abzg.png](https://i.imgur.com/Cs9Abzg.png)](https://i.imgur.com/Cs9Abzg.png)

### Compiling
Note: you will need to edit `development.env` or creation `production.env`.

Run `./compile.sh`

### Testing Locally
1. Requires a python virtual environment
1. Requires a working installation of docker
1. `pip install aws-sam-cli`
1. Run `./local-server.sh`

### Deploying to AWS
1. Run `./deploy.sh`

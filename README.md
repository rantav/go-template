# WARNING
This project is not maintainted. We use it internally at appsflyer, but the external (e.g. this GH version) is not maintainted. 

# Welcome to go-template

This project will jumpstart your Golang project and provide a set of tempaltes you may use to keep your code tidy and
consistent.

## Using Hygen
We use the templating tool called Hygen. https://www.hygen.io

### Installing Hygen
See here https://www.hygen.io/quick-start

But in short:

    $ brew tap jondot/tap
    $ brew install hygen

### Cloning and getting started

Clone the template and update remotes

    # clone the template project
    git clone git@gitlab.appsflyer.com:Architecture/go-template.git my-go-project
    cd my-go-project

    # rename the tempaltes remote
    git remote rename origin template

    # add the project repo
    git remote add origin git@gitlab.appsflyer.com:rantav/my-go-project.git

Use Hygen to initialize your code

    # Learn how to use the tool
    $ hygen init help

    # Templatize
    $ HYGEN_OVERWRITE=1 hygen template init my-go-project --repo_path=gitlab.appsflyer.com/rantav --description="My awesome go project"

    # Build to validate
    make


Add commands:

    $ hygen template new-cmd migrate --desc 'some amazing migration'

    Loaded templates: _templates
        added: cmd/migrate.go

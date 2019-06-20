---
to: README.md
---

# Welcome to <%= name %>

## Using Hygen

Make this your own (call it 'amazing'):

    $ hygen template init amazing
    ✔      exists: go.mod. Overwrite? (y/N):  (y/N) · true
       added: go.mod
    ✔      exists: main.go. Overwrite? (y/N):  (y/N) · true
        added: main.go
    ✔      exists: README.md. Overwrite? (y/N):  (y/N) · true
        added: README.md
    ✔      exists: cmd/root.go. Overwrite? (y/N):  (y/N) · true
        added: cmd/root.go

Add commands:

    $ hygen template new-cmd migrate --desc 'some amazing migration'

    Loaded templates: _templates
        added: cmd/migrate.go

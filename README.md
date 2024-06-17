# Go templ with HTMX

<!-- TODO: add htmx to repo -->
<!-- TODO: add tailwind css to repo -->
<!-- TODO: List of names in a table that are filterable via search bar -->

## Development

### Install `templ` Binary

The `templ` binary must be in your PATH variable for the VS Code extension to work.


1. Download the binary from https://github.com/a-h/templ/tags.

2. Copy it to a PATH folder.
    
    ex: `sudo cp <path-to-templ> /usr/local/bin/`

3. Update the binary permissions.

    `chmod +x ./bin/templ/templ`

### Install `templ` VS Code extension (optional)

1. Install https://marketplace.visualstudio.com/items?itemName=a-h.templ.

2. Update your VS Code `preferences.json`.

    ```json
    "[templ]": {
        "editor.defaultFormatter": "a-h.templ"
    },
    "emmet.includeLanguages": {
        "templ": "html"
    },
    "tailwindCSS.includeLanguages": {
        "templ": "html"
    },
    ```


## Build

### Build, Run and Watch for Changes

For you development convenience.

Requires nodemon, install with `npm i -g nodemon`.

-  `./scripts/watch.sh`

### Build and Run

- `./scripts/run.sh`


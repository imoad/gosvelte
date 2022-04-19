# GoSvelte

[![Golang + Svelte Starter Kit: GoSvelte](/www/public/gosvelte.png?style=centerme)](https://svelte.dev)

# What is GoSvelte?

GoSvelte is starter kit to build modular web applications using Golang and Svelte compiler to power a template engine in a productive approach.

Learn more at the [Golang website](https://go.dev), or stop by the [Svelte website](https://svelte.dev).

# Get started

*Note that you will need to have [Go](https://go.dev) and [Node.js](https://nodejs.org) both installed.*

Clone this repository:

```bash
$ git clone https://github.com/moadkey/gosvelte.git <my-project-name>
$ cd <my-project-name>
```

Then install the project dependencies by doing the following:

```bash
$ yarn                # ( `yarn install` or `npm install`)
```

*Note that GoSvelte comes with svelte script, to change to svelte TypeScript development environment, you can run immediately after cloning:*

```bash
$ yarn gsts           # (or `npm run gsts`)
```
# Development

```bash
$ yarn dev            # (or `npm run dev`)
```
Navigate to [localhost:3000](http://localhost:3000). You should see your gosvelte welcome app running. 

Edit a component file in `resources/views` to see your page changes.

Edit the go file in `modules/welcome` to change your page.

*Note that you need to restart in case go code changed but not if svelte component file: added, modified, renamed or deleted. (hot recompilation & live reloading)*

In case you need to run only svelte compiler or go application:

```bash
$ yarn svelte         # (or `npm run svelte`)
    #or
$ yarn go             # (or `npm run go`)
```

# Deployment

```bash
$ yarn build          # (or `npm run build`)
$ yarn start          # (or `npm run start`)
```

Happy Coding.
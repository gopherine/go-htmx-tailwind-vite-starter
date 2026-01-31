# Golang HTMX Tailwind esbuild Starter

## Overview

Starter template for Go + Echo + Templ with HTMX, Tailwind CSS, and esbuild. It includes:

- File-based routing (`views/pages/**/index.templ`).
- HTMX partial rendering with layout persistence.
- Small JS islands via `data-module` + esbuild code splitting (see the homepage counter).
- Embedded static assets in production builds.

## Project Structure

```bash
.
├── Makefile
├── README.md
├── air.toml
├── go.mod
├── go.sum
├── handlers
│   └── shared.go
├── internal
│   ├── assets
│   └── routes
├── main.go
├── package.json
├── scripts
│   ├── esbuild.mjs
│   └── gen-modules.mjs
├── static_dev.go
├── static_prod.go
├── tailwind.config.js
├── tools
│   ├── tools.go
│   └── generate_routes
├── views
│   ├── components
│   ├── css
│   ├── layouts
│   ├── pages
│   └── scripts
```

## Getting Started

### Prerequisites

- Go 1.24+
- Node.js + npm
- Air (optional, for live reload)
- Templ (`templ generate` is used for build/dev)

### Install

```bash
go mod tidy
pnpm install
# or
npm install
# or
bun install
```

### Development

```bash
make dev
```

This runs:
- `air` for Go reloads (dev build tag)
- `templ generate --watch`
- Tailwind CLI watch
- esbuild watch + manifest generation

### Production

```bash
make build
make run
```

Production builds embed everything under `public/` into the Go binary.

## File-Based Routing

Routes are derived from `views/pages/**/index.templ`:

- `views/pages/index.templ` → `/`
- `views/pages/blog/index.templ` → `/blog`
- `views/pages/blog/[id]/index.templ` → `/blog/:id`
- `views/pages/docs/[...slug]/index.templ` → `/docs/*`

Generate routes with:

```bash
make gen
```

## JS Islands (data-module)

Add a module file in `views/scripts/modules/` and reference it on the page:

```html
<div data-module="module"></div>
```

When adding new module files, regenerate the registry:

```bash
npm run gen:modules
```

esbuild will bundle code-split modules and write `public/manifest.json` for runtime lookup.

## Starter UI

The homepage includes a simple counter island and a lightweight "Teddy" logo to mimic a Next.js-style starter screen.
Edit `views/pages/index.templ` to customize the hero and components.

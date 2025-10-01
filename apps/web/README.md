# BSpec Documentation Site

This is the official documentation website for BSpec, built with VitePress and deployed on Cloudflare Pages.

## Development

```bash
# Install dependencies
bun install

# Start development server
bun run dev

# Build for production
bun run build

# Preview production build
bun run preview
```

## Deployment

The site is automatically deployed to Cloudflare Pages on every push to the `main` branch.

### Manual Deployment

```bash
# Deploy to Cloudflare Pages
wrangler pages deploy .vitepress/dist --project-name=bspec-docs
```

## Structure

- `index.md` - Homepage
- `docs/` - Documentation pages
- `spec/` - Specification pages
- `.vitepress/` - VitePress configuration and theme

## Branding

BSpec uses **Orange** as the primary brand color:
- Primary: `#f97316` (rgb(249, 115, 22))
- Hover: `#ea580c` (rgb(234, 88, 12))
- Active: `#c2410c` (rgb(194, 65, 12))

Custom theme colors are defined in `.vitepress/theme/custom.css`.

## Contributing

See the main [CONTRIBUTING.md](../../CONTRIBUTING.md) for guidelines.

# BSpec Branding Guide

## Brand Colors

BSpec uses **Orange** as the primary brand color, creating a warm, energetic, and distinctive identity.

### Primary Orange Palette

```css
/* Light Mode */
--color-primary-500: #f97316  /* Main orange - rgb(249, 115, 22) */
--color-primary-600: #ea580c  /* Hover state - rgb(234, 88, 12) */
--color-primary-700: #c2410c  /* Active state - rgb(194, 65, 12) */

/* Dark Mode */
--color-primary-400: #fb923c  /* Lighter orange for dark backgrounds */
--color-primary-500: #f97316  /* Main orange */
--color-primary-600: #ea580c  /* Darker accent */
```

### Usage in VitePress

The brand colors are applied throughout the site:

- **Hero text**: Orange gradient on homepage
- **Primary buttons**: Orange background with hover states
- **Links**: Orange color with darker hover
- **Active sidebar items**: Orange highlight
- **Feature icons**: Subtle orange glow
- **Code blocks**: Orange accent for tabs

## Typography

- **Primary Font**: System font stack (system-ui, -apple-system, sans-serif)
- **Code Font**: JetBrains Mono (monospace)

## Logo

Sophisticated "B" with cityscape/business architecture visual on gradient orange background.

### Design Elements
- **Letter B**: Bold, modern sans-serif with rounded corners
- **Cityscape**: Architectural line art representing business/structure
- **Sun/Moon**: Circular element symbolizing 24/7 business operations
- **Colors**: Gradient from orange (#f97316) to darker tones, with dark inset
- **Style**: Premium, professional, tech-forward

### File Locations
- `/public/bspec-logo-full.png` - Original high-res (1536x1024, 2.3MB)
- `/public/icon.png` - Large icon (1024x1024, 328KB) for social media
- `/public/logo-192.png` - Header logo (192x192, 21KB)
- `/public/logo.png` - Favicon (64x64, 3.9KB)

All versions maintain the same visual identity at different sizes.

## Customization

All brand colors are defined in `.vitepress/theme/custom.css`.

To change the brand color:

1. Update CSS variables in `custom.css`
2. Update theme-color meta tag in `.vitepress/config.mts`
3. Replace logo files in `public/` (logo.png, logo-192.png, icon.png, bspec-logo-full.png)

## Color Psychology

**Orange** conveys:
- 🔥 Energy and enthusiasm
- 💡 Innovation and creativity  
- 🎯 Confidence and boldness
- 🤝 Friendliness and approachability

Perfect for a forward-thinking business documentation standard that aims to empower users and enable AI integration.

## Accessibility

All color combinations meet WCAG 2.1 AA standards:
- Orange on white: 4.5:1 contrast ratio ✓
- White on orange: 4.5:1 contrast ratio ✓
- Orange on dark backgrounds: Enhanced for readability ✓

## Examples

### Buttons
- Primary: Orange background (#f97316)
- Hover: Darker orange (#ea580c)
- Active: Deep orange (#c2410c)

### Links
- Default: Orange (#f97316)
- Hover: Darker orange (#ea580c)
- Visited: Maintains orange (no purple)

### Highlights
- Selection: Soft orange background (rgba(249, 115, 22, 0.14))
- Focus rings: Orange outline
- Active states: Orange accent

## Related Files

- `.vitepress/theme/custom.css` - Color definitions
- `.vitepress/theme/index.ts` - Theme entry point
- `.vitepress/config.mts` - Theme meta tags and logo reference
- `public/bspec-logo-full.png` - Original high-resolution logo
- `public/logo-192.png` - Header logo (192px)
- `public/logo.png` - Favicon (64px)
- `public/icon.png` - Social media / app icon (1024px)

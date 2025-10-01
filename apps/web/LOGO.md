# BSpec Logo Integration

The official BSpec logo from `/assets/images/icon.png` has been successfully integrated into the documentation site.

## Source

- **Original**: `/assets/images/icon.png` (1024x1024, 2.2MB)
- **Format**: PNG with transparency
- **Design**: Professional BSpec logo with cityscape architecture

## Optimized Versions Created

From the source file, created multiple optimized sizes for both dark and light mode:

```
public/
# Dark Mode (Original)
├── bspec-logo-full.png    # 1024x1024 (2.2MB) - Original
├── icon.png               # 1024x1024 (2.2MB) - Social media, app icon  
├── logo-192.png          # 192x192   (34KB)  - Header navigation (dark)
└── logo.png              # 64x64     (6.7KB) - Favicon (dark)

# Light Mode (Inverted)
├── bspec-logo-full-light.png  # 1024x1024 (1.8MB) - Inverted original
├── icon-light.png             # 1024x1024 (1.8MB) - Social media (light)  
├── logo-192-light.png        # 192x192   (34KB)  - Header navigation (light)
└── logo-light.png            # 64x64     (6.2KB) - Favicon (light)
```

## Configuration

### `.vitepress/config.mts`

```typescript
themeConfig: {
  logo: {
    light: '/logo-192-light.png',  // Light mode logo
    dark: '/logo-192.png'           // Dark mode logo
  }
}

head: [
  ['link', { rel: 'icon', type: 'image/png', sizes: '64x64', href: '/logo.png' }],
  ['link', { rel: 'icon', type: 'image/png', sizes: '192x192', href: '/logo-192.png' }],
  ['link', { rel: 'icon', type: 'image/png', sizes: '1024x1024', href: '/icon.png' }],
  ['link', { rel: 'apple-touch-icon', href: '/icon.png' }],
  ['meta', { property: 'og:image', content: 'https://bspec.dev/icon.png' }]
]
```

## Where It Appears

1. **Navigation Header**: 192px logo in top-left corner
2. **Browser Tab**: 64px favicon
3. **Social Media**: 1024px high-quality image for link previews
4. **Mobile Home Screen**: 1024px icon when saved as app
5. **Search Results**: Favicon in search engine results

## Build Status

✅ **Build successful** (1.48s)
✅ **All files present** in dist folder
✅ **Logo displays** in navigation and favicons

## File Sizes

### Dark Mode
- **Favicon (64px)**: 6.7KB - Fast loading for browser tabs
- **Header (192px)**: 34KB - Optimal for navigation logo  
- **Icon (1024px)**: 2.2MB - High quality for social/app use
- **Full (1024px)**: 2.2MB - Original source file

### Light Mode (Inverted)
- **Favicon (64px)**: 6.2KB - Fast loading for browser tabs
- **Header (192px)**: 34KB - Optimal for navigation logo  
- **Icon (1024px)**: 1.8MB - High quality for social/app use
- **Full (1024px)**: 1.8MB - Inverted source file

## Maintenance

To update the logo:

1. Replace `/assets/images/icon.png` with new design
2. Regenerate optimized sizes:
   ```bash
   cd apps/web/public
   
   # Copy original
   cp ../../../assets/images/icon.png bspec-logo-full.png
   
   # Create dark mode versions
   sips -z 1024 1024 bspec-logo-full.png --out icon.png
   sips -z 192 192 bspec-logo-full.png --out logo-192.png
   sips -z 64 64 bspec-logo-full.png --out logo.png
   
   # Create inverted light mode versions
   magick bspec-logo-full.png -channel RGB -negate bspec-logo-full-light.png
   sips -z 1024 1024 bspec-logo-full-light.png --out icon-light.png
   sips -z 192 192 bspec-logo-full-light.png --out logo-192-light.png
   sips -z 64 64 bspec-logo-full-light.png --out logo-light.png
   ```
3. Rebuild: `bun run build`

### Requirements
- **sips**: Built-in macOS tool (already available)
- **ImageMagick**: Install with `brew install imagemagick`

## Notes

- PNG format maintains transparency
- All sizes optimized for web performance
- Maintains visual quality at all resolutions
- Source file (1024x1024) provides flexibility for future use
- **Light mode logos**: Automatically inverted using ImageMagick `-channel RGB -negate`
- **Theme switching**: VitePress automatically uses correct logo based on user's theme preference
- **Inverted colors**: RGB channels are negated while preserving alpha transparency

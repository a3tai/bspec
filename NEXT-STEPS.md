# Next Steps for BSpec Deployment

The new documentation site is ready! Here's what to do next to get it live.

## Immediate Steps

### 1. Test the Site Locally ✅
```bash
cd apps/web
bun install
bun run dev
# Visit http://localhost:5173
```

The site builds successfully and is ready for deployment.

### 2. Set Up Cloudflare Pages

You have two options:

#### Option A: Automatic Deployment (Recommended)

1. **Get Cloudflare Credentials**
   - Log into Cloudflare Dashboard
   - Get your Account ID from the dashboard
   - Create an API token with Cloudflare Pages permissions

2. **Add GitHub Secrets**
   - Go to your GitHub repo > Settings > Secrets and Variables > Actions
   - Add two secrets:
     - `CLOUDFLARE_API_TOKEN`
     - `CLOUDFLARE_ACCOUNT_ID`

3. **Push to GitHub**
   ```bash
   git add .
   git commit -m "Rebuild documentation site with VitePress"
   git push origin main
   ```
   
   The GitHub Action will automatically deploy!

#### Option B: Manual Deployment

```bash
cd apps/web
bun run build

# Install wrangler globally if not already installed
npm install -g wrangler

# Login to Cloudflare
wrangler login

# Deploy
wrangler pages deploy .vitepress/dist --project-name=bspec-docs
```

### 3. Configure Custom Domain

1. Go to Cloudflare Pages dashboard
2. Select your project (bspec-docs)
3. Navigate to Custom Domains
4. Add `bspec.dev`
5. Configure DNS:
   ```
   Type: CNAME
   Name: @
   Target: bspec-docs.pages.dev
   Proxy: Enabled (orange cloud)
   ```

### 4. Update GitHub Repository Links

Update all placeholder GitHub links in:
- `index.md`
- `about.md`
- `community.md`
- `docs/overview.md`

Replace `https://github.com/yourusername/bspec` with your actual GitHub URL.

## Content to Add

### High Priority

1. **Specification Pages** (`spec/`)
   - `spec/v1-0-0.md` - Full BSpec v1.0 specification
   - `spec/structure.md` - Document structure details
   - `spec/conformance.md` - Conformance levels

2. **Document Types** (`docs/`)
   - `docs/document-types.md` - Overview of all 82 types
   - `docs/domains/*.md` - Pages for each of 11 domains

3. **Tools Documentation** (`docs/tools/`)
   - `docs/tools/cli.md` - CLI tool guide
   - `docs/tools/typescript.md` - TypeScript SDK
   - `docs/tools/python.md` - Python SDK
   - `docs/tools/go.md` - Go SDK
   - `docs/tools/mcp.md` - MCP server integration

### Medium Priority

4. **Examples**
   - Real-world BSpec implementations
   - Industry-specific examples (SaaS, Physical Products, etc.)
   - Template documents for each type

5. **Guides**
   - Migration guide (if applicable)
   - Best practices
   - Common patterns

### Low Priority

6. **Assets**
   - Logo SVG (`public/logo.svg`)
   - OG image for social sharing (`public/og-image.png`)
   - Favicon

## Verification Checklist

After deployment:

- [ ] Site loads at https://bspec.dev
- [ ] All navigation links work
- [ ] Search functionality works
- [ ] Light/dark mode toggle works
- [ ] Mobile responsive design works
- [ ] HTTPS is enabled
- [ ] No dead links (or update config to remove ignoreDeadLinks)

## Clean Up

Once the new site is live and verified:

```bash
# Remove the old site backup
rm -rf apps/web.backup
```

## Monitoring

Set up monitoring:
1. Cloudflare Analytics (built-in)
2. Google Search Console (optional)
3. Uptime monitoring (optional)

## Future Enhancements

Consider adding:
- [ ] Blog section for updates
- [ ] Interactive document type explorer
- [ ] Conformance level calculator
- [ ] Knowledge graph visualizer
- [ ] Code playground for testing documents
- [ ] Integration examples

## Questions?

If you need help with any step:
1. Check `apps/web/DEPLOY.md` for detailed deployment instructions
2. Review `SITE-REBUILD.md` for technical details
3. Consult VitePress documentation: https://vitepress.dev
4. Cloudflare Pages docs: https://developers.cloudflare.com/pages

---

**You're all set!** The hard work is done. Just follow the steps above to get the site live.

# Deployment Guide

This guide covers deploying the BSpec documentation site to Cloudflare Pages.

## Prerequisites

1. A Cloudflare account
2. A custom domain (bspec.dev) configured in Cloudflare
3. GitHub repository with the BSpec code

## Setup Cloudflare Pages

### Option 1: Automatic GitHub Deployment (Recommended)

1. **Create Cloudflare API Token**
   - Go to Cloudflare Dashboard > My Profile > API Tokens
   - Create a token with "Cloudflare Pages" permissions
   - Save the token securely

2. **Add GitHub Secrets**
   - Go to your GitHub repository settings
   - Navigate to Secrets and Variables > Actions
   - Add these secrets:
     - `CLOUDFLARE_API_TOKEN`: Your Cloudflare API token
     - `CLOUDFLARE_ACCOUNT_ID`: Your Cloudflare account ID

3. **Push to Main Branch**
   - The GitHub Actions workflow will automatically deploy on push
   - Check the Actions tab to monitor deployment

### Option 2: Manual Deployment

1. **Build the Site**
   ```bash
   cd apps/web
   bun install
   bun run build
   ```

2. **Deploy with Wrangler**
   ```bash
   # Install wrangler if not already installed
   npm install -g wrangler

   # Login to Cloudflare
   wrangler login

   # Deploy
   wrangler pages deploy .vitepress/dist --project-name=bspec-docs
   ```

3. **Configure Custom Domain**
   - Go to Cloudflare Pages dashboard
   - Select your project (bspec-docs)
   - Navigate to Custom Domains
   - Add your domain (bspec.dev)

## Domain Configuration

### DNS Settings

Make sure your domain's DNS is configured in Cloudflare:

1. **Add DNS Records** (if not already present)
   ```
   Type: CNAME
   Name: @
   Target: bspec-docs.pages.dev
   Proxy: Enabled (orange cloud)
   ```

2. **SSL/TLS Settings**
   - SSL/TLS encryption mode: Full (strict)
   - Always Use HTTPS: Enabled
   - Automatic HTTPS Rewrites: Enabled

## Verify Deployment

1. Visit https://bspec.dev
2. Check that all pages load correctly
3. Test navigation and search functionality
4. Verify light/dark mode works

## Troubleshooting

### Build Fails
- Check Node.js/Bun version compatibility
- Ensure all dependencies are installed
- Review build logs for specific errors

### Site Not Updating
- Clear Cloudflare cache
- Check GitHub Actions logs
- Verify the correct branch is being deployed

### Custom Domain Not Working
- Verify DNS records are correct
- Wait for DNS propagation (can take up to 24 hours)
- Check Cloudflare Pages dashboard for domain status

## Rollback

If you need to rollback to a previous version:

1. Go to Cloudflare Pages dashboard
2. Select your project
3. Navigate to Deployments
4. Find the previous working deployment
5. Click "Rollback to this deployment"

## Monitoring

Monitor your deployment:
- **Analytics**: Cloudflare Pages dashboard
- **Uptime**: Set up Cloudflare Workers Health Checks
- **Errors**: Check browser console and server logs

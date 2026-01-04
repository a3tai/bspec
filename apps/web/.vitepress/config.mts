import { defineConfig } from 'vitepress'

export default defineConfig({
  title: 'BSpec',
  description: 'Business Specification Standard - A universal language for describing any business',
  
  cleanUrls: true,
  ignoreDeadLinks: true,
  
  themeConfig: {
    logo: {
      light: '/logo-192-light.png',
      dark: '/logo-192.png'
    },
    
    nav: [
      { text: 'Documentation', link: '/docs/overview' },
      { text: 'Specification', link: '/spec/v1-0-0' },
      { text: 'Community', link: '/community' },
      { text: 'About', link: '/about' }
    ],

    sidebar: {
      '/docs/': [
        {
          text: 'Get Started',
          items: [
            { text: 'What is BSpec?', link: '/docs/overview' },
            { text: 'Quick Start', link: '/docs/quickstart' },
            { text: 'Core Concepts', link: '/docs/concepts' }
          ]
        },
        {
          text: 'Document Types',
          items: [
            { text: 'Overview', link: '/docs/document-types' },
            { text: 'Strategic Foundation', link: '/docs/domains/strategic-foundation' },
            { text: 'Market Environment', link: '/docs/domains/market-environment' },
            { text: 'Customer Value', link: '/docs/domains/customer-value' },
            { text: 'Product & Service', link: '/docs/domains/product-service' },
            { text: 'Business Model', link: '/docs/domains/business-model' },
            { text: 'Operations & Execution', link: '/docs/domains/operations-execution' },
            { text: 'Technology & Data', link: '/docs/domains/technology-data' },
            { text: 'Financial & Investment', link: '/docs/domains/financial-investment' },
            { text: 'Risk & Governance', link: '/docs/domains/risk-governance' },
            { text: 'Growth & Innovation', link: '/docs/domains/growth-innovation' },
            { text: 'Learning & Decisions', link: '/docs/domains/learning-decisions' },
            { text: 'Brand & Marketing', link: '/docs/domains/brand-marketing' }
          ]
        },
        {
          text: 'Tools & SDKs',
          items: [
            { text: 'CLI Tool', link: '/docs/tools/cli' },
            { text: 'TypeScript SDK', link: '/docs/tools/typescript' },
            { text: 'Python SDK', link: '/docs/tools/python' },
            { text: 'Go SDK', link: '/docs/tools/go' },
            { text: 'MCP Server', link: '/docs/tools/mcp' }
          ]
        }
      ],
      '/spec/': [
        {
          text: 'Specification',
          items: [
            { text: 'BSpec v1.0.0', link: '/spec/v1-0-0' },
            { text: 'Document Structure', link: '/spec/structure' },
            { text: 'Conformance Levels', link: '/spec/conformance' }
          ]
        }
      ]
    },

    socialLinks: [
      { icon: 'github', link: 'https://github.com/a3tai/bspec' }
    ],

    search: {
      provider: 'local'
    },

    footer: {
      message: 'Released under the MIT License.',
      copyright: 'Copyright © 2025 BSpec'
    }
  },

  head: [
    ['link', { rel: 'icon', type: 'image/png', sizes: '64x64', href: '/logo.png' }],
    ['link', { rel: 'icon', type: 'image/png', sizes: '192x192', href: '/logo-192.png' }],
    ['link', { rel: 'icon', type: 'image/png', sizes: '1024x1024', href: '/icon.png' }],
    ['link', { rel: 'apple-touch-icon', href: '/icon.png' }],
    ['meta', { name: 'theme-color', content: '#f97316' }],
    ['meta', { property: 'og:type', content: 'website' }],
    ['meta', { property: 'og:locale', content: 'en' }],
    ['meta', { property: 'og:site_name', content: 'BSpec' }],
    ['meta', { property: 'og:image', content: 'https://bspec.dev/icon.png' }]
  ]
})

import { themes as prismThemes } from 'prism-react-renderer';

/** @type {import('@docusaurus/types').Config} */
const config = {
  title: 'Polaris',
  tagline: 'High performance workflow orchestrator for Golang',
  favicon: 'img/favicon.ico',

  url: 'https://harshadmanglani.github.io',
  baseUrl: '/polaris',

  organizationName: 'harshadmanglani',
  projectName: 'polaris',

  onBrokenLinks: 'throw',
  onBrokenMarkdownLinks: 'warn',
  trailingSlash: false,

  i18n: {
    defaultLocale: 'en',
    locales: ['en'],
  },

  presets: [
    [
      'classic',
      ({
        docs: {
          routeBasePath: '/',
          sidebarPath: './sidebars.js',
        },
        blog: false,
        theme: {
          customCss: './src/css/custom.css',
        },
        gtag: {
          trackingID: 'G-E1B8MVW2TR',
          anonymizeIP: true,
        },
      })
    ],
  ],

  themeConfig: {
    image: 'img/logo.png',
    navbar: {
      title: 'Polaris',
      logo: {
        alt: 'Polaris Logo',
        src: 'img/light.svg',
        srcDark: 'img/dark.svg'
      },
      items: [
        {
          type: 'doc',
          docId: 'get-started',
          position: 'left',
          label: 'Docs',
        },
        {
          type: 'doc',
          docId: 'usage',
          position: 'left',
          label: `Usage`,
        },
        {
          type: 'doc',
          docId: 'upcoming',
          position: 'left',
          label: `Upcoming`
        },
        {
          position: 'right',
          href: 'https://go.dev',
          html: `<img src="img/go-logo.png" alt="Go" height="45" width="48" style="vertical-align: middle;"></img>`
        },
        {
          href: 'https://github.com/harshadmanglani/polaris',
          position: 'right',
          className: 'header-github-link',
        },
        {
          href: 'https://x.com/polaris_golang/',
          position: 'right',
          className: 'header-twitter-link'
        }
      ],
    },
    footer: {
      style: 'dark',
      links: [
        {
          title: 'Docs',
          items: [
            {
              label: 'Introduction',
              to: '/',
            },
          ],
        },
        {
          title: 'Community',
          items: [
            {
              label: 'Twitter',
              href: 'https://twitter.com/polaris_golang',
            },
          ],
        }
      ],
      copyright: `Copyright © ${new Date().getFullYear()} Polaris, Inc. Built with Docusaurus.`,
    },
    prism: {
      theme: prismThemes.github,
      darkTheme: prismThemes.dracula,
    },
    algolia: {
      appId: 'D94ISND3YS',
      apiKey: '42e3d00f1c1db0e4c63f37ec9ae22cfb',
      indexName: 'polaris',
      insights: true,
      contextualSearch: false,
    }
  },
};

export default config;

import {themes as prismThemes} from 'prism-react-renderer';

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
      }),
    ],
  ],

  themeConfig:
    ({
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
            docId: 'api',
            position: 'left',
            label: `API`,
            class: 'hidden' // TODO: remove this when API reference is done
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
                href: 'https://twitter.com/PolarisGithub',
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
    }),
};

export default config;

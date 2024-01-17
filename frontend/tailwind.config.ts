import type { Config } from 'tailwindcss'
import plugin from "tailwindcss/plugin";

const config: Config = {
  content: [
    './pages/**/*.{js,ts,jsx,tsx,mdx}',
    './components/**/*.{js,ts,jsx,tsx,mdx}',
    './app/**/*.{js,ts,jsx,tsx,mdx}',
  ],
  theme: {
    extend: {
      fontFamily: {
        header: ['Syne', 'sans-serif'],
      },
      backgroundImage: {
        'gradient-radial': 'radial-gradient(var(--tw-gradient-stops))',
        'gradient-conic':
          'conic-gradient(from 180deg at 50% 50%, var(--tw-gradient-stops))',
      },
      colors: {
        "background": "#111111",
        "button": "#DB2777",
        "button-hover": "#EC4899",
        "CAT1": "#DB2877",
        "CAT2": "#C4D600",
        "CAT3": "#00B5E2",
      },
      fill: (theme) => ({
        ...theme.colors
      }),
      stroke: (theme) => ({
        ...theme.colors
      })
    },
  },
  plugins: [
    plugin(function({ addBase, theme }) {
      addBase({
        ':root': {
          '--color-background': theme('colors.background'),
          '--color-button': theme('colors.button'),
          '--color-button-hover': theme('colors.button-hover'),
          '--color-CAT1': theme('colors.CAT1'),
          '--color-CAT2': theme('colors.CAT2'),
          '--color-CAT3': theme('colors.CAT3'),
        },
        ".rotate-x-180": {
          transform: "rotateX(180deg)",
        },
        ".preserve-3d": {
          transformStyle: "preserve-3d",
        },
        ".perspective": {
          perspective: "1000px",
        },
        ".backface-hidden": {
          backfaceVisibility: "hidden",
        },
      });
    })
  ]
}
export default config

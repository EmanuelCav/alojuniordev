import type { Config } from "tailwindcss";

const config: Config = {
  content: [
    "./src/pages/**/*.{js,ts,jsx,tsx,mdx}",
    "./src/components/**/*.{js,ts,jsx,tsx,mdx}",
    "./src/app/**/*.{js,ts,jsx,tsx,mdx}",
  ],
  theme: {
    extend: {
      inset: {
        'eye-left': '92%',
      },
      colors: {
        'background-payload': 'rgba(0, 0, 0, 0.5)',
      },
      fontSize: {
        "lgXl": '1.175rem'
      },
      boxShadow: {
        "gray-shadow": '0 1px 2px 0 #dddddd'
      },
      backgroundImage: {
        "gradient-radial": "radial-gradient(var(--tw-gradient-stops))",
        "gradient-conic":
          "conic-gradient(from 180deg at 50% 50%, var(--tw-gradient-stops))",
      },
    },

  },
  plugins: [],
};
export default config;

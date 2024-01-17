/** @type {import('next').NextConfig} */
const nextConfig = {
  reactStrictMode: true,
};

const env = {};

module.exports = {
  nextConfig,
  env,
  images: {
    domains: ["images.wallpaperscraft.com"],
  },
  webpack(config) {
    config.module.rules.push({
      test: /\.svg$/,
      use: ["@svgr/webpack"],
    });

    return config;
  },
};

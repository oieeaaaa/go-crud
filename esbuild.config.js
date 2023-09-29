import * as esbuild from 'esbuild'
import postCssPlugin from 'esbuild-style-plugin'
import tailwindcss from 'tailwindcss'
import autoprefixer from 'autoprefixer'

console.log("⚡ Client server is initializing...");

let ctx = await esbuild.context({
  entryPoints: ['./client/client.js'],
  bundle: true,
  outdir: './assets',
  plugins: [postCssPlugin({
    postcss: {
      plugins: [
        tailwindcss,
        autoprefixer,
      ],
    }
  })],
})

await ctx.watch();

console.log("🟢 Client server is watching for changes...");

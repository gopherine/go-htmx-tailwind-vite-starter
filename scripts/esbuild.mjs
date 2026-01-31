import { context } from "esbuild";
import { promises as fs } from "fs";
import path from "path";
import { generateModuleRegistry } from "./gen-modules.mjs";

const args = new Set(process.argv.slice(2));
const isWatch = args.has("--watch");
const isMinify = args.has("--minify");

const outdir = path.resolve("public");
const entryPoints = {
  app: path.resolve("views/scripts/index.ts"),
};

async function writeManifest(metafile) {
  const entryNameByPath = new Map(
    Object.entries(entryPoints).map(([name, file]) => [path.resolve(file), name]),
  );

  const manifest = {};
  for (const [outfile, info] of Object.entries(metafile.outputs)) {
    if (!info.entryPoint) continue;
    const entryName = entryNameByPath.get(path.resolve(info.entryPoint));
    if (!entryName) continue;
    const rel = path.relative(outdir, outfile).split(path.sep).join("/");
    manifest[entryName] = rel;
  }

  await fs.mkdir(outdir, { recursive: true });
  await fs.writeFile(
    path.join(outdir, "manifest.json"),
    JSON.stringify(manifest, null, 2),
    "utf8",
  );
}

const manifestPlugin = {
  name: "manifest-writer",
  setup(build) {
    build.onEnd(async (result) => {
      if (result.errors.length > 0 || !result.metafile) {
        return;
      }
      try {
        await writeManifest(result.metafile);
      } catch (err) {
        console.error("Failed to write manifest:", err);
      }
    });
  },
};

async function start() {
  await generateModuleRegistry();

  const ctx = await context({
    entryPoints,
    outdir,
    bundle: true,
    format: "esm",
    splitting: true,
    target: ["es2020"],
    metafile: true,
    minify: isMinify,
    sourcemap: isWatch,
    entryNames: isWatch ? "assets/[name]" : "assets/[name]-[hash]",
    chunkNames: isWatch ? "assets/[name]" : "assets/[name]-[hash]",
    assetNames: isWatch ? "assets/[name]" : "assets/[name]-[hash]",
    plugins: [manifestPlugin],
    logLevel: "info",
  });

  if (isWatch) {
    await ctx.watch();
    fs.watch(path.resolve("views/scripts/modules"), { recursive: false }, async () => {
      try {
        await generateModuleRegistry();
        await ctx.rebuild();
      } catch (err) {
        console.error("Module registry rebuild failed:", err);
      }
    });
    console.log("esbuild watching...");
    return;
  }

  await ctx.rebuild();
  await ctx.dispose();
}

start().catch((err) => {
  console.error(err);
  process.exit(1);
});

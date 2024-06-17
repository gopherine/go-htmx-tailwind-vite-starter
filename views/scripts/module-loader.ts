let manifestCache: any = null;
const loadedModules = new Set<string>();

async function fetchManifest() {
    if (!manifestCache) {
        const response = await fetch(`/public/.vite/manifest.json?v=${new Date().getTime()}`);
        manifestCache = await response.json();
    }
    return manifestCache;
}

export async function loadScriptByManifest(entryName: string) {
    try {
        const manifest = await fetchManifest();
        if (manifest[entryName]) {
            const entry = manifest[entryName];

            // Load CSS files if any
            if (entry.css) {
                entry.css.forEach((cssPath: string) => {
                    const link = document.createElement('link');
                    link.rel = 'stylesheet';
                    link.href = `/public/${cssPath}`;
                    document.head.appendChild(link);
                });
            }

            // Load JS module
            const scriptPath = entry.file;
            const script = document.createElement('script');
            script.type = 'module';
            script.src = `/public/${scriptPath}`;
            document.head.appendChild(script);
        } else {
            console.error(`Entry ${entryName} not found in manifest`);
        }
    } catch (error) {
        console.error('Error loading manifest:', error);
    }
}

export async function loadModule(moduleName: string) {
    if (loadedModules.has(moduleName)) {
        console.log(`Module ${moduleName} already loaded`);
        return;
    }

    try {
        const manifest = await fetchManifest();
        const moduleEntry:any = Object.values(manifest).find((entry: any) => entry.src && entry.src.endsWith(`${moduleName}.ts`));

        if (moduleEntry) {
            // Load CSS files if any
            if (moduleEntry.css) {
                moduleEntry.css.forEach((cssPath: string) => {
                    const link = document.createElement('link');
                    link.rel = 'stylesheet';
                    link.href = `/public/${cssPath}`;
                    document.head.appendChild(link);
                });
            }

            // Load JS module and execute default export if present
            const module = await import(`/public/${moduleEntry.file}?v=${new Date().getTime()}`);
            if (module.default) {
                module.default();
            }

            loadedModules.add(moduleName);
        } else {
            console.error(`Module ${moduleName} not found in manifest`);
        }
    } catch (error) {
        console.error(`Error loading module ${moduleName}:`, error);
    }
}

document.addEventListener('DOMContentLoaded', async () => {
    console.log("index.ts: DOMContentLoaded event triggered");
    await loadScriptByManifest('views/scripts/index.ts');

    // Automatically load modules based on data-module attribute
    const elements = document.querySelectorAll('[data-module]');
    elements.forEach(element => {
        const moduleName = element.getAttribute('data-module');
        if (moduleName) {
            loadModule(moduleName);
        }
    });
});
//@ts-ignore
const modules = import.meta.glob('./modules/*.ts');

async function loadModule(moduleName: string) {
  const modulePath = `./modules/${moduleName}.ts`;
  console.log(`Attempting to load module: ${moduleName} from ${modulePath}`);
  if (modules[modulePath]) {
    try {
      const mod = await modules[modulePath]();
      console.log(`Loaded module: ${moduleName}`, mod);
      if (mod && mod.default) {
        console.log("Executing module default export");
        mod.default();  // Call the default exported function
      } else {
        console.error(`Module ${moduleName} does not have a default export`);
      }
    } catch (error) {
      console.error(`Error loading module ${moduleName}:`, error);
    }
  } else {
    console.error(`Module ${modulePath} not found`);
  }
}

document.addEventListener('DOMContentLoaded', () => {
  console.log("index.ts: DOMContentLoaded event triggered");
  
  // Automatically load modules based on data-module attribute
  const elements = document.querySelectorAll('[data-module]');
  elements.forEach(element => {
    const moduleName = element.getAttribute('data-module');
    if (moduleName) {
      console.log(`Element with data-module found: ${moduleName}`);
      loadModule(moduleName);
    }
  });
});
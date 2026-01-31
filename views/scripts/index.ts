import { moduleLoaders, type ModuleName } from "./modules/registry";

type ModuleInit = (el: HTMLElement) => void;

const initialized = new WeakSet<Element>();

async function initModule(el: HTMLElement, name: ModuleName) {
  const loader = moduleLoaders[name];
  if (!loader) {
    console.error(`Module ${name} not found in registry`);
    return;
  }

  try {
    const mod = await loader();
    if (mod?.default) {
      (mod.default as ModuleInit)(el);
    }
  } catch (error) {
    console.error(`Error loading module ${name}:`, error);
  }
}

function scanAndInit(root: ParentNode = document) {
  const elements = root.querySelectorAll<HTMLElement>("[data-module]");
  elements.forEach((el) => {
    if (initialized.has(el)) return;
    const name = el.dataset.module as ModuleName | undefined;
    if (!name) return;
    initialized.add(el);
    initModule(el, name);
  });
}

document.addEventListener("DOMContentLoaded", () => {
  scanAndInit(document);
});

document.addEventListener("htmx:afterSwap", (event) => {
  const target = (event as CustomEvent).detail?.target as ParentNode | undefined;
  if (target) {
    scanAndInit(target);
  }
});

type CounterState = {
  value: number;
};

function clamp(value: number) {
  return Math.max(-999, Math.min(999, value));
}

export default function initialize(el: HTMLElement) {
  const initial = Number(el.dataset.initial ?? "0");
  const state: CounterState = {
    value: Number.isFinite(initial) ? initial : 0,
  };

  const valueEl = el.querySelector<HTMLElement>("[data-counter-value]");
  const doubleEl = el.querySelector<HTMLElement>("[data-counter-double]");
  const noteEl = el.querySelector<HTMLElement>("[data-counter-note]");

  function render() {
    if (valueEl) valueEl.textContent = String(state.value);
    if (doubleEl) doubleEl.textContent = String(state.value * 2);
    if (noteEl) noteEl.textContent = `Local state is ${state.value}.`;
  }

  el.addEventListener("click", (event) => {
    const target = event.target as HTMLElement | null;
    if (!target) return;
    const action = target.getAttribute("data-action");
    if (!action) return;

    switch (action) {
      case "increment":
        state.value = clamp(state.value + 1);
        break;
      case "decrement":
        state.value = clamp(state.value - 1);
        break;
      case "reset":
        state.value = 0;
        break;
      default:
        return;
    }

    render();
  });

  render();
}

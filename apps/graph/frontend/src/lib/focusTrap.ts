type FocusTrapOptions = {
  onEscape: () => void;
};

const focusableSelector = [
  'a[href]',
  'button:not([disabled])',
  'input:not([disabled])',
  'select:not([disabled])',
  'textarea:not([disabled])',
  '[tabindex]:not([tabindex="-1"])',
].join(',');

export function focusTrap(node: HTMLElement, initialOptions: FocusTrapOptions) {
  let options = initialOptions;
  const previousFocus = document.activeElement instanceof HTMLElement ? document.activeElement : null;
  const backdrop = node.closest('.modal-backdrop');
  const backgroundElements = backdrop?.parentElement
    ? [...backdrop.parentElement.children]
        .filter((element): element is HTMLElement => element instanceof HTMLElement && element !== backdrop)
        .map((element) => ({
          element,
          inert: element.inert,
          ariaHidden: element.getAttribute('aria-hidden'),
        }))
    : [];
  for (const background of backgroundElements) {
    background.element.inert = true;
    background.element.setAttribute('aria-hidden', 'true');
  }

  function focusableElements(): HTMLElement[] {
    return [...node.querySelectorAll<HTMLElement>(focusableSelector)]
      .filter((element) => !element.hidden && element.getAttribute('aria-hidden') !== 'true');
  }

  function handleKeydown(event: KeyboardEvent) {
    if (event.key === 'Escape') {
      event.preventDefault();
      options.onEscape();
      return;
    }
    if (event.key !== 'Tab') return;
    const elements = focusableElements();
    if (elements.length === 0) {
      event.preventDefault();
      node.focus();
      return;
    }
    const first = elements[0];
    const last = elements[elements.length - 1];
    if (event.shiftKey && document.activeElement === first) {
      event.preventDefault();
      last.focus();
    } else if (!event.shiftKey && document.activeElement === last) {
      event.preventDefault();
      first.focus();
    }
  }

  node.addEventListener('keydown', handleKeydown);
  queueMicrotask(() => (focusableElements()[0] ?? node).focus());

  return {
    update(nextOptions: FocusTrapOptions) {
      options = nextOptions;
    },
    destroy() {
      node.removeEventListener('keydown', handleKeydown);
      for (const background of backgroundElements) {
        background.element.inert = background.inert;
        if (background.ariaHidden === null) background.element.removeAttribute('aria-hidden');
        else background.element.setAttribute('aria-hidden', background.ariaHidden);
      }
      previousFocus?.focus();
    },
  };
}

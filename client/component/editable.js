class Editable extends HTMLElement {
  get _oldSay() {
    return document.querySelector('[name="OldSay"]')
  }

  syncHiddenInput = (e) => {
    this._oldSay.value = e.target.outerText
  }

  constructor() {
    super();

    this.addEventListener('input', this.syncHiddenInput);
    this.addEventListener('click', this.syncHiddenInput);
  }
}

if (!customElements.get('x-editable')) {
  customElements.define('x-editable', Editable);
}

---
marp: true
theme: gaia
paginate: true
footer: '#JSBangkok #JSBKK'
---

<style>
  @import url('https://fonts.googleapis.com/css2?family=K2D:ital,wght@0,100;0,200;0,300;0,400;0,500;0,600;0,700;0,800;1,100;1,200;1,300;1,400;1,500;1,600;1,700;1,800&display=swap');

  h1, b, strong,
  h2, h3, h4,
  a, li,
  p {
    font-family: "K2D", sans-serif;
  }

  h1, a, b, strong {
    color: rgb(255 102 51) !important;
  }

  section {
    background: white;
  }

  footer {
    font-size: 20px;
    text-align: left;
  }
</style>


# Develop modern HTML components with Web Components
### A 20 Minutes Introduction

<!--
(Sat) 19 October 2024, 14:35 @Hall 2 at True Digital Park West

Web Components have been around for more than 10 years. They should be one of the concepts we teach to young developers, but somehow, with the rise of promising JavaScript frameworks like Next.js, Nuxt.js, and various others, developing modern frontend applications with Web Components is still not widely known. In this talk, I want to show how we can develop modern JavaScript components with only native APIs that most browsers support. How they can be used with or inside most JavaScript frameworks without problems.
-->
---

# Overview
- HTML components and Web Components
- Core Technologies of Web Components
<!-- - Advantages and Practices -->

<!--
HTML Components are built-in likes <p> <h1> <div> <section>
Web Components are a set of technologies that offer encapsulation and
  reusable components
So, to develop Web Components could said that it is to develop custom HTML components.

-->

---

# Web Components

- Custom Elements
- Shadow DOM
- Templates and Slots

<!--
Introduce Web Components as a way to create encapsulated, reusable elements natively in the browser without frameworks. Mention key technologies and their growing importance in 2024.
-->

---

# Custom Elements

- Custom HTML tags created by developers.
- Extend existing elements or create entirely new ones.
- Lifecycles:
  - `connectedCallback()`: inserted into the DOM.
  - `disconnectedCallback()`: removed from the DOM.
  - `attributeChangedCallback()`: attributes change.

---

## Ex 1: Basic Custom Element
```js
class MyElement extends HTMLElement {
  connectedCallback() {
    this.innerHTML = "<p>Hello, World!</p>";
  }
}
customElements.define('my-element', MyElement);
```
```html
<!-- Usage: -->
<my-element></my-element>
```
> Simple example of a custom element that adds "Hello, World!" to the DOM.

---

## Ex 2: Custom Button with Attributes
```js
class MyButton extends HTMLElement {
  constructor() {
    super();
    this.addEventListener('click', () => alert('Button clicked!'));
  }
  connectedCallback() {
    this.innerHTML = `<button>${this.getAttribute('label')}</button>`;
  }
}
customElements.define('my-button', MyButton);
```
```html
<!-- Usage: -->
<my-button label="Click me"></my-button>
```

---

## Ex 3: Customize Built-in Element

```js
class ClickableParagraph extends HTMLParagraphElement {
  constructor() {
    super();
    this.addEventListener('click', () => alert('Paragraph clicked!'));
  }
  connectedCallback() {
    this.innerHTML = `Click me: ${this.getAttribute('content')}`;
    this.style.cursor = 'pointer'; // Make it clear that the element is clickable.
    this.style.color = 'blue'; // Add some style.
  }
}
customElements.define('clickable-p', MyParagraph, { extends: 'p' });
```

```html
<!-- Usage: -->
<p is="clickable-p" content="Click me for something"></p>
```


<!-- ---

# Custom Elements: Pros

- **Encapsulation:** Custom elements bundle HTML, CSS, and JS behavior.
- **Reusability:** Can be reused across different parts of the application.
- **Framework-agnostic:** Works in any environment (React, Vue, Angular).
- **Native browser support:** Supported without external libraries.

---

# Custom Elements: Cons
- **Browser compatibility:** Some features might require polyfills in older browsers.
- **Learning curve:** Developers need to learn about lifecycle callbacks and custom element APIs.
- **Styling challenges:** Managing scoped styles and integration with external stylesheets can be tricky. -->

---
<!-- Intentional -->
## Ex 5: Custom Modal Element

```html
<!-- Usage: -->
<custom-modal id="myModal">
  <p slot="modal-text">This is your modal text.</p>
</custom-modal>
<button onclick="document.getElementById('myModal').open()">Open Modal</button>
```
---
<!--
<style>
  .modal {
    display: none;
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background-color: rgba(0, 0, 0, 0.5);
  }
  .modal-content {
    background-color: white;
    margin: 15% auto;
    padding: 20px;
    width: 80%;
  }
</style>
-->
```js
class CustomModal extends HTMLElement {
  constructor() {
    super();
    this.attachShadow({ mode: 'open' });
    this.shadowRoot.innerHTML = `
      <style>
        .modal {
          display: none;
          // redact because there is no space here.

        }
        .modal-content {
          background-color: white;
          // redact because there is no space here.
        }
      </style>
      <div class="modal">
        <div class="modal-content">
          <span slot="modal-text">
            This is not a modal text you're looking for.
          </span>
        </div>
      </div>
    `;
  }

  connectedCallback() {
    this.modal = this.shadowRoot.querySelector('.modal');
    this.modal.addEventListener('click', () => this.close());
  }

  open() {
    this.modal.style.display = 'block';
  }

  close() {
    this.modal.style.display = 'none';
  }
}

customElements.define('custom-modal', CustomModal);
```
---

# Shadow DOM

- An encapsulated DOM subtree.
- Isolates styles and markup from the main document.
- Avoid CSS and JavaScript conflicts in complex web applications.
    - Also, make it harder to access Global CSS.

---

## Ex : Basic Shadow DOM
```js
class MyShadowElement extends HTMLElement {
  constructor() {
    super();
    this.attachShadow({ mode: 'open' });
    this.shadowRoot.innerHTML = `
      <style> p { color: blue; } </style>
      <p>Hello from Shadow DOM</p>
    `;
  }
}
customElements.define('my-shadow-element', MyShadowElement);
```
```html
<!-- Usage: -->
<my-shadow-element></my-shadow-element>
```
---

# Declarative Shadow DOM
<!--
- Introduced to make Shadow DOM more declarative in HTML.
- Useful for server-side rendering (SSR).
- Allows attaching shadow roots directly in HTML.
-->

```html
<my-shadow-element>
  <template shadowroot="open">
    <style>
      p { color: blue; }
    </style>
    <p>Hello from Declarative Shadow DOM</p>
  </template>
</my-shadow-element>
```

<!-- # Working with Shadow DOM

- ShadowRoot, shadow trees, and light DOM
- Styling within Shadow DOM
  - Scoping and sharing styles
  - Handling CSS isolation

-->

---

 # Shadow DOM: Pros and Cons

<!-- Pros: -->
- Isolates component styles and DOM structure from the rest of the page.
- Reusable encapsulated components.
- Styles defined within a shadow root are isolated from the global scope.

<!-- Cons: -->
- Developers must learn how to manage scoped styles and lifecycle events.
- Components in different shadow trees cannot easily communicate.

---

# Ex : Custom Card Component

```html
<!-- Usage: -->
<custom-card>
  <h2 slot="title">Card Title</h2>
  <p slot="content">Card content goes here.</p>
</custom-card>
```

---

```js
class CustomCard extends HTMLElement {
  constructor() {
    super();
    this.attachShadow({ mode: 'open' });
    this.shadowRoot.innerHTML = `
      <style>
        .card {
          border: 1px solid #ddd;
          padding: 20px;
          box-shadow: 2px 2px 10px rgba(0, 0, 0, 0.1);
        }
      </style>
      <div class="card">
        <slot name="title"> [object Object] </slot>
        <slot name="content"> [object Object] </slot>
      </div>
    `;
  }
}
customElements.define('custom-card', CustomCard);
```

---

# Templates

- A way to define HTML chunks for later reuse.
- Templates are not rendered when the page loads, only when explicitly instantiated.

---

# Templates Example

```html
<template id="myTemplate">
  <p>This is content from the template.</p>
</template>
<!-- <template> is hidden and won't appear on the page initially. -->


<div id="contentArea"></div>
```

```js
const template = document.getElementById('myTemplate');
const contentArea = document.getElementById('contentArea');

contentArea.appendChild(template.content.cloneNode(true));
```
---

# Templates: Pros and Cons

<!-- Pros -->
- HTML structures defined once and reused multiple times.
- Templates are not displayed until needed, improving performance.

<!-- Cons -->
- Must be activated via JavaScript.
- Templates are static by nature.

---

# Slots

- Mechanism for distributing content inside custom elements.
- Named and default slots for content distribution.

---

# Advantages of Using Web Components

- ...

---

# Advantages of Using Web Components


- Encapsulation of styles and behavior
- Reusability across frameworks (React, Vue, Angular)
- Can fit into most project structures

---

# Where to go next

https://custom-elements-everywhere.com/

(google "github custom elements everywhere")

<!--
Looking at the React full support in 19
Survey Repo: testing in each framework.
-->

---
# Q & A

---

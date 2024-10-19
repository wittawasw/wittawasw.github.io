---
marp: true
theme: gaia
paginate: true
# footer: '#JSBangkok #JSBKK'
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
### A 25 Minutes Introduction

<!--
(Sat) 19 October 2024, 14:35 @Hall 2 at True Digital Park West

Web Components have been around for more than 10 years. They should be one of the concepts we teach to young developers, but somehow, with the rise of promising JavaScript frameworks like Next.js, Nuxt.js, and various others, developing modern frontend applications with Web Components is still not widely known. In this talk, I want to show how we can develop modern JavaScript components with only native APIs that most browsers support. How they can be used with or inside most JavaScript frameworks without problems.
-->
---

### Overview
- Introduction to Web Components
- Custom Elements, Shadow DOM, Templates, Slots
- Advantages and Practices
- Sample and Future of Web Components

<!--
introduce the topics we’ll cover in the next 30 minutes, focusing on how Web Components are revolutionizing modern JavaScript development.
-->

---

# Introduction to Web Components

- Native browser support
- Core technologies: Custom Elements, Shadow DOM, Templates, Slots
- Evolving adoption in 2024

<!--
Introduce Web Components as a way to create encapsulated, reusable elements natively in the browser without frameworks. Mention key technologies and their growing importance in 2024.
-->

---

# Core Technologies: Custom Elements

## What are Custom Elements?
- Custom HTML tags created by developers.
- Extend existing elements or create entirely new ones.
- Lifecycles:
  - `connectedCallback()`: When the element is inserted into the DOM.
  - `disconnectedCallback()`: When the element is removed from the DOM.
  - `attributeChangedCallback()`: When attributes change.

## Example 1: Basic Custom Element
```js
class MyElement extends HTMLElement {
  connectedCallback() {
    this.innerHTML = "<p>Hello, World!</p>";
  }
}
customElements.define('my-element', MyElement);
```
Usage:
```html
<my-element></my-element>
```
- Simple example of a custom element that adds "Hello, World!" to the DOM.

## Example 2: Custom Button with Attributes
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
Usage:
```html
<my-button label="Click me"></my-button>
```

<!--
Explain how Custom Elements can encapsulate behavior and DOM content. In this example, we extend the concept to build an interactive custom button element.
-->

---

# Custom Elements: Pros and Cons

## Pros:
- **Encapsulation:** Custom elements bundle HTML, CSS, and JS behavior.
- **Reusability:** Can be reused across different parts of the application.
- **Framework-agnostic:** Works in any environment (React, Vue, Angular).
- **Native browser support:** Supported without external libraries.

## Cons:
- **Browser compatibility:** Some features might require polyfills in older browsers.
- **Learning curve:** Developers need to learn about lifecycle callbacks and custom element APIs.
- **Styling challenges:** Managing scoped styles and integration with external stylesheets can be tricky.

<!--
Discuss the pros and cons of using Custom Elements. While they offer encapsulation and reusability, they also come with compatibility and learning challenges.
-->

---

# Practical Use Case: Custom Modal Element

## Example: Modal for a Real Website
```js
class CustomModal extends HTMLElement {
  constructor() {
    super();
    this.attachShadow({ mode: 'open' });
    this.shadowRoot.innerHTML = `
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
      <div class="modal">
        <div class="modal-content">
          <slot></slot>
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
Usage in a real website:
```html
<custom-modal id="myModal">
  <p>Modal Content Here</p>
</custom-modal>
<button onclick="document.getElementById('myModal').open()">Open Modal</button>
```

<!--
This is a practical example of how a custom modal could be implemented in a real-world website. It demonstrates the power of Custom Elements to create reusable components like a modal dialog.
-->

---
# Core Technologies: Shadow DOM

## What is Shadow DOM?
- Encapsulated DOM subtree.
- Isolates styles and markup from the main document.
- Helps avoid CSS and JavaScript conflicts in complex web applications.

## Example: Basic Shadow DOM
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
Usage:
```html
<my-shadow-element></my-shadow-element>
```

## Declarative Shadow DOM
- Introduced to make Shadow DOM more declarative in HTML.
- Useful for server-side rendering (SSR).
- Allows attaching shadow roots directly in HTML.

Example:
```html
<template shadowroot="open">
  <p>Declarative Shadow DOM Content</p>
</template>
```

<!--
Explain Shadow DOM’s importance in encapsulation. Discuss both programmatic and declarative approaches, with an example showing how Declarative Shadow DOM enhances server-side rendering.
-->

---

# Shadow DOM: Pros and Cons

## Pros:
- **Encapsulation:** Isolates component styles and DOM structure from the rest of the page.
- **Reusability:** Enables encapsulated, reusable components.
- **Scoped Styles:** Styles defined within a shadow root are isolated from the global scope.

## Cons:
- **Learning curve:** Developers must learn how to manage scoped styles and lifecycle events.
- **SEO Impact:** Declarative Shadow DOM is a step towards addressing SEO challenges but needs server-side support.
- **Limited cross-shadow communication:** Components in different shadow trees cannot easily communicate.

---

# Practical Use Case: Shadow DOM in Real Websites

## Example: Custom Card Component with Shadow DOM
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
        <slot name="title"></slot>
        <slot name="content"></slot>
      </div>
    `;
  }
}
customElements.define('custom-card', CustomCard);
```
Usage in HTML:
```html
<custom-card>
  <h2 slot="title">Card Title</h2>
  <p slot="content">Card content goes here.</p>
</custom-card>
```

<!--
Showcase how Shadow DOM can be used for creating a custom card component in real-world websites. The use of scoped styles ensures the card’s design doesn’t interfere with the page’s global styles.
-->

---

# Core Technologies: Templates

## What are Templates?
- A way to define HTML chunks for later reuse.
- Templates are not rendered when the page loads, only when explicitly instantiated.

Example:
```html
<template id="myTemplate">
  <p>This is template content.</p>
</template>
```
JavaScript:
```js
const template = document.getElementById('myTemplate');
document.body.appendChild(template.content.cloneNode(true));
```

---

# Templates: Pros and Cons

## Pros:
- **Efficiency:** HTML structures defined once and reused multiple times.
- **Deferred Rendering:** Templates are not displayed until needed, improving performance.

## Cons:
- **No styling or functionality on initial render:** Must be activated via JavaScript.
- **Complexity in dynamic content:** Templates are static by nature.

---

# Core Technologies: Slots

## What are Slots?
- Mechanism for distributing content inside custom elements.
- Named and default slots for content distribution.

Example:
```html
<my-element>
  <span slot="header">Header Content</span>
</my-element>
```
JavaScript:
```js
class MyElement extends HTMLElement {
  constructor() {
    super();
    this.attachShadow({ mode: 'open' });
    this.shadowRoot.innerHTML = `
      <slot name="header"></slot>
      <p>Default content goes here.</p>
    `;
  }
}
customElements.define('my-element', MyElement);
```

---

# Slots: Pros and Cons

## Pros:
- **Flexibility:** Allows the user to insert custom content.
- **Modularity:** Encapsulates layout while letting users control the content.

## Cons:
- **Initial complexity:** Developers need to understand how slots work.
- **Content management:** More challenging when working with dynamic content.

---

# Practical Example: Card with Template and Slots

## Customizable Card with Shadow DOM, Template, and Slots
```js
class CustomCard extends HTMLElement {
  constructor() {
    super();
    this.attachShadow({ mode: 'open' });
    this.shadowRoot.innerHTML = `
      <style>
        .card {
          padding: 10px;
          border: 1px solid #000;
        }
      </style>
      <div class="card">
        <slot name="header"></slot>
        <slot></slot>
      </div>
    `;
  }
}
customElements.define('custom-card', CustomCard);
```
Usage:
```html
<custom-card>
  <h1 slot="header">Title</h1>
  <p>Card content goes here.</p>
</custom-card>
```

---


# Advantages of Using Web Components

- Encapsulation of styles and behavior
- Reusability across frameworks (React, Vue, Angular)
- Performance benefits

<!--
Discuss the key advantages of using Web Components such as style encapsulation, cross-framework usage, and performance improvements.
-->

---

# Working with Shadow DOM

- ShadowRoot, shadow trees, and light DOM
- Styling within Shadow DOM
  - Scoping and sharing styles
  - Handling CSS isolation

<!--
Dive deeper into the Shadow DOM, explain how styles work, and give examples of scoped styling. Discuss challenges with style sharing.
-->

---

# custom-elements-everywhere

https://custom-elements-everywhere.com/

---

# Closing

- Recap of advantages and future potential
- Resources for learning and contributing to the ecosystem

<!--
Summarize the key points. Mention the future trends for Web Components.
-->

---

# References

https://eisenbergeffect.medium.com/web-components-2024-winter-update-445f27e7613a
https://web.dev/articles/declarative-shadow-dom
https://utilitybend.com/blog/getting-into-web-components-an-intro/
https://developer.mozilla.org/en-US/docs/Web/API/ShadowRoot

https://stackoverflow.com/questions/34119639/what-is-shadow-root
https://web.dev/articles/shadowdom
https://web.dev/articles/shadowdom-v1
https://web.dev/articles/declarative-shadow-dom

https://developer.mozilla.org/en-US/docs/Web/API/Web_components/Using_shadow_DOM

https://developer.mozilla.org/en-US/docs/Web/API/Web_components#guides
https://developer.mozilla.org/en-US/docs/Web/API/Web_components/Using_custom_elements
https://developer.mozilla.org/en-US/docs/Web/API/Web_components/Using_shadow_DOM
https://developer.mozilla.org/en-US/docs/Web/API/Web_components/Using_templates_and_slots

https://github.com/mdn/web-components-examples

https://custom-elements-everywhere.com/

https://eisenbergeffect.medium.com/using-global-styles-in-shadow-dom-5b80e802e89d

https://eisenbergeffect.medium.com/html-attributes-properties-and-values-752b6eed8c21
https://eisenbergeffect.medium.com/sharing-styles-in-declarative-shadow-dom-c5bf84ffd311

https://stackoverflow.com/questions/11871065/monads-in-javascript
https://en.wikipedia.org/wiki/Function_composition

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

  footer {
    font-size: 20px;
    text-align: left;
  }
</style>

# Exploring Native Web Components

### A 25-minute Guide

---

# What Are Native Web Components?

- Reusable, encapsulated HTML elements
- Works natively in the browser without frameworks
- **Core Features**:
  - Shadow DOM
  - Custom Elements
  - HTML Templates

<!--
Speaker Notes:
- Native Web Components allow developers to create custom HTML tags.
- They are framework-agnostic and work in any browser that supports them.
- Shadow DOM provides encapsulation for style and structure, while Custom Elements enable defining custom HTML tags.
-->

---

# Creating a Simple Web Component

### Steps:
1. Define a class that extends `HTMLElement`.
2. Use `attachShadow()` to encapsulate DOM.
3. Register the element with `customElements.define()`.

```javascript
class MyElement extends HTMLElement {
  constructor() {
    super();
    this.attachShadow({ mode: 'open' });
    this.shadowRoot.innerHTML = `<p>Hello World!</p>`;
  }
}

customElements.define('my-element', MyElement);
```

<!--
Speaker Notes:
- This is a basic Web Component implementation.
- Shadow DOM encapsulates the component's DOM and CSS, ensuring it doesn’t affect the rest of the page.
- \`customElements.define\` is how the browser recognizes your new HTML element.
-->

---

# Class vs Module Pattern

### Class-Based Web Components

```js
class MyComponent extends HTMLElement {
  // Web component logic...
}
```

- **Advantages**: Clear structure, lifecycle methods (\`connectedCallback\`, etc.).
- **Disadvantages**: Can lead to complex inheritance and harder testing.

---

# Class vs Module Pattern

### Module Pattern

```js
function pushRequest() {
  // Function logic...
}
```

- **Advantages**: Functional composition, no inheritance.
- **Disadvantages**: Scattered state, less lifecycle management.

<!--
Speaker Notes:
- Class-based Web Components are the standard as they offer lifecycle methods like \`connectedCallback\`.
- Modules allow for a more functional approach, but state management becomes more difficult, especially without clear lifecycle callbacks.
-->

---

# Plugin Architecture in Web Components

- **Pattern Example**:
  ```js
  Chartkick.use(Chart);
  ```
- **Advantages**:
  - Decouples core functionality from external plugins.
  - Provides flexibility and modularity.
- **Web Component Use Case**:
  - You can allow plugins to extend or modify Web Component functionality.

<!--
Speaker Notes:
- Plugins allow developers to inject or extend functionality in a Web Component.
- This keeps the core small and extensible.
-->

---

# Using Web Components in React

- Import Web Components directly into React JSX:
  ```jsx
  <my-element></my-element>
  ```
- **Passing Props**: Use `ref` and `useEffect` to interact with the Web Component’s attributes.
  ```jsx
  useEffect(() => {
    ref.current.setAttribute('prop-name', 'value');
  }, []);
  ```
- **Handling Events**: Use `ref` to attach custom event listeners:
  ```jsx
  ref.current.addEventListener('custom-event', handler);
  ```

<!--
Speaker Notes:
- Web Components can easily integrate with React by treating them like any other HTML element.
- You’ll need to use \`ref\`s in React to pass data or listen for events.
-->

---

# Best Practices for Web Components

- **Encapsulation**: Use Shadow DOM for style and script encapsulation.
- **Composition**: Build small, reusable components.
- **Interoperability**: Ensure compatibility across frameworks.

<!--
Speaker Notes:
- Encapsulation helps avoid style bleeding into or out of components.
- Compose small components rather than building large, complex ones.
- Always keep interoperability in mind so the components are reusable in different frameworks.
-->

---

# Conclusion and Q&A

- Recap: Flexibility and reusability of Web Components
- Open for questions.

<!--
Speaker Notes:
- Summarize key takeaways from the talk.
- Encourage questions on implementation, architecture, and potential use cases.
-->

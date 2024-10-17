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

<!-- ### Steps: -->
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

# Plugin Architecture

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
<!--
Speaker Notes:
- Web Components can easily integrate with React by treating them like any other HTML element.
- You’ll need to use \`ref\`s in React to pass data or listen for events.
-->

---

# Using Web Components in React

- **Handling Events**: Use `ref` to attach custom event listeners:
  ```jsx
  ref.current.addEventListener('custom-event', handler);
  ```
---

# Practices for Web Components

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

---

# References

https://eisenbergeffect.medium.com/web-components-2024-winter-update-445f27e7613a
https://web.dev/articles/declarative-shadow-dom
https://utilitybend.com/blog/getting-into-web-components-an-intro/
https://developer.mozilla.org/en-US/docs/Web/API/ShadowRoot

https://stackoverflow.com/questions/34119639/what-is-shadow-root
https://web.dev/articles/shadowdom
https://web.dev/articles/shadowdom-v1

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


# New Outline

# Presentation Outline: Developing Modern JavaScript with Web Components (30 minutes)

---

## 1. Introduction to Web Components (3 min)
- Definition and importance of Web Components in modern web development.
- Overview of core technologies: Custom Elements, Shadow DOM, Templates, and Slots.
- Evolution of Web Components and their growing adoption in 2024.

## 2. Core Technologies of Web Components (8 min)
- **Custom Elements**
  - Definition and lifecycle (connectedCallback, disconnectedCallback).
  - Creating and registering custom elements.
- **Shadow DOM**
  - Encapsulation and DOM isolation using Shadow DOM.
  - Declarative Shadow DOM: benefits and usage.
- **Templates and Slots**
  - Defining reusable HTML structures.
  - Using slots for flexible content insertion.

## 3. Advantages of Using Web Components (4 min)
- Encapsulation of styles and behavior.
- Reusability across frameworks (React, Vue, Angular).
- Performance benefits: lightweight, native browser support.

## 4. Working with Shadow DOM (7 min)
- ShadowRoot, shadow trees, and light DOM.
- **Styling within Shadow DOM:**
  - Scoping and sharing styles using global styles.
  - Handling CSS isolation and external stylesheets.
- Real-world examples of Shadow DOM usage.

## 5. Best Practices and Design Patterns (5 min)
- Managing attributes, properties, and events in Web Components.
- Sharing and inheriting styles in declarative Shadow DOM.
- Integration with frameworks and libraries.

## 6. Hands-on Demo and Code Walkthrough (3 min)
- Simple demo: creating a custom button component with Shadow DOM.
- Exploring live examples from GitHub.

## 7. Conclusion and Future of Web Components (2 min)
- Recap of the key advantages and future potential.
- Resources for learning and contributing to the Web Components ecosystem.

---

**Q&A (3 min)**

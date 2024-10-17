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


# Developing Modern JavaScript with Web Components
### A 25 Minutes Introduction

<!--
introduce speaker
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

- Define custom HTML elements
- Lifecycle methods:
  - `connectedCallback`
  - `disconnectedCallback`

<!--
Explain how Custom Elements work and the key lifecycle callbacks. Show the simple code example to demonstrate defining and using custom elements.
-->

---
# Core Technologies: Custom Elements

Example:
```js
class MyElement extends HTMLElement {
  connectedCallback() {
    this.innerHTML = "<p>Hello, World!</p>";
  }
}
customElements.define('my-element', MyElement);
```
---

# Core Technologies: Shadow DOM

- DOM encapsulation and style isolation
- Light DOM vs. Shadow DOM
- Declarative Shadow DOM

Example:
```html
<template id="shadow">
  <style> p { color: red; } </style>
  <p>Shadow DOM content</p>
</template>
```

<!--
Introduce the concept of Shadow DOM for encapsulating DOM and styles. Mention declarative Shadow DOM and its benefits. Briefly show an example of how it works.
-->

---

# Core Technologies: Templates & Slots

- Reusable HTML structures
- Slots for flexible content insertion

Example:
```html
<my-element>
  <span slot="header">Header Content</span>
</my-element>
```

<!--
Explain the importance of templates and slots in allowing content to be injected and reused. Demonstrate with a simple slot example.
-->

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

# Best Practices and Design Patterns

- Managing attributes, properties, and events
- Sharing and inheriting styles in Shadow DOM
- Integration with frameworks

<!--
Cover best practices such as managing component state and properties. Mention design patterns for effective use of Web Components.
-->

---

# Hands-on Demo: Custom Button Component

- Demo: Create a custom button using Shadow DOM
- Walkthrough of GitHub examples

<!--
Show a simple demo of creating a custom button component with Shadow DOM and guide the audience through some real-world GitHub examples.
-->

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

---
marp: true
theme: gaia
math: katex
footer: Built with [marp.app](https://marp.app)
---
<style>
  h1, a, b, strong {
    color: #EE9E0B !important;
  }

  footer {
    font-size: 20px;
    text-align: right;
  }
</style>


<script type="module">
  import mermaid from 'https://cdn.jsdelivr.net/npm/mermaid@10/dist/mermaid.esm.min.mjs';
  mermaid.initialize({ startOnLoad: true });
</script>

# **mermaid**

<div class="mermaid">
  graph LR;
  a --> b;
  b --> c;
  c --> a;
</div>

---

# Table

Fruit | Colour | Amount | Cost
-----|------|:-----:|------:
Banana | Yellow | 4 | £1.00
Apple | Red | 2 | £0.60
Orange | Orange | 10 | £2.50
Coconut | Brown | 1 | £1.50

---

Render inline math such as $ax^2+bc+c$. :smile:

$$ I_{xx}=\int\int_Ry^2f(x,y)\cdot{}dydx $$

$$
f(x) = \int_{-\infty}^\infty
    \hat f(\xi)\,e^{2 \pi i \xi x}
    \,d\xi
$$

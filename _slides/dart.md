---
marp: true
theme: default
class: invert
---

<style>
  :root {
    --color-1: #34ba25;
    --color-2: #6bd385;
    --color-3: #c8fad4;
  }

  h1, b, strong {
    color: var(--color-1) !important;
  }

  a, p {
    color: var(--color-2) !important;
  }

  p {
    color: var(--color-3) !important;
  }

  footer {
    font-size: 20px;
    text-align: right;
  }
</style>

# **Dart**

---
# Dart Runtime

- install (fvm) (homebrew)
  ```sh
  dart run main.dart
  dart create xspring-app
  ```
- [dartpad](https://dartpad.dev/)


---
# Basic File Structure

```
- bin
- lib
- test
- pubspec.yaml
- pubspec.lock
```

---

# Testing in Dart and OOP

- modify test file
- test class `Calculator`
- implement `Calculator` with `plus`
- remove `as`;

---

# OOP cont.

```dart
class Calculator {
  final _a;
  final _b;

  Calculator(this._a, this._b);

  int plus() {
    return _a + _b;
  }
}

```

---
# OOP cont.

```dart
class Calculator {
  final _a;
  final _b;

  Calculator(this._a, this._b);

  int plus() {
    return _a + _b;
  }
}

```

---

# OOP extend

```dart
class SuperCalculator extends Calculator {
  SuperCalculator(int a, int b) : super(a, b);

  @override
  int plus() {
    return super.plus() + 1;
  }
}

```

---

# Async, Await

```dart
Future<int> plusAsync() async {
  return _a + _b;
}

test('plusAsync', () async {
  int a = 1;
  int b = 2;

  Calculator cal = Calculator(a, b);

  expect(await cal.plusAsync(), 3);
});
```

---

# Flutter

- structure (similar to above)
- introduce [GetX](https://pub.dev/packages/get#getxservice)

---

# Flutter

- Initializing
- main.dart -> setup.dart -> app.dart
- Flavor and Target
- null safety `?`, `!`


---

# app.dart

- Route

---

# GetX

- pass params
- GetX service, controller -> Life cycle
  - controller, return mem once unused
  - service, longer live
- `Get.arguments`
- `Get.Put`, `Get.find`
- `Get.toNamed`

# SPEC

**Job:** Build a game in Go out of entities, components and systems, without taking on a
game framework.

**Why:** Every hobby engine rewrites the same entity store and the same system loop.
Writing it once, with no dependencies, lets a game pick its own rendering and input
libraries and keep them.

**Guardrails:**

- Zero third-party dependencies. That is the feature, not a preference.
- No `main` package, no embedded assets, no logging — the consumer decides all three.
- The package doc comment in [doc.go](doc.go) is the long form of *Job*. When it and this
  file disagree, this file wins.
- Named decisions and any waived baseline rule live in [README.md](README.md).
- Still v0. The API may break between minor tags, and the tag message says what broke.

**Done means:**

- `checklists/library.md` in the [baseline](https://github.com/andygeiss/baseline) is
  walked, and every box is checked or waived on the record in the README.
- `make ci` is green on the commit being pushed.

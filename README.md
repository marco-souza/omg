## OMG (Ollama Modelfile Generator)

<p align="center">
<svg width="20" height="20">
  <rect x="3" y="5" width="8" height="8" rx="1" fill="#ffd700" />
  <text x="10" y="9" font-size="10" text-anchor="middle" fill="#333">GitHub Stars</text>
</svg>
<svg width="20" height="20">
  <rect x="3" y="5" width="8" height="8" rx="1" fill="#333" />
  <text x="10" y="9" font-size="10" text-anchor="middle" fill="#fff">GitHub Code</text>
</svg>
<svg width="20" height="20">
  <rect x="3" y="5" width="8" height="8" rx="1" fill="#666" />
  <text x="10" y="9" font-size="10" text-anchor="middle" fill="#fff">GitHub License</text>
</svg>
<svg width="20" height="20">
  <rect x="3" y="5" width="8" height="8" rx="1" fill="#999" />
  <text x="10" y="9" font-size="10" text-anchor="middle" fill="#fff">GitHub Issues</text>
</svg>
<svg width="20" height="20">
  <rect x="3" y="5" width="8" height="8" rx="1" fill="#444" />
  <text x="10" y="9" font-size="10" text-anchor="middle" fill="#fff">GitHub Projects</text>
</svg>
</p>

OMG is a CLI tool written in Go that helps you generate Modelfile for Ollama, given a initial request from the user.

https://www.loom.com/share/d6a1509e0b90474c850e8fb946aa19a6

**Usage:**

```
omg [request]
```

**Options:**

- **Help:** `omg -h`

**Requirements:**

- Go
- Langchain
- Ollama

**Installation:**

```
go install github.com/marco-souza/omg
```

**Example Usage:**

```
omg I want model expert in python which can teach me how to become a great developer
```

This will generate a Modelfile for an Ollama model like the following one

**Additional Notes:**

- The CLI command does not have any options besides help.
- OMG uses langchain for communicating with ollama.
- To use OMG, you must have ollama installed.

**Further Resources:**

- [Ollama Documentation](/docs)
- [Langchain Documentation](/langchain)


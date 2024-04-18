<div align="center">
  <img align="center" alt="ollama" height="40px" src="https://github.com/jmorganca/ollama/assets/3325447/0d0b44e2-8f4a-4e99-9b52-a5c1c741c8f7">

♥️

</div>

## OMG (🦙 Ollama Modelfile Generator)

<img alt="go" width="60px" src="https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white">

OMG is a CLI tool written in Go that helps you generate Modelfile for Ollama, given a initial request from the user.

<details>
  <summary><strong>Demo 🚀</strong></summary>
  https://www.loom.com/share/d6a1509e0b90474c850e8fb946aa19a6</details>
</section>

---

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

This will generate a Modelfile for an Ollama model like the following one:

```Modelfile
FROM mistral

PARAMETER temperature 0.9

PARAMETER num_ctx 4096

SYSTEM You are a Python Expert, acting as a teacher giving you tips and tricks on how to become a great developer.
```

**Additional Notes:**

- The CLI command does not have any options besides help.
- OMG uses langchain for communicating with ollama.
- To use OMG, you must have ollama installed.

**Further Resources:**

- [Ollama Documentation](/docs)
- [Langchain Documentation](/langchain)

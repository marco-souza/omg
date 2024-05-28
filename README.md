<div align="center">
  <img align="center" alt="ollama" height="40px" src="https://github.com/jmorganca/ollama/assets/3325447/0d0b44e2-8f4a-4e99-9b52-a5c1c741c8f7">

♥️

</div>

## OMG (🦙 Ollama Modelfile Generator)

<img alt="go" width="60px" src="https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white">

OMG is a CLI tool written in Go that helps you generate Modelfile for Ollama, given a initial request from the user.

**Features:**

- generate `Modelfile` for Ollama using `llama3` and `langchain`

<details>
  <summary><strong>Demo 🚀</strong></summary>
  https://www.loom.com/share/d6a1509e0b90474c850e8fb946aa19a6</details>
</section>

---

**Usage:**

```
omg -h
omg [request]
omg -o <output file> [request]
```

**Requirements:**

- Go
- Langchain
- Ollama

**Installation:**

```
go install github.com/marco-souza/omg
```

**Example Usage:**

```sh
# usage infos
omg -h

# without output flag
omg I want model expert in python which can teach me how to become a great developer
omg "I want model expert in python which can teach me how to become a great developer" > python.modelfile

# with output flag
omg -o history.modelfile I want to study history
omg -o lang.modelfile I want to how languages where created history


# reading from input
echo "I want to study math" | omg
echo "I want to study math" | omg -o math.modelfile
```

The command above will generate a `Modelfile` for a custom Ollama Model, like the following one:

```Modelfile
FROM mistral

PARAMETER temperature 0.9

PARAMETER num_ctx 4096

SYSTEM You are a Python Expert, acting as a teacher giving you tips and tricks on how to become a great developer.
```

You can use this file to build and run your [custom model](https://github.com/ollama/ollama?tab=readme-ov-file#customize-a-model).

**Additional Notes:**

- To use OMG, you must have ollama installed.

**Further Resources:**

- [Ollama Documentation](https://github.com/ollama/ollama)
- [Langchain Documentation](https://tmc.github.io/langchaingo/docs/)

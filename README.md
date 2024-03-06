## OMG (Ollama Modelfile Generator)

OMG is a CLI tool written in Go that helps you generate Modelfile for Ollama, given a initial request from the user.

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
omg "I want model expert in python which can teach me how to become a great developer"
```

This will generate a Modelfile for an Ollama model like the following one

**Additional Notes:**

- The CLI command does not have any options besides help.
- OMG uses langchain for communicating with ollama.
- To use OMG, you must have ollama installed.

**Further Resources:**

- [Ollama Documentation](/docs)
- [Langchain Documentation](/langchain)


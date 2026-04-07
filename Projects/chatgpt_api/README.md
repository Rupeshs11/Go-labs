# 🤖 ChatGPT API

A simple Go command-line application that demonstrates how to interact with the OpenAI GPT-3 API.

## 📝 Overview

This project initializes an OpenAI GPT-3 client and requests a text completion using the prompt: `"The first thing you should know about javascript is"`. It limits the response to 30 tokens and stops at the first period `.`.

## 🛠️ Setup Instructions

1. Clone the repository and navigate to the project directory:

   ```bash
   cd Go-labs/Projects/chatgpt_api
   ```

2. Open the `.env` file and add your OpenAI API Key:

   ```env
   API_KEY=your_openai_api_key_here
   ```

3. Download dependencies:

   ```bash
   go mod tidy
   ```

4. Run the application:
   ```bash
   go run main.go
   ```

## 📦 Dependencies

- [`godotenv`](https://github.com/joho/godotenv) - To load environment variables from the `.env` file.
- [`go-gpt3`](https://github.com/PullRequestInc/go-gpt3) - Unofficial Go client for the OpenAI GPT-3 API.

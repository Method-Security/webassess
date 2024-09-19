#!/usr/bin/env bash

ollama serve &
ollama list
ollama pull qwen:0.5b-chat
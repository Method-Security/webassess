#!/usr/bin/env bash

ollama serve &
sleep 2
ollama list
ollama pull qwen2.5-coder:1.5b
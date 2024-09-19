0. Test in container manually off whatever image needed
- works by doing serve before install
- then serve just needs to be run as a cmd of like the & syntax below

1. make the docker image and test it
- might need to do multiplee entrypoints with like `ollama serve & <then cmd>`
2. make the model configurable at a root level
3. do checks that ollama is installed
BREAK - write RFC
4. figure out headless browser
5. Need to test with large context windows
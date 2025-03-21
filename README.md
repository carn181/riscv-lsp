# RISC-V LSP
This is a project to learn about LSP's and implement support for LSP features when writing RISC-V code on linux x86 systems  
Work In Progress.

## Guide
The LSP assumes you have a `build` file in the workspace which has the command you use to build your workspace and see errors for.  
Please change `symbolsDir` in `symbols/symbols.go` to the absolute path of where `symbols/riscv-docs` is stored on your system.  
Will fix this by making a separate documentation package  

## Todo
- [x] Document Synchronization
- [x] Workspace
- [ ] Concurrent Handling of Requests
- [ ] Better Error Handling
- [ ] Socket support instead of Stdin/Stdout
- [x] Diagnostic Checking
    - [x] Assembler Errors
    - [ ] Linker Errors
    - [ ] Better Build File System
- [x] Hover Documentation Support
    - [x] Registers/Opcodes Documentation
    - [x] Make Registers/Opcodes Documentation a Separate File
- [ ] TreeSitter
- [ ] Code-Completion
- [ ] Document Symbols
- [ ] Formatting

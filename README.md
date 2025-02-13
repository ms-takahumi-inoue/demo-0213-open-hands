# open-hands-demo-0213
This project is a simple TODO application built with Golang. It supports the following operations:

- Display all TODO items
- Add a new item to the list
- Mark an item as completed
- Remove an item from the list
- Edit the text of an item
- Set a deadline for an item

The application provides an HTTP server with APIs to interact with the TODO list, intended for use with a frontend built using HTML and JavaScript.

## Project Structure

```
ProjectRoot
├── .github
│   └── workflows
├── cmd
│   └── main.go
├── go.mod
├── .gitignore
└── README.md
```

### Data Persistence
Data is stored in memory using a map, without the use of a database.

### Git Workflow
- Use the linked repository for version control.
- The default branch is `main`.
- Create a Pull Request for changes, perform a self-review, and merge into `main` if there are no issues.
- Provide a clear title and summary for each Pull Request.
- Update the README.md with an overview of the application at appropriate milestones.

# Task Tracker CLI (Go)

This is my solution for the **Task Tracker** project from roadmap.sh.  
👉 Project URL: https://roadmap.sh/projects/task-tracker

A simple command-line application written in **Go** that allows you to manage tasks by creating, updating, deleting, and changing their status. All tasks are stored persistently in a local JSON file.

---

## 🚀 Features

- Add new tasks  
- Update an existing task  
- Delete a task  
- Mark a task as:
  - `todo`
  - `in-progress`
  - `done`
- List all tasks  
- Filter tasks by status  
- Automatic creation and updating of `tasks.json`

---

## 🛠️ Technologies Used

- **Go 1.22+**
- Standard library only (`os`, `encoding/json`, `flag`, etc.)
- Local JSON file for data storage

---

## 📦 Installation

```bash
git clone https://github.com/YOUR-USERNAME/YOUR-REPO.git
cd YOUR-REPO

go build -o task-cli
```

---

## 🕹️ Usage

```bash
./task-cli add "Go to the gym"
./task-cli update 1 "Go to the gym and do leg day"
./task-cli delete 1

./task-cli mark-in-progress 2
./task-cli mark-done 2

./task-cli list
./task-cli list todo
./task-cli list in-progress
./task-cli list done


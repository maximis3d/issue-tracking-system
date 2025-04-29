# issue-tracking-system


## Build & Run (Docker)

(Ensure docker desktop is running)

`docker compose up --build`


```
issue-tracking-system
├─ .dockerignore
├─ Dockerfile
├─ Makefile
├─ README.md
├─ cmd
│  ├─ api
│  │  └─ api.go
│  ├─ main.go
│  └─ migrate
│     ├─ main.go
│     └─ migrations
│        ├─ 20250223175431_add-user-table.down.sql
│        ├─ 20250223175431_add-user-table.up.sql
│        ├─ 20250319130716_add-project-table.down.sql
│        ├─ 20250319130716_add-project-table.up.sql
│        ├─ 20250327115317_issues.down.sql
│        ├─ 20250327115317_issues.up.sql
│        ├─ 20250402094431_standups.down.sql
│        ├─ 20250402094431_standups.up.sql
│        ├─ 20250415151455_scopes.down.sql
│        ├─ 20250415151455_scopes.up.sql
│        ├─ 20250415151725_project_scopes.down.sql
│        ├─ 20250415151725_project_scopes.up.sql
│        ├─ 20250428195934_project_assignments.down.sql
│        └─ 20250428195934_project_assignments.up.sql
├─ config
│  └─ env.go
├─ db
│  ├─ db.go
│  └─ issue_tracking_dump.sql
├─ diagrams
│  ├─ Untitled Diagram(1).drawio
│  ├─ Untitled Diagram.drawio
│  ├─ arch_overview(1).drawio
│  ├─ arch_overview.drawio
│  ├─ arch_overview.jpg
│  ├─ db-design.png
│  ├─ overall-architecture.jpg
│  ├─ user-flow-diagram(1).jpg
│  ├─ user-flow-diagram(2).jpg
│  ├─ user-flow-diagram.jpg
│  └─ user-flow-diagram_final.jpg
├─ docker-compose.yml
├─ frontend
│  ├─ Dockerfile
│  ├─ README.md
│  ├─ dist
│  │  ├─ assets
│  │  │  ├─ index-BOrLqaUr.js
│  │  │  └─ index-Bk_gmLgd.css
│  │  ├─ index.html
│  │  └─ vite.svg
│  ├─ eslint.config.js
│  ├─ index.html
│  ├─ package-lock.json
│  ├─ package.json
│  ├─ public
│  │  └─ vite.svg
│  ├─ src
│  │  ├─ App.jsx
│  │  ├─ api
│  │  │  ├─ issue.js
│  │  │  ├─ project.js
│  │  │  ├─ standup.js
│  │  │  └─ user.js
│  │  ├─ assets
│  │  │  └─ react.svg
│  │  ├─ components
│  │  │  ├─ dropDown
│  │  │  │  └─ dropDown.jsx
│  │  │  ├─ issueCreationForm
│  │  │  │  └─ issueCreation.jsx
│  │  │  ├─ loginForm
│  │  │  │  └─ loginForm.jsx
│  │  │  ├─ navbar
│  │  │  │  └─ navbar.jsx
│  │  │  ├─ projectCreationForm
│  │  │  │  └─ projectCreation.jsx
│  │  │  └─ registerForm
│  │  │     └─ registerForm.jsx
│  │  ├─ index.css
│  │  ├─ main.jsx
│  │  └─ pages
│  │     ├─ EditIssue.jsx
│  │     ├─ Issue.jsx
│  │     ├─ LandingPage.jsx
│  │     ├─ Project.jsx
│  │     ├─ ProjectAssignment.jsx
│  │     ├─ Projects.jsx
│  │     └─ StandUp.jsx
│  └─ vite.config.js
├─ go.mod
├─ go.sum
├─ service
│  ├─ auth
│  │  ├─ jwt.go
│  │  ├─ jwt_test.go
│  │  ├─ password.go
│  │  └─ password_test.go
│  ├─ issue
│  │  ├─ routes.go
│  │  ├─ routes_test.go
│  │  └─ store.go
│  ├─ project
│  │  ├─ routes.go
│  │  ├─ routes_test.go
│  │  └─ store.go
│  ├─ project_assignment
│  │  ├─ routes.go
│  │  └─ store.go
│  ├─ project_scopes
│  │  ├─ routes.go
│  │  └─ store.go
│  ├─ standups
│  │  ├─ routes.go
│  │  └─ store.go
│  └─ user
│     ├─ routes.go
│     ├─ routes_test.go
│     └─ store.go
├─ types
│  └─ types.go
└─ utils
   └─ utils.go

```
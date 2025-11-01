# To Do List Application

A full-stack to-do list application built with **Flutter** (frontend) and **Go** (backend), featuring user authentication, task management, and task history tracking.

## 🏗️ Project Structure

```
to_do_list_app/
├── lib/                    # Flutter application source code
│   ├── Authtentication/   # Authentication and main screens
│   │   ├── login.dart     # Login screen
│   │   ├── signup.dart    # Sign up screen
│   │   ├── home.dart      # Main home screen with navigation
│   │   ├── HomeScreen.dart # Task display screen
│   │   ├── CreateTask.dart # Task creation screen
│   │   ├── Histroy.dart   # Task history screen
│   │   └── ProfilePage.dart # User profile screen
│   ├── assets/            # Application assets (images)
│   └── main.dart          # Application entry point
├── myapp/                  # Go backend server
│   ├── server.go          # Main server file with API endpoints
│   ├── go.mod             # Go dependencies
│   └── go.sum             # Go dependencies checksum
├── android/               # Android platform files
├── ios/                   # iOS platform files
├── web/                   # Web platform files
├── windows/               # Windows platform files
├── pubspec.yaml           # Flutter dependencies
└── README.md              # This file
```

## 🛠️ Prerequisites

Before running this application, ensure you have the following installed:

### Required Software:

1. **Flutter SDK** (version 3.3.0 or higher)

   - Download from: https://flutter.dev/docs/get-started/install
   - Verify installation: `flutter doctor`

2. **Go** (version 1.22.0 or higher)

   - Download from: https://go.dev/dl/
   - Verify installation: `go version`

3. **MySQL Server**

   - Download from: https://dev.mysql.com/downloads/mysql/
   - Or use XAMPP/WAMP which includes MySQL

4. **Code Editor** (Optional but recommended)
   - Visual Studio Code with Flutter and Go extensions
   - Or Android Studio

## 📦 Installation & Setup

### Step 1: Clone the Repository

```bash
git clone <repository-url>
cd to_do_list_app
```

### Step 2: Database Setup

1. **Start MySQL Server**

   - Ensure MySQL is running on `localhost:3306`

2. **Create Database and Tables**

   Open MySQL command line or any MySQL client and run:

   ```sql
   CREATE DATABASE todolist;

   USE todolist;

   -- Create user table
   CREATE TABLE user (
       username VARCHAR(255) PRIMARY KEY,
       email VARCHAR(255) NOT NULL,
       password VARCHAR(255) NOT NULL
   );

   -- Create tasks table
   CREATE TABLE tasks (
       id INT AUTO_INCREMENT PRIMARY KEY,
       title VARCHAR(255) NOT NULL,
       description TEXT,
       start_date DATE NOT NULL,
       end_date DATE NOT NULL,
       start_time TIME NOT NULL,
       end_time TIME NOT NULL,
       status VARCHAR(50) DEFAULT 'incomplete'
   );
   ```

3. **Update Database Credentials** (if needed)

   If your MySQL credentials are different, edit `myapp/server.go`:

   ```go
   cfg := mysql.Config{
       User:   "your_username",     // Change if different
       Passwd: "your_password",     // Change if different
       Net:    "tcp",
       Addr:   "localhost:3306",
       DBName: "todolist",
   }
   ```

### Step 3: Backend (Go Server) Setup

1. **Navigate to Go backend directory**

   ```bash
   cd myapp
   ```

2. **Install Go dependencies**

   ```bash
   go mod download
   ```

   Or if you need to tidy dependencies:

   ```bash
   go mod tidy
   ```

3. **Run the Go server**

   ```bash
   go run server.go
   ```

   You should see:

   ```
   Connected!
   ```

   The server will be running on `http://localhost:8080`

   **Note:** Keep this terminal window open. The server needs to be running for the Flutter app to work.

### Step 4: Frontend (Flutter App) Setup

1. **Open a new terminal window** (keep the Go server running in the first terminal)

2. **Navigate back to project root**

   ```bash
   cd ..  # If you're in myapp directory
   # or
   cd to_do_list_app  # If you're in parent directory
   ```

3. **Install Flutter dependencies**

   ```bash
   flutter pub get
   ```

4. **Verify Flutter setup**
   ```bash
   flutter doctor
   ```

## 🚀 Running the Application

### For Development:

1. **Start MySQL** (if not already running)

2. **Start the Go Backend Server**

   ```bash
   cd myapp
   go run server.go
   ```

   Server will be available at: `http://localhost:8080`

3. **Run Flutter App** (in a new terminal)

   ```bash
   flutter run
   ```

   This will:

   - Build the app
   - Launch on your connected device/emulator
   - Show available devices if multiple are connected

### For Specific Platforms:

#### Android:

```bash
flutter run -d android
```

#### iOS:

```bash
flutter run -d ios
```

#### Web:

```bash
flutter run -d chrome
```

#### Windows:

```bash
flutter run -d windows
```

### Important Notes for Mobile/Emulator:

If running on a physical device or emulator, you may need to change `localhost` to your computer's IP address in the Flutter code:

1. Find your local IP address:

   - Windows: `ipconfig` (look for IPv4 Address)
   - Mac/Linux: `ifconfig` or `ip addr`

2. Update API URLs in Flutter files:
   - Replace `http://localhost:8080` with `http://YOUR_IP:8080`
   - Files to update:
     - `lib/Authtentication/sign_in_button.dart`
     - `lib/Authtentication/signup.dart`
     - `lib/Authtentication/HomeScreen.dart`
     - `lib/Authtentication/CreateTask.dart`
     - `lib/Authtentication/Histroy.dart`

## 📡 API Endpoints

The Go backend provides the following REST API endpoints:

| Method | Endpoint          | Description                     |
| ------ | ----------------- | ------------------------------- |
| GET    | `/`               | Health check endpoint           |
| POST   | `/send-data`      | User login authentication       |
| POST   | `/Createaccount`  | Create new user account         |
| GET    | `/tasks`          | Get all incomplete tasks        |
| POST   | `/Create`         | Create a new task               |
| POST   | `/tasks/complete` | Mark a task as completed        |
| POST   | `/tasks/delete`   | Mark a task as deleted          |
| GET    | `/tasks/history`  | Get task history (all statuses) |

## 🎯 Features

- ✅ User Authentication (Login/Signup)
- ✅ Create Tasks with:
  - Title and Description
  - Start Date and End Date
  - Start Time and End Time
- ✅ View Active Tasks (incomplete tasks with valid end dates)
- ✅ Mark Tasks as Complete
- ✅ Delete Tasks
- ✅ View Task History (all tasks with their statuses)
- ✅ User Profile Page
- ✅ Bottom Navigation Bar
- ✅ Search Functionality (UI ready)

## 🗂️ Database Schema

### `user` Table

| Column   | Type         | Description   |
| -------- | ------------ | ------------- |
| username | VARCHAR(255) | Primary Key   |
| email    | VARCHAR(255) | User email    |
| password | VARCHAR(255) | User password |

### `tasks` Table

| Column      | Type         | Description                                |
| ----------- | ------------ | ------------------------------------------ |
| id          | INT          | Primary Key, Auto Increment                |
| title       | VARCHAR(255) | Task title                                 |
| description | TEXT         | Task description                           |
| start_date  | DATE         | Task start date                            |
| end_date    | DATE         | Task end date                              |
| start_time  | TIME         | Task start time                            |
| end_time    | TIME         | Task end time                              |
| status      | VARCHAR(50)  | Task status (incomplete/completed/deleted) |

## 🐛 Troubleshooting

### Backend Issues:

1. **"Connection refused" error**

   - Ensure MySQL server is running
   - Check MySQL credentials in `server.go`
   - Verify MySQL is listening on port 3306

2. **"Database doesn't exist" error**

   - Run the SQL commands from Step 2 to create the database and tables

3. **"Port 8080 already in use"**
   - Change the port in `server.go` (last line): `e.Start(":8080")`
   - Update Flutter API URLs to match the new port

### Frontend Issues:

1. **"Connection failed" error**

   - Ensure Go server is running on port 8080
   - For mobile/emulator: Use your computer's IP instead of `localhost`
   - Check firewall settings

2. **"Package not found" errors**

   - Run `flutter pub get` again
   - Delete `pubspec.lock` and run `flutter pub get`

3. **Build errors**
   - Run `flutter clean`
   - Run `flutter pub get`
   - Try `flutter run` again

### Database Connection Issues:

1. **MySQL authentication error**

   - Update credentials in `myapp/server.go`
   - Ensure MySQL user has proper permissions

2. **Table doesn't exist**
   - Verify tables were created correctly
   - Check database name matches in connection config

## 📝 Dependencies

### Flutter Dependencies:

- `google_nav_bar: ^5.0.6` - Navigation bar
- `intl: ^0.17.0` - Internationalization
- `http: ^1.2.1` - HTTP client for API calls
- `cupertino_icons: ^1.0.6` - iOS style icons

### Go Dependencies:

- `github.com/labstack/echo/v4` - Web framework
- `github.com/go-sql-driver/mysql` - MySQL driver
- `github.com/rs/cors` - CORS middleware

## 🔧 Development

### Project Structure:

- **Frontend**: Flutter app with Material Design
- **Backend**: Go REST API with Echo framework
- **Database**: MySQL relational database

### Code Style:

- Flutter: Follows Dart style guide
- Go: Follows standard Go formatting (`go fmt`)

**Happy Coding! 🎉**

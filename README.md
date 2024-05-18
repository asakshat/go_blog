# Blog Web App in GO (CRUD)

## Instructions to Build and Run the Application

### Prerequisites
- Ensure Docker is installed on your machine. You can download Docker from [here](https://www.docker.com/products/docker-desktop).

### Steps to Run the Application

#### Step 1: Create a `.env` File
In the root directory of your project, create a `.env` file.

#### Step 2: Set Up Environment Variables
Add the following environment variables to the `.env` file:

```plaintext
SECRET_KEY="your-secret-code"
DATABASE_URL="Your-DB-URL"
```
 Example :- 
```bash
DATABASE_URL=postgres://myuser:mypassword@localhost:5432/mydatabase
```






## API Routes

### Authentication Routes

| **Route**              | **Method** | **Controller**       | **Description**          |
|------------------------|------------|----------------------|--------------------------|
| `/api/user/signup`     | `POST`     | `controllers.Signup` | Signup                   |
| `/api/user/login`      | `POST`     | `controllers.Login`  | Login with Cookies       |
| `/api/user/logout`     | `POST`     | `controllers.Logout` | Logout and remove cookies|
| `/api/user`            | `GET`      | `controllers.User`   | Get logged in user       |

### Blog Routes

| **Route**                           | **Method** | **Controller**                   | **Middleware**           | **Description**            |
|-------------------------------------|------------|----------------------------------|--------------------------|----------------------------|
| `/api/blog/all`                     | `GET`      | `controllers.GetAllPosts`        |                          | Get all blog posts         |
| `/api/blog/:post_id`                | `GET`      | `controllers.GetPostWithIdHandler`|                          | Get blog post by ID        |
| `/api/blog/:user_id`                | `POST`     | `controllers.CreateBlog`         | `middlewares.Authenticate`| Create blog post           |
| `/api/blog/edit/:user_id/:post_id`  | `PUT`      | `controllers.EditPost`           | `middlewares.Authenticate`| Edit blog post             |
| `/api/blog/delete/:user_id/:post_id`| `DELETE`   | `controllers.DeletePost`         | `middlewares.Authenticate`| Delete blog post           |

### Like Routes

| **Route**                              | **Method** | **Controller**          | **Middleware**           | **Description**       |
|----------------------------------------|------------|-------------------------|--------------------------|-----------------------|
| `/api/blog/like/:user_id/:post_id`     | `POST`     | `controllers.LikePost`  | `middlewares.Authenticate`| Like a blog post      |
| `/api/blog/unlike/:user_id/:post_id`   | `POST`     | `controllers.UnlikePost`| `middlewares.Authenticate`| Unlike a blog post    |

### Comment Routes

| **Route**                                    | **Method** | **Controller**             | **Middleware**           | **Description**         |
|----------------------------------------------|------------|----------------------------|--------------------------|-------------------------|
| `/api/blog/comment/post/:user_id/:post_id`   | `POST`     | `controllers.PostComment`  | `middlewares.Authenticate`| Post a comment on a blog|
| `/api/blog/comment/edit/:user_id/:comment_id`| `PUT`      | `controllers.EditComment`  | `middlewares.Authenticate`| Edit a comment          |
| `/api/blog/comment/delete/:user_id/:comment_id`| `DELETE`   | `controllers.DeleteComment`| `middlewares.Authenticate`| Delete a comment        |

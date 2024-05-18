# Blog Web App in GO (CRUD)
If you are just looking for routes click here [API Routes](#api-routes)

## Instructions to Build and Run the Application

### Prerequisites
- Ensure Docker is installed on your machine. You can download Docker from [here](https://www.docker.com/products/docker-desktop).

### Steps to Run the Application locally

#### Step 1: Create a `.env` File
In the root directory of your project, create a `.env` file.

#### Step 2: Set Up Environment Variables
Add the following environment variables to the `.env` file:

```plaintext
SECRET_KEY="YOUR-SECRET-CODE"
DATABASE_URL="YOUR-DB-URL"
```
 Example :- 
<br>
- DATABASE_URL according to **docker-compose.yml** file.

```bash
DATABASE_URL=postgres://myuser:mypassword@localhost:5432/mydatabase
```
```bash
SECRET_KEY=asdasd12asd1
```
#### Step 3: Set Up Environment Variables
Make sure you are in the root directory of the project folder.

Now run : <br>

```bash
docker compose up
```
Now docker should create an image and run it as a container .

To stop you the container you can run: <br>
```bash
docker compose down
```
You can also use Docker Desktop to run/stop containers without the CLI.


## API Routes

### Authentication Routes

| **Description**          | **Route**              | **Method** | **Controller**       |
|--------------------------|------------------------|------------|----------------------|
| Signup                   | `/api/user/signup`     | `POST`     | `controllers.Signup` |
| Login with Cookies       | `/api/user/login`      | `POST`     | `controllers.Login`  |
| Logout and remove cookies| `/api/user/logout`     | `POST`     | `controllers.Logout` |
| Get logged in user       | `/api/user`            | `GET`      | `controllers.User`   |

### Blog Routes

| **Description**            | **Route**                           | **Method** | **Controller**                   |
|----------------------------|-------------------------------------|------------|----------------------------------|
| Get all blog posts         | `/api/blog/all`                     | `GET`      | `controllers.GetAllPosts`        |
| Get blog post by ID        | `/api/blog/:post_id`                | `GET`      | `controllers.GetPostWithIdHandler`|
| Create blog post           | `/api/blog/:user_id`                | `POST`     | `controllers.CreateBlog`         |
| Edit blog post             | `/api/blog/edit/:user_id/:post_id`  | `PUT`      | `controllers.EditPost`           |
| Delete blog post           | `/api/blog/delete/:user_id/:post_id`| `DELETE`   | `controllers.DeletePost`         |

### Like Routes

| **Description**       | **Route**                              | **Method** | **Controller**          |
|-----------------------|----------------------------------------|------------|-------------------------|
| Like a blog post      | `/api/blog/like/:user_id/:post_id`     | `POST`     | `controllers.LikePost`  |
| Unlike a blog post    | `/api/blog/unlike/:user_id/:post_id`   | `POST`     | `controllers.UnlikePost`|

### Comment Routes

| **Description**         | **Route**                                    | **Method** | **Controller**             |
|-------------------------|----------------------------------------------|------------|----------------------------|
| Post a comment on a blog| `/api/blog/comment/post/:user_id/:post_id`   | `POST`     | `controllers.PostComment`  |
| Edit a comment          | `/api/blog/comment/edit/:user_id/:comment_id`| `PUT`      | `controllers.EditComment`  |
| Delete a comment        | `/api/blog/comment/delete/:user_id/:comment_id`| `DELETE`   | `controllers.DeleteComment`|
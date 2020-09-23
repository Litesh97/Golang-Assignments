<h2>Simple REST Api Server in GO</h2>

<strong>Problem Statement</strong>

Create a rest server in Go, which hosts APIs:<br>
■ <restserver_url>/article/create: Creates an object of a struct article, and
saves it into DB and returns article id <br>
■ <restserver_url>/article/{article_id}: Fetches the article from DB and
outputs it in the struct format. <br>
■ <restserver_url>/article/delete/{article_id}: Deletes the entry from DB <br>

<strong>End-Points</strong>
1. POST => /article/create => Creates New Article <br>
2. GET => /article/{article_id} => Fetch an article <br>
3. GET =>  /article/allArticles => Returns all articles <br>
4. DELETE => /article/delete/{article_id} => Delete an article <br>


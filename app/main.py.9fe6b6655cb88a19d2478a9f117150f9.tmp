from typing import Optional
from fastapi import FastAPI, HTTPException, status
from pydantic import BaseModel
import psycopg2
from psycopg2.extras import RealDictCursor
import time
from schemas import Post

app = FastAPI()


while True:
    try:
        conn = psycopg2.connect(
            host='localhost',
            database='tutorial',
            user='postgres',
            password='admin',
            cursor_factory=RealDictCursor
        )
        cursor = conn.cursor()
        print("Connection was successful")
        break
    except Exception as error:
        print("Connection failed")
        print("Error:", error)
        time.sleep(60)


@app.get("/")
async def root():
    try:
        cursor.execute("SELECT * FROM posts")
        posts = cursor.fetchall()
        return {"posts": posts}
    except Exception as error:
        raise HTTPException(
            status_code=500, detail=f"Error fetching posts: {error}")


@app.post("/post")
async def create_post(post: Post):
    try:
        cursor.execute(
            """
            INSERT INTO posts (title, content, published)
            VALUES (%s, %s, %s)
            RETURNING *
            """,
            (post.title, post.content, post.published)
        )
        new_post = cursor.fetchone()
        conn.commit()
        return {"data": new_post}
    except Exception as error:
        conn.rollback()
        raise HTTPException(
            status_code=500, detail=f"Error creating post: {error}")


@app.get("/posts/{id}")
async def get_post(id: int):
    try:
        cursor.execute("SELECT * FROM posts WHERE id = %s", (id,))
        post = cursor.fetchone()
        if not post:
            raise HTTPException(status_code=404, detail="Post not found")
        return {"post": post}
    except Exception as error:
        raise HTTPException(
            status_code=500, detail=f"Error fetching post: {error}")


@app.delete("/posts/{id}", status_code=status.HTTP_204_NO_CONTENT)
async def delete_post(id: int):
    cursor.execute("DELETE FROM posts WHERE id = %s RETURNING *", (id,))
    deleted_post = cursor.fetchone()
    conn.commit()
    if not deleted_post:
        raise HTTPException(status_code=404, detail="Post not found")
    return {"message": "Post deleted successfully"}


@app.put("/posts/{id}")
async def update_post(id: int, post: Post):
    cursor.execute(
        """
        UPDATE posts SET title = %s, content = %s, published = %s
        WHERE id = %s RETURNING *
        """,
        (post.title, post.content, post.published, id)
    )
    updated_post = cursor.fetchone()
    conn.commit()
    if not updated_post:
        raise HTTPException(status_code=404, detail="Post not found")
    return {"data": updated_post}

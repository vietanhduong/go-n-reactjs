import React, { useEffect, useState } from "react";
import { Redirect } from "react-router-dom";
import { getPost } from "../utils/api";

const Post = ({ match }) => {
    const {
        params: { id },
    } = match;

    const [post, setPost] = useState({});
    useEffect(() => getPost(id).then(post => setPost(post)).catch(err => setPost(undefined)), [id]);
    if (!post) {
        return (<Redirect to={"/404"} />);
    }
    return (
        <div className="content">
            <article>
                <header>
                    <h1>{post.title}</h1>
                </header>
                {post.content}
            </article>
        </div>
    );
}

export default Post;
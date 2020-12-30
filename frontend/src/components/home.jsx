import React, { useEffect, useState } from "react";
import { Link } from "react-router-dom";
import { fetchPosts } from "../utils/api";

const Home = () => {
    const [posts, setPosts] = useState([]);
    useEffect(() => fetchPosts().then(posts => setPosts(posts)).catch(err => console.log(err)), []);
    return (
        <div className="content">
            <h1>Posts: </h1>
            <table>
                <tbody>
                    {posts.map(e => (
                        <tr key={e.id}><td><Link to={"/posts/" + e.id}>{e.title}</Link></td></tr>
                    ))}
                </tbody>
            </table>
        </div>
    );
}

export default Home;

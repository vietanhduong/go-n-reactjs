import axios from "axios";

const API = process.env.REACT_APP_API || "";

export const fetchPosts = async () => {
    const res = await axios.get(`${API}/posts`);
    const data = await res.data;
    return data.content;
}

export const getPost = async (id) => {
    const res = await axios.get(`${API}/posts/${id}`);
    const data = await res.data;
    return data.content;
}
import React, { useState, useEffect } from "react"
import axios from "axios"

import TodoItem from "./TodoItem"

const TodoList = () => {
    const [inputTitle, setInputTitle] = useState("")
    const [inputContent, setInputContent] = useState("")
    const [todoItems, setTodoItems] = useState([])

    useEffect(() => {
        axios.get("http://localhost:9090/todo-items")
        .then((response) => {
            setTodoItems(response.data)
        })
        .catch((error) => {
            console.log(error)
        })
    }, [])
    const onInputTitle = (e) => setInputTitle(e.target.value)
    const onInputContent = (e) => setInputContent(e.target.value )

    const onAddItem = () => {
        axios.post("http://localhost:9090/todo", {
            title: inputTitle,
            content: inputContent,
        })
        .then(response => {
            if (response.status === 200) {
                const res = response.data
                const item = {
                    id: res.id,
                    title: res.title,
                    content: res.content,
                    date: res.date
                }
                const newItems = todoItems.concat(item)
                setTodoItems(newItems)
            } else {
                console.log(response)
            }
        })
        .catch(error => {
            console.log(error)
        })
    }
    const onDeleteItem = (itemId) => {
        axios.delete("http://localhost:9090/todo", {data: {id: itemId}})
        .then(response => {
            if (response.status === 200) {
                const newItems = []
                todoItems.forEach(item => {
                    if (item.id !== response.data.id) {
                        newItems.push(item)
                    }
                })
                setTodoItems(newItems)
            } else {
                console.log(response)
            }
        })
        .catch(error => {
            console.log(error)
        })
    }
    return (
        <div>
            <p>title: <input type="text" onChange={onInputTitle}/></p>
            <p>content: <input type="text" onChange={onInputContent}/></p>
            <button onClick={onAddItem}>Submit</button>
            <hr/>
            {todoItems && todoItems.map((item, index) => {
                return (
                    <div key={index} >
                        <TodoItem item={item} onDelete={() => onDeleteItem(item.id)}/>
                    </div>
                )
            })}
        </div>
    )
}

export default TodoList

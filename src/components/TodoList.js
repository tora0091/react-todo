import React, { useState, useEffect } from "react"
import axios from "axios"

import TodoItem from "./TodoItem"

const TodoList = () => {
    const [inputTitle, setInputTitle] = useState("")
    const [inputContent, setInputContent] = useState("")
    const [todoItems, setTodoItems] = useState([])
    const [isLoading, setIsLoading] = useState(false)
    const [isError, setIsError] = useState(false)

    const onInputTitle = (e) => setInputTitle(e.target.value)
    const onInputContent = (e) => setInputContent(e.target.value )

    useEffect(() => {
        setIsError(false)

        const fetchData = async () => {
            try {
                setIsLoading(true)
                const response = await axios.get("http://localhost:9090/todo-items")
                setTodoItems(response.data)
            } catch (error) {
                setIsError(true)
            }
        }

        fetchData()
        setIsLoading(false)
    }, [])

    const onAddItem = async () => {
        setIsError(false)

        try {
            const response = await axios.post("http://localhost:9090/todo", {
                title: inputTitle,
                content: inputContent,
            })

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
                setIsError(true)
                console.log(response)
            }
        } catch (error) {
            setIsError(true)
            console.log(error)
        }
    }
    const onDeleteItem = async (itemId) => {
        setIsError(false)

        try {
            const response = await axios.delete("http://localhost:9090/todo", {data: {id: itemId}})
            if (response.status === 200) {
                const newItems = []
                todoItems.forEach(item => {
                    if (item.id !== response.data.id) {
                        newItems.push(item)
                    }
                })
                setTodoItems(newItems)
            } else {
                setIsError(true)
                console.log(response)
            }
        } catch (error) {
            setIsError(true)
            console.log(error)
        }
    }
    return (
        <div>
            <p>title: <input type="text" onChange={onInputTitle}/></p>
            <p>content: <input type="text" onChange={onInputContent}/></p>
            <button onClick={onAddItem}>Submit</button>
            <hr/>

            {isError && <div>Sorry, Something wrong ...</div>}

            {isLoading ? (
                <div>Loading ...</div>
            ) : (
                todoItems && todoItems.map((item, index) => {
                    return (
                        <div key={index} >
                            <TodoItem item={item} onDelete={() => onDeleteItem(item.id)}/>
                        </div>
                    )
                })
            )}
        </div>
    )
}

export default TodoList

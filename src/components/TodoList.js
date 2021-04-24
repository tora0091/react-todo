import React from "react"
import axios from "axios"

import TodoItem from "./TodoItem"

class TodoList extends React.Component {
    constructor() {
        super()
        this.state = {
            todoItems: [],
            title: null,
            content: null,
        }
    }
    componentDidMount() {
        axios.get("http://localhost:9090/todo-items")
        .then((response) => {
            this.setState({todoItems: response.data})
        })
        .catch((error) => {
            console.log(error)            
        })
    }
    onInputTitle = (e) => {
        this.setState({ title: e.target.value })
    }
    onInputContent = (e) => {
        this.setState({ content: e.target.value })
    }
    onAddItem = () => {
        axios.post("http://localhost:9090/todo", { 
            title: this.state.title,
            content: this.state.content,
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
                const newItems = this.state.todoItems.concat(item)
                this.setState({todoItems: newItems})    
            } else {
                console.log(response)
            }
        })
        .catch(error => {
            console.log(error)
        })
    }
    onDeleteItem = (itemId) => {
        axios.delete("http://localhost:9090/todo", {data: {id: itemId}})
        .then(response => {
            if (response.status === 200) {
                const newItems = []
                this.state.todoItems.forEach(item => {
                    if (item.id !== response.data.id) {
                        newItems.push(item)
                    }
                })
                this.setState({todoItems: newItems})        
            } else {
                console.log(response)
            }
        })
        .catch(error => {
            console.log(error)
        })
    }
    render() {
        const todoItems = this.state.todoItems
        return (
            <div>
                <p>title: <input type="text" onInput={this.onInputTitle}/></p>
                <p>content: <input type="text" onInput={this.onInputContent}/></p>
                <button onClick={this.onAddItem}>Submit</button>
                <hr/>
                {todoItems.map((item, index) => {
                    return (
                        <div key={index} >
                            <TodoItem item={item} onDelete={() => this.onDeleteItem(item.id)}/>
                        </div>
                    )
                })}
            </div>
        )
    }
}

export default TodoList

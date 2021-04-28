import React from 'react'

const TodoItem = (props) => {
    const { item, onDelete } = props
    return (
        <ul>
            <li>
                <p>{item.id}: {item.title} : {item.date}</p>
                <p>{item.content}</p>
                <button onClick={() => onDelete()}>Delete</button>
            </li>
        </ul>
    )
}

export default TodoItem

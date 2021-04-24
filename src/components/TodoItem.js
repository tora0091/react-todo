import React from 'react'

class TodoItem extends React.Component {
    render() {
        const { item } = this.props
        return (
            <ul>
                <li>
                    <p>{item.id}: {item.title} : {item.date}</p>
                    <p>{item.content}</p>
                    <button onClick={this.props.onDelete}>Delete</button>
                </li>
            </ul>
        )
    }
}

export default TodoItem

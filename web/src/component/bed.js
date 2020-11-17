import React from 'react';

class Bed extends React.Component {
    constructor(props) {
        super(props)
        this.state = this.props.bedObj;
        this.handleChange = this.handleChange.bind(this)
        this.handleSubmit = this.handleSubmit.bind(this)
        this.postAPIData = this.postAPIData.bind(this)
    }

    async postAPIData() {
        let toSend = this.state
        toSend.IsOut = true
        const url = "/api"
        const response = await fetch(url, {
            method: 'POST',
            body: JSON.stringify(toSend)
        })
        this.setState(response)
    }

    handleChange(event){
        this.setState({
            Comments: event.target.value
        });
    }

    handleSubmit(event) {
        this.postAPIData()
        event.preventDefault();
    }

    render () {
        let someJSX = null
        someJSX = <div>
            <p>id {this.state.Id}</p>
            <p>Room {this.state.Room}</p>
            <p>Side {this.state.Name}</p>
            <p>wing {this.state.Wing}</p>
            <p>time out {this.state.TimeOut}</p>
            <form>
                <label>
                    Comments:
                    <input type="text" value={this.state.Comments} onChange={this.handleChange} />
                </label>
                <input type="submit" value="Submit" />
            </form>
        </div>
        return someJSX;
    }
}

export default Bed;
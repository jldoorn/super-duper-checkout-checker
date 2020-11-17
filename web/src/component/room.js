import React from 'react';
import Bed from './bed'

class Room extends React.Component {
    constructor(props) {
        super(props)
        this.state = {beds: null}
        this.getAPIData = this.getAPIData.bind(this)
    }

    // componentDidMount(){
    //     this.getAPIData(this.props.roomNum)
    // }

    async getAPIData(roomnum){
        const url = "/api/room/" + roomnum
        // const response = await fetch(url, {
        //     method: 'GET',
        //     body: JSON.stringify({RequestType: 1,
        //     RequestBody: String(roomnum)})
        // })
        const response = await (await fetch(url)).json()
        console.log(response)
        let beds = response.Beds.map((item)=><Bed bedObj={item}></Bed>)
        this.setState({beds: beds})
    }

    componentDidMount(){
        this.getAPIData(2)
    }

    render () {
        // let bedObj = {
        //     Id: 3,
        //     Room: 2,
        //     Wing: "ne2",
        //     Name: "left",
        //     IsOut: false,
        //     TimeOut: -1,
        //     Ra1: 0,
        //     Ra2: 0,
        //     Comments: ""
        // }
        // let myBed = <Bed bedObj={bedObj}></Bed>
        return <div>{this.state.beds}</div>;
    }
}

export default Room;
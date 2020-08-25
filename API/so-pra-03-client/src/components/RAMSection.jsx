import React from 'react';
import axios from 'axios';
import {LineChart, Line, YAxis, XAxis, CartesianGrid} from "recharts";
import {Button, message, Statistic} from "antd";

export default class RAMSection extends React.Component {

    lastValue = null;
    charLen = 40;

    constructor(props){
        super(props);

        const data = Array.apply(null, Array(this.charLen))
            .map((x, i) => ({ totalMb: 0, usedMb: 0, total: 0, used: 0, usage: 0 }) );

        this.state = { data: data };
    }

    componentDidMount() {
        this.updateChar();
    }

    updateChar = async () => {

        function sleep(ms) {
            return new Promise(resolve => setTimeout(resolve, ms));
        }

        try {
            while (true){

                const response = await axios.get('/ram');
                await sleep(1000);

                this.setState((state, props) => {
                    const values = state.data.slice(1);

                    const newValue = {};
                    // newValue.used = response.data.TotalConsumida || 4000000;
                    // newValue.total = response.data.TotalServer || 8000000;
                    // newValue.usedMb = newValue.used / 1024;
                    // newValue.totalMb = newValue.total / 1024;
                    // newValue.usage = newValue.used * 100 / newValue.total;
                    newValue.used = response.data.TotalConsumida * 1024;
                    newValue.total = response.data.TotalServer * 1024;
                    newValue.usedMb = response.data.TotalConsumida;
                    newValue.totalMb = response.data.TotalServer;
                    newValue.usage = newValue.usedMb * 100 / newValue.totalMb;

                    values.push(newValue);
                    return { data: values };
                });
            }
        }
        catch (e) {
            console.log('ocurrio un error');
            console.log(e);
            message.error("Ocurrio un error, ya no se pudo actualizar");
        }
    };


    render() {

        const last = this.state.data[this.state.data.length - 1];

        return (
            <div style={{ background: "white", textAlign: "center" }}>

                <Statistic title="Total MB" value={last.totalMb.toFixed(2)}  />
                <Statistic title="Usado MB" value={last.usedMb.toFixed(2)}  />
                <Statistic title="RAM%" value={last.usage.toFixed(2)}  />

                <LineChart style={{ margin: "0 auto" }} width={1200} height={400} data={this.state.data} >
                    <Line type="monotone" dataKey="usedMb" stroke="blue" />
                    <CartesianGrid stroke="#ccc" />
                    <XAxis  />
                    <YAxis domain={[0, parseInt(last.totalMb, 10) ]} />
                </LineChart>

            </div>
        );
    }

}

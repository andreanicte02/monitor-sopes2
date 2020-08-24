import React from 'react';
import axios from 'axios';
import {LineChart, Line, YAxis, XAxis, CartesianGrid} from "recharts";
import {Button, message, Statistic} from "antd";

export default class CPUSection extends React.Component {

    lastValue = null;
    charLen = 40;

    constructor(props){
        super(props);

        const data = Array.apply(null, Array(this.charLen))
            .map((x, i) => ({ uv: 0 }) );

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

                const response = await axios.get('/cpu');
                await sleep(1000);
                // console.log('obtuve del server:', response.data);

                if (this.lastValue) {
                    const cpuDelta = response.data.total - this.lastValue.total;
                    const cpuIdle = response.data.idle - this.lastValue.idle;
                    const cpuUsed = cpuDelta - cpuIdle;
                    const cpuUsage = 100 * cpuUsed / cpuDelta;

                    const cpuValue = { uv: cpuUsage };

                    this.setState((state, props) => {
                        const data = state.data.slice(1);
                        data.push(cpuValue);
                        return { data: data };
                    });
                }

                this.lastValue = response.data;
            }
        }
        catch (e) {
            console.log('ocurrio un error');
            console.log(e);
            message.error("Ocurrio un error, ya no se pudo actualizar");
        }
    };


    render() {

        const cpuUsage = this.state.data[this.state.data.length - 1].uv;

        return (
            <div style={{ background: "white" }}>

                <Statistic title="CPU%" value={cpuUsage.toFixed(2)}  />

                <LineChart style={{ margin: "0 auto" }} width={1200} height={400} data={this.state.data} >
                    <Line type="monotone" dataKey="uv" stroke="blue" />
                    <CartesianGrid stroke="#ccc" />
                    <XAxis  />
                    <YAxis domain={[0, 100]} />
                </LineChart>

            </div>
        );
    }

}
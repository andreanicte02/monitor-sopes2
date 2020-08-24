import React from 'react';
// import logo from './logo.svg';
import './App.css';
import 'antd/dist/antd.css';
import LoginSection from "./components/LoginSection";
import ProcessesSection from "./components/ProcessesSection";
import {Carousel} from "antd";
import CPUSection from "./components/CPUSection";
import RAMSection from "./components/RAMSection";


export default class App extends React.Component {

    constructor(props){
        super(props);
        this.state = { logged: false };
    }

    onLoginSuccess = () => {
        this.setState({ logged: true })
    };

    render() {
        return (
            <div className="App">
                { !this.state.logged && <LoginSection onLoginSuccess={this.onLoginSuccess} /> }

                { this.state.logged && (<Carousel dotPosition="top"  >
                    <div key="1">
                        <h3>Procesos</h3>
                        <ProcessesSection />
                    </div>
                    {/*<div key="2">*/}
                        {/*<h3>CPU</h3>*/}
                        {/*<CPUSection />*/}
                    {/*</div>*/}
                    <div key="3">
                        <h3>Ram</h3>
                        <RAMSection />
                    </div>

                </Carousel>) }




            </div>
        );
    }

}


// function App() {
//   return (
//     <div className="App">
//         <Login/>
//     </div>
//   );
// }
//
// export default App;

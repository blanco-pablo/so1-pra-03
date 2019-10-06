import React from 'react';
// import logo from './logo.svg';
import './App.css';
import 'antd/dist/antd.css';
import LoginSection from "./components/LoginSection";
import ProcessesSection from "./components/ProcessesSection";
import {Carousel} from "antd";


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
                {/*{ !this.state.logged && <LoginSection onLoginSuccess={this.onLoginSuccess} /> }*/}
                {/*{ !this.state.logged && <ProcessesSection /> }*/}

                <Carousel dotPosition="top"  >
                    <div>
                        <h3>Procesos</h3>
                        <ProcessesSection />
                    </div>
                    <div>
                        <h3>2</h3>
                    </div>
                    <div>
                        <h3>3</h3>
                    </div>

                </Carousel>

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

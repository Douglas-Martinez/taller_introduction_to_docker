import React, { Component } from 'react';
import api from '../api';

export class Ejemplo extends Component {

  constructor(props) {
    super(props);

    this.state = {
      data: '',
      error: null
    }
  }

  fetchData = async () => {
    try {
      const txt = await api.getRoot();
      //console.log(txt)

      this.setState({ data: txt.data})
      
    } catch (error) {
      this.setState({ error: error });
      console.error(error);
    }
  }

  componentDidMount() {
    this.fetchData();  }

  render() {
    return (
      <div className="container">
        <br></br>
        <h1>{this.state.data}</h1>
        <h3>Buenas noches</h3>
      </div>
    )
  }
}

export default Ejemplo
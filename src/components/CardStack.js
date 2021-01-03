import React, { Component } from 'react';
import axios from 'axios';

import Hearts from '../Hearts.png'
import Diamonds from '../Diamonds.png'
import Clubs from '../Clubs.png'
import Spades from '../Spades.png'



class CardStack extends Component {
  constructor(props) {
    super(props);

    this.state = {
      suit: '',
      value: '',
      Hearts: ''
    };

    this.nextCard = this.nextCard.bind(this)
  }

  componentDidMount() {
    axios.get('http://localhost:3001/draw')
      .then(response => {
        console.log(response)
        console.log(response.data)
        console.log(response.data.value)
        this.setState({suit: response.data.suit})
        this.setState({value: response.data.value})
        if(response.data.suit === "Heart") {
          document.getElementById('cardSuit').src = Hearts
        } else if(response.data.suit === "Diamond") {
          document.getElementById('cardSuit').src = Diamonds
        } else if(response.data.suit === "Club") {
          document.getElementById('cardSuit').src = Clubs
        } else if(response.data.suit === "Spade") {
          document.getElementById('cardSuit').src = Spades
        }
      })
      .catch(error => {
        console.log(error)
      })
  }

  nextCard() {
    axios.get('http://localhost:3001/draw')
      .then(response => {
        console.log(response)
        console.log(response.data)
        this.setState({suit: response.data.suit})
        this.setState({value: response.data.value})

        if(response.data.suit === "Heart") {
          document.getElementById('cardSuit').src = Hearts
        } else if(response.data.suit === "Diamond") {
          document.getElementById('cardSuit').src = Diamonds
        } else if(response.data.suit === "Club") {
          document.getElementById('cardSuit').src = Clubs
        } else if(response.data.suit === "Spade") {
          document.getElementById('cardSuit').src = Spades
        }
      })
      .catch(error => {
        console.log(error)
      })
  }

  render() {
    const { suit } = this.state
    const { value } = this.state
    return (
      <div className="middle">
        <div className="card">
          <div className={suit}>
            <img src={suit} alt="placeholder text" id="cardSuit"/>
            <div className="value">
              {value}
            </div>
          </div>
        </div>
        <button className='nextCard' onClick={this.nextCard}>Next Card</button>
      </div>
    )
  }
}

export default CardStack

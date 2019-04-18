import React from 'react'
import axios from '../utils/axios'

class Player extends React.Component {
  static async getInitialProps({ query }) {
    if (!query.vid) {
      return { error: true, errorMessage: 'No Vid provided!' }
    }
    try {
    const resp = await axios.get(`/${query.vid}`)
    const m3u8 = resp.data
      return { m3u8 }
    } catch (error) {
      return { error: true, errorMessage: 'Video Not Found.' }
    }
  }
  render() {
    if (this.props.error) {
      return (
        <main>
          <h1>{this.props.errorMessage}</h1>
        </main>
      )
    }
    return (
      <main>
        <pre>
          {this.props.m3u8}
        </pre>
      </main>
    )
  }
}

export default Player
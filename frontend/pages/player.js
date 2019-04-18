import React from 'react'
import Head from 'next/head'
import { backend } from '../config'
import VideoPlayer from '../components/VideoPlayer'

class Player extends React.Component {
  static async getInitialProps({ query }) {
    if (!query.vid) {
      return { error: true, errorMessage: 'No Vid provided!' }
    }
    return { vid: query.vid }
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
        <Head>
          <link rel="stylesheet" href="//vjs.zencdn.net/5.12/video-js.css" />
        </Head>
        <VideoPlayer
          controls={true}
          sources={[{
            src: `${backend}/${this.props.vid}`,
            type: 'application/x-mpegURL',
          }]}
        />
      </main>
    )
  }
}

export default Player
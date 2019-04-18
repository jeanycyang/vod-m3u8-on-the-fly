import Link from 'next/link'

function Home() {
  return (<main>
    <h1>Jean's Video Library</h1>
    <h2>Video List</h2>
    <ul>
      <li><Link href="/player?vid=vod1.mov"><a>vod1.mov</a></Link></li>
      <li><Link href="/player?vid=vod2.mov"><a>vod2.mov</a></Link></li>
    </ul>
  </main>)
}

export default Home
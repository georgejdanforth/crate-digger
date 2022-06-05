import type { NextPage } from 'next'
import Head from 'next/head'

const Home: NextPage = () => {
  return (
    <>
      <Head>
        <title>Crate Digger — Search</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <main>
        <h1 className="h-1 text-xl">Welcome to Crate Digger</h1>
      </main>

      <footer/>
    </>
  )
}

export default Home

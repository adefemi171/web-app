import React from "react";
import Head from 'next/head'
import styles from '../styles/Home.module.css'
import Link from "next/link"

export function Links() {

  return (
    <div className={styles.container}>
    <Head>
      <title>Info Paper</title>
      <link rel="icon" href="/favicon.ico" />
    </Head>

    <main className={styles.main}>
      <h1 className={styles.title}>
        Welcome to Info Paper!
      </h1>
      <Link href="/table">
          <a className={styles.card}>
            <span>View Paper &rarr;</span>
          </a>
      </Link>
    </main>
  </div>
  )
}
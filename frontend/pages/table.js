import React from 'react';
import Head from 'next/head'

import tableStyles from '../styles/Table.module.css'


function Force(props) {
  const keys = Object.entries(props.headers)

  return (
    <>
      <Head>
        <title>Response Header Table</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>
     
      <div className={tableStyles.container}>
        <div className={tableStyles.main}>
          <table className={tableStyles.table}>
              <tbody>
                  <tr className={tableStyles.tr}>
                      <th className={tableStyles.th}>Key</th>
                      <th className={tableStyles.th}>Value</th>
                  </tr>
                  {keys.map((key) => (
                  <tr className={tableStyles.tr}>
                      <td className={tableStyles.td}>{key[0]}</td>
                      <td className={tableStyles.td}>{key[1]}</td>
                  </tr>
                  ))} 
              </tbody>
          </table> 
        </div>
      </div>
    </>
  );
}

export async function getServerSideProps(context) {
  // Fetch data from external API
  const { req } = context;
  console.log({ headers: req.headers });
  // console.log({ context, req });
  const res = await fetch("http://ec2-18-220-106-234.us-east-2.compute.amazonaws.com:8080");
  const data = await res.json();

  // Pass data to the page via props
  return { props: { data, headers: req.headers } };
}

export default Force;
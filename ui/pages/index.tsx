import Head from 'next/head'
import Image from 'next/image'
import LoginForm from '../components/LoginForm';
import Title from '../components/Title';
import Box from '@material-ui/core/Box';
import DevNotes from '../components/DevNotes';
import { Container } from 'next/app';
import { Grid } from '@material-ui/core';

export default function Home({ data }) {

  console.log('From getStaticProps: ', data);

  return (
    <Grid container spacing={3} justifyContent='center'>
      <Grid item xs={6} sm={3} >
        <Box>
          <i>idk what to put here, probably a popular playlist</i>
        </Box>
      </Grid>
      <Grid item xs={6} sm={3} >
        <Box>
          <Title />
          <LoginForm />
        </Box>
      </Grid>
      <Grid item xs={6} sm={3} >
        <Box>
          <DevNotes />
        </Box>
      </Grid>
    </Grid>
  );
}

// Use this for api requests
export async function getStaticProps() {

  /** Why to use this:
   *  Better SEO 🕶
      Performance 🚀
      Can be hosted in CDN 🌍
      Doesn't need to have JavaScript to run (mostly HTML) ⚙️
      Very fewer things to parse from server to client 🌬 
  */


  const res = 'put API query here';

  return {
    props: {
      data: res,
    }
  }
}
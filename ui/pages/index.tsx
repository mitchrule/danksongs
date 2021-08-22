import Head from 'next/head'
import Image from 'next/image'
import LoginForm from '../components/LoginForm';
import Title from '../components/Title';
import Box from '@material-ui/core/Box';
import DevNotes from '../components/DevNotes';
import { Container } from 'next/app';
import { Grid } from '@material-ui/core';

export default function Home() {
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

import Head from 'next/head'
import Image from 'next/image'
import styles from '../styles/Home.module.css'
import LoginForm from '../components/LoginForm';
import Title from '../components/Title';
import Box from '@material-ui/core/Box';
import DevNotes from '../components/DevNotes';

export default function Home() {
  return (
  <Box>
    <Box>
      test
    </Box>
    <Box>
      <Title />
      <LoginForm />
    </Box>
    <Box>
      <DevNotes />
    </Box>
  </Box>
  );
}

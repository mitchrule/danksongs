import { Typography } from "@material-ui/core";
import styles from '../styles/Title.module.css';
import Image from 'next/image';

export default function Title() {
  return (
      <Image 
        src="/../public/logo.gif" 
        alt="Picture of the author" 
        width={500}
        height={500}
       />
  );
}

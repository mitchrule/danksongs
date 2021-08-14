import { Box, Typography, useTheme } from "@material-ui/core";
import styles from '../styles/Title.module.css';

export default function Title() {
  return (
      <Typography className={styles.titleText}>Danksongs</Typography>
  );
}

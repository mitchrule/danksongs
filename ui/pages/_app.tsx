import "../styles/globals.css";
import type { AppProps } from "next/app";
import Title from "../components/Title";

function MyApp({ Component, pageProps }: AppProps) {
  return (
    <>
      <Title></Title>
    </>
  );
}
export default MyApp;
